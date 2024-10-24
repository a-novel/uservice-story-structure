# Service: Story Structure

![GitHub Actions Workflow Status](https://img.shields.io/github/actions/workflow/status/a-novel/uservice-story-structure/main.yaml)
[![codecov](https://codecov.io/gh/a-novel/uservice-story-structure/graph/badge.svg?token=AR0GMEYZ4O)](https://codecov.io/gh/a-novel/uservice-story-structure)

![GitHub repo file or directory count](https://img.shields.io/github/directory-file-count/a-novel/uservice-story-structure)
![GitHub code size in bytes](https://img.shields.io/github/languages/code-size/a-novel/uservice-story-structure)

![Coverage graph](https://codecov.io/gh/a-novel/uservice-story-structure/graphs/sunburst.svg?token=AR0GMEYZ4O)

Story structure microservice.

### Prerequisites

- [Go](https://go.dev/doc/install)
- Make
    - macOS:
      ```bash
      brew install make
      ```
    - Ubuntu:
      ```bash
      sudo apt-get install make
      ```
    - Windows: Install [chocolatey](https://chocolatey.org/install) (from a PowerShell with admin privileges), then run:
      ```bash
      choco install make
      ```

Install the project dependencies.

```bash
go get ./... && go mod tidy
```

## Work on the project

Make sure the project files are properly formatted.

```bash
make format
```

Run tests.

```bash
make test
```

Make sure your code is compliant with the linter.

```bash
make lint
```

If you create / update interfaces signatures, make sure to update the mocks.

```bash
make mocks
```
