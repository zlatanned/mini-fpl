package auth

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
    "github.com/zlatanned/mini-fpl/configs"
)

// RegisterRequest defines expected JSON body for registration
type RegisterRequest struct {
	Username string `json:"username" binding:"required,alphanum,min=3,max=20"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}

// LoginRequest expects username & password
type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// Secret key for JWT signing
var jwtSecret = []byte(configs.GetJWTSecret())

var users = map[string]string{} // in-memory user store: username -> hashed password

func Register(c *gin.Context) {
	var req RegisterRequest

	// Check for any parsing or validation errors related to bindings provided in struct
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// checks if the username already exists in the users map.
	if _, exists := users[req.Username]; exists {
		c.JSON(http.StatusConflict, gin.H{"error": "username already exists"})
		return
	}

	// Hash the password before storing
	hashedPass, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to hash password"})
		return
	}

	users[req.Username] = string(hashedPass) // save hashed password

	c.JSON(http.StatusCreated, gin.H{"message": "user registered successfully"})
}

// Login handler
func Login(c *gin.Context) {
	var req LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	hashedPass, exists := users[req.Username]
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
		return
	}

	// Compare hashed password
	if err := bcrypt.CompareHashAndPassword([]byte(hashedPass), []byte(req.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
		return
	}

	// Create JWT token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": req.Username,
		"exp":      time.Now().Add(time.Hour * 72).Unix(), // expires in 72 hours
	})

	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not generate token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": tokenString})
}
