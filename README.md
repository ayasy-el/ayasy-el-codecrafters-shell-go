# Custom Go Shell - CodeCrafters Challenge

[![Progress Banner](https://backend.codecrafters.io/progress/shell/89ee2e2b-cf8d-4f05-8e57-de62db4caa6d)](https://app.codecrafters.io/users/codecrafters-bot?r=2qF)

## Challenge Overview

This project is a solution to the ["Build Your Own Shell" Challenge](https://app.codecrafters.io/courses/shell/overview) by CodeCrafters. The goal is to build a POSIX-compliant shell capable of:

- Interpreting shell commands
- Running external programs
- Implementing builtin commands (cd, pwd, echo, etc.)

## Project Structure

- **Entry Point**: `cmd/myshell/main.go`
- **Main Implementation**: Custom Go shell with advanced command parsing and execution

## Features

### Built-in Commands

The shell supports the following built-in commands:

1. **exit**: Terminate the shell, with an optional exit code
   - Usage: `exit [code]`
   - Example: `exit 0` or `exit 1`

2. **echo**: Print arguments to the console
   - Usage: `echo [arguments]`
   - Example: `echo Hello, World!`

3. **type**: Identify the type of a command (shell builtin or external command)
   - Usage: `type [command]`
   - Example: `type cd` or `type ls`

4. **pwd**: Print the current working directory
   - Usage: `pwd`

5. **cd**: Change the current working directory
   - Usage: `cd [directory]`
   - Supports home directory expansion (`~`)
   - Example: `cd ~` or `cd /home/user/documents`

### Advanced Parsing Features

- Handles single and double quotes
- Supports escape characters
- Preserves quoted content and complex argument structures
- Searches for executable commands in system PATH

## Getting Started

### Prerequisites

- Go 1.19 or higher
- Git
- Unix-like operating system (Linux, macOS)

### Running the Shell

1. Clone the repository
   ```bash
   git clone https://github.com/ayasy-el/codecrafters-shell-go.git
   cd codecrafters-shell-go
   ```

2. Run the shell
   ```bash
   ./your_program.sh
   ```

### Submitting Solutions

To submit a solution to a stage:

1. Implement the required functionality in `cmd/myshell/main.go`
2. Commit your changes
   ```bash
   git commit -am "pass Xth stage"
   ```
3. Push to master
   ```bash
   git push origin master
   ```

## Development Stages

### Stage 1: Basic Shell Setup
- Create the initial shell program
- Implement basic input loop
- Support simple command parsing

### Stage 2 & Beyond
- Progressively add more complex shell features
- Implement builtin commands
- Enhance command parsing and execution

## Implementation Highlights

- Uses Go's standard library for system interactions
- Sophisticated argument and command parsing
- Flexible command execution model
- Supports both builtin and external commands

## Limitations

- Does not support advanced shell features like piping or redirection
- Minimal error handling for complex scenarios
- No shell scripting capabilities

## Learning Objectives

Through this challenge, you'll gain insights into:
- Shell command parsing techniques
- REPL (Read-Eval-Print Loop) implementation
- System command execution
- Advanced string manipulation in Go

## Contributing

Contributions and improvements are welcome! Feel free to submit pull requests or open issues.


## Acknowledgments

- [CodeCrafters](https://codecrafters.io) for the challenging shell implementation project