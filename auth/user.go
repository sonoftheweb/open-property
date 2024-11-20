package main

import (
	"encoding/json"
	"net/http"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	ID       uint   `gorm:"primaryKey"`
	Username string `gorm:"unique;not null"`
	Email    string `gorm:"unique;not null"`
	Password string `gorm:"not null"`
}

func createUserHandler(db *gorm.DB) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        var input struct {
            Username           string `json:"username"`
            Email              string `json:"email"`
            Password           string `json:"password"`
            PasswordConfirmation string `json:"password_confirmation"`
        }

        w.Header().Set("Content-Type", "application/json")

        if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
            http.Error(w, `{"error": "Invalid input"}`, http.StatusBadRequest)
            return
        }

        if input.Password != input.PasswordConfirmation {
            http.Error(w, `{"error": "Passwords do not match"}`, http.StatusBadRequest)
            return
        }

        hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
        if err != nil {
            http.Error(w, `{"error": "Could not hash password"}`, http.StatusInternalServerError)
            return
        }

        user := User{
            Username: input.Username,
            Email:    input.Email,
            Password: string(hashedPassword),
        }

        if err := db.Create(&user).Error; err != nil {
            http.Error(w, `{"error": "Could not create user"}`, http.StatusInternalServerError)
            return
        }

        w.WriteHeader(http.StatusCreated)
        w.Write([]byte(`{"message": "User created successfully"}`))
    }
}
