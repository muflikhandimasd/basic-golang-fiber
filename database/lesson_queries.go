package database

import (
	"basic-golang-fiber/models"
	"database/sql"
	"fmt"
)

// CreateLesson creates a new lesson
func CreateLesson(lesson *models.Lesson) error {
	query := `INSERT INTO lessons (course_id, title, content, video_url, duration, order_index, is_free)
			  VALUES (?, ?, ?, ?, ?, ?, ?)
			  RETURNING id, created_at, updated_at`

	err := DB.QueryRow(query, lesson.CourseID, lesson.Title, lesson.Content,
		lesson.VideoURL, lesson.Duration, lesson.OrderIndex, lesson.IsFree).
		Scan(&lesson.ID, &lesson.CreatedAt, &lesson.UpdatedAt)

	return err
}

// GetLessonByID retrieves a lesson by ID
func GetLessonByID(id int64) (*models.Lesson, error) {
	lesson := &models.Lesson{}
	query := `SELECT id, course_id, title, content, video_url, duration, order_index, is_free, created_at, updated_at
			  FROM lessons WHERE id = ?`

	err := DB.QueryRow(query, id).Scan(
		&lesson.ID, &lesson.CourseID, &lesson.Title, &lesson.Content,
		&lesson.VideoURL, &lesson.Duration, &lesson.OrderIndex, &lesson.IsFree,
		&lesson.CreatedAt, &lesson.UpdatedAt,
	)

	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("lesson not found")
	}

	return lesson, err
}

// GetLessonsByCourse retrieves all lessons for a specific course
func GetLessonsByCourse(courseID int64) ([]models.Lesson, error) {
	query := `SELECT id, course_id, title, content, video_url, duration, order_index, is_free, created_at, updated_at
			  FROM lessons
			  WHERE course_id = ?
			  ORDER BY order_index ASC`

	rows, err := DB.Query(query, courseID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	lessons := []models.Lesson{}
	for rows.Next() {
		var lesson models.Lesson
		err := rows.Scan(
			&lesson.ID, &lesson.CourseID, &lesson.Title, &lesson.Content,
			&lesson.VideoURL, &lesson.Duration, &lesson.OrderIndex, &lesson.IsFree,
			&lesson.CreatedAt, &lesson.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		lessons = append(lessons, lesson)
	}

	return lessons, nil
}

// UpdateLesson updates a lesson
func UpdateLesson(id int64, updates map[string]interface{}) error {
	if len(updates) == 0 {
		return fmt.Errorf("no updates provided")
	}

	query := "UPDATE lessons SET "
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
		return fmt.Errorf("lesson not found")
	}

	return nil
}

// DeleteLesson deletes a lesson
func DeleteLesson(id int64) error {
	query := `DELETE FROM lessons WHERE id = ?`
	result, err := DB.Exec(query, id)
	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rows == 0 {
		return fmt.Errorf("lesson not found")
	}

	return nil
}
