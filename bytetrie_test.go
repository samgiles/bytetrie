/**
 * The MIT License (MIT)
 *
 * Copyright (c) 2015 Samuel Giles
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy
 * of this software and associated documentation files (the "Software"), to deal
 * in the Software without restriction, including without limitation the rights
 * to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 * copies of the Software, and to permit persons to whom the Software is
 * furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in
 * all copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 * FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 * AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 * LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 * OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
 * THE SOFTWARE.
 */
package bytetrie

import (
	"testing"
)

func TestInsert(t *testing.T) {
	node := new(Node)
	value := []byte{'a','b','c','d','e'}

	node.Insert(value)

	if !node.HasValue(value) {
		t.Fail()
	}
}

func TestRetriveIgnoreSubstrings (t *testing.T) {
	node := new(Node)
	value := []byte{'a', 'b', 'c', 'd', 'e'}

	node.Insert(value)

	if node.HasValue([]byte{'a', 'b', 'c', 'd'}) {
		t.Fail()
	}
}

func TestRetrieve(t *testing.T) {
	node := new(Node)

	node.Insert([]byte{'a', 'b', 'c', 'd', 'e'})
	node.Insert([]byte{'a', 'b', 'c'})

	if node.HasValue([]byte{'a', 'b', 'c', 'd'}) {
		t.Fail()
	}

	if !node.HasValue([]byte{'a', 'b', 'c'}) {
		t.Fail()
	}

	if !node.HasValue([]byte{'a', 'b', 'c', 'd', 'e'}) {
		t.Fail()
	}
}

var benchSearchValues = [][]byte{
	[]byte("abcdefghijklmnopqrstuvwxyz"),
	[]byte("abcdefghijklmnopqrstuvabcd"),
	[]byte("abcdefghijklmnopqrstuvwxyx"),
	[]byte("abcdefghijklmnoqeasdhfallj"),
}

var benchTestValues = [][]byte{
	[]byte("abcdefghijklmnopqrstuvwxyz"),
	[]byte("abcdefghijklmnopqrstuvabcd"),
	[]byte("abcdefghijklmnopqrstuvwxyx"),
	[]byte("abcdefghijklmnoqeasdhfallj"),
	[]byte("abcdefghoqeasdhfallj"),
	[]byte("abcnoqeasdhfallj"),
	[]byte("lmnoqeasdhfallj"),
	[]byte("noqeasdhfallj"),
	[]byte("jklmnoqeasdhfallj"),
	[]byte("mnoqeasdhfallj"),
}

func BenchmarkLookup(b *testing.B) {
	node := new(Node)
	for _, bytearray := range benchTestValues {
		node.Insert(bytearray)
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		for _, bytearray := range benchSearchValues {
			node.HasValue(bytearray)
		}
	}
}
