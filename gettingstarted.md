# Getting Started

## Prerequisites

- Install [Postman](https://www.postman.com/)
- Install [VSCode](https://code.visualstudio.com/)
- Install [Go Language](https://golang.org/)
- Install WSL2 (Optional)
- [Create an account](https://github.com/signup) on GitHub
- [Fork the repository](https://github.com/JRHemmen/software-eng-course/fork) to your own account
- Clone the forked repo to your local machine using VSCodes command palette (Ctrl+Shift+P) and the command `Git: Clone`
- Attempt to run the code

### VSCode Extensions

Please search for and install the following extensions. You can do this by clicking on the Extensions icon on the left side of the window or by pressing `Ctrl+Shift+X` and searching for the extension name.

- Golang
- Powershell
- Live Server
- Markdown All in One
- Prettier
- GitLens
- OpenAPI (Swagger) Editor

## Running an Existing Go Project

```bash
cd /path/to/jrhemmen/software-eng-course/demos/godemo

go build ./path/to/your/main.go
# or to build and run
go run ./path/to/your/main.go
```

## Creating a New Go Project

```bash
go mod init github.com/yourusername/yourproject
```

## Go Project Structure

For a full explanation of the standard Go project structure, see [here](https://github.com/golang-standards/project-layout).

```plaintext
yourproject/
  cmd/
    yourproject/
      main.go
  internal/
    yourproject/
      yourpackage/
        yourpackage.go
  go.mod
  go.sum
```
