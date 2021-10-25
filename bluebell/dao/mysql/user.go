package mysql

import (
	"bluebell/models"
	"crypto/md5"
	"database/sql"
	"encoding/hex"
)

const secret = "liwenzhou.com" //随便写个字符串

// CheckUserExist 检查指定用户名的用户是否存在
func CheckUserExist(username string) (err error) {
	sqlStr := `select count(user_id) from user where username=?`
	var count int
	if err = db.Get(&count, sqlStr, username); err != nil {
		return err
	}
	if count > 0 {
		return ErrorUserExist
	}
	return
}

// InsertUser 向数据库中插入一条新的用户记录
func InsertUser(user *models.User) (err error) {
	//加密密码
	user.Password = encryptPassword(user.Password)
	//执行SQL语句入库
	sqlStr := `insert into user(user_id,username,password) values(?,?,?)`
	_, err = db.Exec(sqlStr, user.UserID, user.Username, user.Password)
	return
}

func encryptPassword(oPassword string) string {
	h := md5.New()
	h.Write([]byte(secret)) //secret作为一个加言的字符串
	return hex.EncodeToString(h.Sum([]byte(oPassword)))
	//EncodeToString将字节转换为16进制的字符串
}
func Login(user *models.User) (err error) {
	oPassword := user.Password
	//查询用户名是否存在
	sqlStr := `select user_id,username,password from user where username=?`
	err = db.Get(user, sqlStr, user.Username)
	if err == sql.ErrNoRows {
		//sql自带的判断字段是否存在
		return ErrorUserNotExist
	}
	if err != nil {
		//查询数据库失败
		return err
	}
	//验证密码是否正确
	password := encryptPassword(oPassword)
	if password != user.Password {
		//return errors.New("密码错误")
		//修改后
		return ErrorInvalidPassword
	}
	return
}
func GetUserById(uid int64) (user *models.User, err error) {
	user = new(models.User)
	sqlStr := `select user_id, username from user where user_id = ?`
	err = db.Get(user, sqlStr, uid)
	return
}
