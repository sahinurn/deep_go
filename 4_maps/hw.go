package main

import (
	"cmp"
)

type BSTreeNode[K cmp.Ordered, V any] struct {
	Key   K
	Value V
	Left  *BSTreeNode[K, V]
	Right *BSTreeNode[K, V]
}

type OrderedMap[K cmp.Ordered, V any] struct {
	root *BSTreeNode[K, V]
	size int
}

func NewOrderedMap[K cmp.Ordered, V any]() OrderedMap[K, V] {
	return OrderedMap[K, V]{
		size: 0,
	}
}

func (m *OrderedMap[K, V]) Insert(key K, value V) {
	if m.root == nil {
		m.size++
		m.root = &BSTreeNode[K, V]{
			Key:   key,
			Value: value,
		}
		return
	}

	n := m.root

LOOP:
	for {
		switch {
		case n.Key == key:
			n.Value = value
			break LOOP
		case n.Key < key:
			if n.Right == nil {
				m.size++
				n.Right = &BSTreeNode[K, V]{
					Key:   key,
					Value: value,
				}
				break LOOP
			}
			n = n.Right
		case n.Key > key:
			if n.Left == nil {
				m.size++
				n.Left = &BSTreeNode[K, V]{
					Key:   key,
					Value: value,
				}
				break LOOP
			}
			n = n.Left
		}
	}
}

func (m *OrderedMap[K, V]) Erase(key K) {
	if m.root == nil {
		return
	}

	n := &m.root

LOOP:
	for {
		switch {
		case (*n).Key == key:
			break LOOP
		case (*n).Key < key:
			if (*n).Right == nil {
				break LOOP
			}
			n = &(*n).Right
		case (*n).Key > key:
			if (*n).Left == nil {
				break LOOP
			}
			n = &(*n).Left
		}
	}

	if (*n).Left == nil && (*n).Right == nil {
		*n = nil
		m.size--
		return
	}

	if (*n).Left == nil && (*n).Right != nil {
		*n = (*n).Right
		m.size--
		return
	}

	if (*n).Left != nil && (*n).Right == nil {
		*n = (*n).Left
		m.size--
		return
	}

	minNode := (*n).Right

	for (*minNode).Left != nil {
		minNode = (*minNode).Left
	}

	minNode.Right = (*n).Right
	*n = minNode

	m.size--

}

func (m *OrderedMap[K, V]) Contains(key K) bool {
	if m.root == nil {
		return false
	}

	n := m.root

	for {
		switch {
		case n.Key == key:
			return true
		case n.Key < key:
			if n.Right == nil {
				return false
			}
			n = n.Right
		case n.Key > key:
			if n.Left == nil {
				return false
			}
			n = n.Left
		}
	}
}

func (m *OrderedMap[K, V]) Size() int {
	return m.size
}

func (m *OrderedMap[K, V]) ForEach(action func(key K, value V)) {
	if m.root == nil {
		return
	}

	forEach(m.root, action)

}

func forEach[K cmp.Ordered, V any](n *BSTreeNode[K, V], action func(key K, value V)) {
	if n == nil {
		return
	}

	if n.Left != nil {
		forEach(n.Left, action)
	}

	action(n.Key, n.Value)

	if n.Right != nil {
		forEach(n.Right, action)
	}
}
