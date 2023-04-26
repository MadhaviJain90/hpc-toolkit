/*
Copyright 2023 Google LLC

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    https://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package cmd

import (
	"os"

	. "gopkg.in/check.v1"
)

func (s *MySuite) TestIsDir(c *C) {
	dir, err := os.MkdirTemp("", "test-*")
	if err != nil {
		c.Fatal(err)
	}
	defer os.RemoveAll(dir)

	err = checkDir(nil, []string{dir})
	c.Assert(err, IsNil)

	os.RemoveAll(dir)
	err = checkDir(nil, []string{dir})
	c.Assert(err, NotNil)

	f, _ := os.CreateTemp("", "test-*")
	err = checkDir(nil, []string{f.Name()})
	c.Assert(err, NotNil)
}

func (s *MySuite) TestRunExport(c *C) {
	dir, err := os.MkdirTemp("", "test-*")
	if err != nil {
		c.Fatal(err)
	}
	defer os.RemoveAll(dir)

	err = runExportCmd(nil, []string{dir})
	c.Assert(err, NotNil)
}
