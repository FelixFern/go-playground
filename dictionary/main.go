package dictionary

import (
	"errors"
	"fmt"
)

var ErrNotFound = errors.New("could not find the word you were looking for")

func (d Dictionary) Search(name string) (Language, error) {
	res, ok := d[name]

	if !ok {
		fmt.Println("Key not found")
		return Language{}, ErrNotFound
	}

	return res, nil
}

func (d Dictionary) Add(key string, translation Language) error {
	d[key] = translation

	return nil
}

func (d Dictionary) Update(key string, translation Language) {
	d[key] = translation

}
