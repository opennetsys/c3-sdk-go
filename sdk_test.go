package sdk

import (
	"fmt"
	"testing"
)

func TestRegisterMethod(t *testing.T) {
	t.Parallel()
	c3 := NewC3()

	err := c3.RegisterMethod("setItem", []string{"string", "string"}, func(key, value string) error {
		fmt.Println("test setItem called with:", key, value)
		return nil
	})
	if err != nil {
		t.Error(err)
	}
}

func TestStore(t *testing.T) {
	t.Parallel()
	c3 := NewC3()

	err := c3.State().Set([]byte("foo"), []byte("bar"))
	if err != nil {
		t.Error(err)
	}
	value, found := c3.State().Get([]byte("foo"))
	if !found {
		t.Error("expected value")
	}
	if string(value) != "bar" {
		t.Error("expected match")
	}
}

func TestState(t *testing.T) {
	t.Parallel()
	c3 := NewC3()

	err := c3.RegisterMethod("setItem", []string{"string", "string"}, func(key, value string) error {
		fmt.Println("test setItem called with:", key, value)
		err := c3.State().Set([]byte(key), []byte(value))
		if err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		t.Error(err)
	}

	err = c3.Process([]byte(`[["setItem", "foo", "bar"],["setItem", "hello", "world"]]`))
	if err != nil {
		t.Error(err)
	}

	value, found := c3.State().Get([]byte("foo"))
	if !found {
		t.Error("expected value")
	}
	if string(value) != "bar" {
		t.Errorf("expected match; got %s", value)
	}

	value, found = c3.State().Get([]byte("hello"))
	if !found {
		t.Error("expected value")
	}
	if string(value) != "world" {
		t.Errorf("expected match; got %s", value)
	}
}
