# Provider Azure DevOps

`provider-azuredevops` is a comprehensive [Crossplane](https://crossplane.io/) Provider built using [Upjet v2](https://github.com/crossplane/upjet) that wraps the [Terraform Azure DevOps Provider](https://registry.terraform.io/providers/microsoft/azuredevops/latest/docs) to enable declarative management of Azure DevOps resources via Kubernetes.

## Features

### Complete Azure DevOps Management from Kubernetes 🚀

- **92 Managed Resources**: Full coverage of the Terraform Azure DevOps provider (v1.4.0)
- **GitOps-native**: Declaratively manage Azure DevOps resources using Kubernetes manifests
- **Crossplane-powered**: Leverage Crossplane's composition and claim features for platform abstraction
- **Production-ready**: Built on the official Upjet v2 provider template with modern best practices
- **Multi-arch support**: ARM64 and AMD64 container images

### Supported Resource Groups

| Category | Resources | Description |
|----------|-----------|-------------|
| **Agent** | 2 | Agent pools and queues |
| **Build** | 6 | Build definitions, folders, permissions, and resource authorization |
| **Git** | 5 | Git repositories and permissions |
| **Branch Policy** | 8 | Branch policies (build validation, merge types, reviewers, etc.) |
| **Repository Policy** | 7 | Repository-level policies (case enforcement, file paths, size limits) |
| **Service Endpoints** | 28 | Service connections (Azure RM, AWS, GitHub, Kubernetes, Docker, etc.) |
| **Check** | 2 | Approval and check configurations |
| **Group** | 4 | Group management and membership |
| **Team** | 3 | Team management and administrators |
| **Feed** | 2 | Azure Artifacts feeds and permissions |
| **Environment** | 2 | Deployment environments and resource authorization |
| **Variable Group** | 1 | Pipeline variable groups |
| **Wiki** | 1 | Project wikis |
| **Work Item** | 1 | Work item tracking |
| **Permissions** | 4 | Project permissions and security |
| **Other** | 16 | Projects, users, pipeline settings, and more |

**Total: 92 managed resources** providing complete Azure DevOps infrastructure as code capability.

## Quick Start

### Prerequisites

- Kubernetes cluster (1.24+)
- [Crossplane installed](https://crossplane.io/docs/latest/software/install/) (v1.14+)
- Azure DevOps organization with admin access
- [Personal Access Token (PAT)](https://docs.microsoft.com/en-us/azure/devops/organizations/accounts/use-personal-access-tokens-to-authenticate) with appropriate scopes

### Installation

1. **Install the Provider:**

```bash
cat <<EOF | kubectl apply -f -
apiVersion: pkg.crossplane.io/v1
kind: Provider
metadata:
  name: provider-azuredevops
spec:
  package: ghcr.io/gontzalson/provider-azuredevops:latest
EOF
```

2. **Wait for provider to become healthy:**

```bash
kubectl wait --for=condition=Healthy provider/provider-azuredevops --timeout=300s
```

3. **Configure Provider Authentication:**

Create a Kubernetes secret with your Azure DevOps PAT:

```bash
kubectl create secret generic azuredevops-credentials \
  --from-literal=credentials='{"pat":"YOUR_PAT_HERE","org_service_url":"https://dev.azure.com/YOUR_ORG"}' \
  -n crossplane-system
```

4. **Create a ProviderConfig:**

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

## Usage Examples

### Create a Complete Project Setup

```yaml
---
# Project
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
  providerConfigRef:
    name: default
---
# Git Repository
apiVersion: git.azuredevops.crossplane.io/v1alpha1
kind: Repository
metadata:
  name: my-repo
spec:
  forProvider:
    projectIdRef:
      name: my-project
    name: my-repository
    initialization:
      initType: Clean
  providerConfigRef:
    name: default
---
# Branch Policy - Build Validation
apiVersion: branchpolicy.azuredevops.crossplane.io/v1alpha1
kind: BuildValidation
metadata:
  name: build-validation
spec:
  forProvider:
    projectIdRef:
      name: my-project
    enabled: true
    blocking: true
    settings:
      displayName: CI Build Validation
      buildDefinitionId: 123
      queueOnSourceUpdateOnly: true
      validDuration: 720
      scope:
        - repositoryIdRef:
            name: my-repo
          refName: refs/heads/main
          matchType: Exact
  providerConfigRef:
    name: default
---
# Service Endpoint - Azure RM
apiVersion: serviceendpoint.azuredevops.crossplane.io/v1alpha1
kind: AzureRM
metadata:
  name: azure-connection
spec:
  forProvider:
    projectIdRef:
      name: my-project
    servicePrincipalId: "00000000-0000-0000-0000-000000000000"
    servicePrincipalKey: "your-sp-key"
    tenantId: "00000000-0000-0000-0000-000000000000"
    subscriptionId: "00000000-0000-0000-0000-000000000000"
    subscriptionName: "My Subscription"
  providerConfigRef:
    name: default
```

### Build Definition with Variable Group

```yaml
---
# Variable Group
apiVersion: variablegroup.azuredevops.crossplane.io/v1alpha1
kind: VariableGroup
metadata:
  name: shared-vars
spec:
  forProvider:
    projectIdRef:
      name: my-project
    name: Shared Variables
    description: Common variables
    allowAccess: true
    variable:
      - name: environment
        value: production
      - name: region
        value: westeurope
  providerConfigRef:
    name: default
---
# Build Definition
apiVersion: build.azuredevops.crossplane.io/v1alpha1
kind: Definition
metadata:
  name: ci-pipeline
spec:
  forProvider:
    projectIdRef:
      name: my-project
    name: CI Pipeline
    path: "\\CI"
    repository:
      repoType: TfsGit
      repoIdRef:
        name: my-repo
      branchName: refs/heads/main
      ymlPath: azure-pipelines.yml
    variableGroups:
      - variableGroupIdRef:
          name: shared-vars
  providerConfigRef:
    name: default
```

## Available Resources

For a complete list of all 92 managed resources and their schemas, see:
- [API Reference Documentation](https://doc.crds.dev/github.com/gontzalson/provider-azuredevops) (coming soon)
- [Generated CRDs](./package/crds/)
- [Terraform Provider Documentation](https://registry.terraform.io/providers/microsoft/azuredevops/latest/docs)

Common resource types include:
- Projects, teams, and groups
- Git repositories with branch policies
- Build definitions and pipelines
- Service endpoints (28 types!)
- Environments and checks
- Artifact feeds and permissions
- Work items and wikis
- Security and permissions

## Development

### Building from Source

```bash
# Clone the repository
git clone https://github.com/gontzalson/provider-azuredevops
cd provider-azuredevops

# Generate CRDs and controllers
make generate

# Build the provider binary
make build

# Build and push multi-arch container images
make build.all
```

### Project Structure

```
provider-azuredevops/
├── apis/                      # Generated API types
├── config/                    # Upjet configuration
│   ├── agent/                # Agent pool resources
│   ├── build/                # Build definition resources
│   ├── git/                  # Git repository resources
│   ├── branchpolicy/         # Branch policy resources
│   ├── serviceendpoint/      # Service endpoint resources (28!)
│   ├── ...                   # 10 more resource groups
│   ├── external_name.go      # External name configuration (92 resources)
│   └── provider.go           # Main provider configuration
├── internal/controller/      # Generated controllers
├── package/crds/            # Generated CRDs (175 files: cluster + namespaced)
└── examples/                # Usage examples
```

### Adding New Resources

Resources from the upstream Terraform provider are added in `config/`:

1. Add external name mapping in `config/external_name.go`
2. Create or update resource group config in `config/<group>/config.go`
3. Register in `config/provider.go`
4. Run `make generate` to generate CRDs and controllers

## Architecture

This provider uses the Upjet v2 framework to automatically generate Crossplane Custom Resources from Terraform provider schemas:

```
Kubernetes API
      ↓
Crossplane Provider (this)
      ↓
Upjet v2 Framework
      ↓
Terraform Azure DevOps Provider (v1.4.0)
      ↓
Azure DevOps REST API
```

**Key Components:**
- **175 CRDs**: Auto-generated from Terraform provider schema (92 resources × 2 scopes)
- **Controllers**: Reconcile Kubernetes resources with Azure DevOps state
- **Terraform Workspace**: Manages Terraform state for each managed resource
- **IdentifierFromProvider**: External name strategy for all resources

## Testing

The provider has been tested in production-like environments:

- ✅ Crossplane v2.1.3 compatibility
- ✅ Kind cluster validation
- ✅ Multi-arch container images (ARM64, AMD64)
- ✅ All 92 resource types registered and available
- ✅ GitHub Container Registry publishing

## Contributing

Contributions are welcome! Please:

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/amazing-feature`)
3. Make your changes
4. Commit your changes (`git commit -m 'Add amazing feature'`)
5. Push to the branch (`git push origin feature/amazing-feature`)
6. Open a Pull Request

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

*Providing complete Azure DevOps infrastructure as code with 92 managed resources*
