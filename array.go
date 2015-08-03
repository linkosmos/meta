package meta

import "fmt"

// Array - stores map of links and count
type Array struct {
	Data []string
}

// NewArray - returns new instance to meta arrau
func NewArray(size int) *Array {
	return &Array{
		Data: make([]string, 0, size),
	}
}

// Add - adds link if not exist OR increments existing
func (a *Array) Add(s string) {
	a.Data = append(a.Data, s)
}

// Range - ranges over Data array returns index and existance
func (a *Array) Range(s string) (index int, exist bool) {
	for index, value := range a.Data {
		if value == s {
			return index, true
		}
	}
	return -1, false
}

// Index - item index in Data array
// -1 => not exist
func (a *Array) Index(s string) int {
	index, _ := a.Range(s)
	return index
}

// Exist - checks whether URL exists
func (a *Array) Exist(s string) bool {
	_, exist := a.Range(s)
	return exist
}

// Values - return map values as array
func (a *Array) Values() []string {
	return a.Data
}

// Count - returns map length
func (a *Array) Count() int {
	return len(a.Data)
}

// Dump - map to function
func (a *Array) Dump(f func(string)) {
	for _, value := range a.Data {
		f(value)
	}
}

// String - string interface
func (a *Array) String() (output string) {
	a.Dump(func(value string) {
		output += fmt.Sprintf("%s, \n", value)
	})
	return output
}
