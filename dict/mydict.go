package dict

import "errors"

// Dictionary is a map of string to string
type Dictionary map[string]string

var errNotFound = errors.New("Word not found")

// Search searches a word in the dictionary
func (d Dictionary) Search(word string) (string, error) {
  value, exists := d[word]
  if exists {
    return value, nil
  }
  return "", errNotFound 
}
