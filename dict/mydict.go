package dict

import "errors"

// Dictionary is a map of string to string
type Dictionary map[string]string

var errNotFound = errors.New("Word not found")
var errWordExists = errors.New("Word already exists")

// Search searches a word in the dictionary
func (d Dictionary) Search(word string) (string, error) {
  value, exists := d[word]
  if exists {
    return value, nil
  }
  return "", errNotFound 
}

// Add adds a word to the dictionary
func (d Dictionary) Add(word string, definition string) error {
  _, exists := d[word]
  if exists {
    return errWordExists
  }
  d[word] = definition
  return nil
}
