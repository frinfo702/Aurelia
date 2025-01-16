package postgresql

import (
	"Aurelia/internal/domain/models"
	"database/sql"
	"log"

	"golang.org/x/crypto/bcrypt"
)

type PostgresUserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *PostgresUserRepository {
	return &PostgresUserRepository{db: db}
}

func (r *PostgresUserRepository) FindAll() ([]models.User, error) {
	// query to database
	query := "SELECT user_id, user_name, user_address, user_mail, user_password FROM users"
	rows, err := r.db.Query(query)
	if err != nil {
		log.Println()
		return nil, err
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		var user models.User
		err := rows.Scan(
			&user.UserID,
			&user.UserName,
			&user.UserAddress,
			&user.UserEmail,
			&user.UserPassword, // password_hash
		)
		if err != nil {
			log.Println("Error while scannig rows", err)
			return nil, err
		}
		users = append(users, user)

	}

	return users, nil
}

func (r *PostgresUserRepository) FindByEmail(email string) (models.User, error) {
	// query to database
	query := `SELECT user_id, user_name, user_address, user_mail, user_password 
			  FROM users 
			  WHERE user_mail = $1`
	// fetch

	var user models.User
	row := r.db.QueryRow(query, email)
	err := row.Scan(&user.UserID,
		&user.UserName,
		&user.UserAddress,
		&user.UserEmail,
		&user.UserPassword,
	)
	if err != nil {
		log.Println("error while scannig row: ", err)
		return models.User{}, err
	}
	return user, nil
}

func (r *PostgresUserRepository) Insert(user models.User) error {
	hashed, err := bcrypt.GenerateFromPassword([]byte(user.UserPassword), bcrypt.DefaultCost)
	if err != nil {
		log.Println("error while crypting user's password", err)
	}

	query := `INSERT INTO users (user_name, user_address, email, password_hash)
              VALUES ($1, $2, $3, $4) RETURNING user_id`

	err = r.db.QueryRow(query,
		user.UserName,
		user.UserAddress,
		user.UserEmail,
		string(hashed),
	).Scan(&user.UserID)
	if err != nil {
		log.Println("error while inserting user data", err)
	}
	return nil
}

func (r *PostgresUserRepository) CheckPassword(hash, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return nil == err
}
