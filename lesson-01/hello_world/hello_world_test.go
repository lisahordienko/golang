package main

import "testing"

func TestHelloWorldMessage(t *testing.T) {
    got := helloWorldMessage()
    want := "Hello, World!"

    if got != want {
        t.Fatalf("helloWorldMessage() = %q, want %q", got, want)
    }
}
