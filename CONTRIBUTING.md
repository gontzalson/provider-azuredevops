# Contributing to provider-azuredevops

Thank you for your interest in contributing to provider-azuredevops! This guide will help you get started with contributing to this Crossplane provider for Azure DevOps.

## Table of Contents

- [Code of Conduct](#code-of-conduct)
- [Getting Started](#getting-started)
- [Development Environment Setup](#development-environment-setup)
- [Building and Testing](#building-and-testing)
- [Codebase Architecture](#codebase-architecture)
- [Adding New Resources](#adding-new-resources)
- [Testing Guidelines](#testing-guidelines)
- [Pull Request Process](#pull-request-process)
- [Commit Message Guidelines](#commit-message-guidelines)
- [Community and Support](#community-and-support)

## Code of Conduct

This project adheres to the [Code of Conduct](CODE_OF_CONDUCT.md). By participating, you are expected to uphold this code.

## Getting Started

Before you begin:
- Familiarize yourself with [Crossplane](https://crossplane.io/docs)
- Review the [Upjet documentation](https://github.com/crossplane/upjet)
- Understand the [Terraform Azure DevOps Provider](https://registry.terraform.io/providers/microsoft/azuredevops/latest/docs)
- Read our [README](README.md) to understand the project's goals

## Development Environment Setup

### Prerequisites

- **Go**: Version 1.24 or later ([installation guide](https://golang.org/doc/install))
- **Docker**: For building container images ([installation guide](https://docs.docker.com/get-docker/))
- **Kubernetes cluster**: For testing (use [kind](https://kind.sigs.k8s.io/) for local development)
- **kubectl**: Kubernetes command-line tool
- **make**: Build automation tool
- **Git**: Version control

### Clone the Repository

```bash
git clone https://github.com/gontzalson/provider-azuredevops.git
cd provider-azuredevops
```

### Install Dependencies

The project uses Go modules for dependency management:

```bash
go mod download
```

### Set up Build Tools

The first time you run `make`, it will set up the necessary build submodules:

```bash
make
```

## Building and Testing

### Local Build

Build the provider binary:

```bash
make build
```

Build Docker images:

```bash
make build.all
```

### Generate Code

After modifying API types or adding new resources, regenerate the code:

```bash
make generate
```

### Linting and Formatting

Ensure your code follows the project's coding standards:

```bash
# Run linter
make lint

# Format code and run all reviewable checks
make reviewable
```

### Running Tests

```bash
# Run all unit tests
make test

# Run tests with coverage
make test TESTARGS="-cover"

# Run tests for a specific package
go test -v ./internal/clients/...
```

### Local Deployment

Deploy the provider to a local Kubernetes cluster:

```bash
# Create a kind cluster (if you don't have one)
kind create cluster

# Install Crossplane
kubectl create namespace crossplane-system
helm repo add crossplane-stable https://charts.crossplane.io/stable
helm install crossplane crossplane-stable/crossplane --namespace crossplane-system

# Build and load the provider image
make build.all
kind load docker-image gontzalson/provider-azuredevops:latest

# Install the provider
kubectl apply -f examples/install.yaml
```

## Codebase Architecture

### Directory Structure

```
provider-azuredevops/
├── apis/                      # API definitions
│   ├── cluster/              # Cluster-scoped resources
│   │   └── v1beta1/          # ProviderConfig API
│   ├── namespaced/           # Namespaced resources
│   │   └── v1beta1/          # Namespaced ProviderConfig API
│   └── <group>/              # Resource group APIs (e.g., project)
│       └── v1alpha1/         # Resource CRDs
├── cmd/
│   ├── generator/            # Code generator for CRDs
│   └── provider/             # Provider controller entrypoint
├── config/                   # Provider configuration
│   ├── external_name.go      # External name configurations
│   ├── provider.go           # Provider configuration
│   └── <group>/              # Group-specific configurations
│       └── config.go         # Resource configurations
├── internal/
│   ├── clients/              # Terraform client setup
│   ├── controller/           # Controller implementations
│   └── version/              # Version information
├── package/                  # Crossplane package metadata
├── cluster/                  # Docker image configuration
└── examples/                 # Example manifests
```

### Key Components

1. **Code Generator** (`cmd/generator`): Generates Crossplane CRDs from Terraform provider schemas
2. **Controllers** (`internal/controller`): Reconcile Kubernetes resources with Azure DevOps
3. **Configuration** (`config/`): Maps Terraform resources to Crossplane CRDs
4. **Client Setup** (`internal/clients`): Configures Terraform provider credentials

### How It Works

```
Kubernetes API
     ↓
Crossplane Provider (this project)
     ↓
Terraform Azure DevOps Provider
     ↓
Azure DevOps REST API
```

The provider:
1. Watches for changes to custom resources in Kubernetes
2. Uses Upjet to manage Terraform workspaces for each resource
3. Translates CRD spec to Terraform configuration
4. Applies changes via the Terraform Azure DevOps provider
5. Updates resource status based on Terraform state

## Adding New Resources

### Step 1: Add External Name Configuration

Edit `config/external_name.go` to add your resource:

```go
var ExternalNameConfigs = map[string]config.ExternalName{
    // Existing resources...
    "azuredevops_git_repository": config.IdentifierFromProvider,
}
```

External name strategies:
- `config.IdentifierFromProvider`: Use Terraform provider's ID
- `config.NameAsIdentifier`: Use the resource name as external ID
- `config.ExternalName{...}`: Custom configuration

### Step 2: Create Group Configuration

Create a new file `config/<group>/config.go`:

```go
package repository

import "github.com/crossplane/upjet/v2/pkg/config"

// Configure configures the repository group
func Configure(p *config.Provider) {
    p.AddResourceConfigurator("azuredevops_git_repository", func(r *config.Resource) {
        r.ShortGroup = "repository"
        r.Kind = "GitRepository"
        // Add custom configuration if needed
    })
}
```

### Step 3: Register the Group

In `config/provider.go`, import and call your configure function:

```go
import (
    // ... other imports
    "github.com/gontzalson/provider-azuredevops/config/repository"
)

func GetProvider() *config.Provider {
    pc := config.NewProvider(/* ... */)

    // ... existing configurations
    repository.Configure(pc)

    return pc
}
```

### Step 4: Generate Resources

Run the generator to create CRDs and controllers:

```bash
go run cmd/generator/main.go .
```

### Step 5: Generate Kubernetes Artifacts

Generate the final manifests:

```bash
make generate
```

### Step 6: Test Your Resource

Create an example manifest in `examples/<group>/<resource>/`:

```yaml
apiVersion: repository.azuredevops.crossplane.io/v1alpha1
kind: GitRepository
metadata:
  name: example-repo
spec:
  forProvider:
    projectIdRef:
      name: my-project
    name: example-repository
    initialization:
      initType: Clean
  providerConfigRef:
    name: default
```

## Testing Guidelines

### Unit Tests

- Write unit tests for custom logic in `internal/clients/`
- Use table-driven tests for multiple scenarios
- Mock external dependencies

Example:
```go
func TestTerraformSetupBuilder(t *testing.T) {
    tests := []struct {
        name    string
        input   map[string]string
        wantErr bool
    }{
        // test cases
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            // test logic
        })
    }
}
```

### Integration Tests

- Test actual resource creation/update/deletion
- Use the `uptest` framework for end-to-end tests
- Ensure tests clean up resources

### Manual Testing

1. Deploy provider to a test cluster
2. Create ProviderConfig with valid credentials
3. Apply example manifests
4. Verify resources are created in Azure DevOps
5. Update and delete resources to test full lifecycle

## Pull Request Process

### Before Submitting

1. **Create an issue** (if one doesn't exist) to discuss the change
2. **Fork the repository** and create a feature branch from `main`
3. **Make your changes** following the coding standards
4. **Run tests and linters**:
   ```bash
   make reviewable
   make test
   make lint
   ```
5. **Update documentation** (README, examples, etc.)
6. **Commit with sign-off**: `git commit -s`

### PR Guidelines

1. **Title**: Use [Conventional Commits](https://www.conventionalcommits.org/) format
   - `feat: add Git repository resource`
   - `fix(project): resolve nil pointer in project controller`
   - `docs: update contributing guide`

2. **Description**:
   - Reference the issue: `Fixes #123`
   - Explain what changed and why
   - Include testing details

3. **Checklist**: Complete the PR template checklist

4. **Size**: Keep PRs focused and reasonably sized
   - Large changes should be broken into multiple PRs
   - Discuss architectural changes before implementing

### Review Process

1. Automated CI checks must pass
2. At least one maintainer approval required
3. Address review feedback promptly
4. Maintainers may request changes or tests
5. Once approved, maintainers will merge

## Commit Message Guidelines

We follow [Conventional Commits](https://www.conventionalcommits.org/):

### Format

```
<type>(<scope>): <subject>

<body>

<footer>
```

### Types

- `feat`: New feature
- `fix`: Bug fix
- `docs`: Documentation changes
- `refactor`: Code refactoring
- `test`: Adding tests
- `chore`: Maintenance tasks
- `ci`: CI/CD changes
- `perf`: Performance improvements

### Examples

```
feat(repository): add Git repository resource

Implements support for azuredevops_git_repository resource,
enabling declarative management of Git repositories.

Fixes #42
```

```
fix(project): prevent nil pointer when features not specified

Added nil check before accessing project features to prevent
panic when features field is not provided in the spec.

Fixes #56
```

### Sign-off Requirement

All commits must be signed off to certify you agree to the [Developer Certificate of Origin (DCO)](https://developercertificate.org/):

```bash
git commit -s -m "feat: add new feature"
```

## Community and Support

### Communication Channels

- **GitHub Issues**: Bug reports and feature requests
- **GitHub Discussions**: Questions and discussions
- **Crossplane Slack**: [#providers channel](https://crossplane.slack.com)

### Getting Help

- Check existing [issues](https://github.com/gontzalson/provider-azuredevops/issues)
- Review [Crossplane documentation](https://crossplane.io/docs)
- Ask in [GitHub Discussions](https://github.com/gontzalson/provider-azuredevops/discussions)

### Reporting Bugs

When filing an issue, include:
- Provider version
- Kubernetes version
- Crossplane version
- Steps to reproduce
- Expected vs actual behavior
- Relevant logs or error messages

### Suggesting Features

Before suggesting a feature:
- Check if it already exists in the Terraform Azure DevOps provider
- Search for existing feature requests
- Describe the use case and benefits
- Consider implementation complexity

## Additional Resources

- [Crossplane Documentation](https://crossplane.io/docs)
- [Upjet Documentation](https://github.com/crossplane/upjet)
- [Terraform Azure DevOps Provider](https://registry.terraform.io/providers/microsoft/azuredevops/latest/docs)
- [Azure DevOps REST API](https://docs.microsoft.com/en-us/rest/api/azure/devops/)
- [Crossplane Contributing Guide](https://github.com/crossplane/crossplane/blob/master/CONTRIBUTING.md)

## License

By contributing, you agree that your contributions will be licensed under the Apache 2.0 License.

---

Thank you for contributing to provider-azuredevops!
