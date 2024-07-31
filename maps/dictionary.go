package main

import "errors"

type Dictionary map[string]string

var (
	ErrNotFound   = errors.New("could not find the word you were looking for")
	ErrWordExists = errors.New("cannot add word because it already exists")
)

func (d Dictionary) Search(word string) (string, error) {
	// In order to make this pass, we are using an interesting property of the map lookup.
	// It can return 2 values. The second value is a boolean which indicates if the key was found successfully.
	definition, ok := d[word]
	if !ok {
		return "", ErrNotFound
	}

	return definition, nil
}


func (d Dictionary) Add(word, definition string) error {
	// An interesting property of maps is that you can modify them without passing as an address to it (e.g &myMap)
	// When you pass a map to a function/method, you are indeed copying it, but just the pointer part, not the underlying data structure that contains the data.
	_, err := d.Search(word)

	switch err {
	case ErrNotFound:
		d[word] = definition
	case nil:
		return ErrWordExists
	default:
		return err
	}

	return nil
}

func (d Dictionary) Delete(word string) {
	delete(d, word)
}