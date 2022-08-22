package services

import (
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/mickattapat/atp-labs/repositories"
	"golang.org/x/crypto/bcrypt"
)

const jwtSecret = "tesAtplabs"

type userService struct {
	userRepo repositories.UserRepository
}

func NewUserService(userRepo repositories.UserRepository) UserService {
	return userService{userRepo: userRepo}
}

func (s userService) Login(request UserLoginRequest) (string, error) {
	user := repositories.User{
		Username: request.Username,
		Password: request.Password,
	}

	userLogin, err := s.userRepo.Login(user)
	if err != nil {
		return "", err
	}
	match := CheckPasswordHash(request.Password, userLogin.Password)
	if !match {
		return "", errors.New("password do not found !!")
	}

	claims := jwt.StandardClaims{
		Issuer:    strconv.Itoa(userLogin.UserID),
		ExpiresAt: time.Now().Add(time.Hour * 1).Unix(),
	}

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	token, err := jwtToken.SignedString([]byte(jwtSecret))
	if err != nil {
		return "", err
	}

	return token, nil
}

func (s userService) GetUsers() ([]UserReponse, error) {
	users, err := s.userRepo.GetAll()
	if err != nil {
		return nil, err
	}

	usersResp := []UserReponse{}

	for _, user := range users {
		response := UserReponse{
			UserID:   user.UserID,
			Username: user.Username,
			Name:     user.Name,
			LastName: user.LastName,
			ExpDate:  user.ExpDate.Format("2006-01-02 15:04:05"),
			Role:     user.Role,
		}
		usersResp = append(usersResp, response)
	}

	return usersResp, nil
}

func (s userService) GetUser(id int) (*UserReponse, error) {
	user, err := s.userRepo.GetByID(id)
	if err != nil {
		return nil, err
	}
	userResp := UserReponse{
		UserID:   user.UserID,
		Username: user.Username,
		Name:     user.Name,
		LastName: user.LastName,
		ExpDate:  user.ExpDate.Format("2006-01-02 15:04:05"),
	}
	return &userResp, nil
}

func (s userService) CreateUser(request NewUserRequest) (*NewUserResponse, error) {

	password := request.Password
	hash, err := HashPassword(password)
	if err != nil {
		return nil, errors.New("Cannot hash password !")
	}

	expDate, err := time.Parse("2006-01-2 15:04:05", request.ExpDate)
	if err != nil {
		return nil, errors.New("Time no match !")
	}
	setStartDate := time.Now()
	startDate := setStartDate.Format("2006-01-02 15:04:05")
	fmt.Println(startDate)
	if err != nil {
		return nil, errors.New("Time no match !")
	}
	account := repositories.User{
		Username:  request.Username,
		Password:  hash,
		Name:      request.Name,
		LastName:  request.LastName,
		StartDate: startDate,
		ExpDate:   expDate,
		Status:    1,
	}
	newAccount, err := s.userRepo.Create(account)
	if err != nil {
		return nil, err
	}

	response := NewUserResponse{
		Username: newAccount.Username,
		Name:     newAccount.Name,
		LastName: newAccount.LastName,
		ExpDate:  newAccount.ExpDate,
	}

	return &response, nil
}

func (s userService) UpdateUser(id int, request UpdateUserRequest) error {
	var hash string
	var err error

	if request.Password != "" {
		password := request.Password
		hash, err = HashPassword(password)
		// hash = password
		if err != nil {
			return errors.New("Cannot hash password !")
		}
	}

	expDate, err := time.Parse("2006-01-2 15:04:05", request.ExpDate)
	if err != nil {
		return errors.New("Time no match !")
	}

	var account repositories.User

	if hash != "" {
		account = repositories.User{
			Password: hash,
			Name:     request.Name,
			LastName: request.LastName,
			ExpDate:  expDate,
			Status:   1,
		}
	} else {
		account = repositories.User{
			Name:     request.Name,
			LastName: request.LastName,
			ExpDate:  expDate,
			Status:   1,
		}
	}

	err = s.userRepo.Update(id, account)
	if err != nil {
		return err
	}

	return nil
}

func (s userService) DeleteUser(id int) error {
	err := s.userRepo.Delete(id)
	if err != nil {
		return err
	}
	return nil
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
