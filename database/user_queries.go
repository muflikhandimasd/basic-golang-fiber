package database

import (
	"basic-golang-fiber/models"
	"database/sql"
	"fmt"
)

// CreateUser creates a new user
func CreateUser(user *models.User) error {
	query := `INSERT INTO users (email, password, full_name, role)
			  VALUES (?, ?, ?, ?)
			  RETURNING id, created_at, updated_at`

	err := DB.QueryRow(query, user.Email, user.Password, user.FullName, user.Role).
		Scan(&user.ID, &user.CreatedAt, &user.UpdatedAt)

	return err
}

// GetUserByID retrieves a user by ID
func GetUserByID(id int64) (*models.User, error) {
	user := &models.User{}
	query := `SELECT id, email, password, full_name, role, created_at, updated_at
			  FROM users WHERE id = ?`

	err := DB.QueryRow(query, id).Scan(
		&user.ID, &user.Email, &user.Password, &user.FullName,
		&user.Role, &user.CreatedAt, &user.UpdatedAt,
	)

	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("user not found")
	}

	return user, err
}

// GetUserByEmail retrieves a user by email
func GetUserByEmail(email string) (*models.User, error) {
	user := &models.User{}
	query := `SELECT id, email, password, full_name, role, created_at, updated_at
			  FROM users WHERE email = ?`

	err := DB.QueryRow(query, email).Scan(
		&user.ID, &user.Email, &user.Password, &user.FullName,
		&user.Role, &user.CreatedAt, &user.UpdatedAt,
	)

	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("user not found")
	}

	return user, err
}

// GetAllUsers retrieves all users with pagination
func GetAllUsers(limit, offset int) ([]models.User, error) {
	query := `SELECT id, email, full_name, role, created_at, updated_at
			  FROM users
			  ORDER BY id DESC
			  LIMIT ? OFFSET ?`

	rows, err := DB.Query(query, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	users := []models.User{}
	for rows.Next() {
		var user models.User
		err := rows.Scan(
			&user.ID, &user.Email, &user.FullName,
			&user.Role, &user.CreatedAt, &user.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

// UpdateUser updates a user's information
func UpdateUser(id int64, updates map[string]interface{}) error {
	if len(updates) == 0 {
		return fmt.Errorf("no updates provided")
	}

	query := "UPDATE users SET "
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
		return fmt.Errorf("user not found")
	}

	return nil
}

// DeleteUser deletes a user
func DeleteUser(id int64) error {
	query := `DELETE FROM users WHERE id = ?`
	result, err := DB.Exec(query, id)
	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rows == 0 {
		return fmt.Errorf("user not found")
	}

	return nil
}

// CountUsers returns total number of users
func CountUsers() (int, error) {
	var count int
	query := `SELECT COUNT(*) FROM users`
	err := DB.QueryRow(query).Scan(&count)
	return count, err
}
