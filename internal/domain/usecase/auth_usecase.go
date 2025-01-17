package usecase

import (
	"Aurelia/internal/domain/models"
	"Aurelia/internal/domain/repository/postgresql"
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type authUsecase struct {
	userRepo postgresql.UserRepository
}

func NewAuthUsecase(u postgresql.UserRepository) *authUsecase {
	return &authUsecase{userRepo: u}
}

var jwtSecret = []byte("SUPER_SECRET_KEY")

func (a *authUsecase) SignUp(user *models.User) error {
	existing, err := a.userRepo.FindByEmail(user.UserEmail)
	if err != nil {
		return err
	}
	if existing != nil {
		return errors.New("email already in use")
	}

	// register user
	err = a.userRepo.Insert(user)
	if err != nil {
		return err
	}

	return nil
}

func (a *authUsecase) Login(email, password string) (string, error) {
	// emailからユーザーを取得
	user, err := a.userRepo.FindByEmail(email)
	if err != nil {
		return "", err
	}
	if user == nil {
		return "", errors.New("not found user")
	}

	// パスワードと符号するかチェック
	if ok := a.userRepo.CheckPassword(user.UserPassword, password); !ok {
		return "", errors.New("invalid email or password")
	}

	// ここまで大丈夫ならJWT tokenを発行
	tokenString, err := a.generateJWT(user)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (a *authUsecase) generateJWT(user *models.User) (string, error) {

	expirationTime := time.Now().Add(time.Hour)
	// claimに含めたい情報を書く
	claims := jwt.MapClaims{
		"user_id":  user.UserID,
		"email":    user.UserEmail,
		"exp":      expirationTime.Unix(),
		"issuedAt": time.Now().Unix(),
	}
	// そのclaimでtokenを作成
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// 署名された内容を返す
	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", fmt.Errorf("failed to sign token: %w", err)
	}

	return tokenString, nil
}
