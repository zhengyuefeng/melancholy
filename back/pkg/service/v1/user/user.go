package user

import (
	model "github.com/HaHadaxigua/melancholy/pkg/model/user"
	"github.com/HaHadaxigua/melancholy/pkg/msg"
	"github.com/HaHadaxigua/melancholy/pkg/store/user"
)

//CreateUser 请求创建用户
func CreateUser(r *msg.UserRequest) (*model.User, error) {
	valid, err := VerifyReq(r)
	if !valid && err != nil {
		return nil, err
	}
	// todo: 判断用户是否存在时， 出现问题
	eu, err := user.GetUserByEmail(r.Email)
	if eu != nil {
		return nil, msg.UserHasExistedErr
	}

	newUser, err := model.NewUser(r.Username, r.Password, r.Email)
	err = user.CreateUser(newUser)
	if err != nil {
		return nil, msg.UserCreateErr
	}
	return newUser, nil
}

//FindUserByUsername 根据用户名找学生
func FindUserByUsername(r *msg.UserRequest) (*model.User, error) {
	if !CheckUsername(r.Username) {
		return nil, msg.UserNameIllegalErr
	}
	tu, err := user.GetUserByName(r.Username)
	if err != nil {
		return nil, err
	}
	return tu, nil
}

//ListAllUser 列出所有的用户
func ListAllUser() ([]*model.User, error) {
	users, err := user.GetAllUsers()
	if err != nil {
		return nil, err
	}
	return users, nil
}

// VerifyReq 验证请求合法性
func VerifyReq(r *msg.UserRequest) (bool, error) {
	if !CheckUsername(r.Username) {
		return false, msg.UserNameOrPwdIncorrectlyErr
	}
	if !CheckPassword(r.Password) {
		return false, msg.UserPwdIllegalErr
	}
	if !CheckEmail(r.Email) {
		return false, msg.UserEmailIllegalErr
	}
	return true, msg.OK
}

// todo 用户名合法性
func CheckUsername(username string) bool {
	return true
}

// todo 密码合法性
func CheckPassword(password string) bool {
	return true
}

// todo 邮箱合法性
func CheckEmail(email string) bool {
	return true
}