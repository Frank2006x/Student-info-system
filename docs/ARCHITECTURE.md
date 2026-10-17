# Student Information System Architecture

## 1. Target Topology

```text
User browser
    |
    v
Vercel: Next.js frontend
    |
    | HTTPS API requests
    v
EC2: Go API service
    |\
    | \-- write connection: primary PostgreSQL/RDS
    |
    \---- read connection: RDS read replica
             |
             \-- reporting, dashboards, exports, safe list reads
```

The browser communicates with the Go API only. It never connects directly to PostgreSQL or RDS.

## 2. Current State

### Frontend

- Next.js 16 with React 19 and TypeScript.
- App Router structure in `frontend/src/app`.
- Tailwind CSS 4 is available.
- The current page is still the default starter page.
- Deployment target: Vercel.

### Backend

- Go module in `backend`.
- Local PostgreSQL 18 Compose service.
- `migrate` migration files under `backend/db/migration`.
- `sqlc` configuration under `backend/sqlc.yaml`.
- API implementation is not yet complete.
- Deployment target: EC2.

### Database

The current migration creates:

- `departments`
- `students`
- `faculty`
- `courses`
- `enrollments`
- `attendance`

The migration must be tested from an empty database before every production release path.

## 3. Deployment Environments

### Local

- Frontend: Next.js development server.
- Backend: Go process or Docker.
- Database: PostgreSQL via `docker compose`.
- Migrations: local `migrate` CLI.
- Generated queries: local `sqlc` CLI.

### Preview

- Frontend: Vercel preview deployment.
- Backend: shared development or staging API endpoint.
- Database: isolated staging PostgreSQL database.
- Never allow a preview deployment to use production write credentials.

### Production

- Frontend: Vercel production deployment.
- Backend: EC2 service behind HTTPS and a stable API hostname.
- Database: RDS PostgreSQL primary.
- Optional RDS read replica for reporting and read-heavy operations.
- Secrets: AWS Secrets Manager or SSM Parameter Store, not committed files.

## 4. Vercel Frontend Configuration

Recommended environment variables:

```text
NEXT_PUBLIC_API_BASE_URL=https://api.example.com
```

Only values intended for browser use may use the `NEXT_PUBLIC_` prefix. Do not expose database URLs, signing secrets, private keys, or administrative credentials.

Frontend deployment requirements:

- Configure production and preview environment variables separately.
- Use the API hostname, not an EC2 private address.
- Configure backend CORS for the exact Vercel production and preview origins required.
- Use HTTPS in every non-local environment.
- Add a deployment check that runs lint and the production build.

## 5. EC2 Backend Deployment

### Initial approach

1. Provision an EC2 instance in a private or restricted network position.
2. Install Docker or build the Go binary in CI.
3. Run the API under systemd or Docker with automatic restart.
4. Put Nginx, an AWS Application Load Balancer, or another HTTPS reverse proxy in front of the API.
5. Expose only HTTP/HTTPS and restricted administration ports in the security group.
6. Allow database access only from the backend security group.
7. Send application and system logs to CloudWatch.

### Backend process requirements

- Read configuration from environment variables or AWS secret storage.
- Listen on a configurable port.
- Provide `/health/live` for process liveness.
- Provide `/health/ready` for database and dependency readiness.
- Shut down gracefully and drain connections.
- Limit request body size and request duration.
- Use a bounded database connection pool.

### Minimum production configuration

```text
APP_ENV=production
HTTP_ADDR=:8080
DATABASE_PRIMARY_URL=<secret>
DATABASE_READ_URL=<secret-or-primary-during-phase-one>
CORS_ALLOWED_ORIGINS=https://app.example.com
LOG_LEVEL=info
```

Do not put real values in committed documentation or `.env` files.

## 6. Database Connection Strategy

Use two logical database handles in the backend:

- `primaryDB`: all inserts, updates, deletes, transactions, and reads requiring read-after-write consistency.
- `readDB`: reporting, dashboards, exports, and safe read-only list queries.

During the first deployment, both handles may point to the primary database. When an RDS read replica is available, change only the read connection configuration and repository routing.

### Primary reads

Use the primary for:

- Login and account state.
- Immediately refreshed student or enrollment details after a mutation.
- Writes and transactions.
- Any workflow where stale data could cause a duplicate or invalid decision.

### Replica reads

Use the read replica for:

- Aggregated reports.
- Historical dashboards.
- Large exports.
- Search/list views that tolerate small replication delay.
- Analytics queries that should not compete with transactional writes.

The API should make the consistency choice in code. A frontend query must not choose a database directly.

## 7. API Contract Direction

Use versioned endpoints under `/api/v1`.

Example resources:

```text
POST   /api/v1/auth/login
GET    /api/v1/me
GET    /api/v1/departments
POST   /api/v1/departments
GET    /api/v1/students
POST   /api/v1/students
GET    /api/v1/students/{id}
PATCH  /api/v1/students/{id}
GET    /api/v1/faculty
GET    /api/v1/courses
POST   /api/v1/courses
GET    /api/v1/enrollments
POST   /api/v1/enrollments
PATCH  /api/v1/enrollments/{id}
GET    /api/v1/attendance
POST   /api/v1/attendance
PATCH  /api/v1/attendance/{id}
GET    /api/v1/reports/enrollment-summary
GET    /api/v1/reports/attendance-summary
```

List responses should have a consistent shape:

```json
{
  "data": [],
  "pagination": {
    "page": 1,
    "page_size": 25,
    "total": 0
  }
}
```

Errors should have a stable shape with a machine-readable code and user-safe message:

```json
{
  "error": {
    "code": "validation_failed",
    "message": "The request contains invalid fields.",
    "fields": {
      "email": "must be a valid email address"
    }
  }
}
```

## 8. Backend Package Direction

A practical Go structure is:

```text
backend/
  cmd/api/main.go
  internal/
    auth/
    config/
    controller/
    middleware/
    service/
    repository/
    router/
  db/
    migration/
    query/
    sqlc/
```

Responsibilities:

- `controller`: HTTP parsing, response formatting, and status codes.
- `service`: administrator-only business rules and transactions.
- `repository`: database access interfaces and read/write routing.
- `middleware`: authentication, request ID, logging, and recovery.
- `db/sqlc`: generated query code; do not edit generated files manually.
- `db/migration`: versioned schema changes.

## 9. Security Baseline

- Use TLS for browser-to-API and API-to-managed services where applicable.
- Store secrets outside Git.
- Use least-privilege database credentials.
- Restrict RDS inbound access to the backend security group.
- Restrict EC2 administrative access by IP, VPN, or a managed access path.
- Hash passwords with a modern password hashing algorithm if accounts are application-managed.
- Require an authenticated administrator for every resource mutation and sensitive read.
- Avoid logging passwords, tokens, student personal data, or full request bodies.
- Rate-limit login and other abuse-sensitive endpoints.
- Audit administrative changes and attendance corrections.
- Encrypt RDS storage and backups.
- Define backup retention and test restore procedures.

## 10. Observability and Operations

Track at minimum:

- API request count, latency, status code, and route.
- Database query latency and connection pool usage.
- Primary and replica health.
- RDS replica lag.
- EC2 CPU, memory, disk, and process restarts.
- Migration version at deployment time.
- Authentication failures and rejected administrator requests.
- Report and export duration.

Alerts should cover API unavailability, high error rate, database connection exhaustion, storage pressure, failed backups, and excessive replica lag.

## 11. Migration and Release Safety

1. Build and test the API and frontend.
2. Validate migrations against a fresh PostgreSQL database.
3. Validate migrations against a sanitized staging database.
4. Back up production before destructive changes.
5. Apply backward-compatible migrations before deploying code that depends on them.
6. Deploy the backend.
7. Verify readiness, migration version, and smoke-test endpoints.
8. Deploy the frontend.
9. Monitor errors and rollback using the documented release procedure.

Avoid using `migrate force` as a normal deployment step. Use it only to repair a known migration tracking state after checking the database manually.

## 12. AWS Read-Scalability Sequence

### Stage A: EC2 plus one PostgreSQL primary

- Finish the API and deploy it to EC2.
- Use one primary connection for all reads and writes.
- Establish metrics and query indexes.
- Measure actual read/write workload.

### Stage B: RDS primary

- Provision RDS PostgreSQL with private networking.
- Import or recreate schema through migrations.
- Validate backups, restore, security groups, and connection limits.
- Point the backend primary connection to RDS.

### Stage C: RDS read replica

- Create a replica from the RDS primary.
- Configure a read-only database credential.
- Route only explicitly safe read workloads to it.
- Monitor replica lag and fail back to primary when lag exceeds the product tolerance.

### Stage D: Reporting isolation

- Move heavy reports to dedicated read queries.
- Add report-specific indexes or pre-aggregations based on measured plans.
- Add export job limits and asynchronous processing if reports become long-running.

## 13. Recommended First Technical Slice

The highest-value next slice is:

1. Implement the Go server and `/health/live` endpoint.
2. Configure sqlc queries for departments and students.
3. Add a simple authenticated API boundary.
4. Replace the Next.js starter page with a dashboard shell.
5. Connect one paginated student list to the API.
6. Run the same flow locally before adding AWS resources.

This gives the deployment plan a real vertical path while keeping AWS infrastructure replaceable during development.
