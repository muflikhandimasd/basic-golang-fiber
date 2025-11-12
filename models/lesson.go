package models

import "time"

type Lesson struct {
	ID         int64     `json:"id"`
	CourseID   int64     `json:"course_id"`
	Title      string    `json:"title"`
	Content    string    `json:"content"`
	VideoURL   string    `json:"video_url"`
	Duration   int       `json:"duration"` // in seconds
	OrderIndex int       `json:"order_index"`
	IsFree     bool      `json:"is_free"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

type CreateLessonSchema struct {
	CourseID   int64  `json:"course_id" validate:"required"`
	Title      string `json:"title" validate:"required"`
	Content    string `json:"content"`
	VideoURL   string `json:"video_url"`
	Duration   int    `json:"duration" validate:"min=0"`
	OrderIndex int    `json:"order_index" validate:"min=0"`
	IsFree     bool   `json:"is_free"`
}

type UpdateLessonSchema struct {
	Title      string `json:"title"`
	Content    string `json:"content"`
	VideoURL   string `json:"video_url"`
	Duration   int    `json:"duration" validate:"min=0"`
	OrderIndex int    `json:"order_index" validate:"min=0"`
	IsFree     *bool  `json:"is_free"`
}
