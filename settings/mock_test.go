package settings

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSingleWildcard(t *testing.T) {
	m := Mock{
		Path: "/user/{userid}",
	}
	b, s := m.Wildcard()
	assert.True(t, b, "should be true")
	assert.Len(t, s, 1, "should have length 1")
	for _, v := range s {
		assert.Equal(t, "userid", v, "should be equal")
	}
}

func TestMultipleWildcard(t *testing.T) {
	m := Mock{
		Path: "/user/{userid}/{more}",
	}
	b, s := m.Wildcard()
	assert.True(t, b)
	assert.Len(t, s, 2)
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
