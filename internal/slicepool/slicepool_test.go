// Copyright 2020 The Outline Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package slicepool

import (
	"testing"
)

func TestPool(t *testing.T) {
	pool := MakePool(10)
	slice := pool.LazySlice()
	buf := slice.Acquire()
	if len(buf) != 10 {
		t.Errorf("Wrong slice length: %d", len(buf))
	}
	slice.Release()
}

func BenchmarkPool(b *testing.B) {
	pool := MakePool(10)
	for i := 0; i < b.N; i++ {
		slice := pool.LazySlice()
		slice.Acquire()
		slice.Release()
	}
}
