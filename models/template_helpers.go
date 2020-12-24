// pmm-managed
// Copyright (C) 2017 Percona LLC
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program. If not, see <https://www.gnu.org/licenses/>.

package models

import (
	"time"

	"github.com/percona-platform/saas/pkg/alert"
	"github.com/pkg/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gopkg.in/reform.v1"
)

func checkUniqueTemplateName(q *reform.Querier, name string) error {
	if name == "" {
		panic("empty template name")
	}

	template := &Template{Name: name}
	switch err := q.Reload(template); err {
	case nil:
		return status.Errorf(codes.AlreadyExists, "Template with name %q already exists.", name)
	case reform.ErrNoRows:
		return nil
	default:
		return errors.WithStack(err)
	}
}

// FindTemplates returns saved notification rule templates.
func FindTemplates(q *reform.Querier) ([]Template, error) {
	structs, err := q.SelectAllFrom(TemplateTable, "")
	if err != nil {
		return nil, errors.Wrap(err, "failed to select notification rule templates")
	}

	templates := make([]Template, len(structs))
	for i, s := range structs {
		c := s.(*Template)

		templates[i] = *c
	}

	return templates, nil
}

// FindTemplateByName finds template by name.
func FindTemplateByName(q *reform.Querier, name string) (*Template, error) {
	if name == "" {
		return nil, status.Error(codes.InvalidArgument, "Empty template name.")
	}

	template := &Template{Name: name}
	switch err := q.Reload(template); err {
	case nil:
		return template, nil
	case reform.ErrNoRows:
		return nil, status.Errorf(codes.NotFound, "Template with name %q not found.", name)
	default:
		return nil, errors.WithStack(err)
	}
}

// CreateTemplateParams are params for creating new rule template.
type CreateTemplateParams struct {
	Template *alert.Template
	Yaml     string
	Source   Source
}

// CreateTemplate creates rule template.
func CreateTemplate(q *reform.Querier, params *CreateTemplateParams) (*Template, error) {
	template := params.Template
	if err := template.Validate(); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "Invalid rule template: %v.", err)
	}

	if err := checkUniqueTemplateName(q, params.Template.Name); err != nil {
		return nil, err
	}

	p, err := convertTemplateParams(params.Template.Params)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "Invalid rule template parameters: %v.", err)
	}

	row := &Template{
		Name:     template.Name,
		Version:  template.Version,
		Summary:  template.Summary,
		Tiers:    template.Tiers,
		Expr:     template.Expr,
		Params:   p,
		For:      time.Duration(template.For),
		Severity: convertSeverity(template.Severity),
		Source:   params.Source,
		Yaml:     params.Yaml,
	}

	if err := row.SetLabels(template.Labels); err != nil {
		return nil, err
	}

	if err := row.SetAnnotations(template.Annotations); err != nil {
		return nil, err
	}

	err = q.Insert(row)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create rule template")
	}

	return row, nil
}

// ChangeTemplateParams is params for changing existing rule template.
type ChangeTemplateParams struct {
	Template *alert.Template
	Yaml     string
}

// ChangeTemplate updates existing rule template.
func ChangeTemplate(q *reform.Querier, params *ChangeTemplateParams) (*Template, error) {
	row, err := FindTemplateByName(q, params.Template.Name)
	if err != nil {
		return nil, err
	}

	template := params.Template
	if err := template.Validate(); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "Invalid rule template: %v.", err)
	}

	p, err := convertTemplateParams(params.Template.Params)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "Invalid rule template parameters: %v.", err)
	}

	row.Name = template.Name
	row.Version = template.Version
	row.Summary = template.Summary
	row.Tiers = template.Tiers
	row.Expr = template.Expr
	row.Params = p
	row.For = time.Duration(template.For)
	row.Severity = convertSeverity(template.Severity)
	row.Yaml = params.Yaml

	if err := row.SetLabels(template.Labels); err != nil {
		return nil, err
	}

	if err := row.SetAnnotations(template.Annotations); err != nil {
		return nil, err
	}

	err = q.Update(row)
	if err != nil {
		return nil, errors.Wrap(err, "failed to update rule template")
	}

	return row, nil
}

// RemoveTemplate removes rule template with specified name.
func RemoveTemplate(q *reform.Querier, name string) error {
	if _, err := FindTemplateByName(q, name); err != nil {
		return err
	}

	rules, err := FindRules(q)
	if err != nil {
		return err
	}

	for _, rule := range rules {
		if name == rule.TemplateName {
			return errors.Errorf("failed to delete rule template, as it is being used by rule: %s", rule.ID)
		}
	}

	err = q.Delete(&Template{Name: name})
	if err != nil {
		return errors.Wrap(err, "failed to delete rule template")
	}
	return nil
}

func convertTemplateParams(params []alert.Parameter) (TemplateParams, error) {
	res := make(TemplateParams, len(params))
	for i, param := range params {
		t, err := convertParamType(param.Type)
		if err != nil {
			return nil, err
		}

		res[i] = TemplateParam{
			Name:    param.Name,
			Summary: param.Summary,
			Unit:    param.Unit,
			Type:    t,
		}

		switch param.Type {
		case alert.Float:
			var fp FloatParam
			var err error
			fp.Default, err = param.GetValueForFloat()
			if err != nil {
				return nil, errors.Wrap(err, "failed to parse param value")
			}

			fp.Min, fp.Max, err = param.GetRangeForFloat()
			if err != nil {
				return nil, errors.Wrap(err, "failed to parse param range")
			}

			res[i].FloatParam = &fp
		}
	}

	return res, nil
}

func convertParamType(paramType alert.Type) (ParamType, error) {
	switch paramType {
	case alert.Float:
		return Float, nil
	default:
		return "", errors.Errorf("UnknownSeverity parameter type %s", paramType)
	}
}
