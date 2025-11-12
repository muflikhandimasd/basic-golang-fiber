# eCourses API - Golang with SQLite

A complete e-learning platform API built with Golang, Fiber framework, and SQLite database using raw SQL queries. Includes a built-in admin panel for managing courses, users, and content.

## Features

- **Authentication & Authorization**: JWT-based authentication with role-based access control (Admin, Instructor, Student)
- **Course Management**: Create, update, delete, and publish courses
- **Lesson Management**: Organize course content into lessons with video support
- **Category System**: Organize courses by categories
- **Enrollment System**: Students can enroll in courses and track progress
- **Admin Panel**: Web-based admin interface for managing the platform
- **Raw SQL Queries**: Direct SQLite queries for optimal performance
- **RESTful API**: Clean and well-documented API endpoints

## Tech Stack

- **Backend**: Go 1.24+ with Fiber v2
- **Database**: SQLite3 with raw SQL queries
- **Authentication**: JWT (golang-jwt/jwt/v5)
- **Password Hashing**: bcrypt
- **Template Engine**: Fiber HTML templates

## Getting Started

### Prerequisites

- Go 1.24 or higher
- SQLite3

### Installation

1. Clone the repository
2. Install dependencies:
```bash
go mod download
```

3. Run the application:
```bash
go run main.go
```

The server will start on `http://localhost:9000`

### Default Admin Credentials

- **Email**: admin@ecourses.com
- **Password**: admin123

## Quick Start

### 1. Start the Server
```bash
go run main.go
```

### 2. Access the Admin Panel
Open your browser and navigate to: `http://localhost:9000/admin`

### 3. Test the API

#### Register a new user:
```bash
curl -X POST http://localhost:9000/api/auth/register \
  -H "Content-Type: application/json" \
  -d '{"email":"test@example.com","password":"password123","full_name":"Test User","role":"student"}'
```

#### Login:
```bash
curl -X POST http://localhost:9000/api/auth/login \
  -H "Content-Type: application/json" \
  -d '{"email":"admin@ecourses.com","password":"admin123"}'
```

## API Documentation

### Base URL
```
http://localhost:9000/api
```

### Authentication Endpoints

#### Register
```http
POST /api/auth/register
Content-Type: application/json

{
  "email": "user@example.com",
  "password": "password123",
  "full_name": "John Doe",
  "role": "student"
}
```

#### Login
```http
POST /api/auth/login
Content-Type: application/json

{
  "email": "user@example.com",
  "password": "password123"
}
```

#### Get Profile
```http
GET /api/auth/profile
Authorization: Bearer <token>
```

### Category Endpoints

- `GET /api/categories` - Get all categories
- `GET /api/categories/:id` - Get category by ID
- `POST /api/categories` - Create category (Admin only)
- `PUT /api/categories/:id` - Update category (Admin only)
- `DELETE /api/categories/:id` - Delete category (Admin only)

### Course Endpoints

- `GET /api/courses` - Get all courses
- `GET /api/courses/:id` - Get course by ID
- `GET /api/courses/:id/lessons` - Get course lessons
- `POST /api/courses` - Create course (Instructor/Admin)
- `GET /api/courses/my/courses` - Get my courses (Instructor)
- `PUT /api/courses/:id` - Update course (Instructor/Admin)
- `DELETE /api/courses/:id` - Delete course (Instructor/Admin)

### Lesson Endpoints

- `GET /api/lessons/:id` - Get lesson by ID
- `POST /api/lessons` - Create lesson (Instructor/Admin)
- `PUT /api/lessons/:id` - Update lesson (Instructor/Admin)
- `DELETE /api/lessons/:id` - Delete lesson (Instructor/Admin)

### Enrollment Endpoints

- `POST /api/enrollments` - Enroll in course
- `GET /api/enrollments/my` - Get my enrollments
- `PUT /api/enrollments/:id/progress` - Update progress
- `GET /api/courses/:courseId/enrollments` - Get course enrollments (Instructor/Admin)

## Admin Panel

Access the admin panel at: `http://localhost:9000/admin`

### Features:
- **Dashboard**: Platform statistics overview
- **Courses Management**: View and manage all courses
- **Users Management**: View all registered users
- **Categories**: Manage course categories

## Database Schema

The application uses SQLite with the following tables:

- **users**: User accounts with roles (student, instructor, admin)
- **categories**: Course categories
- **courses**: Course information and metadata
- **lessons**: Course lessons with content and videos
- **enrollments**: Student course enrollments and progress
- **reviews**: Course reviews and ratings
- **lesson_progress**: Individual lesson completion tracking

## Security Features

- Password hashing with bcrypt
- JWT-based authentication (24-hour token expiry)
- Role-based access control (Student, Instructor, Admin)
- Protected routes with middleware

## Project Structure

```
basic-golang-fiber/
├── controllers/     # API and Admin controllers
├── database/       # Database layer with raw SQL
├── middleware/     # Auth middleware
├── models/         # Data models
├── utils/          # Helper functions
├── views/          # HTML templates
└── main.go         # Application entry point
```

## License

MIT License
