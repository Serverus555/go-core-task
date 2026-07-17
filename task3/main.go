package main

import "hash/crc32"

// StringIntMap Простая реализация Map на односвязных списках без усложнений, по KISS
type StringIntMap struct {
	buckets []*Node
	length  int
}

type Node struct {
	next  *Node
	key   string
	value int
}

func NewStringIntMap() *StringIntMap {
	return &StringIntMap{
		buckets: []*Node{nil},
		length:  0,
	}
}

const loadFactor = 0.75

func (m *StringIntMap) Add(key string, value int) {
	if float64(m.length) >= float64(len(m.buckets))*loadFactor {
		m.evacuate()
	}

	bucketInd, current, prev := m.locate(key)

	if current == nil {
		// New key
		m.length++
		current = &Node{
			key:   key,
			value: value,
		}
		if prev == nil {
			// Head
			m.buckets[bucketInd] = current
		} else {
			// Tail
			prev.next = current
		}
	} else {
		// Existing key
		current.value = value
	}
}

// Remove No resize back
func (m *StringIntMap) Remove(key string) {
	bucketInd, current, prev := m.locate(key)

	if current == nil {
		return
	}

	m.length--
	if prev == nil {
		// Head
		m.buckets[bucketInd] = current.next
		return
	}
	// Body/Tail
	prev.next = current.next
}

func (m *StringIntMap) Copy() *StringIntMap {
	newMap := &StringIntMap{
		buckets: make([]*Node, len(m.buckets)),
		length:  m.length,
	}

	var newTail *Node = nil
	for i, bucket := range m.buckets {
		oldCurrent := bucket
		for oldCurrent != nil {
			newCurrent := &Node{
				key:   oldCurrent.key,
				value: oldCurrent.value,
			}

			if newTail == nil {
				// Head
				newMap.buckets[i] = newCurrent
			} else {
				// Tail
				newTail.next = newCurrent
			}

			newTail = newCurrent
			oldCurrent = oldCurrent.next
		}
	}

	return newMap
}

func (m *StringIntMap) Exist(key string) bool {
	_, current, _ := m.locate(key)
	return current != nil
}

func (m *StringIntMap) Get(key string) (int, bool) {
	_, current, _ := m.locate(key)
	if current != nil {
		return current.value, true
	}
	return 0, false
}

func calcHash(key string, divider uint32) uint32 {
	return crc32.ChecksumIEEE([]byte(key)) % divider
}

func (m *StringIntMap) locate(key string) (uint32, *Node, *Node) {
	bucketInd := calcHash(key, uint32(len(m.buckets)))

	current := m.buckets[bucketInd]
	var prev *Node = nil

	// Find node
	for current != nil {
		if key == current.key {
			break
		}
		prev = current
		current = current.next
	}
	return bucketInd, current, prev
}

func (m *StringIntMap) evacuate() {
	newMap := &StringIntMap{
		buckets: make([]*Node, len(m.buckets)*2),
		length:  m.length,
	}

	for _, bucket := range m.buckets {
		current := bucket
		for current != nil {
			newMap.Add(current.key, current.value)
			current = current.next
		}
	}
	m.buckets = newMap.buckets
}
