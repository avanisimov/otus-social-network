package user

import (
	"context"
	"database/sql"
)

type Repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		db: db,
	}
}

func (r *Repository) SearchUsers(context context.Context, first_name string, second_name string, limit int) ([]User, error) {
	rows, err := r.db.QueryContext(context, `SELECT id, first_name, second_name, birthdate, biography, city FROM users WHERE first_name LIKE $1 AND second_name LIKE $2 LIMIT $3`, "%"+first_name+"%", "%"+second_name+"%", limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var u User
		err := rows.Scan(&u.ID, &u.FirstName, &u.SecondName, &u.Birthdate, &u.Biography, &u.City)
		if err != nil {
			return nil, err
		}
		users = append(users, u)
	}

	return users, nil
}



func (r *Repository) GetUserByID(ctx context.Context, userID string) (*User, error) {
	u := User{}
	err := r.db.QueryRowContext(ctx, `SELECT id, first_name, second_name, birthdate, biography, city FROM users WHERE id=$1`, userID).
		Scan(&u.ID, &u.FirstName, &u.SecondName, &u.Birthdate, &u.Biography, &u.City)

	if err != nil {
		return nil, err
	}

	return &u, nil
}

func (r *Repository) GetUserByIAndPassword(ctx context.Context, userID string, passwordHash string) (*User, error) {
	u := User{}
	err := r.db.QueryRowContext(ctx, `SELECT id, first_name, second_name, birthdate, biography, city FROM users WHERE id=$1 AND password_hash=$2`, userID, passwordHash).
		Scan(&u.ID, &u.FirstName, &u.SecondName, &u.Birthdate, &u.Biography, &u.City)

	if err != nil {
		return nil, err
	}

	return &u, nil
}

func (r *Repository) CreateUser(ctx context.Context, u *User) (string, error) {
	query := `
		INSERT INTO users (first_name, second_name, birthdate, biography, city, password_hash)
		VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING id
	`
	var id string
	err := r.db.QueryRowContext(ctx, query,
		u.FirstName, u.SecondName, u.Birthdate, u.Biography, u.City, u.PasswordHash,
	).Scan(&id)
	// id := fmt.Sprintf("user-%d", len(r.Users))
	// user.ID = id
	// r.Users[id] = &user
	return id, err
}
