# Student Information System Project Plan

## 1. Objective

Build a simple Student Information System with:

- A Next.js frontend deployed to Vercel.
- A Go Fiber v3 backend deployed to an AWS EC2 instance.
- PostgreSQL as the database.
- One administrator who manages all records.
- CRUD features for departments, students, faculty, and courses.

The first release should focus on reliable record management, clear validation, and a simple administrator workflow. Advanced reporting, multiple user roles, read replicas, and student or faculty logins are outside the initial scope.

## 2. Administrator Workflow

The administrator can:

- Sign in and sign out.
- Create, view, edit, and delete departments.
- Create, view, edit, and delete students.
- Create, view, edit, and delete faculty.
- Create, view, edit, and delete courses.
- Search and filter records.
- View relationships between departments, students, faculty, and courses.

All protected API endpoints require administrator authentication. There are no separate registrar, faculty, or student accounts in the first release.

## 3. Core CRUD Features

### Departments

Fields:

- Department code
- Department name
- Created and updated timestamps

Features:

- List departments
- View one department
- Create a department
- Edit a department
- Delete a department when it has no dependent records
- Validate unique department code and name

### Students

Fields:

- Student code
- First name
- Last name
- Email
- Department
- Enrollment date
- Status

Features:

- List students with pagination
- Search by student code, name, or email
- Filter by department and status
- View one student
- Create a student
- Edit a student
- Delete a student
- Validate unique student code and email
- Validate that the selected department exists

### Faculty

Fields:

- Faculty code
- First name
- Last name
- Email
- Department
- Designation
- Specialization
- Hire date
- Status

Features:

- List faculty with pagination
- Search by faculty code, name, or email
- Filter by department and status
- View one faculty member
- Create a faculty member
- Edit a faculty member
- Delete a faculty member
- Validate unique faculty code and email
- Validate that the selected department exists

### Courses

Fields:

- Course code
- Course name
- Department
- Faculty
- Credits
- Semester
- Academic year

Features:

- List courses with pagination
- Search by course code or course name
- Filter by department, faculty member, semester, and academic year
- View one course
- Create a course
- Edit a course
- Delete a course
- Validate unique course code
- Validate credits from 1 to 10
- Validate semester from 1 to 12
- Validate that the selected department and faculty member exist

## 4. API Contract

Use versioned endpoints under `/api/v1`.

```text
GET     /api/v1/departments
POST    /api/v1/departments
GET     /api/v1/departments/:id
PATCH   /api/v1/departments/:id
DELETE  /api/v1/departments/:id

GET     /api/v1/students
POST    /api/v1/students
GET     /api/v1/students/:id
PATCH   /api/v1/students/:id
DELETE  /api/v1/students/:id

GET     /api/v1/faculty
POST    /api/v1/faculty
GET     /api/v1/faculty/:id
PATCH   /api/v1/faculty/:id
DELETE  /api/v1/faculty/:id

GET     /api/v1/courses
POST    /api/v1/courses
GET     /api/v1/courses/:id
PATCH   /api/v1/courses/:id
DELETE  /api/v1/courses/:id
```

The API should return JSON with consistent validation, not-found, conflict, and server-error responses.

## 5. Frontend Pages

Build the frontend in Next.js and deploy it to Vercel.

```text
/login
/dashboard
/departments
/departments/new
/departments/:id
/departments/:id/edit
/students
/students/new
/students/:id
/students/:id/edit
/faculty
/faculty/new
/faculty/:id
/faculty/:id/edit
/courses
/courses/new
/courses/:id
/courses/:id/edit
```

Each resource page should provide:

- A searchable and paginated list.
- A create button and form.
- An edit action and form.
- A detail view.
- A delete action with confirmation.
- Loading, empty, validation-error, and API-error states.

The browser must never receive database credentials or private AWS secrets.

## 6. Backend Structure

The Go Fiber v3 backend should follow this structure:

```text
backend/
├── cmd/api/main.go
├── internals/
│   ├── handler/
│   │   ├── department.handler.go
│   │   ├── student.handler.go
│   │   ├── faculty.handler.go
│   │   └── course.handler.go
│   ├── middleware/
│   ├── model/
│   ├── router/
│   │   ├── department.router.go
│   │   ├── student.router.go
│   │   ├── faculty.router.go
│   │   └── course.router.go
│   ├── service/
│   └── store/
├── db/
│   ├── migration/
│   ├── query/
│   └── sqlc/
├── Makefile
└── sqlc.yaml
```

Responsibilities:

- `handler`: parse requests and return HTTP responses.
- `service`: validation and application rules.
- `store`: database access through sqlc-generated queries.
- `model`: request, response, and domain types.
- `router`: Fiber route registration.
- `middleware`: authentication, request IDs, logging, and recovery.
- `db/migration`: versioned PostgreSQL schema changes.
- `db/sqlc`: generated database code. Do not edit generated files manually.

## 7. Database Scope

The first schema contains:

- `departments`
- `students`
- `faculty`
- `courses`

The project should use `faculty` consistently in the API, Go model, SQL queries, and frontend.

Use `migrate` for schema changes and `sqlc` for typed PostgreSQL queries. Test every migration from an empty database.

## 8. Testing Requirements

### Backend

- Unit-test validation rules for every resource.
- Test create, list, get, update, and delete behavior.
- Test invalid IDs and malformed request bodies.
- Test not-found and duplicate-record errors.
- Test foreign-key validation for departments and courses.
- Test Fiber routes with `httptest`.

### Frontend

- Test each CRUD form's required fields and validation errors.
- Test list loading, empty, error, and success states.
- Test create, edit, and delete interactions.
- Run lint and production build before deployment.

### Database

- Run migrations against a fresh PostgreSQL database.
- Verify tables, foreign keys, unique constraints, and check constraints.
- Run sqlc generation after query changes.

## 9. Delivery Phases

### Phase 1: Backend foundation

- Complete Fiber server and health endpoints.
- Organize handlers and routers by resource.
- Connect PostgreSQL through sqlc.
- Implement department CRUD.
- Add unit and HTTP tests.

### Phase 2: Core CRUD

- Implement student CRUD.
- Implement faculty CRUD.
- Implement course CRUD.
- Add validation and relationship checks.
- Add seed data for local development.

### Phase 3: Frontend CRUD

- Build the administrator dashboard.
- Build department, student, faculty, and course list pages.
- Build create, detail, edit, and delete flows.
- Connect all pages to the Fiber API.
- Add frontend tests and responsive states.

### Phase 4: Deployment

- Deploy the Next.js frontend to Vercel.
- Deploy the Go API to EC2.
- Configure HTTPS, CORS, environment variables, and health checks.
- Connect the backend to PostgreSQL or RDS.
- Run migrations through a controlled deployment step.

## 10. MVP Completion Checklist

- The administrator can sign in.
- The administrator can perform full CRUD on departments.
- The administrator can perform full CRUD on students.
- The administrator can perform full CRUD on faculty.
- The administrator can perform full CRUD on courses.
- Search, filters, pagination, and validation work for each list.
- Invalid relationships and duplicate records are rejected clearly.
- Backend unit and HTTP tests pass.
- Frontend lint and production build pass.
- Database migrations and sqlc generation work from a clean setup.
- The frontend deploys to Vercel and the backend runs on EC2.
