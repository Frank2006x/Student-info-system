const baseUrl = process.env.API_URL ?? "http://localhost:8080";
const runId = Date.now().toString();
const unusedId = "9223372036854775807";

type Json = Record<string, unknown> | unknown[];

async function request(method: string, path: string, body?: Json) {
  const response = await fetch(`${baseUrl}${path}`, {
    method,
    headers: body ? { "Content-Type": "application/json" } : undefined,
    body: body ? JSON.stringify(body) : undefined,
  });

  const text = await response.text();
  let data: unknown = text;
  if (text) {
    try {
      data = JSON.parse(text);
    } catch {
      // Keep non-JSON responses as text.
    }
  }

  console.log(`${method.padEnd(6)} ${path.padEnd(32)} ${response.status}`);
  if (!response.ok && response.status !== 204) {
    throw new Error(`${method} ${path} failed: ${response.status} ${text}`);
  }
  return { status: response.status, data: data as Record<string, unknown> };
}

function idFrom(data: Record<string, unknown>, resource: string): string {
  const id = data.id;
  if (typeof id !== "number" && typeof id !== "string") {
    throw new Error(`${resource} response did not contain an id`);
  }
  return String(id);
}

const health = await request("GET", "/health");
if ((health.data as Record<string, unknown>).status !== "ok") {
  throw new Error("health check did not return status=ok");
}

const department = await request("POST", "/api/v1/departments", {
  department_code: `BUN-${runId}`,
  department_name: `Bun Test Department ${runId}`,
});
const departmentId = idFrom(department.data, "department");
await request("GET", "/api/v1/departments");
await request("GET", `/api/v1/departments/${departmentId}`);
await request("PATCH", `/api/v1/departments/${departmentId}`, {
  department_code: `BUN-UP-${runId}`,
  department_name: `Updated Bun Department ${runId}`,
});
await request("DELETE", `/api/v1/departments/${unusedId}`);

const student = await request("POST", "/api/v1/students", {
  student_code: `BUN-ST-${runId}`,
  first_name: "Bun",
  last_name: "Student",
  email: `bun-student-${runId}@example.com`,
  department_id: Number(departmentId),
  enrollment_date: "2026-10-17",
  status: "active",
});
const studentId = idFrom(student.data, "student");
await request("GET", "/api/v1/students");
await request("GET", `/api/v1/students/${studentId}`);
await request("PATCH", `/api/v1/students/${studentId}`, {
  student_code: `BUN-UP-ST-${runId}`,
  first_name: "Updated Bun",
  last_name: "Student",
  email: `bun-updated-student-${runId}@example.com`,
  department_id: Number(departmentId),
  enrollment_date: "2026-10-17",
  status: "active",
});
await request("DELETE", `/api/v1/students/${unusedId}`);

const faculty = await request("POST", "/api/v1/faculty", {
  faculty_code: `BUN-FC-${runId}`,
  first_name: "Bun",
  last_name: "Faculty",
  email: `bun-faculty-${runId}@example.com`,
  department_id: Number(departmentId),
  designation: "Professor",
  specialization: "API testing",
  hire_date: "2026-10-17",
  status: "active",
});
const facultyId = idFrom(faculty.data, "faculty");
await request("GET", "/api/v1/faculty");
await request("GET", `/api/v1/faculty/${facultyId}`);
await request("PATCH", `/api/v1/faculty/${facultyId}`, {
  faculty_code: `BUN-UP-FC-${runId}`,
  first_name: "Updated Bun",
  last_name: "Faculty",
  email: `bun-updated-faculty-${runId}@example.com`,
  department_id: Number(departmentId),
  designation: "Senior Professor",
  specialization: "API testing",
  hire_date: "2026-10-17",
  status: "active",
});
await request("DELETE", `/api/v1/faculty/${unusedId}`);

const course = await request("POST", "/api/v1/courses", {
  course_code: `BUN-CO-${runId}`,
  course_name: "Bun API Testing",
  department_id: Number(departmentId),
  faculty_id: Number(facultyId),
  credits: 3,
  semester: 1,
  academic_year: "2026-2027",
});
const courseId = idFrom(course.data, "course");
await request("GET", "/api/v1/courses");
await request("GET", `/api/v1/courses/${courseId}`);
await request("PATCH", `/api/v1/courses/${courseId}`, {
  course_code: `BUN-UP-CO-${runId}`,
  course_name: "Updated Bun API Testing",
  department_id: Number(departmentId),
  faculty_id: Number(facultyId),
  credits: 4,
  semester: 2,
  academic_year: "2026-2027",
});
await request("DELETE", `/api/v1/courses/${unusedId}`);

console.log("\nAll endpoint calls passed.");
console.log(`Test entities were preserved: department=${departmentId}, student=${studentId}, faculty=${facultyId}, course=${courseId}`);
export { };

