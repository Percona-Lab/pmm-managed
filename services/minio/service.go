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

package minio

import "github.com/minio/minio-go"

type Service struct {
}

// BucketExists return true if bucket can be accessed with provided credentials and exists.
func (s *Service) BucketExists(host string, secure bool, accessKey, secretKey, name string) (bool, error) {
	minioClient, err := minio.New(host, accessKey, secretKey, secure)
	if err != nil {
		return false, err
	}
	return minioClient.BucketExists(name)
}

// GetBucketLocation retrieves bucket location by specified bucket name.
func (s *Service) GetBucketLocation(host string, secure bool, accessKey, secretKey, name string) (string, error) {
	minioClient, err := minio.New(host, accessKey, secretKey, secure)
	if err != nil {
		return "", err
	}
	return minioClient.GetBucketLocation(name)
}
