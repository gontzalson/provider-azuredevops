<!--
Thank you for contributing to provider-azuredevops!

Please read through https://git.io/fj2m9 if this is your first time opening a
Crossplane provider pull request.
-->

## Description

<!--
Briefly describe what this pull request does. Be sure to direct your reviewers'
attention to anything that needs special consideration.

We love pull requests that resolve an open issue. If yours does, you
can uncomment the below line to indicate which issue your PR fixes:
-->

Fixes #

## Type of Change

<!-- Mark the relevant option with an [x] -->

- [ ] Bug fix (non-breaking change that fixes an issue)
- [ ] New feature (non-breaking change that adds functionality)
- [ ] Breaking change (fix or feature that would cause existing functionality to change)
- [ ] Documentation update
- [ ] Refactoring (no functional changes)
- [ ] CI/CD changes
- [ ] Dependency update

## Checklist

<!-- Mark completed items with an [x] -->

### Code Quality
- [ ] My code follows the project's coding standards
- [ ] I have run `make reviewable` to ensure code is properly formatted
- [ ] I have run `make lint` and fixed any linting issues
- [ ] I have added/updated comments where necessary

### Testing
- [ ] I have run `make test` and all tests pass
- [ ] I have added tests that prove my fix/feature works
- [ ] New and existing unit tests pass locally with my changes

### Documentation
- [ ] I have updated the documentation accordingly (if applicable)
- [ ] I have updated the examples (if applicable)

### Commit Standards
- [ ] My commits are signed off (DCO): `git commit -s`
- [ ] My PR title follows [Conventional Commits](https://www.conventionalcommits.org/)
  - Examples: `feat: add new resource`, `fix(api): resolve nil pointer`, `docs: update README`

## How Has This Been Tested?

<!--
Describe the tests you ran to verify your changes. Provide instructions so
reviewers can reproduce. Include any relevant details about your test configuration.
-->

- [ ] Unit tests
- [ ] Local deployment (`make local-deploy`)
- [ ] End-to-end tests (`/test-examples` comment)
- [ ] Manual testing

**Test Configuration:**
- Go version:
- Kubernetes version:

## Additional Notes

<!--
Add any additional notes, screenshots, or context about the PR here.
-->

---

[contribution process]: https://git.io/fj2m9
