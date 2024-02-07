# Getting Started

## Prerequisites

- Install [VSCode](https://code.visualstudio.com/) and the following extensions (`Ctrl+Shift+X`)
  - [Go](https://marketplace.visualstudio.com/items?itemName=golang.Go) for Go language support
  - [Powershell](https://marketplace.visualstudio.com/items?itemName=ms-vscode.PowerShell) for Powershell support
  - [Live Server](https://marketplace.visualstudio.com/items?itemName=ritwickdey.LiveServer) for HTML/CSS/JS live preview
  - [Markdown All in One](https://marketplace.visualstudio.com/items?itemName=yzhang.markdown-all-in-one) for Markdown support
  - [Prettier](https://marketplace.visualstudio.com/items?itemName=esbenp.prettier-vscode) for code formatting/linting
  - [GitLens](https://marketplace.visualstudio.com/items?itemName=eamodio.gitlens) for inline commit history
  - [OpenAPI (Swagger) Editor](https://marketplace.visualstudio.com/items?itemName=42Crunch.vscode-openapi) for OpenAPI support
  - [Remote Development Extension Pack](https://marketplace.visualstudio.com/items?itemName=ms-vscode-remote.vscode-remote-extensionpack) for opening code on a remote machine, container, or WSL
- Install [Go Language](https://golang.org/)
- Install [Postman](https://www.postman.com/) for API testing
- Install [Docker Desktop](https://www.docker.com/products/docker-desktop) for quickly testing dbs and apis
- Install [WSL2](https://learn.microsoft.com/en-us/windows/wsl/install) (Windows only, **optional**)
- [Create an account](https://github.com/signup) on GitHub
- [Create a Personal Access Token](https://github.com/settings/tokens/new) on GitHub
  - Select the `repo` scope
  - Copy the token to a safe place
- [Fork the repository](https://github.com/JRHemmen/software-eng-course/fork) to your own account
- Clone the forked repo to your local machine
  - Open the VSCode command pallette `(Ctrl+Shift+P)` and use command `Git: Clone`
  - Use the URL of your forked repository
  - Choose a location on your local machine
  - Provide your Github username and Token (**not password**) when prompted
- Attempt to run the code

## Running an Existing Go Project

To avoid any issues with the Go module system, it is recommended to open the specific demo folder (`$reporoot/demos/godemo`) in VSCode before running the following commands. This is due to the fact that the Go module system requires the `go.mod` file to be in the root of the project (Not nested inside a demos folder).

```bash
# to compile
go build cmd/godemo/main.go
# or to compile and run
go run cmd/godemo/main.go
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
