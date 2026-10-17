"use client";

import { useEffect, useState, type FormEvent } from "react";
import type { Department, Faculty, Resource } from "@/lib/api";
import { useAcademicStore } from "@/store/useAcademicStore";

type FormValues = Record<string, string>;
type RecordRow = Record<string, unknown>;
type ResourceConfig = { label: string; singular: string; columns: { key: string; label: string }[] };

const resources: Resource[] = ["departments", "students", "faculty", "courses"];
const configs: Record<Resource, ResourceConfig> = {
  departments: { label: "Departments", singular: "Department", columns: [{ key: "department_code", label: "Code" }, { key: "department_name", label: "Department name" }] },
  students: { label: "Students", singular: "Student", columns: [{ key: "student_code", label: "Student code" }, { key: "first_name", label: "Name" }, { key: "email", label: "Email" }, { key: "department_id", label: "Department" }, { key: "status", label: "Status" }] },
  faculty: { label: "Faculty", singular: "Faculty member", columns: [{ key: "faculty_code", label: "Faculty code" }, { key: "first_name", label: "Name" }, { key: "email", label: "Email" }, { key: "designation", label: "Designation" }, { key: "status", label: "Status" }] },
  courses: { label: "Courses", singular: "Course", columns: [{ key: "course_code", label: "Course code" }, { key: "course_name", label: "Course name" }, { key: "department_id", label: "Department" }, { key: "faculty_id", label: "Faculty" }, { key: "credits", label: "Credits" }, { key: "semester", label: "Semester" }] },
};

const emptyForms: Record<Resource, FormValues> = {
  departments: { department_code: "", department_name: "" },
  students: { student_code: "", first_name: "", last_name: "", email: "", department_id: "", enrollment_date: "", status: "active" },
  faculty: { faculty_code: "", first_name: "", last_name: "", email: "", department_id: "", designation: "", specialization: "", hire_date: "", status: "active" },
  courses: { course_code: "", course_name: "", department_id: "", faculty_id: "", credits: "3", semester: "1", academic_year: "" },
};

export default function Home() {
  const [activeResource, setActiveResource] = useState<Resource>("departments");
  const [query, setQuery] = useState("");
  const [form, setForm] = useState<FormValues>(emptyForms.departments);
  const [editingId, setEditingId] = useState<number | null>(null);
  const [formOpen, setFormOpen] = useState(false);
  const [formError, setFormError] = useState("");
  const { departments, students, faculty, courses, isLoading, error, fetchAll, createEntity, updateEntity, deleteEntity } = useAcademicStore();
  const recordsByResource = { departments, students, faculty, courses };
  const records = recordsByResource[activeResource] as unknown as RecordRow[];
  const config = configs[activeResource];
  const visibleRecords = records.filter((record) => JSON.stringify(record).toLowerCase().includes(query.toLowerCase()));

  useEffect(() => { void fetchAll(); }, [fetchAll]);

  function selectResource(resource: Resource) {
    setActiveResource(resource);
    setQuery("");
    setFormError("");
  }

  function openCreate() {
    setEditingId(null);
    setForm({ ...emptyForms[activeResource] });
    setFormError("");
    setFormOpen(true);
  }

  function openEdit(record: RecordRow) {
    const values = Object.fromEntries(Object.keys(emptyForms[activeResource]).map((key) => [key, record[key] == null ? "" : String(record[key])]));
    setEditingId(Number(record.id));
    setForm(values);
    setFormError("");
    setFormOpen(true);
  }

  async function submitForm(event: FormEvent<HTMLFormElement>) {
    event.preventDefault();
    setFormError("");
    const payload: Record<string, unknown> = { ...form };
    if (["students", "faculty", "courses"].includes(activeResource)) payload.department_id = Number(form.department_id);
    if (activeResource === "courses") {
      payload.credits = Number(form.credits);
      payload.semester = Number(form.semester);
      payload.faculty_id = form.faculty_id ? Number(form.faculty_id) : null;
    }
    try {
      if (editingId) await updateEntity(activeResource, editingId, payload);
      else await createEntity(activeResource, payload);
      setFormOpen(false);
    } catch {
      setFormError(useAcademicStore.getState().error);
    }
  }

  async function removeRecord(record: RecordRow) {
    if (!window.confirm(`Delete this ${config.singular.toLowerCase()}?`)) return;
    await deleteEntity(activeResource, Number(record.id));
  }

  function displayValue(record: RecordRow, key: string) {
    const value = record[key];
    if (key === "department_id") return departments.find((item) => item.id === Number(value))?.department_name ?? "Unknown";
    if (key === "faculty_id") return faculty.find((item) => item.id === Number(value))?.last_name ?? "Unassigned";
    if (key === "first_name") return `${value} ${record.last_name ?? ""}`;
    return value == null || value === "" ? "-" : String(value);
  }

  return (
    <main className="shell">
      <aside className="sidebar">
        <div className="brand"><span className="brand-mark">VM</span><span>Vmate</span></div>
        <div className="sidebar-label">Workspace</div>
        <nav className="nav-list" aria-label="Resources">
          {resources.map((resource) => <button className={`nav-item ${activeResource === resource ? "active" : ""}`} key={resource} onClick={() => selectResource(resource)}><span className={`nav-dot ${resource}`} />{configs[resource].label}<span className="nav-count">{recordsByResource[resource].length}</span></button>)}
        </nav>
        <div className="sidebar-footer"><span className="status-pulse" />API connected<div>Administrator workspace</div></div>
      </aside>

      <section className="content">
        <header className="topbar"><div><p className="eyebrow">Administration / Records</p><h1>Academic overview</h1></div><div className="profile"><span className="avatar">AD</span><span><strong>Admin</strong><small>Records manager</small></span></div></header>
        <div className="content-inner">
          <div className="welcome"><div><span className="kicker">Monday, October 17, 2026</span><h2>Keep your campus moving.</h2><p>A clear view of the academic records that power your day.</p></div><button className="primary-button" onClick={openCreate}>+ Add {config.singular}</button></div>
          <section className="stats" aria-label="Record totals"><Stat label="Departments" value={departments.length} tone="blue" /><Stat label="Students" value={students.length} tone="teal" /><Stat label="Faculty" value={faculty.length} tone="orange" /><Stat label="Courses" value={courses.length} tone="violet" /></section>
          <section className="workspace-panel">
            <div className="panel-heading"><div><span className="kicker">Directory</span><h2>{config.label}</h2></div><span className="record-total">{visibleRecords.length} records</span></div>
            <div className="toolbar"><div className="tabs">{resources.map((resource) => <button className={activeResource === resource ? "selected" : ""} key={resource} onClick={() => selectResource(resource)}>{configs[resource].label}</button>)}</div><label className="search"><span>Search</span><input value={query} onChange={(event) => setQuery(event.target.value)} placeholder={`Search ${config.label.toLowerCase()}...`} /></label></div>
            {error && <div className="error-banner">{error}</div>}
            <div className="table-wrap">{isLoading ? <div className="empty-state"><div className="loader" />Loading records</div> : visibleRecords.length === 0 ? <div className="empty-state"><strong>No records found</strong><span>Try another search or add a new {config.singular.toLowerCase()}.</span></div> : <table><thead><tr>{config.columns.map((column) => <th key={column.key}>{column.label}</th>)}<th><span className="sr-only">Actions</span></th></tr></thead><tbody>{visibleRecords.map((record) => <tr key={String(record.id)}>{config.columns.map((column) => <td key={column.key}>{column.key === "status" ? <span className={`status ${String(record[column.key])}`}>{displayValue(record, column.key)}</span> : displayValue(record, column.key)}</td>)}<td className="actions"><button onClick={() => openEdit(record)}>Edit</button><button className="delete-action" onClick={() => void removeRecord(record)}>Delete</button></td></tr>)}</tbody></table>}</div>
          </section>
        </div>
      </section>

      {formOpen && <div className="modal-backdrop" role="presentation" onMouseDown={(event) => { if (event.target === event.currentTarget) setFormOpen(false); }}><section className="modal" role="dialog" aria-modal="true" aria-labelledby="form-title"><div className="modal-header"><div><span className="kicker">{editingId ? "Edit record" : "New record"}</span><h2 id="form-title">{editingId ? `Edit ${config.singular}` : `Add ${config.singular}`}</h2></div><button className="close-button" onClick={() => setFormOpen(false)} aria-label="Close form">x</button></div><form onSubmit={submitForm}><Fields resource={activeResource} form={form} setForm={setForm} departments={departments} faculty={faculty} /><p className="form-error">{formError}</p><div className="form-actions"><button type="button" className="secondary-button" onClick={() => setFormOpen(false)}>Cancel</button><button type="submit" className="primary-button">{editingId ? "Save changes" : "Create record"}</button></div></form></section></div>}
    </main>
  );
}

function Stat({ label, value, tone }: { label: string; value: number; tone: string }) { return <div className="stat-card"><span className={`stat-icon ${tone}`} /><div><strong>{value}</strong><span>{label}</span></div></div>; }

function Fields({ resource, form, setForm, departments, faculty }: { resource: Resource; form: FormValues; setForm: (form: FormValues) => void; departments: Department[]; faculty: Faculty[] }) {
  const update = (key: string, value: string) => setForm({ ...form, [key]: value });
  const input = (key: string, label: string, type = "text", required = true) => <label className="field" key={key}><span>{label}</span><input type={type} value={form[key] ?? ""} onChange={(event) => update(key, event.target.value)} required={required} /></label>;
  const select = (key: string, label: string, options: { value: string; label: string }[]) => <label className="field" key={key}><span>{label}</span><select value={form[key] ?? ""} onChange={(event) => update(key, event.target.value)} required><option value="">Select {label.toLowerCase()}</option>{options.map((option) => <option value={option.value} key={option.value}>{option.label}</option>)}</select></label>;
  if (resource === "departments") return <div className="form-grid">{input("department_code", "Department code")}{input("department_name", "Department name")}</div>;
  if (resource === "students") return <div className="form-grid">{input("student_code", "Student code")}{select("department_id", "Department", departments.map((item) => ({ value: String(item.id), label: item.department_name })))}{input("first_name", "First name")}{input("last_name", "Last name")}{input("email", "Email", "email")}{input("enrollment_date", "Enrollment date", "date")}{select("status", "Status", [{ value: "active", label: "Active" }, { value: "graduated", label: "Graduated" }, { value: "withdrawn", label: "Withdrawn" }])}</div>;
  if (resource === "faculty") return <div className="form-grid">{input("faculty_code", "Faculty code")}{select("department_id", "Department", departments.map((item) => ({ value: String(item.id), label: item.department_name })))}{input("first_name", "First name")}{input("last_name", "Last name")}{input("email", "Email", "email")}{input("designation", "Designation")}{input("specialization", "Specialization", "text", false)}{input("hire_date", "Hire date", "date")}{select("status", "Status", [{ value: "active", label: "Active" }, { value: "on_leave", label: "On leave" }, { value: "retired", label: "Retired" }])}</div>;
  return <div className="form-grid">{input("course_code", "Course code")}{input("course_name", "Course name")}{select("department_id", "Department", departments.map((item) => ({ value: String(item.id), label: item.department_name })))}{select("faculty_id", "Faculty", faculty.map((item) => ({ value: String(item.id), label: `${item.first_name} ${item.last_name}` })))}{input("credits", "Credits", "number")}{input("semester", "Semester", "number")}{input("academic_year", "Academic year")}</div>;
}