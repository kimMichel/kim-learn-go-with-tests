package main

type Dictionary map[string]string

type ErrDictionary string

const (
	ErrNotFound        = ErrDictionary("the word that was searched for could not be found")
	ErrExistentWord    = ErrDictionary("it is not possible to add the word because it already exists")
	ErrNotExistentWord = ErrDictionary("it is not possible to update the word because it not exists")
)

func (d Dictionary) Search(word string) (string, error) {
	definition, exist := d[word]
	if !exist {
		return "", ErrNotFound
	}

	return definition, nil
}

func (d Dictionary) Add(key, value string) error {
	_, err := d.Search(key)
	switch err {
	case ErrNotFound:
		d[key] = value
	case nil:
		return ErrExistentWord
	default:
		return err
	}

	return nil
}

func (d Dictionary) Update(word, newDefinition string) error {
	_, err := d.Search(word)
	switch err {
	case ErrNotFound:
		return ErrNotExistentWord
	case nil:
		d[word] = newDefinition
	default:
		return err
	}

	return nil
}

func (d Dictionary) Delete(word string) {
	delete(d, word)
}

func (e ErrDictionary) Error() string {
	return string(e)
}
