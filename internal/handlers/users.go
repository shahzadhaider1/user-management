package handlers

import (
	"fmt"
	"math/rand"
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type User struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	PhoneNumber string `json:"phone_number"`
}

type OTP struct {
	PhoneNumber string    `json:"phone_number"`
	Code        string    `json:"code"`
	Expiration  time.Time `json:"expiration"`
}

var (
	users []User
	otps  []OTP
	mu    sync.Mutex
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// CreateUser creates a new user
func CreateUser(c *gin.Context) {
	var user User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	mu.Lock()
	defer mu.Unlock()

	for _, u := range users {
		if u.PhoneNumber == user.PhoneNumber {
			c.JSON(http.StatusBadRequest, gin.H{"error": "phone number already exists"})
			return
		}
	}

	user.ID = uuid.New().String()
	users = append(users, user)
	c.JSON(http.StatusCreated, user)
}

// GenerateOTP generates a 4-digit OTP for the given phone number
func GenerateOTP(c *gin.Context) {
	var request struct {
		PhoneNumber string `json:"phone_number"`
	}
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	mu.Lock()
	defer mu.Unlock()

	for _, u := range users {
		if u.PhoneNumber == request.PhoneNumber {
			otp := OTP{
				PhoneNumber: request.PhoneNumber,
				Code:        generateRandomOTP(),
				Expiration:  time.Now().Add(5 * time.Minute),
			}
			otps = append(otps, otp)
			c.JSON(http.StatusOK, gin.H{"message": "OTP generated", "otp": otp.Code})
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "phone number not found"})
}

// VerifyOTP verifies the OTP for the given phone number
func VerifyOTP(c *gin.Context) {
	var request struct {
		PhoneNumber string `json:"phone_number"`
		OTP         string `json:"otp"`
	}
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	mu.Lock()
	defer mu.Unlock()

	for i, otp := range otps {
		if otp.PhoneNumber == request.PhoneNumber && otp.Code == request.OTP {
			if otp.Expiration.After(time.Now()) {
				otps = append(otps[:i], otps[i+1:]...)
				c.JSON(http.StatusOK, gin.H{"message": "OTP verified"})
				return
			}
			c.JSON(http.StatusUnauthorized, gin.H{"error": "OTP expired"})
			return
		}
	}

	c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid OTP"})
}

func generateRandomOTP() string {
	return fmt.Sprintf("%04d", rand.Intn(10000))
}
