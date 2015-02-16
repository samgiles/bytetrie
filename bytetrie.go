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

 type Node struct {
	IsLeaf bool
	Children [255]*Node
 }

func (n *Node) HasChild(b byte) bool {
	return n.Children[b] != nil
}

func (n *Node) AddChild(b byte) (*Node) {
	var newNode =  &Node{}

	n.Children[b] = newNode
	return newNode;
}

func (n *Node) GetChild(b byte) (*Node) {
	return n.Children[b]
}

func (n* Node) HasValue(byteArray []byte) bool {
	var currentTestNode = n;
	for _, b := range byteArray {
		if (!currentTestNode.HasChild(b)) {
			return false
		} else {
			currentTestNode = currentTestNode.GetChild(b)
		}
	}

	return currentTestNode.IsLeaf
}

func (n *Node) Insert(search []byte) {
	var currentNode = n
	var lastIndex = len(search) - 1

	for index, b := range search {
		if currentNode.HasChild(b) {
			currentNode = currentNode.GetChild(b)
		} else {
			currentNode = currentNode.AddChild(b)
		}

		if index == lastIndex {
			currentNode.IsLeaf = true
		}
	}
}

func (n *Node) Accepts(b byte) (*Node, bool) {
	if !n.HasChild(b) {
		return nil, false
	} else {
		return n.GetChild(b), true
	}
}
