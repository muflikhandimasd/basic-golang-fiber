package models

import "time"

type Category struct {
	ID          int64     `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Slug        string    `json:"slug"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type CreateCategorySchema struct {
	Name        string `json:"name" validate:"required"`
	Description string `json:"description"`
	Slug        string `json:"slug" validate:"required"`
}

type UpdateCategorySchema struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Slug        string `json:"slug"`
}
