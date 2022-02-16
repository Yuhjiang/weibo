package models

import (
	"errors"
	"fmt"
	orm "github.com/Yuhjiang/weibo/database"
	"github.com/Yuhjiang/weibo/utils"
	"github.com/golang-jwt/jwt"
	"log"
	"time"
)

type User struct {
	Id         int64     `json:"id" gorm:"primaryKey"`
	Username   string    `json:"username" binding:"required,min=5,max=10"`
	Password   string    `json:"password" binding:"required,min=5,max=20"`
	Status     int       `json:"status"`
	CreateTime time.Time `json:"createTime" gorm:"autoCreateTime"`
}

func (User) TableName() string {
	return "user"
}

// InsertUser 注册用户，需要对密码进行Md5加密，避免泄漏
func InsertUser(user *User) error {
	user.Password = utils.MD5(user.Password)
	result := orm.DB.Create(user)
	if result.Error != nil {
		return result.Error
	} else {
		return nil
	}
}

// GetUserByUsername 通过用户名查询用户
func GetUserByUsername(username string) (User, bool) {
	user := User{}
	res := orm.DB.Where("username = ?", username).First(&user)
	if res.Error != nil {
		return user, false
	}
	return user, true
}

func GetUserList() []User {
	var users []User
	orm.DB.Find(&users)
	return users
}

func LoginUser(user *User) bool {
	user.Password = utils.MD5(user.Password)
	res := orm.DB.Select("id", "username", "status",
		"create_time").Where(
		user).Where("status = ?", 0).First(&user)
	if res.Error != nil {
		return false
	} else {
		return true
	}
}

var secret = []byte("hello family")

type CustomClaims struct {
	Id       int64  `json:"id"`
	Username string `json:"username"`
	jwt.StandardClaims
}

func CreateToken(user User) (string, error) {
	claims := CustomClaims{
		user.Id,
		user.Username,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Unix() + int64(10),
			Issuer:    "admin",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(secret)
	if err != nil {
		log.Fatal("token生成失败")
	}
	return tokenString, err
}

func ValidateToken(tokenString string) (User, error) {
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return secret, nil
	})
	if err != nil {
		log.Println("token解析失败")
	}
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		fmt.Println(claims)
		return User{Id: claims.Id, Username: claims.Username}, nil
	} else {
		fmt.Println("失败处理")
		return User{}, errors.New("token解析错误")
	}
}
