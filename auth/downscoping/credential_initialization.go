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

// package downscopingoverview contains Google Cloud auth snippets showing how to
// downscope credentials with Credential Access Boundaries.
// https://cloud.google.com/iam/docs/downscoping-short-lived-credentials
package downscopingoverview

// [START auth_overview_downscoped_credential_initialization]

import (
	"context"
	"fmt"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"golang.org/x/oauth2/google/downscope"
)

// initializeCredentials will return a downscoped token using the provided
// context and Access Boundary Rules.
func initializeCredentials(ctx context.Context, accessBoundary []downscope.AccessBoundaryRule) (*oauth2.Token, error) {
	// This Source can be initialized in multiple ways; the following example uses
	// Application Default Credentials.
	var rootSource oauth2.TokenSource

	// You must provide the "https://www.googleapis.com/auth/cloud-platform" scope.
	rootSource, err := google.DefaultTokenSource(ctx, "https://www.googleapis.com/auth/cloud-platform")
	if err != nil {
		return nil, fmt.Errorf("failed to generate rootSource: %v", err)
	}

	// downscope.NewTokenSource constructs the token source with the configuration provided.
	dts, err := downscope.NewTokenSource(ctx, downscope.DownscopingConfig{RootSource: rootSource, Rules: accessBoundary})
	if err != nil {
		return nil, fmt.Errorf("failed to generate downscoped token source: %v", err)
	}
	// Token() uses the previously declared TokenSource to generate a downscoped token.
	tok, err := dts.Token()
	if err != nil {
		return nil, fmt.Errorf("failed to generate token: %v", err)
	}
	return tok, nil
}

// [END auth_overview_downscoped_credential_initialization]
