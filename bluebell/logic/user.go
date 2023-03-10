package logic

import (
	"errors"
	"go-web/bluebell/dao/mysql"
	"go-web/bluebell/models"
	"go-web/bluebell/pkg/jwt"
	"go-web/bluebell/pkg/snowflake"
)

// 存放业务逻辑的代码
func SignUp(p *models.ParamSignUp) (err error) {
	//1. 判断用户存不存在
	if err = mysql.CheckUserExist(p.Username); err != nil {
		return err
	}

	//2. 生成UID
	userID := snowflake.GenID()
	if err != nil {
		return errors.New("生成用户ID错误")
	}
	// 构造一个User实例
	user := &models.User{
		UserID:   userID,
		UserName: p.Username,
		Password: p.Password,
	}

	//3. 保存进数据库
	return mysql.InsertUser(user)
}

func Login(p *models.ParamLogin) (token string, err error) {
	user := &models.User{
		UserName: p.Username,
		Password: p.Password,
	}
	// 传递的是指针，就能拿到user.UserID
	if err := mysql.Login(user); err != nil {
		return "", err
	}
	// 生成JWT
	return jwt.GenToken(user.UserID, user.UserName)

}
