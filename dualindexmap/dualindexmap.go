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

// This package implements a DualIndexMap as an example to go along
// with the deepmesa blog post. For more see:
// http://deepmesa.com/blog/rust-storing-data-in-multiple-containers
package dualindexmap

type Key struct {
	k string
}

type Val struct {
	v int
}

type DualIndexMap struct {
	kvMap map[Key]Val
	vkMap map[Val]Key
}

func NewDualIndexMap() *DualIndexMap {
	return &DualIndexMap {
		kvMap: make(map[Key]Val),
		vkMap: make(map[Val]Key),
	}
}

func (m *DualIndexMap) Put(key Key, val Val) {
	m.kvMap[key] = val
	m.vkMap[val] = key
}

func (m *DualIndexMap) GetByKey(key Key) Val {
	return m.kvMap[key]
}

func (m *DualIndexMap) GetByVal(val Val) Key {
	return m.vkMap[val]
}

