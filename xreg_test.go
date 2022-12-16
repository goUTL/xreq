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
	"net/http"
	"reflect"
	"testing"
)

func TestNewRequest(t *testing.T) {

	r := NewRequest("host", "path")

	if reflect.TypeOf(r) != reflect.TypeOf(&Request{}) {
		t.Fatalf("xreq.NewRequest() = %T, want *xreq.Request", r)
	} else if r == nil {
		t.Fatal("xreq.NewRequest() = <nil>, want &xreq.Request{}")
	}

	t.Run("req", func(t *testing.T) {
		if reflect.TypeOf(r.req) != reflect.TypeOf(&http.Request{}) {
			t.Fatalf("xreq.NewRequest().req = %T, want *http.Request", r.req)
		} else if r.req == nil {
			t.Fatal("xreq.NewRequest().req = <nil>, want &http.Request{}")
		}

		t.Run("method", func(t *testing.T) {
			if r.req.Method != http.MethodGet {
				t.Errorf("xreq.NewRequest().req.Method = \"%s\", want \"GET\"", r.req.Method)
			}
		})

		t.Run("scheme", func(t *testing.T) {
			if r.req.URL.Scheme != "https" {
				t.Errorf("xreq.NewRequest().req.URL.Scheme = \"%s\", want \"https\"", r.req.URL.Scheme)
			}
		})
	})

	t.Run("qry", func(t *testing.T) {
		w := map[string]*string{}
		if r.qry == nil || reflect.TypeOf(r.qry) != reflect.TypeOf(w) {
			t.Errorf("xreq.NewRequest().qry = %#v, want %#v", r.qry, w)
		}
	})

	t.Run("hdr", func(t *testing.T) {
		w := map[string]*string{}
		if r.hdr == nil || reflect.TypeOf(r.hdr) != reflect.TypeOf(w) {
			t.Errorf("xreq.NewRequest().hdr = %#v, want %#v", r.hdr, w)
		}
	})
}

func TestRequest_Query(t *testing.T) {

	r := NewRequest("host", "path")

	parameter := "QP"
	query := "QV"

	// Add query.
	r.Query(parameter, &query)
	if _, ok := r.qry[parameter]; !ok {
		t.Errorf("this query %p was not added", &query)
	}

	// Checking pointer.
	if _, ok := r.qry[parameter]; ok && r.qry[parameter] != &query {
		t.Errorf("this query %p, want %p", r.qry[parameter], &query)
	}

	// Delete query.
	r.Query(parameter, nil)
	if _, ok := r.qry[parameter]; ok {
		t.Errorf("this query %p was not deleted", &query)
	}
}

func TestRequest_Header(t *testing.T) {

	r := NewRequest("host", "path")

	name := "HN"
	header := "HV"

	// Add header.
	r.Header(name, &header)
	if _, ok := r.hdr[name]; !ok {
		t.Errorf("this header %p was not added", &header)
	}

	// Checking pointer.
	if _, ok := r.hdr[name]; ok && r.hdr[name] != &header {
		t.Errorf("this header %p, want %p", r.hdr[name], &header)
	}

	// Delete header.
	r.Header(name, nil)
	if _, ok := r.hdr[name]; ok {
		t.Errorf("this header %p was not deleted", &header)
	}
}
