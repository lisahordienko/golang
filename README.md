# Go + GenAI Course

This repository contains the homework for the Go + Generative AI course.

## Lesson 01

The lesson includes:
- a Hello World program
- a manual greeting program that reads a username from command-line arguments
- an AI-generated version of the same greeting program
- a short comparison report

### Run the examples

From the project root:

```bash
cd lesson-01/hello_world
go run .

cd ../greeting
go run . Alice

cd ../ai_greeting
go run . Bob
```

Each program is stored in its own package, which is the idiomatic Go way to keep multiple independent entry points in one lesson.
