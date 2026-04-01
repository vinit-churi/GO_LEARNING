# Go Modules - Reference

## What is a Module?

A module is a collection of related Go packages that are versioned together as a single unit. The module is defined by a `go.mod` file in its root directory.

## The go.mod File

Example `go.mod` file:

```
module github.com/username/project

go 1.21

require (
    github.com/pkg/errors v0.9.1
)
```

Components:
- **module directive:** Declares the module path (its import path)
- **go directive:** Specifies the minimum Go version required
- **require directive:** Lists module dependencies with their versions

## Initializing a Module

```bash
go mod init example.com/myproject
```

This creates a `go.mod` file with:
- The module path you specified
- The Go version from your installed Go toolchain

## Module Paths

A module path is the import prefix for all packages within the module.

**Common patterns:**
- `github.com/username/repository` - for GitHub projects
- `example.com/project` - for examples and learning
- `company.com/team/project` - for company projects

**Rules:**
- Must be globally unique if you plan to publish
- Cannot contain spaces or special characters (except `/`, `.`, `-`, `_`)
- Case-sensitive

## Packages vs Modules

**Package:**
- A collection of `.go` files in the same directory
- Declared with `package packagename` at the top of each file
- Smallest unit of code organization

**Module:**
- A collection of packages
- Has a `go.mod` file
- Unit of versioning and distribution

Example structure:
```
mymodule/
  go.mod                    # module declaration
  main.go                   # package main
  utils/
    helpers.go              # package utils
  internal/
    config/
      config.go             # package config
```

## The main Package

The `main` package is special:
- Must contain a `func main()` - the program entry point
- Compiles to an executable binary
- Cannot be imported by other packages

## Importing Packages

**Standard library:**
```go
import "fmt"
import "strings"
```

**Your own packages:**
```go
import "github.com/username/project/utils"
```

**External packages:**
```go
import "github.com/pkg/errors"
```

## Package Names and Paths

The last element of the import path is typically the package name:

```go
import "github.com/username/project/utils"  // package name is usually "utils"
```

But the actual package name is declared in the `.go` files:

```go
package utils  // this is what you use in your code
```

## Semantic Versioning

Go modules use semantic versioning (semver):

```
v1.2.3
│ │ │
│ │ └─── PATCH: backwards-compatible bug fixes
│ └───── MINOR: backwards-compatible new features
└─────── MAJOR: incompatible API changes
```

**Rules:**
- v0.x.x is for initial development (no stability guarantees)
- v1.0.0 is the first stable release
- Major version ≥2 must be part of the module path: `github.com/user/project/v2`

## Common Commands

```bash
go mod init <module-path>   # Initialize a new module
go mod tidy                 # Add missing and remove unused dependencies
go mod vendor              # Copy dependencies to vendor directory
go get <package>           # Add or update a dependency
go list -m all            # List all module dependencies
```

## Why Modules?

**Before modules:**
- GOPATH was required
- No built-in versioning
- Difficult dependency management

**With modules:**
- Work anywhere (no GOPATH needed)
- Reproducible builds
- Clear dependency versioning
- Easier project organization

## Common Mistakes

1. **Forgetting to initialize:** Running `go` commands without `go.mod` can cause confusion
2. **Wrong module path:** Using spaces or invalid characters
3. **Mixing packages:** Putting multiple main packages in the same module root
4. **Not running go mod tidy:** Leaving unused dependencies in go.mod
5. **Confusion between package and module names:** They're related but distinct concepts

## When To Use

- **Always use modules for new projects** - it's the standard Go dependency management
- **Initialize at project start** - run `go mod init` before writing code
- **One module per project** - unless you're building a multi-module workspace (advanced)

## Best Practices

1. **Module path should reflect hosting:** Use github.com/username/repo for GitHub
2. **Keep go.mod clean:** Run `go mod tidy` regularly
3. **Version tags on releases:** Tag commits like `v1.0.0` when releasing
4. **Vendor if needed:** Use `go mod vendor` for controlled builds
5. **Document package purpose:** Add package comments in each package

## Additional Resources

- Go Modules Reference: https://go.dev/ref/mod
- Module Tutorial: https://go.dev/doc/tutorial/create-module
- go.mod file reference: https://go.dev/doc/modules/gomod-ref
