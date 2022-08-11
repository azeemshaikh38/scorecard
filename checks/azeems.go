// Copyright 2021 Security Scorecard Authors
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
	"github.com/ossf/scorecard/v4/checker"
)

// CheckBinaryArtifacts is the exported name for Binary-Artifacts check.
const CheckAzeems string = "Azeems-Awesome-Check"

// nolint
func init() {
	supportedRequestTypes := []checker.RequestType{
		checker.CommitBased,
	}
	if err := registerCheck(CheckAzeems, AzeemsAwesomeCheck, supportedRequestTypes); err != nil {
		// this should never happen
		panic(err)
	}
}

// BinaryArtifacts  will check the repository contains binary artifacts.
func AzeemsAwesomeCheck(c *checker.CheckRequest) checker.CheckResult {
	return checker.CreateMaxScoreResult(CheckAzeems, "because Azeem is awesome")
}
