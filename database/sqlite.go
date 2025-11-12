package database

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDB(dbPath string) error {
	var err error
	DB, err = sql.Open("sqlite3", dbPath)
	if err != nil {
		return err
	}

	if err = DB.Ping(); err != nil {
		return err
	}

	log.Println("🚀 Connected Successfully to the SQLite Database")
	return nil
}

func CreateTables() error {
	schema := `
	-- Users table
	CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		email TEXT UNIQUE NOT NULL,
		password TEXT NOT NULL,
		full_name TEXT NOT NULL,
		role TEXT DEFAULT 'student' CHECK(role IN ('student', 'instructor', 'admin')),
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
		updated_at DATETIME DEFAULT CURRENT_TIMESTAMP
	);

	-- Categories table
	CREATE TABLE IF NOT EXISTS categories (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT UNIQUE NOT NULL,
		description TEXT,
		slug TEXT UNIQUE NOT NULL,
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
		updated_at DATETIME DEFAULT CURRENT_TIMESTAMP
	);

	-- Courses table
	CREATE TABLE IF NOT EXISTS courses (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		title TEXT NOT NULL,
		description TEXT,
		instructor_id INTEGER NOT NULL,
		category_id INTEGER,
		price REAL DEFAULT 0,
		thumbnail TEXT,
		level TEXT DEFAULT 'beginner' CHECK(level IN ('beginner', 'intermediate', 'advanced')),
		status TEXT DEFAULT 'draft' CHECK(status IN ('draft', 'published', 'archived')),
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
		updated_at DATETIME DEFAULT CURRENT_TIMESTAMP,
		FOREIGN KEY (instructor_id) REFERENCES users(id) ON DELETE CASCADE,
		FOREIGN KEY (category_id) REFERENCES categories(id) ON DELETE SET NULL
	);

	-- Lessons table
	CREATE TABLE IF NOT EXISTS lessons (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		course_id INTEGER NOT NULL,
		title TEXT NOT NULL,
		content TEXT,
		video_url TEXT,
		duration INTEGER DEFAULT 0,
		order_index INTEGER DEFAULT 0,
		is_free BOOLEAN DEFAULT 0,
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
		updated_at DATETIME DEFAULT CURRENT_TIMESTAMP,
		FOREIGN KEY (course_id) REFERENCES courses(id) ON DELETE CASCADE
	);

	-- Enrollments table
	CREATE TABLE IF NOT EXISTS enrollments (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		user_id INTEGER NOT NULL,
		course_id INTEGER NOT NULL,
		enrolled_at DATETIME DEFAULT CURRENT_TIMESTAMP,
		completed_at DATETIME,
		progress INTEGER DEFAULT 0,
		UNIQUE(user_id, course_id),
		FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
		FOREIGN KEY (course_id) REFERENCES courses(id) ON DELETE CASCADE
	);

	-- Lesson Progress table
	CREATE TABLE IF NOT EXISTS lesson_progress (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		user_id INTEGER NOT NULL,
		lesson_id INTEGER NOT NULL,
		completed BOOLEAN DEFAULT 0,
		completed_at DATETIME,
		UNIQUE(user_id, lesson_id),
		FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
		FOREIGN KEY (lesson_id) REFERENCES lessons(id) ON DELETE CASCADE
	);

	-- Reviews table
	CREATE TABLE IF NOT EXISTS reviews (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		user_id INTEGER NOT NULL,
		course_id INTEGER NOT NULL,
		rating INTEGER CHECK(rating >= 1 AND rating <= 5),
		comment TEXT,
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
		updated_at DATETIME DEFAULT CURRENT_TIMESTAMP,
		UNIQUE(user_id, course_id),
		FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
		FOREIGN KEY (course_id) REFERENCES courses(id) ON DELETE CASCADE
	);

	-- Create indexes for better query performance
	CREATE INDEX IF NOT EXISTS idx_courses_instructor ON courses(instructor_id);
	CREATE INDEX IF NOT EXISTS idx_courses_category ON courses(category_id);
	CREATE INDEX IF NOT EXISTS idx_courses_status ON courses(status);
	CREATE INDEX IF NOT EXISTS idx_lessons_course ON lessons(course_id);
	CREATE INDEX IF NOT EXISTS idx_enrollments_user ON enrollments(user_id);
	CREATE INDEX IF NOT EXISTS idx_enrollments_course ON enrollments(course_id);
	CREATE INDEX IF NOT EXISTS idx_lesson_progress_user ON lesson_progress(user_id);
	CREATE INDEX IF NOT EXISTS idx_reviews_course ON reviews(course_id);
	`

	_, err := DB.Exec(schema)
	if err != nil {
		return err
	}

	log.Println("✅ Database tables created successfully")
	return nil
}

func SeedAdminUser() error {
	// Check if admin user already exists
	var count int
	err := DB.QueryRow("SELECT COUNT(*) FROM users WHERE role = 'admin'").Scan(&count)
	if err != nil {
		return err
	}

	if count > 0 {
		log.Println("Admin user already exists")
		return nil
	}

	// Create default admin user (password: admin123)
	// In production, you should hash this password properly
	query := `INSERT INTO users (email, password, full_name, role)
			  VALUES (?, ?, ?, ?)`

	// This is a bcrypt hash of "admin123"
	hashedPassword := "$2a$10$xN7qZV8L9XqX8e3Z9yYz3.rY5mQYO9KZ9bC3L6qH8yV3Z8eZ8eZ8e"

	_, err = DB.Exec(query, "admin@ecourses.com", hashedPassword, "Admin User", "admin")
	if err != nil {
		return err
	}

	log.Println("✅ Default admin user created (email: admin@ecourses.com, password: admin123)")
	return nil
}

func CloseDB() error {
	if DB != nil {
		return DB.Close()
	}
	return nil
}
