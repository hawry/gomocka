package settings

import "testing"

func TestSingleWildcard(t *testing.T) {
	m := Mock{
		Path: "/user/{userid}",
	}
	b, s := m.Wildcard()
	if !b || (len(s) != 1) {
		t.Fail()
	}
	for _, v := range s {
		if v != "userid" {
			t.Fail()
		}
	}
}

func TestMultipleWildcard(t *testing.T) {
	m := Mock{
		Path: "/user/{userid}/{more}",
	}
	b, s := m.Wildcard()
	if !b || (len(s) != 2) {
		t.Fail()
	}
	for _, v := range s {
		if v != "userid" && v != "more" {
			t.Fail()
		}
	}
}

func TestNoWildcard(t *testing.T) {
	m := Mock{
		Path: "/user/nothing",
	}
	if b, m := m.Wildcard(); b || (len(m) != 0) {
		t.Fail()
	}
}
