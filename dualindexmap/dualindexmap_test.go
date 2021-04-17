/*
   Example golang code for the deepmesa project. 

   Copyright 2021 "Rahul Singh <rsingh@arrsingh.com>"

   Licensed under the Apache License, Version 2.0 (the "License");
   you may not use this file except in compliance with the License.
   You may obtain a copy of the License at

       http://www.apache.org/licenses/LICENSE-2.0

   Unless required by applicable law or agreed to in writing, software
   distributed under the License is distributed on an "AS IS" BASIS,
   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
   See the License for the specific language governing permissions and
   limitations under the License.
*/

// This package contains tests for the DualIndexMap whish is an
// example that goes along with the deepmesa blog post. For more see:
// http://deepmesa.com/blog/rust-storing-data-in-multiple-containers
package dualindexmap

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestDualIndexMap(t *testing.T) {
	m := NewDualIndexMap()
	key := Key{k: "testkey"} 
	val := Val{v: 10}
	m.Put(key, val)

	assert.True(t, m.GetByKey(key) == val)
	assert.True(t, m.GetByVal(val) == key)
}
