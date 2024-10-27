package auth

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/squishydal/MAGANG/initializers"
	"github.com/squishydal/MAGANG/models"
	"golang.org/x/crypto/bcrypt"
)

// Secret key for signing JWT tokens
var jwtKey = []byte("your_secret_key")

// AuthenticateUser checks if the username and password are correct
func AuthenticateUser(username, password string) (string, error) {
	var user models.User

	// Find user by username
	if err := initializers.DB.Where("username = ?", username).First(&user).Error; err != nil {
		return "", errors.New("user not found")
	}

	// Check if password is correct
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return "", errors.New("incorrect password")
	}

	// Generate and return JWT token
	token, err := generateJWT(user.UserID.String())
	if err != nil {
		return "", err
	}

	return token, nil
}

// generateJWT creates a JWT token for authenticated users
func generateJWT(userID string) (string, error) {
	claims := &jwt.RegisteredClaims{
		Subject:   userID,
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)), // Token valid for 24 hours
		IssuedAt:  jwt.NewNumericDate(time.Now()),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}
