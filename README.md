# Secret

This project is my way yo learn golang and get more familiar with k8s


## Roadmap

### Phase 1 — Project Setup
- [x] Initialize Go module
- [x] Create project folder structure
- [x] Setup configuration management
- [x] Add environment variable support
- [x] Setup basic HTTP server
- [x] Add graceful shutdown handling
- [x] Setup logging

---

### Phase 2 — Database Layer
- [ ] Integrate SQLite
- [ ] Create database connection manager
- [ ] Add migrations support
- [ ] Create `secrets` table
- [ ] Implement storage repository layer
- [ ] Add CRUD database operations

---

### Phase 3 — Encryption Layer
- [ ] Implement AES-GCM encryption
- [ ] Generate secure nonces
- [ ] Implement secret encryption
- [ ] Implement secret decryption
- [ ] Secure master key handling
- [ ] Add encryption unit tests

---

### Phase 4 — API Development
- [ ] Create REST API routes
- [ ] Implement `POST /secrets`
- [ ] Implement `GET /secrets/{name}`
- [ ] Implement `DELETE /secrets/{name}`
- [ ] Add request validation
- [ ] Add JSON response formatting
- [ ] Add error handling

---

### Phase 5 — Authentication
- [ ] Implement API token authentication
- [ ] Add auth middleware
- [ ] Protect all endpoints
- [ ] Add unauthorized access handling
- [ ] Add token configuration via environment variables

---

### Phase 6 — Service Layer Refactor
- [ ] Separate business logic from handlers
- [ ] Create service interfaces
- [ ] Improve dependency injection
- [ ] Decouple storage and crypto layers
- [ ] Improve package organization

---

### Phase 7 — Testing
- [ ] Add unit tests for crypto layer
- [ ] Add unit tests for storage layer
- [ ] Add API handler tests
- [ ] Add authentication tests
- [ ] Add integration tests
- [ ] Improve test coverage

---

### Phase 8 — Developer Experience
- [ ] Add Makefile
- [ ] Add Docker support
- [ ] Add Swagger/OpenAPI documentation
- [ ] Improve README documentation
- [ ] Add example curl commands
- [ ] Add local development guide

---

### Phase 9 — Security Improvements
- [ ] Add audit logging
- [ ] Add request rate limiting
- [ ] Add secret access logging
- [ ] Add secure key rotation support
- [ ] Add secret expiration (TTL)
- [ ] Add secret versioning

---

### Phase 10 — Kubernetes Preparation
- [ ] Abstract authentication provider
- [ ] Add JWT validation support
- [ ] Prepare Kubernetes Service Account authentication
- [ ] Add configuration for cluster environments
- [ ] Add container health checks

---

### Phase 11 — Kubernetes Integration
- [ ] Validate Kubernetes Service Account tokens
- [ ] Integrate with Kubernetes Secrets
- [ ] Add CSI driver compatibility
- [ ] Add sidecar integration support
- [ ] Add Kubernetes deployment manifests
- [ ] Add Helm chart support

---

### Phase 12 — Production Readiness
- [ ] Replace SQLite with PostgreSQL
- [ ] Add structured logging
- [ ] Add metrics and monitoring
- [ ] Add Prometheus integration
- [ ] Add TLS support
- [ ] Add CI/CD pipeline
- [ ] Add deployment automation

