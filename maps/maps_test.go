package main

import (
	"testing"
)

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": "this is just a test"}

	t.Run("found word", func(t *testing.T) {
		result, _ := dictionary.Search("test")
		expect := "this is just a test"

		compareString(t, result, expect)
	})

	t.Run("not found word", func(t *testing.T) {
		_, result := dictionary.Search("unknown")

		compareError(t, result, ErrNotFound)
	})
}

func TestAdd(t *testing.T) {

	t.Run("new word", func(t *testing.T) {
		dictionary := Dictionary{}
		word := "test"
		definition := "this is just a test"

		dictionary.Add(word, definition)

		compareDefinition(t, dictionary, word, definition)
	})

	t.Run("already existent word", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"
		dictionary := Dictionary{word: definition}
		err := dictionary.Add(word, "new test")

		compareError(t, err, ErrNotFound)
		compareDefinition(t, dictionary, word, definition)
	})

}

func TestUpdate(t *testing.T) {

	t.Run("already existent word", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"
		dictionary := Dictionary{word: definition}
		newDefinition := "new definition"
		err := dictionary.Update(word, newDefinition)

		compareError(t, err, nil)
		compareDefinition(t, dictionary, word, newDefinition)
	})

	t.Run("new word", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"
		dictionary := Dictionary{}

		err := dictionary.Update(word, definition)

		compareError(t, err, ErrNotExistentWord)
	})

}

func TestDelete(t *testing.T) {
	word := "test"
	dictionary := Dictionary{word: "definition of test"}

	dictionary.Delete(word)

	_, err := dictionary.Search(word)
	if err != ErrNotFound {
		t.Errorf("It is expected that '%s' will be deleted", word)
	}
}

func compareString(t *testing.T, result, expect string) {
	t.Helper()

	if result != expect {
		t.Errorf("result '%s', expect '%s', passed '%s'", result, expect, "test")
	}
}

func compareError(t *testing.T, result, expect error) {
	t.Helper()

	if result != expect {
		t.Errorf("result error '%s', expected '%s'", result, expect)
	}
}

func compareDefinition(t *testing.T, dictionary Dictionary, word, definition string) {
	t.Helper()

	result, err := dictionary.Search(word)

	if err != nil {
		t.Fatal("should have found added word", err)
	}

	if definition != result {
		t.Errorf("result '%s', expected '%s'", result, definition)
	}
}
