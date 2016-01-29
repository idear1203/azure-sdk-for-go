package cdn

// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by Microsoft (R) AutoRest Code Generator 0.14.0.0
// Changes may cause incorrect behavior and will be lost if the code is
// regenerated.

import (
	"fmt"
)

const (
	major = "0"
	minor = "4"
	patch = "0"
	// Always begin a "tag" with a dash (as per http://semver.org)
	tag             = "-beta"
	semVerFormat    = "%s.%s.%s%s"
	userAgentFormat = "Azure-SDK-for-Go/%s;Package arm/%s;API %s"
)

// UserAgent returns the UserAgent string to use when sending http.Requests.
func UserAgent() string {
	return fmt.Sprintf(userAgentFormat, Version(), "cdn", "2015-06-01")
}

// Version returns the semantic version (see http://semver.org) of the client.
func Version() string {
	return fmt.Sprintf(semVerFormat, major, minor, patch, tag)
}
