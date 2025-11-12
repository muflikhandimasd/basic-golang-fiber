import axios, { AxiosInstance } from 'axios';
import {
  AuthResponse,
  LoginRequest,
  RegisterRequest,
  User,
  Course,
  CourseWithDetails,
  Lesson,
  Category,
  Enrollment,
  EnrollmentWithDetails,
  ApiResponse,
} from '@/types';

const API_URL = process.env.NEXT_PUBLIC_API_URL || 'http://localhost:9000/api';

class ApiService {
  private api: AxiosInstance;

  constructor() {
    this.api = axios.create({
      baseURL: API_URL,
      headers: {
        'Content-Type': 'application/json',
      },
    });

    // Add token to requests if available
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
    if (typeof window !== 'undefined') {
      return localStorage.getItem('token');
    }
    return null;
  }

  setToken(token: string): void {
    if (typeof window !== 'undefined') {
      localStorage.setItem('token', token);
    }
  }

  removeToken(): void {
    if (typeof window !== 'undefined') {
      localStorage.removeItem('token');
      localStorage.removeItem('user');
    }
  }

  setUser(user: User): void {
    if (typeof window !== 'undefined') {
      localStorage.setItem('user', JSON.stringify(user));
    }
  }

  getUser(): User | null {
    if (typeof window !== 'undefined') {
      const user = localStorage.getItem('user');
      return user ? JSON.parse(user) : null;
    }
    return null;
  }

  // Authentication
  async register(data: RegisterRequest): Promise<AuthResponse> {
    const response = await this.api.post<AuthResponse>('/auth/register', data);
    if (response.data.data) {
      this.setToken(response.data.data.token);
      this.setUser(response.data.data.user);
    }
    return response.data;
  }

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

  async getMyCourses(): Promise<CourseWithDetails[]> {
    const response = await this.api.get<ApiResponse<CourseWithDetails[]>>('/courses/my/courses');
    return response.data.data || [];
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
  async enrollCourse(courseId: number): Promise<Enrollment> {
    const response = await this.api.post<ApiResponse<Enrollment>>('/enrollments', {
      course_id: courseId,
    });
    return response.data.data!;
  }

  async getMyEnrollments(): Promise<EnrollmentWithDetails[]> {
    const response = await this.api.get<ApiResponse<EnrollmentWithDetails[]>>('/enrollments/my');
    return response.data.data || [];
  }

  async updateEnrollmentProgress(id: number, progress: number): Promise<Enrollment> {
    const response = await this.api.put<ApiResponse<Enrollment>>(`/enrollments/${id}/progress`, { progress });
    return response.data.data!;
  }

  async getCourseEnrollments(courseId: number): Promise<EnrollmentWithDetails[]> {
    const response = await this.api.get<ApiResponse<EnrollmentWithDetails[]>>(`/courses/${courseId}/enrollments`);
    return response.data.data || [];
  }
}

export const apiService = new ApiService();
