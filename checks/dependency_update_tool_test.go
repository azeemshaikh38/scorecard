// Copyright 2020 Security Scorecard Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package checks

import (
	"testing"

	"github.com/golang/mock/gomock"

	"github.com/ossf/scorecard/v4/checker"
	mockrepo "github.com/ossf/scorecard/v4/clients/mockclients"
	scut "github.com/ossf/scorecard/v4/utests"
)

// TestDependencyUpdateTool tests the DependencyUpdateTool checker.
func TestDependencyUpdateTool(t *testing.T) {
	t.Parallel()
	//nolint
	tests := []struct {
		name     string
		wantErr  bool
		files    []string
		want     checker.CheckResult
		expected scut.TestReturn
	}{
		{
			name:    "dependency yml",
			wantErr: false,
			files: []string{
				".github/dependabot.yml",
			},
			expected: scut.TestReturn{
				NumberOfInfo: 1,
				Score:        10,
			},
		},
		{
			name:    "dependency yaml ",
			wantErr: false,
			files: []string{
				".github/dependabot.yaml",
			},
			expected: scut.TestReturn{
				NumberOfInfo: 1,
				Score:        10,
			},
		},
		{
			name:    "foo bar",
			wantErr: false,
			files: []string{
				".github/foobar.yml",
			},
			expected: scut.TestReturn{
				NumberOfWarn: 2,
			},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			ctrl := gomock.NewController(t)
			mockRepo := mockrepo.NewMockRepoClient(ctrl)
			mockRepo.EXPECT().ListFiles(gomock.Any()).Return(tt.files, nil)
			dl := scut.TestDetailLogger{}
			c := &checker.CheckRequest{
				RepoClient: mockRepo,
				Dlogger:    &dl,
			}
			res := DependencyUpdateTool(c)

			if !scut.ValidateTestReturn(t, tt.name, &tt.expected, &res, &dl) {
				t.Fail()
			}
		})
	}
}
