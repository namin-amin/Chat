package users

import (
	"Chat/auth"
	"Chat/users/models"
	"context"
	"errors"
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

// IUserService
//
// methods involved to manipulate Users
type IUserService interface {
	IUserRepo
	Create(t models.SignUpDto, c context.Context) (models.SignUpRespDto, error) //service creates new user with proper validation
	Change(t models.User, c context.Context) error                              //Service changes user details with proper validation
	SignIn(t models.SigInDto, c context.Context) (models.SignUpRespDto, error)
}

// UserService
//
// Implementation of IUserService
type UserService struct {
	IUserRepo
}

func (u *UserService) Create(t models.SignUpDto, c context.Context) (models.SignUpRespDto, error) {
	var returnData models.SignUpRespDto
	validate := validator.New()
	err := validate.Struct(t)
	if err != nil {
		if _, ok := err.(*validator.InvalidValidationError); !ok {
			fmt.Println(err)
			return returnData, err
		}
	}

	hashedPassword, err := hashPassword(t.PassWord)

	if err != nil {
		return returnData, errors.New("could not create user try again")
	}

	newUser := &models.User{
		Name:     t.Name,
		Email:    t.Email,
		PassWord: hashedPassword,
		Role:     models.Normal,
	}
	newUser.Id = uuid.NewString()

	user, err := u.CreateNew(c, *newUser)

	if err != nil {
		return returnData, err
	}

	tk, err := auth.NewToken(user.Name, false, newUser.Id)

	if err != nil {
		return models.SignUpRespDto{}, errors.New("registering user failed")
	}
	returnData = models.SignUpRespDto{
		User:  user,
		Token: tk,
	}

	return returnData, nil

}

func (u *UserService) SignIn(t models.SigInDto, c context.Context) (models.SignUpRespDto, error) {

	var returnData models.SignUpRespDto
	hashedPassword, err := hashPassword(t.PassWord)

	if err != nil {
		return returnData, errors.New("could not login the user try again")
	}

	user := u.GetWithEmail(t.Email, hashedPassword)

	if user.Id == "" {
		return returnData, errors.New("user not found")
	}

	if !doPasswordsMatch(user.PassWord, t.PassWord) {
		return returnData, errors.New("wrong password")
	}

	tk, err := auth.NewToken(user.Name, false, user.Id)

	if err != nil {
		return models.SignUpRespDto{}, errors.New("registering user failed")
	}
	returnData = models.SignUpRespDto{
		User:  user,
		Token: tk,
	}

	return returnData, nil

}

func hashPassword(password string) (string, error) {
	// Convert password string to byte slice
	var passwordBytes = []byte(password)

	// Hash password with bcrypt min cost
	hashedPasswordBytes, err := bcrypt.
		GenerateFromPassword(passwordBytes, bcrypt.DefaultCost)

	return string(hashedPasswordBytes), err
}

func doPasswordsMatch(hashedPassword, currentPassword string) bool {
	err := bcrypt.CompareHashAndPassword(
		[]byte(hashedPassword), []byte(currentPassword))
	return err == nil
}

func (u UserService) Change(t models.User, c context.Context) error {
	return nil
}

// NewUserService
//
// Returns the instance of UserServices
func NewUserService(userRepo IUserRepo) *UserService {
	return &UserService{
		IUserRepo: userRepo,
	}
}
