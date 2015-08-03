package meta

import "fmt"

// Map - stores map of links and count
type Map struct {
	Data map[string]int
}

// NewMap - returns instance to new meta Map
func NewMap(size int) *Map {
	return &Map{
		Data: make(map[string]int, size),
	}
}

// Add - adds link if not exist OR increments existing
func (m *Map) Add(s string) bool {
	_, ok := m.Data[s]
	m.Data[s]++
	return ok
}

// Exist - checks whether URL exists
func (m *Map) Exist(s string) bool {
	_, exist := m.Data[s]
	return exist
}

// Value - returns key value
func (m *Map) Value(s string) int {
	if value, exist := m.Data[s]; exist {
		return value
	}
	return 0
}

// Keys - return map values as array
func (m *Map) Keys() []string {
	output := make([]string, 0, m.Count())
	m.Dump(func(key string, value int) {
		output = append(output, key)
	})
	return output
}

// Sum - sums all values
func (m *Map) Sum() (sum int) {
	m.Dump(func(key string, value int) {
		sum += value
	})
	return sum
}

// Count - returns map length
func (m *Map) Count() int {
	return len(m.Data)
}

// Dump - map to function
func (m *Map) Dump(f func(string, int)) {
	for value, count := range m.Data {
		f(value, count)
	}
}

// String - string interface
func (m *Map) String() (output string) {
	m.Dump(func(value string, count int) {
		output += fmt.Sprintf("%s => %d\n", value, count)
	})
	return output
}
