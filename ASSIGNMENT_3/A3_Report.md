 Assignment 3 Report: Performance & E2E Testing

## Part A: Performance Testing (Summary)
- Performance baseline established using k6.
- Load, stress, spike, and soak tests conducted.
- Bottlenecks identified in database queries and article creation endpoints.
- Optimizations implemented: database indexing, eager loading, query optimization.
- Post-optimization results: 20-30% faster response times, lower CPU/memory usage.

## Part B: Cypress E2E Testing
- Authentication flows tested: registration, login, logout, validation errors.
- Article management tested: create, read, update, delete, tag management.
- Comments functionality verified.
- User profiles and feeds verified.
- Complete user workflows executed successfully.
- Cross-browser testing conducted across Chrome, Firefox, Edge, Electron.

## Key Learnings
- E2E testing helps catch real-world user issues.
- Using Cypress fixtures and commands simplifies test code.
- Performance and functional testing together provide a robust QA coverage.

## Evidence
- Videos and screenshots from Cypress.
- Cross-browser testing report.
- k6 performance graphs and terminal outputs.
