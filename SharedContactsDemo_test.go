// Copyright 2016 Takbok, Inc. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package demo

import (
	"encoding/xml"
	"testing"
)

const testFeed = `<?xml version='1.0' encoding='UTF-8'?>
<feed xmlns='http://www.w3.org/2005/Atom'
      xmlns:gContact='http://schemas.google.com/contact/2008'
      xmlns:gd='http://schemas.google.com/g/2005'
      xmlns:batch='http://schemas.google.com/gdata/batch'>
  <category scheme='http://schemas.google.com/g/2005#kind'
      term='http://schemas.google.com/g/2008#contact' />
  <entry>
    <batch:id>1</batch:id>
    <batch:operation type='insert' />
    <category scheme='http://schemas.google.com/g/2005#kind'
      term='http://schemas.google.com/g/2008#contact'/>
    <gd:name>
      <gd:givenName>Contact</gd:givenName>
      <gd:familyName>One</gd:familyName>
    </gd:name>
    <gd:email rel='http://schemas.google.com/g/2005#home'
      address='contact1@example.com' primary='true'/>
  </entry>
  <entry>
    <batch:id>2</batch:id>
    <batch:operation type='insert' />
    <category scheme='http://schemas.google.com/g/2005#kind'
      term='http://schemas.google.com/g/2008#contact'/>
    <gd:name>
      <gd:givenName>Contact</gd:givenName>
      <gd:familyName>Two</gd:familyName>
    </gd:name>
    <gd:email rel='http://schemas.google.com/g/2005#home'
      address='contact2@example.com'
      primary='true'/>
  </entry>
</feed>`

func TestFeed(t *testing.T) {
	var feed Feed

	if err := xml.Unmarshal([]byte(testFeed), &feed); err != nil {
		t.Errorf("unmarshal feed: %v", err)
		return
	}

	if n := len(feed.Entry); n != 2 {
		t.Errorf("len(feed.Entry) != 2")
		return
	}

	e1 := &feed.Entry[0]
	e2 := &feed.Entry[1]

	if s := e1.Name.GivenName; s != "Contact" {
		t.Errorf("e1.Name.GivenName != Contact (%v)", s)
		return
	}
	if s := e1.Name.FamilyName; s != "One" {
		t.Errorf("e1.Name.FamilyName != One (%v)", s)
		return
	}
	if s := e2.Name.GivenName; s != "Contact" {
		t.Errorf("e2.Name.GivenName != Contact (%v)", s)
		return
	}
	if s := e2.Name.FamilyName; s != "Two" {
		t.Errorf("e2.Name.FamilyName != Two (%v)", s)
		return
	}

	if n := len(e1.Email); n != 1 {
		t.Errorf("len(e1.Email) != 1 (%v)", n)
		return
	}
	if n := len(e2.Email); n != 1 {
		t.Errorf("len(e2.Email) != 1 (%v)", n)
		return
	}

	em1 := e1.Email[0]
	em2 := e2.Email[0]

	if s := em1.Address; s != "contact1@example.com" {
		t.Errorf("em1.Address != contact1@example.com (%v)", s)
		return
	}
	if s := em2.Address; s != "contact2@example.com" {
		t.Errorf("em2.Address != contact2@example.com (%v)", s)
		return
	}

	if b := em1.Primary; !b {
		t.Errorf("!em1.Primary (%v)", b)
		return
	}
	if b := em2.Primary; !b {
		t.Errorf("!em2.Primary (%v)", b)
		return
	}
}
