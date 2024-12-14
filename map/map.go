package main

type Dictionary map[string]string

const (
	ErrNotFound         = DictionaryErr("could not find the word you were looking for")
	ErrWordExists       = DictionaryErr("cannot add word because it already exists")
	ErrWordDoesNotExist = DictionaryErr("cannot perform operation on word because it does not exist")
)

type DictionaryErr string

func (e DictionaryErr) Error() string {
	return string(e)
}

func (d Dictionary) Search(word string) (string, error) {
	definition, ok := d[word]
	if !ok {
		return "", ErrNotFound
	}
	return definition, nil
}

func (d Dictionary) Add(index, value string) error {
	_, err := d.Search(index)

	switch err {
	case ErrNotFound:
		d[index] = value
	case nil:
		return ErrWordExists
	default:
		return err
	}

	return nil
}

func (d Dictionary) Update(index, value string) error {
	_, err := d.Search(index)

	switch err {
	case ErrNotFound:
		return ErrWordDoesNotExist
	case nil:
		d[index] = value
	default:
		return err
	}

	return nil
}

func (d Dictionary) Delete(word string) error {
	_, err := d.Search(word)

	switch err {
	case ErrNotFound:
		return ErrWordDoesNotExist
	case nil:
		delete(d, word)
	default:
		return err
	}

	return nil
}
