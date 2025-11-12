package models

import "time"

type Enrollment struct {
	ID          int64      `json:"id"`
	UserID      int64      `json:"user_id"`
	CourseID    int64      `json:"course_id"`
	EnrolledAt  time.Time  `json:"enrolled_at"`
	CompletedAt *time.Time `json:"completed_at"`
	Progress    int        `json:"progress"` // 0-100
}

type EnrollmentWithDetails struct {
	Enrollment
	CourseName string `json:"course_name"`
	UserName   string `json:"user_name"`
}

type CreateEnrollmentSchema struct {
	UserID   int64 `json:"user_id" validate:"required"`
	CourseID int64 `json:"course_id" validate:"required"`
}

type Review struct {
	ID        int64     `json:"id"`
	UserID    int64     `json:"user_id"`
	CourseID  int64     `json:"course_id"`
	Rating    int       `json:"rating"` // 1-5
	Comment   string    `json:"comment"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type CreateReviewSchema struct {
	CourseID int64  `json:"course_id" validate:"required"`
	Rating   int    `json:"rating" validate:"required,min=1,max=5"`
	Comment  string `json:"comment"`
}
