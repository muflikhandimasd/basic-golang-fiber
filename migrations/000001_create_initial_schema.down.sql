-- Drop triggers
DROP TRIGGER IF EXISTS update_reviews_updated_at ON reviews;
DROP TRIGGER IF EXISTS update_lessons_updated_at ON lessons;
DROP TRIGGER IF EXISTS update_courses_updated_at ON courses;
DROP TRIGGER IF EXISTS update_categories_updated_at ON categories;
DROP TRIGGER IF EXISTS update_users_updated_at ON users;

-- Drop trigger function
DROP FUNCTION IF EXISTS update_updated_at_column();

-- Drop indexes
DROP INDEX IF EXISTS idx_reviews_course;
DROP INDEX IF EXISTS idx_lesson_progress_user;
DROP INDEX IF EXISTS idx_enrollments_course;
DROP INDEX IF EXISTS idx_enrollments_user;
DROP INDEX IF EXISTS idx_lessons_course;
DROP INDEX IF EXISTS idx_courses_status;
DROP INDEX IF EXISTS idx_courses_category;
DROP INDEX IF EXISTS idx_courses_instructor;

-- Drop tables
DROP TABLE IF EXISTS reviews CASCADE;
DROP TABLE IF EXISTS lesson_progress CASCADE;
DROP TABLE IF EXISTS enrollments CASCADE;
DROP TABLE IF EXISTS lessons CASCADE;
DROP TABLE IF EXISTS courses CASCADE;
DROP TABLE IF EXISTS categories CASCADE;
DROP TABLE IF EXISTS users CASCADE;
