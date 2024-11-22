package main

import (
	"bytes"
	"testing"
)

func TestMain(t *testing.T) {
	t.Run("Buffer returns initial bytes", func(t *testing.T) {
		initialData := []byte("hello")
		b := NewBuffer(initialData)

		got := b.Bytes()

		if !bytes.Equal(got, initialData) {
			t.Errorf("got %v, expected %v", got, initialData)
		}
	})

	t.Run("Buffer returns initial plus new bytes", func(t *testing.T) {
		initialData := []byte("hello")
		b := NewBuffer(initialData)

		newData := []byte(" World")
		b.Write(newData)

		got := b.Bytes()
		expected := append(initialData, newData...)
		if !bytes.Equal(got, expected) {
			t.Errorf("got %v, expected %v", got, expected)
		}
	})

	t.Run("Read buffer in slice if it's big enough", func(t *testing.T) {
		initialData := []byte("hello")
		b := bytes.NewBuffer(initialData)

		buf := make([]byte, len(initialData))
		n, err := b.Read(buf)
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}

		if n != len(initialData) {
			t.Errorf("length: got %v, expected %v", n, len(initialData))
		}

		if !bytes.Equal(buf, initialData) {
			t.Errorf("data: got %v, expected %v", buf, initialData)
		}
	})

	t.Run("Read parts of butter into slice", func(t *testing.T) {
		initialData := []byte("hello World")
		b := bytes.NewBuffer(initialData)

		buf := make([]byte, 5)
		n, err := b.Read(buf)
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}

		buf2 := make([]byte, 6)
		n2, err := b.Read(buf2)
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}

		if n != 5 {
			t.Errorf("length: got %v, expected %v", n, 5)
		}
		if n2 != 6 {
			t.Errorf("length: got %v, expected %v", n2, 6)
		}

		bufs := append(buf, buf2...)
		if !bytes.Equal(bufs, initialData) {
			t.Errorf("data: got %v, expected %v", bufs, initialData)
		}
	})
}
