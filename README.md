# Go CLI Template

A GitHub template for Go CLI tools — cobra commands, viper config, structured output, tests, Docker, CI, and release automation included.

## Features

- **CLI framework** with [cobra](https://github.com/spf13/cobra) + [viper](https://github.com/spf13/viper)
- **Structured output** — JSON via `--json` flag, table via `--output table`
- **Configuration** via config file, env vars, and flags (viper handles precedence)
- **Tests** with table-driven patterns
- **Docker** multi-stage build (scratch image, ~8MB)
- **CI** via GitHub Actions (test, vet, build on push/PR)
- **Release** via [GoReleaser](https://goreleaser.com/) on tag push
- **Makefile** for common tasks

## Usage

1. Click **"Use this template"** on GitHub to create a new repo
2. Run the setup script:
   ```sh
   ./setup.sh mytool github.com/you/mytool
   ```
3. Add your commands in `cmd/` and logic in `internal/`
4. Delete the example command once you have your own

## Project Structure

```
.
├── cmd/
│   ├── root.go           # Root command, global flags, config init
│   ├── example.go         # Example subcommand (delete me)
│   └── version.go         # Version command
├── internal/
│   ├── config/
│   │   └── config.go      # Viper config binding
│   └── example/
│       └── example.go      # Example business logic (delete me)
├── .github/
│   └── workflows/
│       └── ci.yml
├── .goreleaser.yml
├── Dockerfile
├── Makefile
├── go.mod
├── main.go
├── setup.sh
└── README.md
```

## Quick Start

```sh
# Run locally
make run

# Run tests
make test

# Build binary
make build

# Build Docker image
make docker
```

## Adding a New Command

```go
// cmd/mycommand.go
package cmd

import (
	"github.com/spf13/cobra"
)

var myCmd = &cobra.Command{
	Use:   "mycommand",
	Short: "Does something useful",
	RunE: func(cmd *cobra.Command, args []string) error {
		// Your logic here
		return nil
	},
}

func init() {
	rootCmd.AddCommand(myCmd)
}
```

## Container Images

CI builds and pushes a container image to GHCR on every push to any branch.

```sh
# Pull the latest image
docker pull ghcr.io/<owner>/go-cli-template:latest

# Pull a specific commit
docker pull ghcr.io/<owner>/go-cli-template:<sha>

# Run
docker run ghcr.io/<owner>/go-cli-template:latest
```

Replace `<owner>` with your GitHub username or org. Images are tagged with both `latest` and the commit SHA.

## License

MIT