package meta

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var addTests = []struct {
	input    string
	expected bool
	count    int
}{
	{"awd", false, 1},
	{"awd", true, 2},
	{"j1", false, 1},
	{"33", false, 1},
	{"j1", true, 2},
	{"awd", true, 3},
}

func TestAdd(t *testing.T) {
	m := NewMap(len(addTests))
	for _, test := range addTests {
		got := m.Add(test.input)
		if got != test.expected {
			t.Errorf("Expected %t, got %t, for %s", test.expected, got, test.input)
		}
		gotValue := m.Value(test.input)
		if gotValue != test.count {
			t.Errorf("Expected %d, got %d, for %s", test.count, gotValue, test.input)
		}
	}
}

func TestSum(t *testing.T) {
	m := NewMap(2)
	m.Data["1"] = 100
	m.Data["2"] = 3
	got := m.Sum()
	if got != 103 {
		t.Errorf("Expected sum %d, got %d", 103, got)
	}
}

func TestContains(t *testing.T) {
	m := NewMap(2)
	m.Data["zz@zz"] = 100
	m.Data["bbbb"] = 3
	assert.True(t, m.Contains("@"))
	assert.False(t, m.Contains("!"))
}
