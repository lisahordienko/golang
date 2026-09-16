package main

import "testing"

func TestAIGreetingMessage(t *testing.T) {
    t.Run("valid name", func(t *testing.T) {
        got := aiGreetingMessage("Bob")
        want := "Hi, Bob! Welcome!"

        if got != want {
            t.Fatalf("aiGreetingMessage(%q) = %q, want %q", "Bob", got, want)
        }
    })

    t.Run("empty name", func(t *testing.T) {
        got := aiGreetingMessage(" ")
        want := "Please provide a username."

        if got != want {
            t.Fatalf("aiGreetingMessage(%q) = %q, want %q", " ", got, want)
        }
    })
}
