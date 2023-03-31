package main

import "errors"

type Dictionary map[string]string

var ErrNotFound = errors.New("the word that was searched for could not be found")

func (d Dictionary) Search(word string) (string, error) {
	definition, exist := d[word]
	if !exist {
		return "", ErrNotFound
	}

	return definition, nil
}
