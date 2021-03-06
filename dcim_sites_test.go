// Copyright 2016 The go-netbox Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package netbox

import (
	"net/http"
	"reflect"
	"testing"
)

func TestClientDCIMGetSite(t *testing.T) {
	wantSite := testSite(1)

	c, done := testClient(t, testHandler(t, http.MethodGet, "/api/dcim/sites/1", wantSite))
	defer done()

	gotSite, err := c.DCIM.GetSite(wantSite.ID)
	if err != nil {
		t.Fatalf("unexpected error from Client.DCIM.GetSite: %v", err)
	}

	if want, got := *wantSite, *gotSite; !reflect.DeepEqual(want, got) {
		t.Fatalf("unexpected Site:\n- want: %v\n-  got: %v", want, got)
	}
}

func TestClientDCIMListSites(t *testing.T) {
	wantSites := []*Site{
		testSite(1),
		testSite(2),
		testSite(3),
	}

	c, done := testClient(t, testHandler(t, http.MethodGet, "/api/dcim/sites/", wantSites))
	defer done()

	gotSites, err := c.DCIM.ListSites()
	if err != nil {
		t.Fatalf("unexpected error from Client.DCIM.ListSites: %v", err)
	}

	if want, got := derefSites(wantSites), derefSites(gotSites); !reflect.DeepEqual(want, got) {
		t.Fatalf("unexpected Sites:\n- want: %v\n-  got: %v", want, got)
	}
}

// derefSites is used to print values of Sites in slice, instead of memory addresses.
func derefSites(sites []*Site) []Site {
	s := make([]Site, len(sites))
	for i := range sites {
		s[i] = *sites[i]
	}

	return s
}
