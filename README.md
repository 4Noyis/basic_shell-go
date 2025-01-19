# Custom Shell

Welcome to my custom shell project. This shell is built from scratch using Go and includes implementations of several basic shell commands. Additionally, it can execute external programs with arguments, making it a versatile tool for command-line interactions.

---

## Features

### Built-in Commands

The shell comes with the following built-in commands:

- **`ls`**: List directory contents.
- **`cd`**: Change the current working directory.
- **`cat`**: Display the contents of a file.
- **`echo`**: Print text to the console.
- **`type`**: Display whether a command is built-in or external.
- **`exit`**: Exit the shell.
- **`clear`**: Clear the terminal screen.
- **`pwd`**: Print the current working directory.

### External Program Execution

The shell supports running external programs with arguments. For example:

```bash
$ gcc --version
$ python3 script.py
```

---

## Getting Started

### Prerequisites

- Go (Golang) installed on your system.
- A Unix-like environment for best compatibility.

### Installation

1. Clone this repository:
   ```bash
   git clone <repository_url>
   cd <repository_name>
   ```
2. Build the project:
   ```bash
   go build -o my_shell
   ```
3. Run the shell:
   ```bash
   ./my_shell
   ```

---

## Code Structure

- **`main.go`**: The entry point of the application.
- **`models/`**: Contains structures and utilities for handling commands.
- **`commands/`**: Implements the built-in commands.
- **`utils/`**: Helper functions for common operations.

---
