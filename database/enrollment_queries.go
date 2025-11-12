package database

import (
	"basic-golang-fiber/models"
	"database/sql"
	"fmt"
)

// CreateEnrollment creates a new enrollment
func CreateEnrollment(enrollment *models.Enrollment) error {
	query := `INSERT INTO enrollments (user_id, course_id, progress)
			  VALUES (?, ?, ?)
			  RETURNING id, enrolled_at`

	err := DB.QueryRow(query, enrollment.UserID, enrollment.CourseID, enrollment.Progress).
		Scan(&enrollment.ID, &enrollment.EnrolledAt)

	return err
}

// GetEnrollmentByID retrieves an enrollment by ID
func GetEnrollmentByID(id int64) (*models.Enrollment, error) {
	enrollment := &models.Enrollment{}
	query := `SELECT id, user_id, course_id, enrolled_at, completed_at, progress
			  FROM enrollments WHERE id = ?`

	var completedAt sql.NullTime
	err := DB.QueryRow(query, id).Scan(
		&enrollment.ID, &enrollment.UserID, &enrollment.CourseID,
		&enrollment.EnrolledAt, &completedAt, &enrollment.Progress,
	)

	if completedAt.Valid {
		enrollment.CompletedAt = &completedAt.Time
	}

	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("enrollment not found")
	}

	return enrollment, err
}

// GetEnrollmentByUserAndCourse checks if user is enrolled in a course
func GetEnrollmentByUserAndCourse(userID, courseID int64) (*models.Enrollment, error) {
	enrollment := &models.Enrollment{}
	query := `SELECT id, user_id, course_id, enrolled_at, completed_at, progress
			  FROM enrollments WHERE user_id = ? AND course_id = ?`

	var completedAt sql.NullTime
	err := DB.QueryRow(query, userID, courseID).Scan(
		&enrollment.ID, &enrollment.UserID, &enrollment.CourseID,
		&enrollment.EnrolledAt, &completedAt, &enrollment.Progress,
	)

	if completedAt.Valid {
		enrollment.CompletedAt = &completedAt.Time
	}

	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("enrollment not found")
	}

	return enrollment, err
}

// GetEnrollmentsByUser retrieves all enrollments for a user
func GetEnrollmentsByUser(userID int64) ([]models.EnrollmentWithDetails, error) {
	query := `SELECT
				e.id, e.user_id, e.course_id, e.enrolled_at, e.completed_at, e.progress,
				c.title as course_name,
				u.full_name as user_name
			  FROM enrollments e
			  LEFT JOIN courses c ON e.course_id = c.id
			  LEFT JOIN users u ON e.user_id = u.id
			  WHERE e.user_id = ?
			  ORDER BY e.enrolled_at DESC`

	rows, err := DB.Query(query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	enrollments := []models.EnrollmentWithDetails{}
	for rows.Next() {
		var enrollment models.EnrollmentWithDetails
		var completedAt sql.NullTime

		err := rows.Scan(
			&enrollment.ID, &enrollment.UserID, &enrollment.CourseID,
			&enrollment.EnrolledAt, &completedAt, &enrollment.Progress,
			&enrollment.CourseName, &enrollment.UserName,
		)
		if err != nil {
			return nil, err
		}

		if completedAt.Valid {
			enrollment.CompletedAt = &completedAt.Time
		}

		enrollments = append(enrollments, enrollment)
	}

	return enrollments, nil
}

// GetEnrollmentsByCourse retrieves all enrollments for a course
func GetEnrollmentsByCourse(courseID int64) ([]models.EnrollmentWithDetails, error) {
	query := `SELECT
				e.id, e.user_id, e.course_id, e.enrolled_at, e.completed_at, e.progress,
				c.title as course_name,
				u.full_name as user_name
			  FROM enrollments e
			  LEFT JOIN courses c ON e.course_id = c.id
			  LEFT JOIN users u ON e.user_id = u.id
			  WHERE e.course_id = ?
			  ORDER BY e.enrolled_at DESC`

	rows, err := DB.Query(query, courseID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	enrollments := []models.EnrollmentWithDetails{}
	for rows.Next() {
		var enrollment models.EnrollmentWithDetails
		var completedAt sql.NullTime

		err := rows.Scan(
			&enrollment.ID, &enrollment.UserID, &enrollment.CourseID,
			&enrollment.EnrolledAt, &completedAt, &enrollment.Progress,
			&enrollment.CourseName, &enrollment.UserName,
		)
		if err != nil {
			return nil, err
		}

		if completedAt.Valid {
			enrollment.CompletedAt = &completedAt.Time
		}

		enrollments = append(enrollments, enrollment)
	}

	return enrollments, nil
}

// UpdateEnrollmentProgress updates enrollment progress
func UpdateEnrollmentProgress(id int64, progress int) error {
	query := `UPDATE enrollments SET progress = ? WHERE id = ?`

	result, err := DB.Exec(query, progress, id)
	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rows == 0 {
		return fmt.Errorf("enrollment not found")
	}

	// If progress is 100%, mark as completed
	if progress >= 100 {
		query = `UPDATE enrollments SET completed_at = CURRENT_TIMESTAMP WHERE id = ?`
		_, err = DB.Exec(query, id)
		if err != nil {
			return err
		}
	}

	return nil
}

// DeleteEnrollment deletes an enrollment
func DeleteEnrollment(id int64) error {
	query := `DELETE FROM enrollments WHERE id = ?`
	result, err := DB.Exec(query, id)
	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rows == 0 {
		return fmt.Errorf("enrollment not found")
	}

	return nil
}
