package dkey

import "errors"

var ErrorNoSuchKey = errors.New("no such key")

type DKey struct {
	store map[string]string
}

func NewDKey() (DKey, error) {
	return DKey{store: make(map[string]string)}, nil
}

func (dk *DKey) Get(key string) (string, error) {
	val, ok := dk.store[key]
	if !ok {
		return "", ErrorNoSuchKey
	}
	return val, nil
}

func (dk *DKey) Put(key string, value string) error {
	dk.store[key] = value
	return nil
}

func (dk *DKey) Delete(key string) error {
	delete(dk.store, key)
	return nil
}
