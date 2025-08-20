# Comet CLI

A powerful command-line interface tool for creating and managing Comet Framework projects. Streamline your development workflow with automated project scaffolding and component generation.

![Go Version](https://img.shields.io/badge/Go-1.18%2B-blue)
![License](https://img.shields.io/badge/License-MIT-green)
![CLI](https://img.shields.io/badge/Type-CLI-orange)

## 📦 Installation

### Prerequisites
- Go 1.18 or higher
- Git

### Method 1: Install from source (Recommended)
```bash
go install github.com/ramoncl001/comet-cli@latest
```

### Method 2: Build from source
```bash
git clone https://github.com/ramoncl001/comet-cli.git
cd comet-cli
go build -o comet ./main.go
sudo mv comet /usr/local/bin/
```

## 🚀 Quickstart
```bash
# Create a new Comet project
comet-cli new my-project module.name

# Navigate to your project
cd my-project

# Add a controller
comet-cli add controller User controllers

# Add a service
comet-cli add service User services

# Add middleware
comet-cli add middleware Auth middlewares

# Run your project
go run .
```

## 📋 Commands
`comet-cli new [project-name] [module-name]`

Creates a new Comet project with the specified name and Go module path.

#### Arguments
* `project-name`: The directory name for your new project
* `module-name`: The Go module path (e.g., `github.com/username/project`)

#### Example:
```bash
comet-cli new awesome-api github.com/yourusername/awesome-api
```

`comet-cli add [type] [name] [location]`

Generates new components for your Comet project.

#### Arguments
* `type`: The type of component to create (`controller`, `service`, or `middleware`)
* `name`: The name of the component (e.g., `User`)
* `location`: The directory where the component should be created

#### Examples:
```bash
# Create a controller in the controllers directory
comet-cli add controller User controllers

# Create a service in the services directory
comet-cli add service User services

# Create middleware in the middleware directory
comet-cli add middleware Auth middlewares

# Create component in current directory
comet-cli add controller Health .
```

## 📁 Project Structure
When you create a new project with `comet-cli new`, it generates the following structure:

```text
your-project
├── go.mod
├── go.sum
├── infrastructure
├── main.go
├── middlewares
│   └── your_middleware.go
└── modules
    └── foo
        ├── controllers
        │   └── foo_controller.go
        ├── domain
        └── services
            └── foo_service.go
```

## 🧩 Generated Components
### Controllers

Handle HTTP requests and responses. Generated with proper Comet Framework structure including:

* Request handling

* Response generation

* Dependency injection support

* Route registration

### Services

Business logic components with:

* Dependency injection support

* Interface implementation

* Scoped lifecycle management

### Middleware

HTTP request processing middleware with:

* Standard Comet middleware interface

* Request/response interception

* Error handling

## 🔧 Requirements

* #### Go 1.18 or higher

* #### Comet Framework (automatically included as dependency)

## 📝 License
This project is licensed under the MIT License - see the LICENSE file for details.

## 🐛 Troubleshooting
### Common issues

#### 1. Command not found
```bash
# Ensure Go bin directory is in your PATH
export PATH=$PATH:$(go env GOPATH)/bin
```

#### 2. Permission denied
```bash
# On Unix systems, you might need to make the binary executable
chmod +x $(go env GOPATH)/bin/comet-cli
```

#### 3. Dependency errors
```bash
# Clean module cache and rebuild
go clean -modcache
go mod tidy
```