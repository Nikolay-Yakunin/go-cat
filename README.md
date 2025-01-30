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

## Adding Binary to PATH
To make it easier to use go-cat from anywhere, you can add the binary to your system's PATH. We provide scripts for both Bash (for Unix-like systems) and PowerShell (for Windows).

For Unix-like Systems (Bash Script)
Ensure the script has execute permissions:
```sh
chmod +x install.sh
```
Run the script as root:
```sh
sudo ./install.sh
```
This script will build the binary, copy it to /usr/local/bin, and ensure it is available in your PATH.

## For Windows Systems (PowerShell Script) if you have scripts allowed to run
Open PowerShell as Administrator.
Run the script:
```powershell
.\install.ps1
```
This script will build the binary, copy it to a directory like C:\Program Files\mycat, and add this directory to your system's PATH. You may need to restart your terminal or system for the changes to take effect.

## How to scripts allowed to run
Opening PowerShell with administrator rights :
Make sure that you have opened PowerShell with administrator rights. To do this, press Win + X and select "Windows PowerShell (Admin)" or "Terminal (Admin)" if you have Windows 10/11.
Checking the current execution policy :
You can check the current execution policy by running the following command:
```powershell
Get-ExecutionPolicy
```
If the result is Restricted, it means that scripting is prohibited.
Changing the execution policy :
Temporarily change the execution policy to a more permissive one, for example, RemoteSigned or Bypass. Run one of the following commands in PowerShell with administrator privileges:
To allow the execution of local scripts:
```powershell
Set-ExecutionPolicy RemoteSigned -Scope Process
```
To completely disable script execution verification (only for the current process):
```powershell
Set-ExecutionPolicy Bypass -Scope Process
```
After executing one of these commands, try running the script again.:
powershell
```Copy
.\install.ps1
```
Return to the original execution policy :
After you finish working with the scripts, you can return the execution policy to its original state by running:
```powershell
Set-ExecutionPolicy Restricted -Scope Process
```

Or set a policy at the user or system level, if necessary.:
```powershell
Set-ExecutionPolicy Restricted -Scope CurrentUser
```
or
```powershell
Set-ExecutionPolicy Restricted -Scope LocalMachine
```
Safety Notes
- RemoteSigned : This option allows you to execute all local scripts and remote scripts that have a valid digital signature. This is the safest option for most users.
- Bypass : This option completely disables script execution verification. It is only used for temporary use in a trusted environment.

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
