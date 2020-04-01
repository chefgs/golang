// Simple Hi World unit test program
package main

import "testing"

func TestHello(t *testing.T) {
    expected := "Go hiworld program"
    if got := hifun(); got != expected {
        t.Errorf("hifun() = %q, want %q", got, expected)
    }
}
