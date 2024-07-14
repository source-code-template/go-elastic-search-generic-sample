// Licensed to Elasticsearch B.V. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Elasticsearch B.V. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

// Code generated from the elasticsearch-specification DO NOT EDIT.
// https://github.com/elastic/elasticsearch-specification/tree/b7d4fb5356784b8bcde8d3a2d62a1fd5621ffd67

// Package migrationstatus
package migrationstatus

import "strings"

// https://github.com/elastic/elasticsearch-specification/blob/b7d4fb5356784b8bcde8d3a2d62a1fd5621ffd67/specification/migration/get_feature_upgrade_status/GetFeatureUpgradeStatusResponse.ts#L30-L35
type MigrationStatus struct {
	Name string
}

var (
	NOMIGRATIONNEEDED = MigrationStatus{"NO_MIGRATION_NEEDED"}

	MIGRATIONNEEDED = MigrationStatus{"MIGRATION_NEEDED"}

	INPROGRESS = MigrationStatus{"IN_PROGRESS"}

	ERROR = MigrationStatus{"ERROR"}
)

func (m MigrationStatus) MarshalText() (text []byte, err error) {
	return []byte(m.String()), nil
}

func (m *MigrationStatus) UnmarshalText(text []byte) error {
	switch strings.ReplaceAll(strings.ToLower(string(text)), "\"", "") {

	case "NO_MIGRATION_NEEDED":
		*m = NOMIGRATIONNEEDED
	case "MIGRATION_NEEDED":
		*m = MIGRATIONNEEDED
	case "IN_PROGRESS":
		*m = INPROGRESS
	case "ERROR":
		*m = ERROR
	default:
		*m = MigrationStatus{string(text)}
	}

	return nil
}

func (m MigrationStatus) String() string {
	return m.Name
}
