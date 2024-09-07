package service

import (
	"To-Do/internal/models"
	"errors"
	"github.com/dgrijalva/jwt-go"
	"gorm.io/gorm"
	"time"
)

const (
	salt = "123sdaio213iojioj324432jio"
	signingKey = "asda12#3"
	tokenTTL = 12 * time.Hour
	)



//func GetLastIDByUser(db *gorm.DB) uint{
//	var lastUser models.Users
//
//	err := db.Order("id desc").First(&lastUser).Error
//	if err != nil {
//		if err == gorm.ErrRecordNotFound {
//
//			return 0
//		}
//		return 0
//	}
//	return lastUser.ID
//}


type tokenClimes struct {
	jwt.StandardClaims
	UserID uint `json:"user_id"`
}

func GenerateToken(db *gorm.DB, login, password string) (string, error){
	var user models.Users
	user.Login = login
	user.Password = password
	if err := db.Model(&models.Users{}).Find(&user).Error; err != nil {
		return "", err
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClimes{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenTTL).Unix(),
			IssuedAt: time.Now().Unix(),
		},
		UserID: user.ID,
	})
	return token.SignedString([]byte(signingKey))
}


func RegistrationUsers(db *gorm.DB, login, password, email string) error {
	var user models.Users

	if err := db.Where("login = ?", login).First(&user).Error; err == nil {
		return errors.New("User already exists")
	} else if !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}

	user.Login = login
	user.Password = password
	user.Email = email
	if err := db.Create(&user).Error; err != nil {
		return err
	}
	return nil
}