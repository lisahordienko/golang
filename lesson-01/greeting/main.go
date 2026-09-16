package main

import (
    "fmt"
    "os"
    "strings"
)

func greetingMessage(name string) string {
    if strings.TrimSpace(name) == "" {
        return "Please provide a username."
    }

    return fmt.Sprintf("Hello, %s! Nice to meet you!", name)
}

func main() {
    if len(os.Args) < 2 {
        fmt.Println("Usage: go run . <username>")
        return
    }

    name := os.Args[1]
    fmt.Println(greetingMessage(name))
}
