<!--
SYNC IMPACT REPORT
Version: 0.0.0 -> 1.0.0
Modified Principles:
- Added I. Code Quality
- Added II. Testing Standards
- Added III. User Experience Consistency
- Added IV. Performance Requirements
- Removed unused Principle 5 slot
Templates requiring updates:
- .specify/templates/plan-template.md (✅ Generic "Constitution Check" is compatible)
- .specify/templates/spec-template.md (✅ Requirements section compatible)
- .specify/templates/tasks-template.md (✅ Testing section compatible)
-->
# neodb Constitution

## Core Principles

### I. Code Quality
Adherence to strict language idioms is mandatory: idiomatic Go for backend, TypeScript/Angular style guide for frontend. Zero linting errors (ESLint, golangci-lint) and strict formatting (Prettier, gofmt) are enforced by CI. Code must be readable, maintainable, and self-documenting; prefer simplicity over cleverness.

### II. Testing Standards
Testing is non-negotiable. All business logic must be covered by unit tests (Vitest for frontend, `testing` package for Go). Critical user journeys require integration tests. Tests must be deterministic, fast, and written alongside production code (TDD encouraged). "It works on my machine" is not an acceptable validation.

### III. User Experience Consistency
The user interface must remain consistent across the entire Single Page Application (SPA). Reusable components and a unified design system must be used to ensure predictable interactions. Accessibility (a11y) standards must be met, and the application should be responsive and intuitive on supported devices.

### IV. Performance Requirements
Performance is a feature. Backend API response times should aim for <100ms (p95) for standard operations. Frontend bundle sizes must be optimized (lazy loading) to ensure fast First Contentful Paint (FCP). Database queries must be indexed and optimized to prevent scaling bottlenecks.

## Security & Compliance
All inputs must be validated and sanitized to prevent injection attacks (SQLi, XSS). Authentication and authorization checks are mandatory for all protected endpoints. Sensitive data must be encrypted at rest and in transit.

## Development Workflow
All changes require a Pull Request (PR) with passing CI checks. PRs must be reviewed by at least one other engineer. New features must start with a Design Spec (`/speckit.plan`) derived from this constitution before implementation begins.

## Governance
This Constitution is the supreme authority for technical decision-making in the neodb project. Amendments require a formalized proposal and team consensus. Non-compliant code will be rejected at the review or CI stage.

**Version**: 1.0.0 | **Ratified**: 2026-02-09 | **Last Amended**: 2026-02-09