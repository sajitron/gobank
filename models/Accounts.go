package models

import (
	"os"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/jinzhu/gorm"
	"github.com/sajicode/gobank/email"
	u "github.com/sajicode/gobank/utils"
	"golang.org/x/crypto/bcrypt"
)

type Token struct {
	UserId string
	jwt.StandardClaims
}

//* struct to represent user account
type Account struct {
	Base
	AvatarUrl string `sql:"type:VARCHAR(255);not null;DEFAULT:'https://res.cloudinary.com/sajicode/image/upload/v1549973773/avatar.png'"json:"avatar_url"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Role      string `sql:"type:VARCHAR(25);DEFAULT:'customer'"json:"role"`
	Token     string `json:"token";sql:"-"`
}

//* validate incoming user details
func (account *Account) Validate() (map[string]interface{}, bool) {
	if !strings.Contains(account.Email, "@") {
		return u.Message(false, "Email address is required"), false
	}

	if len(account.Password) < 6 {
		return u.Message(false, "Password is required"), false
	}

	if len(account.FirstName) < 2 {
		return u.Message(false, "Enter a valid firstname"), false
	}

	if len(account.LastName) < 2 {
		return u.Message(false, "Enter a valid lastname"), false
	}

	//* Email must be unique
	temp := &Account{}

	//* check for errors & duplicate emails
	err := GetDB().Table("accounts").Where("email = ?", account.Email).First(temp).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return u.Message(false, "Connection error. Please retry"), false
	}

	if temp.Email != "" {
		return u.Message(false, "Email already in use"), false
	}

	return u.Message(false, "Requirement passed"), true
}

func (account *Account) Create() (map[string]interface{}, bool) {
	if resp, ok := account.Validate(); !ok {
		//* return true if there is an error & false if none
		return resp, true
	}

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(account.Password), bcrypt.DefaultCost)
	account.Password = string(hashedPassword)

	GetDB().Create(account)

	if account.ID == "" {
		return u.Message(false, "Failed to create account, connection error."), true
	}

	//* create new JWT token for newly registered account
	tk := &Token{UserId: account.ID}
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
	tokenString, _ := token.SignedString([]byte(os.Getenv("token_password")))
	account.Token = tokenString

	account.Password = "" //* delete password

	response := u.Message(true, "Account has been created")
	response["account"] = account

	email.Mailer([]string{account.Email})

	return response, false
}

func Login(email, password string) (map[string]interface{}, bool) {
	account := &Account{}
	err := GetDB().Table("accounts").Where("email = ?", email).First(account).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return u.Message(false, "Email address not found"), true
		}
		return u.Message(false, "Connection error. Please retry"), true
	}

	err = bcrypt.CompareHashAndPassword([]byte(account.Password), []byte(password))

	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return u.Message(false, "Invalid login credentials. Please try again"), true
	}

	//* worked! logged in
	account.Password = ""

	//* Create JWT token
	tk := &Token{UserId: account.ID}
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
	tokenString, _ := token.SignedString([]byte(os.Getenv("token_password")))
	account.Token = tokenString //* store token in response

	resp := u.Message(true, "Logged In")
	resp["account"] = account
	return resp, false
}

func GetUser(u uint) *Account {
	acc := &Account{}
	GetDB().Table("accounts").Where("id = ?", u).First(acc)
	if acc.Email == "" { //* user not found
		return nil
	}

	acc.Password = ""
	return acc
}
