// Copyright 2025 Akamai Technologies, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package policy

import (
	"context"

	"github.com/minio/minio-go/v7"
)

type Client struct {
	policy GetSetter
}

type GetSetter interface {
	SetBucketPolicy(ctx context.Context, bucketName string, policy string) error
	GetBucketPolicy(ctx context.Context, bucketName string) (string, error)
}

var _ GetSetter = (*minio.Client)(nil)

func NewClient(gs GetSetter) *Client {
	return &Client{policy: gs}
}

func (c *Client) Get(ctx context.Context) error
func (c *Client) Set(ctx context.Context) error
