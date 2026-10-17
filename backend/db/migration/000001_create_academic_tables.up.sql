-- Departments
CREATE TABLE departments (
	id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
	department_code VARCHAR(30) NOT NULL UNIQUE,
	department_name VARCHAR(150) NOT NULL UNIQUE,
	created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
	updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- Students
CREATE TABLE students (
	id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
	student_code VARCHAR(30) NOT NULL UNIQUE,
	first_name VARCHAR(100) NOT NULL,
	last_name VARCHAR(100) NOT NULL,
	email VARCHAR(255) NOT NULL UNIQUE,
	department_id BIGINT NOT NULL REFERENCES departments(id)
		ON UPDATE CASCADE ON DELETE RESTRICT,
	enrollment_date DATE NOT NULL DEFAULT CURRENT_DATE,
	status VARCHAR(20) NOT NULL DEFAULT 'active'
		CHECK (status IN ('active', 'graduated', 'withdrawn')),
	created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
	updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE INDEX idx_students_department
	ON students(department_id);

CREATE INDEX idx_students_status
	ON students(status);

-- Faculty members
CREATE TABLE faculty (
	id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
	faculty_code VARCHAR(30) NOT NULL UNIQUE,
	first_name VARCHAR(100) NOT NULL,
	last_name VARCHAR(100) NOT NULL,
	email VARCHAR(255) NOT NULL UNIQUE,
	department_id BIGINT NOT NULL REFERENCES departments(id)
		ON UPDATE CASCADE ON DELETE RESTRICT,
	designation VARCHAR(100) NOT NULL,
	specialization VARCHAR(150),
	hire_date DATE NOT NULL DEFAULT CURRENT_DATE,
	status VARCHAR(20) NOT NULL DEFAULT 'active'
		CHECK (status IN ('active', 'on_leave', 'retired')),
	created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
	updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE INDEX idx_faculty_department
	ON faculty(department_id);

CREATE INDEX idx_faculty_status
	ON faculty(status);

-- Courses taught by faculty
CREATE TABLE courses (
	id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
	course_code VARCHAR(30) NOT NULL UNIQUE,
	course_name VARCHAR(150) NOT NULL,
	department_id BIGINT NOT NULL REFERENCES departments(id)
		ON UPDATE CASCADE ON DELETE RESTRICT,
	faculty_id BIGINT REFERENCES faculty(id)
		ON UPDATE CASCADE ON DELETE SET NULL,
	credits SMALLINT NOT NULL CHECK (credits BETWEEN 1 AND 10),
	semester SMALLINT NOT NULL CHECK (semester BETWEEN 1 AND 12),
	academic_year VARCHAR(9) NOT NULL,
	created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE INDEX idx_courses_department
	ON courses(department_id);

CREATE INDEX idx_courses_faculty
	ON courses(faculty_id);

-- Students registered for courses
CREATE TABLE enrollments (
	id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
	student_id BIGINT NOT NULL REFERENCES students(id)
		ON DELETE CASCADE,
	course_id BIGINT NOT NULL REFERENCES courses(id)
		ON DELETE CASCADE,
	enrollment_date DATE NOT NULL DEFAULT CURRENT_DATE,
	status VARCHAR(20) NOT NULL DEFAULT 'enrolled'
		CHECK (status IN ('enrolled', 'completed', 'dropped')),
	created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
	UNIQUE (student_id, course_id)
);

CREATE INDEX idx_enrollments_course
	ON enrollments(course_id);

CREATE INDEX idx_enrollments_status
	ON enrollments(status);

-- Attendance records
CREATE TABLE attendance (
	id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
	enrollment_id BIGINT NOT NULL REFERENCES enrollments(id)
		ON DELETE CASCADE,
	attendance_date DATE NOT NULL DEFAULT CURRENT_DATE,
	status VARCHAR(10) NOT NULL
		CHECK (status IN ('present', 'absent', 'late', 'excused')),
	remarks VARCHAR(255),
	created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
	UNIQUE (enrollment_id, attendance_date)
);

CREATE INDEX idx_attendance_date
	ON attendance(attendance_date);

CREATE INDEX idx_attendance_enrollment_date
	ON attendance(enrollment_id, attendance_date);
