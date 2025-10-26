# Golan - A Go Test Project

This is a test project written in Go, demonstrating basic Go project structure, testing, and best practices.

## Project Structure

```
golan/
├── cmd/
│   └── main.go          # Main application entry point
├── math.go              # Math utility functions
├── math_test.go         # Unit tests for math functions
├── go.mod              # Go module definition
└── README.md           # This file
```

## Features

- Basic mathematical operations (Add, Subtract, Multiply, Divide)
- Comprehensive unit tests
- Example application demonstrating the usage

## Getting Started

### Prerequisites

- Go 1.24 or higher

### Installation

Clone the repository:
```bash
git clone https://github.com/IndWit/golan.git
cd golan
```

### Running the Application

```bash
go run cmd/main.go
```

### Running Tests

```bash
go test -v
```

### Building the Application

```bash
go build -o golan cmd/main.go
./golan
```

## Usage

The project provides simple math operations:

```go
import "github.com/IndWit/golan"

result := golan.Add(5, 3)        // Returns 8
result = golan.Subtract(5, 3)    // Returns 2
result = golan.Multiply(5, 3)    // Returns 15
result = golan.Divide(6, 3)      // Returns 2
```

## Testing

The project includes comprehensive unit tests for all functions. Run them with:

```bash
go test -v
```

## License

This is a test project for learning purposes.