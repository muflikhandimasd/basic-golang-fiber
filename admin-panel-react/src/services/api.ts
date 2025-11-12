import axios, { AxiosInstance } from 'axios';
import {
  User,
  Course,
  CourseWithDetails,
  Lesson,
  Category,
  Enrollment,
  EnrollmentWithDetails,
  ApiResponse,
  LoginRequest,
  AuthResponse,
} from '../types';

const API_URL = import.meta.env.VITE_API_URL || 'http://localhost:9000/api';

class AdminApiService {
  private api: AxiosInstance;

  constructor() {
    this.api = axios.create({
      baseURL: API_URL,
      headers: {
        'Content-Type': 'application/json',
      },
    });

    // Add token to requests
    this.api.interceptors.request.use((config) => {
      const token = this.getToken();
      if (token) {
        config.headers.Authorization = `Bearer ${token}`;
      }
      return config;
    });
  }

  // Token management
  getToken(): string | null {
    return localStorage.getItem('admin_token');
  }

  setToken(token: string): void {
    localStorage.setItem('admin_token', token);
  }

  removeToken(): void {
    localStorage.removeItem('admin_token');
    localStorage.removeItem('admin_user');
  }

  setUser(user: User): void {
    localStorage.setItem('admin_user', JSON.stringify(user));
  }

  getUser(): User | null {
    const user = localStorage.getItem('admin_user');
    return user ? JSON.parse(user) : null;
  }

  // Authentication
  async login(data: LoginRequest): Promise<AuthResponse> {
    const response = await this.api.post<AuthResponse>('/auth/login', data);
    if (response.data.data) {
      this.setToken(response.data.data.token);
      this.setUser(response.data.data.user);
    }
    return response.data;
  }

  async getProfile(): Promise<ApiResponse<User>> {
    const response = await this.api.get<ApiResponse<User>>('/auth/profile');
    if (response.data.data) {
      this.setUser(response.data.data);
    }
    return response.data;
  }

  logout(): void {
    this.removeToken();
  }

  // Users
  async getUsers(limit = 100, offset = 0): Promise<User[]> {
    // Note: This endpoint doesn't exist in backend, using workaround
    const response = await this.api.get<ApiResponse<User[]>>('/users', {
      params: { limit, offset },
    });
    return response.data.data || [];
  }

  // Categories
  async getCategories(): Promise<Category[]> {
    const response = await this.api.get<ApiResponse<Category[]>>('/categories');
    return response.data.data || [];
  }

  async getCategory(id: number): Promise<Category> {
    const response = await this.api.get<ApiResponse<Category>>(`/categories/${id}`);
    return response.data.data!;
  }

  async createCategory(data: Partial<Category>): Promise<Category> {
    const response = await this.api.post<ApiResponse<Category>>('/categories', data);
    return response.data.data!;
  }

  async updateCategory(id: number, data: Partial<Category>): Promise<Category> {
    const response = await this.api.put<ApiResponse<Category>>(`/categories/${id}`, data);
    return response.data.data!;
  }

  async deleteCategory(id: number): Promise<void> {
    await this.api.delete(`/categories/${id}`);
  }

  // Courses
  async getCourses(params?: { limit?: number; offset?: number; status?: string }): Promise<CourseWithDetails[]> {
    const response = await this.api.get<ApiResponse<CourseWithDetails[]>>('/courses', { params });
    return response.data.data || [];
  }

  async getCourse(id: number): Promise<CourseWithDetails> {
    const response = await this.api.get<ApiResponse<CourseWithDetails>>(`/courses/${id}`);
    return response.data.data!;
  }

  async createCourse(data: Partial<Course>): Promise<Course> {
    const response = await this.api.post<ApiResponse<Course>>('/courses', data);
    return response.data.data!;
  }

  async updateCourse(id: number, data: Partial<Course>): Promise<Course> {
    const response = await this.api.put<ApiResponse<Course>>(`/courses/${id}`, data);
    return response.data.data!;
  }

  async deleteCourse(id: number): Promise<void> {
    await this.api.delete(`/courses/${id}`);
  }

  // Lessons
  async getCourseLessons(courseId: number): Promise<Lesson[]> {
    const response = await this.api.get<ApiResponse<Lesson[]>>(`/courses/${courseId}/lessons`);
    return response.data.data || [];
  }

  async getLesson(id: number): Promise<Lesson> {
    const response = await this.api.get<ApiResponse<Lesson>>(`/lessons/${id}`);
    return response.data.data!;
  }

  async createLesson(data: Partial<Lesson>): Promise<Lesson> {
    const response = await this.api.post<ApiResponse<Lesson>>('/lessons', data);
    return response.data.data!;
  }

  async updateLesson(id: number, data: Partial<Lesson>): Promise<Lesson> {
    const response = await this.api.put<ApiResponse<Lesson>>(`/lessons/${id}`, data);
    return response.data.data!;
  }

  async deleteLesson(id: number): Promise<void> {
    await this.api.delete(`/lessons/${id}`);
  }

  // Enrollments
  async getCourseEnrollments(courseId: number): Promise<EnrollmentWithDetails[]> {
    const response = await this.api.get<ApiResponse<EnrollmentWithDetails[]>>(`/courses/${courseId}/enrollments`);
    return response.data.data || [];
  }
}

export const adminApi = new AdminApiService();
