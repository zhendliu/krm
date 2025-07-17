package auth

import (
	"crypto/md5"
	"fmt"
	"io"
	"krm-backend/config"
	"krm-backend/utils/jwtutil"
	"krm-backend/utils/logs"

	"github.com/gin-gonic/gin"
)

type UserInfo struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// 登录的逻辑
func Login(r *gin.Context) {
	fmt.Println("登陆 - Login")
	// 1. 获取前端传递用户名和密码
	userInfo := UserInfo{}
	returnData := config.NewReturnData()
	if err := r.ShouldBindJSON(&userInfo); err != nil {
		returnData.Status = 401
		returnData.Message = err.Error()
		r.JSON(200, returnData)
		return
	}
	logs.Info(map[string]interface{}{"用户名": userInfo.Username, "密码": userInfo.Password}, "开始验证登录信息")
	// 验证用户名和密码是否正确
	// 数据库 环境变量

	// md5  config.Username

	// 打印 Md5String(config.Username)
	fmt.Printf(" Md5String(userInfo.Username): %s\n", Md5String(userInfo.Username))
	// userInfo.Username
	fmt.Printf("config.Username: %s\n", config.Username)
	//  Md5String(config.Password)

	fmt.Printf("Md5String(userInfo.Password): %s\n", Md5String(userInfo.Password))
	// userInfo.Password
	fmt.Printf("config.Password: %s\n", config.Password)

	if Md5String(userInfo.Username) == config.Username && Md5String(userInfo.Password) == config.Password {
		// 认证成功
		// 生成JWT的Token
		ss, err := jwtutil.GenToken(userInfo.Username)
		if err != nil {
			logs.Error(map[string]interface{}{"用户名": userInfo.Username, "错误信息": err.Error()},
				"用户名密码正确,但生成token失败")
			r.JSON(200, gin.H{
				"status":  401,
				"message": "生成token失败",
			})
			return
		}
		// token正常生成，返回给前端
		logs.Info(map[string]interface{}{"用户名": userInfo.Username}, "登录成功")
		returnData.Message = "登录成功"
		returnData.Data["token"] = ss
		r.JSON(200, returnData)
		return
	} else {
		// 用户名和密码错误
		r.JSON(200, gin.H{
			"status":  401,
			"message": "用户名或密码错误",
		})
		return
	}
}

// 退出的逻辑
func Logout(r *gin.Context) {
	// 退出
	// 实现退出逻辑
	r.JSON(200, gin.H{
		"message": "退出成功",
		"status":  200,
	})
	logs.Debug(nil, "用户已退出")
	// 验证用户名和密码是否正确
}

func Md5String(input string) string {
	// 创建一个新的 MD5 哈希对象
	hash := md5.New()

	// 将输入字符串写入哈希对象
	io.WriteString(hash, input)

	// 计算哈希值并格式化为 16 进制字符串返回
	return fmt.Sprintf("%x", hash.Sum(nil))
}
