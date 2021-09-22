// Copyright 2021 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go_gapic. DO NOT EDIT.

package resourcemanager_test

import (
	"context"

	resourcemanager "cloud.google.com/go/resourcemanager/apiv3"
	"google.golang.org/api/iterator"
	resourcemanagerpb "google.golang.org/genproto/googleapis/cloud/resourcemanager/v3"
	iampb "google.golang.org/genproto/googleapis/iam/v1"
)

func ExampleNewTagKeysClient() {
	ctx := context.Background()
	c, err := resourcemanager.NewTagKeysClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	// TODO: Use client.
	_ = c
}

func ExampleTagKeysClient_ListTagKeys() {
	ctx := context.Background()
	c, err := resourcemanager.NewTagKeysClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &resourcemanagerpb.ListTagKeysRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/resourcemanager/v3#ListTagKeysRequest.
	}
	it := c.ListTagKeys(ctx, req)
	for {
		resp, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			// TODO: Handle error.
		}
		// TODO: Use resp.
		_ = resp
	}
}

func ExampleTagKeysClient_GetTagKey() {
	ctx := context.Background()
	c, err := resourcemanager.NewTagKeysClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &resourcemanagerpb.GetTagKeyRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/resourcemanager/v3#GetTagKeyRequest.
	}
	resp, err := c.GetTagKey(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleTagKeysClient_CreateTagKey() {
	ctx := context.Background()
	c, err := resourcemanager.NewTagKeysClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &resourcemanagerpb.CreateTagKeyRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/resourcemanager/v3#CreateTagKeyRequest.
	}
	op, err := c.CreateTagKey(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}

	resp, err := op.Wait(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleTagKeysClient_UpdateTagKey() {
	ctx := context.Background()
	c, err := resourcemanager.NewTagKeysClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &resourcemanagerpb.UpdateTagKeyRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/resourcemanager/v3#UpdateTagKeyRequest.
	}
	op, err := c.UpdateTagKey(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}

	resp, err := op.Wait(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleTagKeysClient_DeleteTagKey() {
	ctx := context.Background()
	c, err := resourcemanager.NewTagKeysClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &resourcemanagerpb.DeleteTagKeyRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/resourcemanager/v3#DeleteTagKeyRequest.
	}
	op, err := c.DeleteTagKey(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}

	resp, err := op.Wait(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleTagKeysClient_GetIamPolicy() {
	ctx := context.Background()
	c, err := resourcemanager.NewTagKeysClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &iampb.GetIamPolicyRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/iam/v1#GetIamPolicyRequest.
	}
	resp, err := c.GetIamPolicy(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleTagKeysClient_SetIamPolicy() {
	ctx := context.Background()
	c, err := resourcemanager.NewTagKeysClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &iampb.SetIamPolicyRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/iam/v1#SetIamPolicyRequest.
	}
	resp, err := c.SetIamPolicy(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleTagKeysClient_TestIamPermissions() {
	ctx := context.Background()
	c, err := resourcemanager.NewTagKeysClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &iampb.TestIamPermissionsRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/iam/v1#TestIamPermissionsRequest.
	}
	resp, err := c.TestIamPermissions(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}
