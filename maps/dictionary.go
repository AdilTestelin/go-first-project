package maps

import "errors"

type Dictionary map[string]string

var (
	errWordNotFound      = errors.New("could not find the word you were looking for")
	errWordAlreadyExists = errors.New("this word is already in the dictionary")
)

func (dictionary Dictionary) Search(word string) (string, error) {
	definition, found := dictionary[word]
	if !found {
		return "", errWordNotFound
	}

	return definition, nil
}

func (d Dictionary) Add(word, definition string) error {
	_, err := d.Search(word)

	switch err {
	case errWordNotFound:
		d[word] = definition
	case nil:
		return errWordAlreadyExists
	default:
		return err
	}

	return nil
}
