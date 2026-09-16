package main

import "testing"

func TestGreetingMessage(t *testing.T) {
    t.Run("valid name", func(t *testing.T) {
        got := greetingMessage("Alice")
        want := "Hello, Alice! Nice to meet you!"

        if got != want {
            t.Fatalf("greetingMessage(%q) = %q, want %q", "Alice", got, want)
        }
    })

    t.Run("empty name", func(t *testing.T) {
        got := greetingMessage(" ")
        want := "Please provide a username."

        if got != want {
            t.Fatalf("greetingMessage(%q) = %q, want %q", " ", got, want)
        }
    })
}
