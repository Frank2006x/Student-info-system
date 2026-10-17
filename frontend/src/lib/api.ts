import axios from "axios";

export const api = axios.create({
  baseURL: process.env.NEXT_PUBLIC_API_URL ?? "http://localhost:8080/api/v1",
  headers: { "Content-Type": "application/json" },
});

export type Department = {
  id: number;
  department_code: string;
  department_name: string;
};

export type Student = {
  id: number;
  student_code: string;
  first_name: string;
  last_name: string;
  email: string;
  department_id: number;
  enrollment_date: string;
  status: string;
};

export type Faculty = {
  id: number;
  faculty_code: string;
  first_name: string;
  last_name: string;
  email: string;
  department_id: number;
  designation: string;
  specialization: string | null;
  hire_date: string;
  status: string;
};

export type Course = {
  id: number;
  course_code: string;
  course_name: string;
  department_id: number;
  faculty_id: number | null;
  credits: number;
  semester: number;
  academic_year: string;
};

export type Resource = "departments" | "students" | "faculty" | "courses";
export type Entity = Department | Student | Faculty | Course;

export const resourceApi = {
  list: (resource: Resource) => api.get<Entity[]>(`/${resource}`),
  create: (resource: Resource, payload: Record<string, unknown>) => api.post<Entity>(`/${resource}`, payload),
  update: (resource: Resource, id: number, payload: Record<string, unknown>) => api.patch<Entity>(`/${resource}/${id}`, payload),
  remove: (resource: Resource, id: number) => api.delete(`/${resource}/${id}`),
};
