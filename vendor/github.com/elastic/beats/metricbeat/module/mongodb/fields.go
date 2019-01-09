// Licensed to Elasticsearch B.V. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Elasticsearch B.V. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

// Code generated by beats/dev-tools/cmd/asset/asset.go - DO NOT EDIT.

package mongodb

import (
	"github.com/elastic/beats/libbeat/asset"
)

func init() {
	if err := asset.SetFields("metricbeat", "mongodb", asset.ModuleFieldsPri, Asset); err != nil {
		panic(err)
	}
}

// Asset returns asset data
func Asset() string {
	return "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"
}
