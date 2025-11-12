package models

import "time"

type Course struct {
	ID           int64     `json:"id"`
	Title        string    `json:"title"`
	Description  string    `json:"description"`
	InstructorID int64     `json:"instructor_id"`
	CategoryID   *int64    `json:"category_id"`
	Price        float64   `json:"price"`
	Thumbnail    string    `json:"thumbnail"`
	Level        string    `json:"level"` // beginner, intermediate, advanced
	Status       string    `json:"status"` // draft, published, archived
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type CourseWithDetails struct {
	Course
	InstructorName string `json:"instructor_name"`
	CategoryName   string `json:"category_name,omitempty"`
	LessonCount    int    `json:"lesson_count"`
	EnrollmentCount int   `json:"enrollment_count"`
}

type CreateCourseSchema struct {
	Title        string  `json:"title" validate:"required"`
	Description  string  `json:"description" validate:"required"`
	InstructorID int64   `json:"instructor_id" validate:"required"`
	CategoryID   *int64  `json:"category_id"`
	Price        float64 `json:"price" validate:"min=0"`
	Thumbnail    string  `json:"thumbnail"`
	Level        string  `json:"level" validate:"omitempty,oneof=beginner intermediate advanced"`
	Status       string  `json:"status" validate:"omitempty,oneof=draft published archived"`
}

type UpdateCourseSchema struct {
	Title       string  `json:"title"`
	Description string  `json:"description"`
	CategoryID  *int64  `json:"category_id"`
	Price       float64 `json:"price" validate:"min=0"`
	Thumbnail   string  `json:"thumbnail"`
	Level       string  `json:"level" validate:"omitempty,oneof=beginner intermediate advanced"`
	Status      string  `json:"status" validate:"omitempty,oneof=draft published archived"`
}
