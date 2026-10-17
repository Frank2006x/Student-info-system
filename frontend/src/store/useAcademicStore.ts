import { create } from "zustand";
import { api, type Course, type Department, type Faculty, type Resource, type Student, resourceApi } from "@/lib/api";

type AcademicState = {
  departments: Department[];
  students: Student[];
  faculty: Faculty[];
  courses: Course[];
  isLoading: boolean;
  error: string;
  fetchAll: () => Promise<void>;
  createEntity: (resource: Resource, payload: Record<string, unknown>) => Promise<void>;
  updateEntity: (resource: Resource, id: number, payload: Record<string, unknown>) => Promise<void>;
  deleteEntity: (resource: Resource, id: number) => Promise<void>;
};

function getMessage(error: unknown) {
  if (error && typeof error === "object" && "response" in error) {
    const response = (error as { response?: { data?: { error?: string } } }).response;
    return response?.data?.error ?? "The API request could not be completed.";
  }
  return "The API request could not be completed.";
}

export const useAcademicStore = create<AcademicState>((set) => ({
  departments: [],
  students: [],
  faculty: [],
  courses: [],
  isLoading: false,
  error: "",

  fetchAll: async () => {
    set({ isLoading: true, error: "" });
    try {
      const [departments, students, faculty, courses] = await Promise.all([
        api.get<Department[]>("/departments"),
        api.get<Student[]>("/students"),
        api.get<Faculty[]>("/faculty"),
        api.get<Course[]>("/courses"),
      ]);
      set({
        departments: departments.data,
        students: students.data,
        faculty: faculty.data,
        courses: courses.data,
        isLoading: false,
      });
    } catch (error) {
      set({ isLoading: false, error: getMessage(error) });
    }
  },

  createEntity: async (resource, payload) => {
    try {
      await resourceApi.create(resource, payload);
      await useAcademicStore.getState().fetchAll();
    } catch (error) {
      set({ error: getMessage(error) });
      throw error;
    }
  },

  updateEntity: async (resource, id, payload) => {
    try {
      await resourceApi.update(resource, id, payload);
      await useAcademicStore.getState().fetchAll();
    } catch (error) {
      set({ error: getMessage(error) });
      throw error;
    }
  },

  deleteEntity: async (resource, id) => {
    try {
      await resourceApi.remove(resource, id);
      await useAcademicStore.getState().fetchAll();
    } catch (error) {
      set({ error: getMessage(error) });
      throw error;
    }
  },
}));
