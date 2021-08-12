package main

import (
	"errors"
	"testing"
)

func TestCreateDKey(t *testing.T) {
	_, err := NewDKey()
	if err != nil {
		t.Errorf("error when creating new DKey: %e", err)
	}
}

func TestPut(t *testing.T) {
	dk, err := NewDKey()
	if err != nil {
		t.Errorf("error when creating new DKey: %e", err)
	}

	err = dk.Put("test", "test")
	if err != nil {
		t.Errorf("error when putting into DKey: %e", err)
	}
}

func TestGet(t *testing.T) {

	testKey, testVal := "test-key", "test-value"

	dk, err := NewDKey()
	if err != nil {
		t.Errorf("error when creating new DKey: %e", err)
	}

	// This should return sentinel error `ErrorNoSuchKey` because
	// the key does not exist.
	v, err := dk.Get(testKey)
	if !errors.Is(err, ErrorNoSuchKey) && err != nil {
		t.Errorf("unexpected error getting from empty DKey: %e", err)
	}
	if err == nil {
		t.Errorf("no sentinel-error returned when GETting from empty  result: \"%v\"", v)
	}

	err = dk.Put(testKey, testVal)
	if err != nil {
		t.Errorf("error PUTting to DKey: %e", err)
	}

	v, err = dk.Get(testKey)
	if err != nil {
		t.Errorf("unexpected error getting value from non-empty DKey: %e", err)
	}

	if v != testVal {
		t.Errorf("unexpected value returned from GET. expected=\"%v\" got=\"%v\"", testVal, v)
	}
}
