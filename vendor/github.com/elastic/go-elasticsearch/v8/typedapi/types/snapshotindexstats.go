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

package types

// SnapshotIndexStats type.
//
// https://github.com/elastic/elasticsearch-specification/blob/b7d4fb5356784b8bcde8d3a2d62a1fd5621ffd67/specification/snapshot/_types/SnapshotIndexStats.ts#L25-L29
type SnapshotIndexStats struct {
	Shards      map[string]SnapshotShardsStatus `json:"shards"`
	ShardsStats SnapshotShardsStats             `json:"shards_stats"`
	Stats       SnapshotStats                   `json:"stats"`
}

// NewSnapshotIndexStats returns a SnapshotIndexStats.
func NewSnapshotIndexStats() *SnapshotIndexStats {
	r := &SnapshotIndexStats{
		Shards: make(map[string]SnapshotShardsStatus, 0),
	}

	return r
}