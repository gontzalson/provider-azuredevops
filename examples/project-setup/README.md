# Complete Project Setup Example

This example demonstrates setting up a complete Azure DevOps project with:

- Project creation
- Git repository with branch policies
- Build definition
- Service endpoint (Azure RM)
- Variable group
- Team and permissions

## Prerequisites

1. Install Crossplane in your Kubernetes cluster
2. Install provider-azuredevops
3. Configure authentication (see [../install.yaml](../install.yaml))

## Usage

```bash
# Apply all resources
kubectl apply -f .

# Watch resources being created
kubectl get managed

# Check project status
kubectl get project.project.azuredevops.crossplane.io

# Check all Azure DevOps resources
kubectl get managed -l example=project-setup
```

## Resources Created

- **Project**: "My Platform Project"
- **Git Repository**: "infrastructure"
- **Branch Policy**: Build validation on main branch
- **Service Endpoint**: Azure RM connection
- **Variable Group**: Shared variables
- **Build Definition**: CI pipeline
- **Team**: Platform Team
- **Group Membership**: Team assignments

## Cleanup

```bash
kubectl delete -f .
```

Note: Crossplane will remove the resources from Azure DevOps when you delete the Kubernetes objects.
