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

package rendersearchtemplate

import (
	"encoding/json"
)

// Response holds the response body struct for the package rendersearchtemplate
//
// https://github.com/elastic/elasticsearch-specification/blob/b7d4fb5356784b8bcde8d3a2d62a1fd5621ffd67/specification/_global/render_search_template/RenderSearchTemplateResponse.ts#L23-L25
type Response struct {
	TemplateOutput map[string]json.RawMessage `json:"template_output"`
}

// NewResponse returns a Response
func NewResponse() *Response {
	r := &Response{
		TemplateOutput: make(map[string]json.RawMessage, 0),
	}
	return r
}