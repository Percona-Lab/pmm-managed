package models

import (
	"encoding/json"

	"github.com/pkg/errors"
	"gopkg.in/reform.v1"
)

// SaveChannel persists notification channel.
func SaveChannel(q reform.DBTX, c *Channel) error {
	if err := ValidateChannel(c); err != nil {
		return errors.Wrap(err, "channel validation failed")
	}

	b, err := json.Marshal(c)
	if err != nil {
		return errors.Wrap(err, "failed to marshall notification channel")
	}

	_, err = q.Exec("INSERT INTO notification_channels (id, channel) VALUES ($1, $2)", c.Id, b)
	if err != nil {
		return errors.Wrap(err, "failed to create notifications channel")
	}

	return nil
}

// UpdateChannel updates existing notifications channel.
func UpdateChannel(q reform.DBTX, c *Channel) error {
	if err := ValidateChannel(c); err != nil {
		return errors.Wrap(err, "channel validation failed")
	}
	b, err := json.Marshal(c)
	if err != nil {
		return errors.Wrap(err, "failed to marshall notification channel")
	}

	_, err = q.Exec("UPDATE notification_channels SET channel=$1 WHERE id=$2", b, c.Id)
	if err != nil {
		return errors.Wrap(err, "failed to create notifications channel")
	}

	return nil
}

// RemoveChannel removes notification channel with specified id.
func RemoveChannel(q reform.DBTX, id string) error {
	_, err := q.Exec("DELETE FROM notification_channels WHERE id=$1", id)
	if err != nil {
		return errors.Wrap(err, "failed to delete notifications channel")
	}
	return nil
}

// GetChannels returns saved notification channels configuration.
func GetChannels(q reform.DBTX) ([]Channel, error) {
	rows, err := q.Query("SELECT channel FROM notification_channels")
	if err != nil {
		return nil, errors.Wrap(err, "failed to select notification channels")
	}

	var channels []Channel
	for rows.Next() {
		var b []byte
		if err = rows.Scan(&b); err != nil {
			break
		}

		var channel Channel
		if err = json.Unmarshal(b, &channel); err != nil {
			break
		}
		channels = append(channels, channel)
	}

	if closeErr := rows.Close(); closeErr != nil {
		return nil, errors.Wrap(err, "failed to close rows")
	}

	if err != nil {
		return nil, errors.Wrap(err, "failed to read notification channels")
	}

	if err = rows.Err(); err != nil {
		return nil, errors.Wrap(err, "failed to scan rows")
	}

	return channels, nil
}

// ValidateChannel validates notification channel.
func ValidateChannel(ch *Channel) error {
	if ch.Id == "" {
		return errors.New("notification channel id is empty")
	}

	switch ch.Type {
	case Email:
		if ch.SlackConfig != nil || ch.WebHookConfig != nil {
			return errors.New("email channel should has only email configuration")
		}

		return validateEmailConfig(ch.EmailConfig)
	case Slack:
		if ch.EmailConfig != nil || ch.WebHookConfig != nil {
			return errors.New("slack channel should has only slack configuration")
		}

		return validateSlackConfig(ch.SlackConfig)
	case WebHook:
		if ch.SlackConfig != nil || ch.EmailConfig != nil {
			return errors.New("webhook channel should has only webhook configuration")
		}

		return validateWebHookConfig(ch.WebHookConfig)
	case "":
		return errors.New("notification channel type is empty")
	default:
		return errors.Errorf("unknown channel type %s", ch.Type)
	}
}

func validateEmailConfig(c *EmailConfig) error {
	if c == nil {
		return errors.New("email config is empty")
	}

	if len(c.To) == 0 {
		return errors.New("email to field is empty")
	}

	return nil
}

func validateSlackConfig(c *SlackConfig) error {
	if c == nil {
		return errors.New("slack config is empty")
	}

	if c.Channel == "" {
		return errors.New("slack channel field is empty")
	}

	return nil
}

func validateWebHookConfig(c *WebHookConfig) error {
	if c == nil {
		return errors.New("webhook config is empty")
	}

	if c.Url == "" {
		return errors.New("webhook url field is empty")
	}

	return nil
}
