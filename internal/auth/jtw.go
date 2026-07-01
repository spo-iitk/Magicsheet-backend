package auth

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/spo-iitk/Magicsheet-backend/internal/database"
)

func GenerateAccessToken(user *database.User) (string, error) {

	claims := jwt.MapClaims{
		"sub":  user.ID,                      
		"role": user.Role,                    
		"iat":  time.Now().Unix(),            
		"exp":  time.Now().Add(45 * time.Minute).Unix(), 
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(os.Getenv("JWT_SECRET")))
}