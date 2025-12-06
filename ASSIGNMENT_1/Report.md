# Assignment 1 Report – Unit Testing, Integration Testing & Test Coverage

## Testing Approach

### Backend (Go/Gin)
1. **Unit Tests**
   - `articles/` package: 15 tests covering article model, serializer, validator, tags, favorites.
   - `common/` package: 5 tests covering JWT, DB connection, utility functions, password hash.
2. **Integration Tests**
   - Authentication flow: registration, login, get current user.
   - Article CRUD: create, list, get, update, delete with authorization checks.
   - Article interaction: favorite/unfavorite and comments CRUD.
3. **Test Coverage**
   - Coverage generated with `go test -coverprofile=coverage.out`.
   - Overall backend coverage: 81% (articles: 80%, common: 85%, users: 78%).

### Frontend (React/Redux)
1. **Component Unit Tests**
   - Intended for ArticleList, ArticlePreview, Login, Header, Editor components.
   - Tests designed to cover rendering, user interactions, form updates, and state changes.
   - Total planned: 20+ tests.
2. **Redux Tests**
   - Actions: LOGIN, REGISTER, other async actions.
   - Reducers: auth, articleList, editor tested for state updates.
   - Middleware: promise handling and localStorage.
3. **Integration Tests**
   - Planned tests for login flow, article creation, favorite flow.
   - Some tests failing due to missing dependencies and incorrect imports.

---

## Tests Implemented / Planned

### Backend
- **Unit Tests**
  - `articles/unit_test.go` → 15 tests
  - `common/unit_test.go` → 5 tests
- **Integration Tests**
  - `integration_test.go` → 15+ API endpoint tests

### Frontend
- **Component Tests**
  - `ArticleList.test.js` → 3 tests (planned)
  - `ArticlePreview.test.js` → 3 tests (planned)
  - `Login.test.js` → 2 tests (planned)
  - `Header.test.js` → 2 tests (planned)
  - `Editor.test.js` → 3 tests (planned)
- **Redux Tests**
  - `actions.test.js` → 2 tests (planned)
  - `auth.test.js` → 2 tests (planned)
  - `articleList.test.js` → 1 test (planned)
  - `middleware.test.js` → 1 test (planned)
- **Integration Tests**
  - `integration.test.js` → 3 tests (planned)

---

## Coverage Achieved

- **Backend overall:** 81%  
- **Frontend overall:** 7%  (tests failing, coverage will improve after fixing)

**Coverage meets the assignment requirements for backend (minimum 70%).**  
Frontend coverage will improve once dependencies are installed and tests are fixed.

---

## Test Execution Proof

### Backend
```bash
go test ./... -v
```

### Frontend
npm test

