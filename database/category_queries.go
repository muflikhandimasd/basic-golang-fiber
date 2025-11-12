package database

import (
	"basic-golang-fiber/models"
	"database/sql"
	"fmt"
)

// CreateCategory creates a new category
func CreateCategory(category *models.Category) error {
	query := `INSERT INTO categories (name, description, slug)
			  VALUES (?, ?, ?)
			  RETURNING id, created_at, updated_at`

	err := DB.QueryRow(query, category.Name, category.Description, category.Slug).
		Scan(&category.ID, &category.CreatedAt, &category.UpdatedAt)

	return err
}

// GetCategoryByID retrieves a category by ID
func GetCategoryByID(id int64) (*models.Category, error) {
	category := &models.Category{}
	query := `SELECT id, name, description, slug, created_at, updated_at
			  FROM categories WHERE id = ?`

	err := DB.QueryRow(query, id).Scan(
		&category.ID, &category.Name, &category.Description,
		&category.Slug, &category.CreatedAt, &category.UpdatedAt,
	)

	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("category not found")
	}

	return category, err
}

// GetAllCategories retrieves all categories
func GetAllCategories() ([]models.Category, error) {
	query := `SELECT id, name, description, slug, created_at, updated_at
			  FROM categories
			  ORDER BY name ASC`

	rows, err := DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	categories := []models.Category{}
	for rows.Next() {
		var category models.Category
		err := rows.Scan(
			&category.ID, &category.Name, &category.Description,
			&category.Slug, &category.CreatedAt, &category.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		categories = append(categories, category)
	}

	return categories, nil
}

// UpdateCategory updates a category
func UpdateCategory(id int64, updates map[string]interface{}) error {
	if len(updates) == 0 {
		return fmt.Errorf("no updates provided")
	}

	query := "UPDATE categories SET "
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
		return fmt.Errorf("category not found")
	}

	return nil
}

// DeleteCategory deletes a category
func DeleteCategory(id int64) error {
	query := `DELETE FROM categories WHERE id = ?`
	result, err := DB.Exec(query, id)
	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rows == 0 {
		return fmt.Errorf("category not found")
	}

	return nil
}
