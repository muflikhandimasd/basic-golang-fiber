package database

import (
	"basic-golang-fiber/models"
	"database/sql"
	"fmt"
)

// CreateCourse creates a new course
func CreateCourse(course *models.Course) error {
	query := `INSERT INTO courses (title, description, instructor_id, category_id, price, thumbnail, level, status)
			  VALUES (?, ?, ?, ?, ?, ?, ?, ?)
			  RETURNING id, created_at, updated_at`

	err := DB.QueryRow(query, course.Title, course.Description, course.InstructorID,
		course.CategoryID, course.Price, course.Thumbnail, course.Level, course.Status).
		Scan(&course.ID, &course.CreatedAt, &course.UpdatedAt)

	return err
}

// GetCourseByID retrieves a course by ID
func GetCourseByID(id int64) (*models.Course, error) {
	course := &models.Course{}
	query := `SELECT id, title, description, instructor_id, category_id, price, thumbnail, level, status, created_at, updated_at
			  FROM courses WHERE id = ?`

	var categoryID sql.NullInt64
	err := DB.QueryRow(query, id).Scan(
		&course.ID, &course.Title, &course.Description, &course.InstructorID,
		&categoryID, &course.Price, &course.Thumbnail, &course.Level,
		&course.Status, &course.CreatedAt, &course.UpdatedAt,
	)

	if categoryID.Valid {
		course.CategoryID = &categoryID.Int64
	}

	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("course not found")
	}

	return course, err
}

// GetCourseWithDetails retrieves a course with instructor and category details
func GetCourseWithDetails(id int64) (*models.CourseWithDetails, error) {
	course := &models.CourseWithDetails{}
	query := `SELECT
				c.id, c.title, c.description, c.instructor_id, c.category_id,
				c.price, c.thumbnail, c.level, c.status, c.created_at, c.updated_at,
				u.full_name as instructor_name,
				COALESCE(cat.name, '') as category_name,
				(SELECT COUNT(*) FROM lessons WHERE course_id = c.id) as lesson_count,
				(SELECT COUNT(*) FROM enrollments WHERE course_id = c.id) as enrollment_count
			  FROM courses c
			  LEFT JOIN users u ON c.instructor_id = u.id
			  LEFT JOIN categories cat ON c.category_id = cat.id
			  WHERE c.id = ?`

	var categoryID sql.NullInt64
	err := DB.QueryRow(query, id).Scan(
		&course.ID, &course.Title, &course.Description, &course.InstructorID,
		&categoryID, &course.Price, &course.Thumbnail, &course.Level,
		&course.Status, &course.CreatedAt, &course.UpdatedAt,
		&course.InstructorName, &course.CategoryName, &course.LessonCount, &course.EnrollmentCount,
	)

	if categoryID.Valid {
		course.CategoryID = &categoryID.Int64
	}

	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("course not found")
	}

	return course, err
}

// GetAllCourses retrieves all courses with pagination
func GetAllCourses(limit, offset int, status string) ([]models.CourseWithDetails, error) {
	query := `SELECT
				c.id, c.title, c.description, c.instructor_id, c.category_id,
				c.price, c.thumbnail, c.level, c.status, c.created_at, c.updated_at,
				u.full_name as instructor_name,
				COALESCE(cat.name, '') as category_name,
				(SELECT COUNT(*) FROM lessons WHERE course_id = c.id) as lesson_count,
				(SELECT COUNT(*) FROM enrollments WHERE course_id = c.id) as enrollment_count
			  FROM courses c
			  LEFT JOIN users u ON c.instructor_id = u.id
			  LEFT JOIN categories cat ON c.category_id = cat.id`

	args := []interface{}{}
	if status != "" {
		query += " WHERE c.status = ?"
		args = append(args, status)
	}

	query += " ORDER BY c.id DESC LIMIT ? OFFSET ?"
	args = append(args, limit, offset)

	rows, err := DB.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	courses := []models.CourseWithDetails{}
	for rows.Next() {
		var course models.CourseWithDetails
		var categoryID sql.NullInt64

		err := rows.Scan(
			&course.ID, &course.Title, &course.Description, &course.InstructorID,
			&categoryID, &course.Price, &course.Thumbnail, &course.Level,
			&course.Status, &course.CreatedAt, &course.UpdatedAt,
			&course.InstructorName, &course.CategoryName, &course.LessonCount, &course.EnrollmentCount,
		)
		if err != nil {
			return nil, err
		}

		if categoryID.Valid {
			course.CategoryID = &categoryID.Int64
		}

		courses = append(courses, course)
	}

	return courses, nil
}

// GetCoursesByInstructor retrieves courses by instructor ID
func GetCoursesByInstructor(instructorID int64, limit, offset int) ([]models.CourseWithDetails, error) {
	query := `SELECT
				c.id, c.title, c.description, c.instructor_id, c.category_id,
				c.price, c.thumbnail, c.level, c.status, c.created_at, c.updated_at,
				u.full_name as instructor_name,
				COALESCE(cat.name, '') as category_name,
				(SELECT COUNT(*) FROM lessons WHERE course_id = c.id) as lesson_count,
				(SELECT COUNT(*) FROM enrollments WHERE course_id = c.id) as enrollment_count
			  FROM courses c
			  LEFT JOIN users u ON c.instructor_id = u.id
			  LEFT JOIN categories cat ON c.category_id = cat.id
			  WHERE c.instructor_id = ?
			  ORDER BY c.id DESC
			  LIMIT ? OFFSET ?`

	rows, err := DB.Query(query, instructorID, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	courses := []models.CourseWithDetails{}
	for rows.Next() {
		var course models.CourseWithDetails
		var categoryID sql.NullInt64

		err := rows.Scan(
			&course.ID, &course.Title, &course.Description, &course.InstructorID,
			&categoryID, &course.Price, &course.Thumbnail, &course.Level,
			&course.Status, &course.CreatedAt, &course.UpdatedAt,
			&course.InstructorName, &course.CategoryName, &course.LessonCount, &course.EnrollmentCount,
		)
		if err != nil {
			return nil, err
		}

		if categoryID.Valid {
			course.CategoryID = &categoryID.Int64
		}

		courses = append(courses, course)
	}

	return courses, nil
}

// UpdateCourse updates a course
func UpdateCourse(id int64, updates map[string]interface{}) error {
	if len(updates) == 0 {
		return fmt.Errorf("no updates provided")
	}

	query := "UPDATE courses SET "
	args := []interface{}{}
	i := 0

	for key, value := range updates {
		if i > 0 {
			query += ", "
		}
		query += key + " = ?"
		args = append(args, value)
		i++
	}

	query += ", updated_at = CURRENT_TIMESTAMP WHERE id = ?"
	args = append(args, id)

	result, err := DB.Exec(query, args...)
	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rows == 0 {
		return fmt.Errorf("course not found")
	}

	return nil
}

// DeleteCourse deletes a course
func DeleteCourse(id int64) error {
	query := `DELETE FROM courses WHERE id = ?`
	result, err := DB.Exec(query, id)
	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rows == 0 {
		return fmt.Errorf("course not found")
	}

	return nil
}

// CountCourses returns total number of courses
func CountCourses() (int, error) {
	var count int
	query := `SELECT COUNT(*) FROM courses`
	err := DB.QueryRow(query).Scan(&count)
	return count, err
}
