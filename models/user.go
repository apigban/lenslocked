package models

import (
	"database/sql"
	"fmt"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID           int
	Email        string
	PasswordHash string
}

type UserService struct {
	DB *sql.DB
}

func (us *UserService) Create(email, password string) (*User, error) {
	email = strings.ToLower(email)
	hashedBytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, fmt.Errorf("create user: %w", err)
	}
	passwordHash := string(hashedBytes)

	user := User{
		Email:        email,
		PasswordHash: passwordHash,
	}

	row := us.DB.QueryRow(`
	INSERT INTO users (email, password_hash)
	VALUES ($1, $2) RETURNING id`, email, passwordHash)

	err = row.Scan(&user.ID)
	if err != nil {
		return nil, fmt.Errorf("create user: %w", err)
	}

	return &user, nil
}

func (us *UserService) Authenticate(email, password string) (*User, error) {
	// Normalize email address to lowercase
	email = strings.ToLower(email)

	// Create User instance for customer that wants to authenticate
	user := User{
		Email: email,
	}

	// Query the database with information given by the customer
	row := us.DB.QueryRow(
		`SELECT id, password_hash
		FROM users WHERE email=$1`, email)

	//Scan the result - access the db row and assign to struct fields
	err := row.Scan(&user.ID, &user.PasswordHash)
	if err != nil {
		//TODO: handle error when an email address doesn't have an account
		//TODO: in other words, the result of QueryRow (deferred to Scan method) is 0 row
		//TODO: Rows.Scan() produces ErrNoRows if QueryRow finds nothing
		return nil, fmt.Errorf("authenticate: %w", err)
	}

	//Validate the email and password_hash provided by the customer by comparing it to the database entry
	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password))
	if err != nil {
		fmt.Printf("Password is invalid: %v\n", password)
		return nil, fmt.Errorf("authenticate: %w, err")
	}
	fmt.Println("Password is correct!")
	return &user, nil
}
