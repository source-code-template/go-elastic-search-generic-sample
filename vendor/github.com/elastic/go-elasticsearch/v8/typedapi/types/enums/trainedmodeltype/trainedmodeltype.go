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

// Package trainedmodeltype
package trainedmodeltype

import "strings"

// https://github.com/elastic/elasticsearch-specification/blob/b7d4fb5356784b8bcde8d3a2d62a1fd5621ffd67/specification/ml/_types/TrainedModel.ts#L257-L271
type TrainedModelType struct {
	Name string
}

var (
	Treeensemble = TrainedModelType{"tree_ensemble"}

	Langident = TrainedModelType{"lang_ident"}

	Pytorch = TrainedModelType{"pytorch"}
)

func (t TrainedModelType) MarshalText() (text []byte, err error) {
	return []byte(t.String()), nil
}

func (t *TrainedModelType) UnmarshalText(text []byte) error {
	switch strings.ReplaceAll(strings.ToLower(string(text)), "\"", "") {

	case "tree_ensemble":
		*t = Treeensemble
	case "lang_ident":
		*t = Langident
	case "pytorch":
		*t = Pytorch
	default:
		*t = TrainedModelType{string(text)}
	}

	return nil
}

func (t TrainedModelType) String() string {
	return t.Name
}