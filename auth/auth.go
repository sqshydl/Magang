// auth/auth.go
package auth

import (
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/squishydal/MAGANG/initializers"
	"github.com/squishydal/MAGANG/models"
	"golang.org/x/crypto/bcrypt"
)

// Claims struct for JWT
type Claims struct {
	UserID string `json:"user_id"`
	Role   string `json:"role"`
	jwt.RegisteredClaims
}

// Get JWT key from environment variable
func getJWTKey() []byte {
	key := os.Getenv("JWT_SECRET")
	if key == "" {
		key = "your_secret_key" // fallback to default key
	}
	return []byte(key)
}

// HashPassword hashes a plain-text password
func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

// CheckPasswordHash compares a hashed password with a plain-text password
func CheckPasswordHash(password, hash string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
}

// AuthenticateUser checks if the username and password are correct
func AuthenticateUser(username, password string) (string, *models.User, error) {
	var user models.User

	// Find user by username
	if err := initializers.DB.Where("username = ?", username).First(&user).Error; err != nil {
		return "", nil, errors.New("user not found")
	}

	// Check if password is correct
	if err := CheckPasswordHash(password, user.Password); err != nil {
		return "", nil, errors.New("incorrect password")
	}

	// Generate and return JWT token
	token, err := generateJWT(user.UserID.String(), user.Role)
	if err != nil {
		return "", nil, err
	}

	return token, &user, nil
}

// generateJWT creates a JWT token for authenticated users
func generateJWT(userID string, role string) (string, error) {
	claims := &Claims{
		UserID: userID,
		Role:   role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(getJWTKey())
}

// ValidateToken validates the JWT token and returns the claims
func ValidateToken(tokenString string) (*Claims, error) {
	claims := &Claims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return getJWTKey(), nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, errors.New("invalid token")
	}

	return claims, nil
}
