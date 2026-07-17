package main

import "testing"

func Test_AddGet(t *testing.T) {
	m := NewStringIntMap()
	m.Add("qwe", 1)
	m.Add("asd", 2)

	if val, ok := m.Get("qwe"); !ok || val != 1 {
		t.Error()
	}
	if val, ok := m.Get("asd"); !ok || val != 2 {
		t.Error()
	}
}

func Test_Remove(t *testing.T) {
	m := NewStringIntMap()
	m.Add("qwe", 1)
	m.Add("asd", 2)

	m.Remove("qwe")
	if val, ok := m.Get("qwe"); ok || val != 0 {
		t.Error()
	}
	m.Remove("asd")
	if val, ok := m.Get("asd"); ok || val != 0 {
		t.Error()
	}
}

func Test_Copy(t *testing.T) {
	m := NewStringIntMap()
	m.Add("qwe", 1)
	m.Add("asd", 2)

	newM := m.Copy()
	m.Remove("qwe")
	if val, ok := newM.Get("qwe"); !ok || val != 1 {
		t.Error()
	}
	if val, ok := newM.Get("asd"); !ok || val != 2 {
		t.Error()
	}
}

func Test_Exist(t *testing.T) {
	m := NewStringIntMap()
	m.Add("qwe", 1)

	if !m.Exist("qwe") {
		t.Error()
	}
	if m.Exist("asd") {
		t.Error()
	}
}
