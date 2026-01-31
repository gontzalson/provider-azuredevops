# Provider Azure DevOps

`provider-azuredevops` is a [Crossplane](https://crossplane.io/) Provider built using [Upjet v2](https://github.com/crossplane/upjet) that wraps the [Terraform Azure DevOps Provider](https://registry.terraform.io/providers/microsoft/azuredevops/latest/docs) to enable declarative management of Azure DevOps resources via Kubernetes.

## Features

### Manage Azure DevOps from Kubernetes 🚀

- **GitOps-native**: Declaratively manage Azure DevOps resources using Kubernetes manifests
- **Crossplane-powered**: Leverage Crossplane's composition and claim features for platform abstraction
- **Production-ready**: Built on the official Upjet v2 provider template with modern best practices

### Currently Supported Resources

**Project Management:**
- `Project` - Azure DevOps projects
- `ProjectFeatures` - Feature toggles for projects (Boards, Repos, Pipelines, Test Plans, Artifacts)
- `ProjectPipelineSettings` - Pipeline configuration and security settings

**Coming Soon:**
- Git Repositories & Branch Policies
- Build Pipelines & Variable Groups
- Service Connections (Azure RM, GitHub, Kubernetes, Docker Registry)
- And more!

## Quick Start

### Prerequisites

- Kubernetes cluster (1.24+)
- [Crossplane installed](https://crossplane.io/docs/latest/software/install/)
- Azure DevOps organization with admin access
- [Personal Access Token (PAT)](https://docs.microsoft.com/en-us/azure/devops/organizations/accounts/use-personal-access-tokens-to-authenticate)

### Installation

1. **Install the Provider:**

```bash
cat <<EOF | kubectl apply -f -
apiVersion: pkg.crossplane.io/v1
kind: Provider
metadata:
  name: provider-azuredevops
spec:
  package: gontzalson/provider-azuredevops:v0.1.0
EOF
```

2. **Configure Provider Authentication:**

Create a Kubernetes secret with your Azure DevOps PAT:

```bash
kubectl create secret generic azuredevops-credentials \
  --from-literal=credentials='{"pat":"YOUR_PAT_HERE","org_service_url":"https://dev.azure.com/YOUR_ORG"}' \
  -n crossplane-system
```

3. **Create a ProviderConfig:**

```yaml
apiVersion: azuredevops.crossplane.io/v1beta1
kind: ProviderConfig
metadata:
  name: default
spec:
  credentials:
    source: Secret
    secretRef:
      namespace: crossplane-system
      name: azuredevops-credentials
      key: credentials
```

### Usage Examples

#### Create a Project

```yaml
apiVersion: project.azuredevops.crossplane.io/v1alpha1
kind: Project
metadata:
  name: my-project
spec:
  forProvider:
    name: My Awesome Project
    description: Managed by Crossplane
    visibility: private
    versionControl: Git
    workItemTemplate: Agile
    features:
      boards: enabled
      repositories: enabled
      pipelines: enabled
      testplans: disabled
      artifacts: enabled
  providerConfigRef:
    name: default
```

#### Configure Project Features

```yaml
apiVersion: project.azuredevops.crossplane.io/v1alpha1
kind: Features
metadata:
  name: my-project-features
spec:
  forProvider:
    projectIdRef:
      name: my-project
    features:
      boards: enabled
      repositories: enabled
      pipelines: enabled
      testplans: disabled
      artifacts: enabled
  providerConfigRef:
    name: default
```

#### Set Pipeline Settings

```yaml
apiVersion: project.azuredevops.crossplane.io/v1alpha1
kind: PipelineSettings
metadata:
  name: my-project-pipeline-settings
spec:
  forProvider:
    projectIdRef:
      name: my-project
    enforceSettability: true
    publishPipelineMetadata: true
    statusBadgesArePrivate: false
  providerConfigRef:
    name: default
```

## Development

### Building from Source

```bash
# Clone the repository
git clone https://github.com/gontzalson/provider-azuredevops
cd provider-azuredevops

# Install dependencies
go mod download

# Generate CRDs and controllers
export PATH="$HOME/go/bin:$PATH"
go run cmd/generator/main.go .

# Build the provider
make build

# Build Docker image
make build.all
```

### Adding New Resources

1. Add the resource to `config/external_name.go`
2. Create a config file in `config/<group>/config.go`
3. Run the generator: `go run cmd/generator/main.go .`
4. Generate Kubernetes artifacts: `make generate`

See the [Upjet documentation](https://github.com/crossplane/upjet) for details.

## Architecture

This provider uses the Upjet framework to automatically generate Crossplane Custom Resources from Terraform provider schemas:

```
Kubernetes API
      ↓
Crossplane Provider (this)
      ↓
Terraform Azure DevOps Provider
      ↓
Azure DevOps REST API
```

**Key Components:**
- **CRDs**: Auto-generated from Terraform provider schema
- **Controllers**: Reconcile Kubernetes resources with Azure DevOps state
- **Terraform Workspace**: Manages Terraform state for each managed resource

## Roadmap

- [x] Project resources
- [ ] Git repository resources
- [ ] Pipeline & build resources
- [ ] Service endpoint resources
- [ ] Work item tracking
- [ ] Test plans
- [ ] Wiki resources
- [ ] Comprehensive examples
- [ ] E2E tests
- [ ] CI/CD pipeline

## Contributing

Contributions are welcome! Please:

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/amazing-feature`)
3. Make your changes
4. Run tests: `make test`
5. Commit your changes (`git commit -m 'Add amazing feature'`)
6. Push to the branch (`git push origin feature/amazing-feature`)
7. Open a Pull Request

## Community & Support

- **Issues**: [GitHub Issues](https://github.com/gontzalson/provider-azuredevops/issues)
- **Discussions**: [GitHub Discussions](https://github.com/gontzalson/provider-azuredevops/discussions)
- **Crossplane Slack**: [#providers channel](https://crossplane.slack.com)

## Resources

- [Crossplane Documentation](https://crossplane.io/docs)
- [Upjet Documentation](https://github.com/crossplane/upjet)
- [Terraform Azure DevOps Provider](https://registry.terraform.io/providers/microsoft/azuredevops/latest/docs)
- [Azure DevOps REST API](https://docs.microsoft.com/en-us/rest/api/azure/devops/)

## License

Apache 2.0 - see [LICENSE](LICENSE) file for details.

---

**Built with ❤️ using Crossplane and Upjet v2**
