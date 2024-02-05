# Getting Started

## Prerequisites

- Install [Postman](https://www.postman.com/)
- Install [VSCode](https://code.visualstudio.com/)
- Install [Go Language](https://golang.org/)
- Install WSL2 (Optional)
- Create an account on Github.com
- Clone the repository using VSCode
- Attempt to run the code

### VSCode Extensions

Please search for and install the following extensions:

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
