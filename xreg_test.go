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
	"net/textproto"
	"reflect"
	"testing"
)

func TestNewRequest(t *testing.T) {

	r := NewRequest("host", "path")

	if reflect.TypeOf(r) != reflect.TypeOf(&Request{}) {
		t.Fatalf("NewRequest() = %T, want *Request", r)
	}

	if reflect.TypeOf(r.r) != reflect.TypeOf(&http.Request{}) {
		t.Fatalf("Request{r: %T}, want Request{r: *http.Request}", r.r)
	}

	if reflect.TypeOf(r.q) != reflect.TypeOf(map[string]*string{}) {
		t.Fatalf("Request{q: %T}, want Request{q: map[string]*string}", r.q)
	}

	if reflect.TypeOf(r.h) != reflect.TypeOf(map[string]*string{}) {
		t.Fatalf("Request{h: %T}, want Request{h: map[string]*string}", r.h)
	}

	if r == nil {
		t.Fatal("NewRequest() = <nil>, want *Request{}")
	}

	if r.r == nil {
		t.Fatalf("Request{r: %v}, want Request{r: *http.Request{}}", r.r)
	}

	if r.q == nil {
		t.Fatalf("Request{q: %v}, want Request{q: map[string]*string{}}", r.q)
	}

	if r.h == nil {
		t.Fatalf("Request{h: %v}, want Request{h: map[string]*string{}}", r.h)
	}

	if r.r.Method != http.MethodGet {
		t.Fatalf("Method = %s, want GET", r.r.Method)
	}

	if r.r.URL.String() != "https://host/path" {
		t.Fatalf("URL = %s, want https://host/path", r.r.URL.String())
	}
}

func TestRequest_Query(t *testing.T) {

	r := NewRequest("host", "path")

	parameter := "QP"
	query := "QV"

	// Add query.
	r.Query(parameter, &query)
	if _, ok := r.q[parameter]; !ok {
		t.Errorf("this query %p was not added", &query)
	}

	// Checking pointer.
	if _, ok := r.q[parameter]; ok && r.q[parameter] != &query {
		t.Errorf("this query %p, want %p", r.q[parameter], &query)
	}

	// Delete query.
	r.Query(parameter, nil)
	if _, ok := r.q[parameter]; ok {
		t.Errorf("this query %p was not deleted", &query)
	}
}

func TestRequest_Header(t *testing.T) {

	var header any
	var pointer *string
	r := NewRequest("host", "path")

	// New header.
	if header = r.Header("name", "value"); reflect.TypeOf(header) != reflect.TypeOf(pointer) {
		t.Fatalf("Request.Header() = %T, want *string", header)
	} else if header.(*string) == nil {
		t.Fatal("Request.Header() = <nil>, want *string()")
	} else {
		pointer = header.(*string)
	}

	// Header was is added.
	if len(r.r.Header.Values("name")) == 0 {
		t.Fatal("header was not added")
	}

	// Pointer comparison which has been added to the map header and with the returned method.
	if h := &r.r.Header[textproto.CanonicalMIMEHeaderKey("name")][0]; pointer != h {
		t.Fatalf("Request.Header() = %p, want %p", pointer, h)
	}
}
