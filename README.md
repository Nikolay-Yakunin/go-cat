# go-cat

[![Go Report Card](https://goreportcard.com/badge/github.com/Nikolay-Yakunin/go-cat)](https://goreportcard.com/report/github.com/Nikolay-Yakunin/go-cat)
[![License](https://img.shields.io/badge/license-MIT-blue.svg)](LICENSE)

A Go implementation of the Unix `cat` command with additional features.

## Features

- Supports various flags for customizing output:
  - `--number-nonblank`: Number non-empty output lines.
  - `--squeeze-blank`: Squeeze multiple adjacent blank lines.
  - `--number`: Number all output lines.
  - Short flags: `-b`, `-E`, `-e`, `-n`, `-s`, `-T`, `-t`, `-v`.

## Installation

### Prerequisites

- [Go](https://golang.org/dl/) version 1.22.6 or higher.

### Build from Source

1. Clone the repository:
   ```sh
   git clone https://github.com/Nikolay-Yakunin/go-cat.git
   cd go-cat
   ```

2. Build the project:
   ```sh
   make
   ```

   This will build binaries for both Linux and Windows in the `build` directory.

3. Run the executable:
   - For Linux:
     ```sh
     ./build/linux/go-cat --number file.txt
     ```
   - For Windows:
     ```sh
     .\build\windows\go-cat.exe --number file.txt
     ```

## Usage

```sh
go-cat [OPTION]... [FILE]...
```

### Examples

- Display the contents of a file:
  ```sh
  go-cat file.txt
  ```

- Number all output lines:
  ```sh
  go-cat -n file.txt
  ```

- Number non-empty output lines:
  ```sh
  go-cat -b file.txt
  ```

- Squeeze multiple adjacent blank lines:
  ```sh
  go-cat -s file.txt
  ```

## Contributing

Contributions are welcome! Please follow these steps:

1. Fork the repository.
2. Create a new branch (`git checkout -b feature/new-feature`).
3. Commit your changes (`git commit -am 'Add some feature'`).
4. Push to the branch (`git push origin feature/new-feature`).
5. Create a new Pull Request.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Acknowledgments

- Thanks to the Go community for providing such a powerful language and tools.
