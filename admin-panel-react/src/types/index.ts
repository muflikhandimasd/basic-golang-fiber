export interface User {
  id: number;
  email: string;
  full_name: string;
  role: 'student' | 'instructor' | 'admin';
  created_at: string;
  updated_at: string;
}

export interface Course {
  id: number;
  title: string;
  description: string;
  instructor_id: number;
  category_id?: number;
  price: number;
  thumbnail: string;
  level: 'beginner' | 'intermediate' | 'advanced';
  status: 'draft' | 'published' | 'archived';
  created_at: string;
  updated_at: string;
}

export interface CourseWithDetails extends Course {
  instructor_name: string;
  category_name?: string;
  lesson_count: number;
  enrollment_count: number;
}

export interface Lesson {
  id: number;
  course_id: number;
  title: string;
  content: string;
  video_url: string;
  duration: number;
  order_index: number;
  is_free: boolean;
  created_at: string;
  updated_at: string;
}

export interface Category {
  id: number;
  name: string;
  description: string;
  slug: string;
  created_at: string;
  updated_at: string;
}

export interface Enrollment {
  id: number;
  user_id: number;
  course_id: number;
  enrolled_at: string;
  completed_at?: string;
  progress: number;
}

export interface EnrollmentWithDetails extends Enrollment {
  course_name: string;
  user_name: string;
}

export interface DashboardStats {
  total_courses: number;
  total_users: number;
  total_enrollments: number;
  published_courses: number;
}

export interface ApiResponse<T> {
  status: boolean;
  data?: T;
  message?: string;
}

export interface LoginRequest {
  email: string;
  password: string;
}

export interface AuthResponse {
  status: boolean;
  data: {
    user: User;
    token: string;
  };
}
