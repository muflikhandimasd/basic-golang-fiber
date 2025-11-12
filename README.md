# eCourses API - Golang with PostgreSQL

A complete e-learning platform API built with Golang, Fiber framework, and PostgreSQL database using raw SQL queries with database migrations. Includes a built-in admin panel for managing courses, users, and content.

## Features

- **Authentication & Authorization**: JWT-based authentication with role-based access control (Admin, Instructor, Student)
- **Course Management**: Create, update, delete, and publish courses
- **Lesson Management**: Organize course content into lessons with video support
- **Category System**: Organize courses by categories
- **Enrollment System**: Students can enroll in courses and track progress
- **Admin Panel**: Web-based admin interface for managing the platform
- **Raw SQL Queries**: Direct PostgreSQL queries for optimal performance
- **Database Migrations**: Automated schema migrations with golang-migrate
- **RESTful API**: Clean and well-documented API endpoints

## Tech Stack

- **Backend**: Go 1.24+ with Fiber v2
- **Database**: PostgreSQL with raw SQL queries
- **Migrations**: golang-migrate/migrate
- **Authentication**: JWT (golang-jwt/jwt/v5)
- **Password Hashing**: bcrypt
- **Template Engine**: Fiber HTML templates

## Getting Started

### Prerequisites

- Go 1.24 or higher
- PostgreSQL 12+ installed and running

### PostgreSQL Setup

1. Install PostgreSQL (if not already installed):
```bash
# macOS
brew install postgresql@15
brew services start postgresql@15

# Ubuntu/Debian
sudo apt-get install postgresql postgresql-contrib
sudo systemctl start postgresql

# Windows
Download from https://www.postgresql.org/download/windows/
```

2. Create database:
```bash
# Connect to PostgreSQL
psql -U postgres

# Create database
CREATE DATABASE ecourses_db;

# Create user (optional)
CREATE USER ecourses_user WITH PASSWORD 'your_password';
GRANT ALL PRIVILEGES ON DATABASE ecourses_db TO ecourses_user;

# Exit
\q
```

### Installation

1. Clone the repository
2. Install dependencies:
```bash
go mod download
```

3. Configure environment variables:
Edit `app.env` file with your PostgreSQL credentials:
```env
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=postgres
DB_DB=ecourses_db
DB_SSLMODE=disable
```

4. Run the application:
```bash
go run main.go
```

The application will automatically:
- Connect to PostgreSQL
- Run database migrations
- Seed initial data (admin user and categories)

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

## Database Migrations

The application uses `golang-migrate` for database migrations. Migrations are located in the `migrations/` directory.

### Migration Files

- `000001_create_initial_schema.up.sql` - Creates all database tables
- `000001_create_initial_schema.down.sql` - Drops all database tables
- `000002_seed_admin_user.up.sql` - Seeds admin user and categories
- `000002_seed_admin_user.down.sql` - Removes seeded data

### Automatic Migrations

Migrations run automatically when you start the application. The app will:
1. Connect to PostgreSQL
2. Check for pending migrations
3. Apply all pending migrations in order
4. Start the server

### Manual Migration Commands

You can also run migrations manually using the golang-migrate CLI:

```bash
# Install migrate CLI
go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

# Run all migrations
migrate -path migrations -database "postgresql://postgres:postgres@localhost:5432/ecourses_db?sslmode=disable" up

# Rollback last migration
migrate -path migrations -database "postgresql://postgres:postgres@localhost:5432/ecourses_db?sslmode=disable" down 1

# Check migration version
migrate -path migrations -database "postgresql://postgres:postgres@localhost:5432/ecourses_db?sslmode=disable" version
```

### Creating New Migrations

To create a new migration:

```bash
migrate create -ext sql -dir migrations -seq your_migration_name
```

This creates two files:
- `NNNNNN_your_migration_name.up.sql` - Apply changes
- `NNNNNN_your_migration_name.down.sql` - Revert changes

## Admin Panel

Access the admin panel at: `http://localhost:9000/admin`

### Features:
- **Dashboard**: Platform statistics overview
- **Courses Management**: View and manage all courses
- **Users Management**: View all registered users
- **Categories**: Manage course categories

## Database Schema

The application uses PostgreSQL with the following tables:

- **users**: User accounts with roles (student, instructor, admin)
- **categories**: Course categories
- **courses**: Course information and metadata
- **lessons**: Course lessons with content and videos
- **enrollments**: Student course enrollments and progress
- **reviews**: Course reviews and ratings
- **lesson_progress**: Individual lesson completion tracking

All tables include:
- Primary keys with SERIAL auto-increment
- Timestamps (created_at, updated_at)
- Foreign key constraints with CASCADE/SET NULL
- Indexes for performance optimization
- Automatic updated_at triggers

## Security Features

- Password hashing with bcrypt
- JWT-based authentication (24-hour token expiry)
- Role-based access control (Student, Instructor, Admin)
- Protected routes with middleware

## Project Structure

```
basic-golang-fiber/
├── controllers/          # API and Admin controllers
├── database/            # Database layer with raw SQL
│   ├── postgres.go      # PostgreSQL connection
│   ├── migrate.go       # Migration runner
│   ├── user_queries.go  # User CRUD operations
│   ├── course_queries.go
│   ├── lesson_queries.go
│   ├── enrollment_queries.go
│   └── category_queries.go
├── migrations/          # Database migration files
│   ├── 000001_create_initial_schema.up.sql
│   ├── 000001_create_initial_schema.down.sql
│   ├── 000002_seed_admin_user.up.sql
│   └── 000002_seed_admin_user.down.sql
├── middleware/          # Auth middleware
├── models/              # Data models
├── utils/               # Helper functions
├── views/               # HTML templates
├── initializers/        # Config loaders
├── app.env             # Environment configuration
└── main.go             # Application entry point
```

## License

MIT License
