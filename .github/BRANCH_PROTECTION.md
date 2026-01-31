# Branch Protection Configuration

This document describes the recommended branch protection rules for the `main` branch.
These settings should be configured in GitHub repository settings under **Settings > Branches > Branch protection rules**.

## Recommended Settings for `main` Branch

### Protect matching branches

- [x] **Require a pull request before merging**
  - [x] Require approvals: **1** (minimum)
  - [x] Dismiss stale pull request approvals when new commits are pushed
  - [x] Require review from Code Owners
  - [ ] Restrict who can dismiss pull request reviews (optional)
  - [x] Require approval of the most recent reviewable push

### Required Status Checks

Enable **Require status checks to pass before merging** and add the following checks:

#### CI Checks (from `ci.yml`)
| Check Name | Required | Description |
|------------|----------|-------------|
| `lint` | ✅ Yes | Go code linting with golangci-lint |
| `check-diff` | ✅ Yes | Verify code formatting (goimports) |
| `unit-tests` | ✅ Yes | Unit test execution and coverage |
| `container-build` | ✅ Yes | Container image build validation |
| `local-deploy` | ✅ Yes | Kubernetes local deployment test |
| `report-breaking-changes` | ✅ Yes | CRD breaking change detection |

#### Security Checks (from `security.yml`)
| Check Name | Required | Description |
|------------|----------|-------------|
| `codeql` | ✅ Yes | CodeQL static analysis |
| `trivy-fs` | ✅ Yes | Filesystem vulnerability scan |
| `trivy-config` | ✅ Yes | Configuration security scan |
| `dependency-review` | ✅ Yes | Dependency license and vulnerability review |
| `secrets-scan` | ✅ Yes | Secret detection scan |

#### PR Validation Checks (from `pr-validation.yml`)
| Check Name | Required | Description |
|------------|----------|-------------|
| `pr-title` | ✅ Yes | Conventional commit format validation |
| `dco` | ✅ Yes | Developer Certificate of Origin check |
| `branch-naming` | ⚠️ Optional | Branch naming convention validation |

### Additional Settings

- [x] **Require status checks to be up-to-date before merging**
  - Ensures the branch is tested with the latest `main` branch changes

- [x] **Require conversation resolution before merging**
  - All review comments must be resolved before merge

- [x] **Require signed commits** (recommended)
  - Ensures all commits are cryptographically signed

- [x] **Require linear history**
  - Prevents merge commits, requires squash or rebase merges

- [x] **Do not allow bypassing the above settings**
  - Even administrators must follow the rules

- [x] **Restrict who can push to matching branches**
  - Only allow maintainers listed in CODEOWNERS

### Merge Settings

Configure under **Settings > General > Pull Requests**:

- [x] **Allow squash merging** (recommended as default)
  - Default commit message: Pull request title and description
- [ ] **Allow merge commits** (disable for cleaner history)
- [x] **Allow rebase merging** (optional)
- [x] **Automatically delete head branches**

## Setting Up Branch Protection via GitHub CLI

You can configure branch protection using the GitHub CLI:

```bash
# Set up branch protection for main
gh api repos/{owner}/{repo}/branches/main/protection \
  --method PUT \
  --field required_status_checks='{"strict":true,"checks":[{"context":"lint"},{"context":"check-diff"},{"context":"unit-tests"},{"context":"container-build"},{"context":"local-deploy"},{"context":"codeql"},{"context":"trivy-fs"},{"context":"dependency-review"},{"context":"pr-title"},{"context":"dco"}]}' \
  --field enforce_admins=true \
  --field required_pull_request_reviews='{"required_approving_review_count":1,"dismiss_stale_reviews":true,"require_code_owner_reviews":true}' \
  --field restrictions=null \
  --field required_linear_history=true \
  --field allow_force_pushes=false \
  --field allow_deletions=false
```

## Rulesets (Modern Alternative)

GitHub Rulesets provide a more flexible alternative to branch protection. Create a ruleset:

1. Go to **Settings > Rules > Rulesets**
2. Click **New ruleset > New branch ruleset**
3. Configure:
   - Name: `main-protection`
   - Enforcement: Active
   - Target branches: `main`, `release-*`
   - Apply the same rules as documented above

## Verifying Configuration

After setting up branch protection, verify by:

1. Creating a test PR
2. Confirming all required checks run
3. Verifying CODEOWNERS review is required
4. Attempting to push directly to `main` (should be blocked)

## Additional Security Recommendations

1. **Enable GitHub Advanced Security** for enhanced CodeQL and secret scanning
2. **Enable Dependabot security updates** (already configured via Renovate)
3. **Require two-factor authentication** for all organization members
4. **Enable audit logging** for security events
5. **Set up security advisories** for responsible disclosure
