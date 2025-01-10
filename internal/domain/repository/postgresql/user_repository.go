package postgresql

import (
	"database/sql"
	"log"
)

type PostgresUserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *PostgresUserRepository {
	return &PostgresUserRepository{db: db}
}

func (r *PostgresUserRepository) FindAll() {
	log.Println("user repository FindAll")
}

func (r *PostgresUserRepository) FindByID(id int) {
	log.Println("user repository FindByID")
}

func (r *PostgresUserRepository) Insert() {
	log.Println("user repository Insert")
}

func (r *PostgresUserRepository) Delete(id int) {
	log.Println("user repository Delete")
}
