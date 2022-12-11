// Copyright 2022 Vasiliy Vdovin

// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at

// http://www.apache.org/licenses/LICENSE-2.0

// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package xreq

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/goUTL/xlog"
)

type Request struct {
	r *http.Request      // HTTP request
	q map[string]*string // URL queries
}

// NewRequest wraps http.NewRequestWith.
func NewRequest(host, path string) *Request {
	r, e := http.NewRequest(http.MethodGet, fmt.Sprintf("https://%s/%s", host, strings.TrimPrefix(path, "/")), nil)
	xlog.Fatalln(e)
	return &Request{r: r, q: map[string]*string{}}
}

// Query adds the value to parameter.
func (r *Request) Query(parameter string, value *string) {
	if value != nil {
		r.q[parameter] = value
	} else {
		delete(r.q, parameter)
	}
}
