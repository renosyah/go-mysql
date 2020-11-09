package model

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
)

type (
	User struct {
		ID    int64  `json:"id"`
		Name  string `json:"name"`
		Email string `json:"email"`
	}

	UserResponse struct {
		ID    int64  `json:"id"`
		Name  string `json:"name"`
		Email string `json:"email"`
	}

	AllUser struct {
		SearchBy    string `json:"search_by"`
		SearchValue string `json:"search_value"`
		OrderBy     string `json:"order_by"`
		OrderDir    string `json:"order_dir"`
		Offset      int    `json:"offset"`
		Limit       int    `json:"limit"`
	}
)

func (u *User) Response() UserResponse {
	return UserResponse{
		ID:    u.ID,
		Name:  u.Name,
		Email: u.Email,
	}
}

func (u *User) Add(ctx context.Context, db *sql.DB) error {
	query := `INSERT INTO user (name,email) VALUES (?,?)`
	result, err := db.ExecContext(ctx, fmt.Sprintf(query), u.Name, u.Email)
	if err != nil {
		return err
	}

	u.ID, err = result.LastInsertId()
	if err != nil {
		return err
	}

	return nil
}

func (u *User) All(ctx context.Context, db *sql.DB, param AllUser) ([]User, error) {
	all := []User{}

	query := `SELECT id,name,email FROM user WHERE %s LIKE ? ORDER BY %s %s LIMIT ? OFFSET ?`
	rows, err := db.QueryContext(ctx, fmt.Sprintf(query, param.SearchBy, param.OrderBy, param.OrderDir), "%"+param.SearchValue+"%", param.Limit, param.Offset)
	if err != nil {
		return all, err
	}

	defer rows.Close()

	for rows.Next() {
		one := User{}
		err = rows.Scan(
			&one.ID, &one.Name, &one.Email,
		)
		if err != nil {
			return all, err
		}
		all = append(all, one)
	}

	return all, nil
}

func (u *User) One(ctx context.Context, db *sql.DB) (User, error) {
	one := User{}

	query := `SELECT id,name,email FROM user WHERE id = ? LIMIT 1`
	err := db.QueryRowContext(ctx, fmt.Sprintf(query), u.ID).Scan(
		&one.ID, &one.Name, &one.Email,
	)
	if err != nil {
		return one, err
	}

	return one, nil
}

func (u *User) Update(ctx context.Context, db *sql.DB) error {

	query := `UPDATE user SET name = ?, email = ? WHERE id = ?`
	result, err := db.ExecContext(ctx, fmt.Sprintf(query), u.Name, u.Email, u.ID)
	if err != nil {
		return err
	}

	num, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if num == 0 {
		return errors.New("no update is performed!")
	}

	return nil
}

func (u *User) Delete(ctx context.Context, db *sql.DB) error {

	query := `DELETE FROM user WHERE id = ?`
	result, err := db.ExecContext(ctx, fmt.Sprintf(query), u.ID)
	if err != nil {
		return err
	}

	num, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if num == 0 {
		return errors.New("no delete is performed!")
	}

	return nil
}
