# Snyk Security Remediation Plan



## Overview

This document outlines the prioritized remediation plan for all security vulnerabilities identified by Snyk scans. The plan is organized by severity and estimated time to completion.

---

## Priority Matrix

| Priority | Severity | Count | ETA |
|----------|----------|-------|-----|
| P0 | Critical | 1 | 2 hours |
| P1 | High | 2 | 4 hours |
| P2 | Medium | 5 | 3 hours |
| **Total** | | **8** | **9 hours** |

---

## Phase 1: Critical Issues (P0) - Must Fix Immediately

### 🔴 CRITICAL-1: form-data Predictable Boundary Values

**Vulnerability:** SNYK-JS-FORMDATA-10841150  
**Package:** form-data@2.3.3  
**Severity:** Critical (CVSS 9.4)  
**EPSS:** 0.053%

#### Issue Summary
HTTP request boundaries use predictable `Math.random()` values, allowing parameter pollution attacks.

#### Remediation Steps

1. **Update superagent** (which brings updated form-data)
   ```bash
   cd react-redux-realworld-example-app
   npm install superagent@10.2.2
   ```

2. **Verify the fix**
   ```bash
   snyk test
   ```

3. **Test Impact Areas**
   - User registration
   - Login functionality
   - Article creation with images
   - Profile updates
   - Any file upload features

#### Breaking Changes

**superagent v3 → v10 Changes:**
- `.end()` callback syntax changed
- Promise-based API now default
- Error handling structure updated

**Migration Example:**

**Before (v3):**
```javascript
superagent
  .post('/api/users')
  .send(data)
  .end((err, res) => {
    if (err) return handleError(err);
    handleSuccess(res.body);
  });
```

**After (v10):**
```javascript
try {
  const res = await superagent
    .post('/api/users')
    .send(data);
  handleSuccess(res.body);
} catch (err) {
  handleError(err);
}
```

#### Estimated Time
- **Update:** 30 minutes
- **Testing:** 1 hour
- **Code Migration:** 30 minutes
- **Total:** 2 hours

#### Risk Assessment
- **Risk Level:** High
- **Breaking Change:** Yes
- **Rollback Plan:** Revert package-lock.json and run `npm install`

---

## Phase 2: High Priority Issues (P1) - Fix Within 24 Hours

### 🟠 HIGH-1: JWT Access Restriction Bypass

**Vulnerability:** SNYK-GOLANG-GITHUBCOMDGRIJALVAJWTGO-596515  
**Package:** github.com/dgrijalva/jwt-go@3.2.0  
**Severity:** High (CVSS 7.5)  
**CVE:** CVE-2020-26160

#### Issue Summary
Empty audience array bypasses JWT audience verification, allowing unauthorized access.

#### Remediation Options

**Option A (Recommended): Migrate to golang-jwt**
```bash
cd golang-gin-realworld-example-app
go get github.com/golang-jwt/jwt/v5
```

Update imports in all files:
```go
// Old import
import "github.com/dgrijalva/jwt-go"

// New import
import jwt "github.com/golang-jwt/jwt/v5"
```

**Option B: Upgrade to v4**
```bash
go get github.com/dgrijalva/jwt-go@v4.0.0-preview1
```

#### Files to Update
- `common/jwt.go` - Token generation/validation
- `users/middlewares.go` - Auth middleware
- Any files importing jwt-go

#### Breaking Changes

**API Changes in v4/v5:**
- `ParseWithClaims` signature changed
- `StandardClaims` renamed to `RegisteredClaims`
- Error handling improvements

**Migration Example:**

**Before:**
```go
token, err := jwt.ParseWithClaims(tokenString, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
    return []byte(secret), nil
})
```

**After (v5):**
```go
token, err := jwt.ParseWithClaims(tokenString, &jwt.RegisteredClaims{}, func(token *jwt.Token) (interface{}, error) {
    return []byte(secret), nil
})
```

#### Estimated Time
- **Update Package:** 15 minutes
- **Code Migration:** 1 hour
- **Testing:** 1 hour
- **Total:** 2 hours 15 minutes

#### Testing Checklist
- [ ] User registration creates valid tokens
- [ ] Login returns valid tokens
- [ ] Protected routes validate tokens correctly
- [ ] Token expiration works
- [ ] Invalid tokens are rejected

---

### 🟠 HIGH-2: SQLite3 Heap Buffer Overflow

**Vulnerability:** SNYK-GOLANG-GITHUBCOMMATTNGOSQLITE3-6139875  
**Package:** github.com/mattn/go-sqlite3@1.14.15  
**Severity:** High (CVSS 7.3)  
**CVE:** CVE-2023-7104

#### Issue Summary
Heap-based buffer overflow in sessionReadRecord function can cause crashes or arbitrary code execution.

#### Remediation Steps

**Option A: Force Direct Upgrade**
```bash
cd golang-gin-realworld-example-app
go get github.com/mattn/go-sqlite3@v1.14.18
```

**Option B: Update GORM (recommended for long-term)**
```bash
# This will also update sqlite3
go get -u gorm.io/gorm
go get -u gorm.io/driver/sqlite
```

#### Code Changes (if migrating to GORM v2)

**Current (GORM v1):**
```go
import "github.com/jinzhu/gorm"
import _ "github.com/jinzhu/gorm/dialects/sqlite"

db, err := gorm.Open("sqlite3", "./gorm.db")
```

**New (GORM v2):**
```go
import "gorm.io/gorm"
import "gorm.io/driver/sqlite"

db, err := gorm.Open(sqlite.Open("./gorm.db"), &gorm.Config{})
```

#### Estimated Time
- **Option A (Force Update):** 15 minutes + 30 minutes testing = 45 minutes
- **Option B (GORM v2):** 1 hour migration + 1 hour testing = 2 hours
- **Recommended:** Option A for now, Option B as future improvement

#### Testing Checklist
- [ ] Database connection works
- [ ] User CRUD operations
- [ ] Article CRUD operations
- [ ] Comments functionality
- [ ] Database migrations work
- [ ] No data loss

---

## Phase 3: Medium Priority Issues (P2) - Fix Within 1 Week

### 🟡 MEDIUM-1 to MEDIUM-5: marked ReDoS Vulnerabilities

**Package:** marked@0.3.19  
**Target Version:** marked@4.0.10  
**Combined CVSS:** 5.3 - 5.9  

#### Issue Summary
Five separate Regular Expression Denial of Service vulnerabilities in different regex patterns.

#### Single Remediation for All Five

```bash
cd react-redux-realworld-example-app
npm install marked@4.0.10
```

#### Breaking Changes

**marked v0.3 → v4.0 Major Changes:**

1. **Renderer API Changes**
2. **Extension System Introduced**
3. **Options Structure Updated**

**Migration Example:**

**Before (v0.3):**
```javascript
import marked from 'marked';

marked.setOptions({
  gfm: true,
  breaks: true,
  sanitize: true
});

const html = marked(markdown);
```

**After (v4.0):**
```javascript
import { marked } from 'marked';

marked.setOptions({
  gfm: true,
  breaks: true
});

// Note: sanitize option removed - use DOMPurify instead
import DOMPurify from 'dompurify';
const html = DOMPurify.sanitize(marked.parse(markdown));
```

#### Additional Security: Add DOMPurify

```bash
npm install dompurify
npm install --save-dev @types/dompurify  # if using TypeScript
```

#### Files to Update
- `src/components/Article/index.js` - Article rendering
- `src/components/ArticlePreview.js` - Preview rendering
- Any component using markdown

#### Estimated Time
- **Update Package:** 10 minutes
- **Code Migration:** 1 hour
- **Add DOMPurify:** 30 minutes
- **Testing:** 1 hour
- **Total:** 2 hours 40 minutes

#### Testing Checklist
- [ ] Article content displays correctly
- [ ] Markdown formatting works (bold, italic, lists)
- [ ] Code blocks render properly
- [ ] Links work correctly
- [ ] Images display
- [ ] No XSS vulnerabilities

---

## Workaround Strategies (If Upgrade Not Possible)

### Backend Workarounds

#### JWT Workaround
```go
// Add explicit audience validation
func validateAudience(token *jwt.Token) error {
    claims := token.Claims.(jwt.MapClaims)
    aud, ok := claims["aud"]
    if !ok || aud == "" {
        return errors.New("invalid audience")
    }
    // Additional validation
    return nil
}
```

#### SQLite Workaround
- Use PostgreSQL or MySQL instead of SQLite
- Limit user input to database
- Add input validation layers

### Frontend Workarounds

#### form-data Workaround
- Implement custom boundary generation
- Add request validation middleware

#### marked Workaround
```javascript
// Add timeout wrapper
function safeMarkdown(md) {
  return Promise.race([
    marked.parse(md),
    new Promise((_, reject) => 
      setTimeout(() => reject(new Error('Timeout')), 5000)
    )
  ]);
}
```

---

## Testing Strategy

### Pre-Upgrade Testing
1. Create backup of current state
2. Run existing test suite
3. Document current behavior

### Post-Upgrade Testing

#### Backend Tests
```bash
cd golang-gin-realworld-example-app
go test ./...
```

Manual tests:
- [ ] User registration
- [ ] User login
- [ ] Create article
- [ ] Update article
- [ ] Delete article
- [ ] Follow/unfollow users
- [ ] Favorite articles

#### Frontend Tests
```bash
cd react-redux-realworld-example-app
npm test
```

Manual tests:
- [ ] Register new user
- [ ] Login
- [ ] View home feed
- [ ] Create article with markdown
- [ ] Edit article
- [ ] View user profile
- [ ] Settings page

### Integration Testing
- [ ] End-to-end user flows
- [ ] API contract compatibility
- [ ] Error handling scenarios

---

## Rollback Plan

### Backend Rollback
```bash
cd golang-gin-realworld-example-app
git checkout go.mod go.sum
go mod download
go build
```

### Frontend Rollback
```bash
cd react-redux-realworld-example-app
git checkout package.json package-lock.json
npm install
```

---

## Implementation Timeline

### Week 1 (Current)
- **Day 1:** Fix CRITICAL-1 (form-data) - 2 hours
- **Day 1-2:** Fix HIGH-1 (JWT) - 2 hours
- **Day 2:** Fix HIGH-2 (SQLite) - 1 hour
- **Day 3-4:** Fix MEDIUM (marked) - 3 hours
- **Day 5:** Complete testing and documentation

### Total Estimated Time: 9 hours spread over 5 days

---

## Success Criteria

- [ ] All critical vulnerabilities resolved
- [ ] All high vulnerabilities resolved
- [ ] All medium vulnerabilities resolved
- [ ] Zero regression in functionality
- [ ] All tests passing
- [ ] Documentation updated
- [ ] Snyk scan shows 0 high/critical issues

---

## Post-Remediation Actions

### Continuous Monitoring
1. Enable Snyk monitoring
   ```bash
   snyk monitor
   ```

2. Set up GitHub integration for automatic scans

3. Configure Snyk alerts

### Prevention
1. Add pre-commit hooks for security checks
2. Implement dependency update policy
3. Schedule monthly security reviews
4. Add security testing to CI/CD pipeline

---

## Sign-off

| Role | Name | Date | Signature |
|------|------|------|-----------|
| Developer | | | |
| Security Review | | | |
| Project Manager | | | |

---

## Next Steps

Proceed to implementation phase. Track progress in `snyk-fixes-applied.md`.