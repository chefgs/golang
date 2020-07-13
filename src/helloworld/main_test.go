package main

import "testing"

func TestHello(t *testing.T) {
    expected := "Hello World!"
    if got := print(); got != expected {
        t.Errorf("print() = %q, want %q", got, expected)
    }
}
