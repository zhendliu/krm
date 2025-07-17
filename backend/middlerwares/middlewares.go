// 中间件层
package middlerwares

import (
	"bytes"
	"io/ioutil"
	"time"
	"krm-backend/config"
	"krm-backend/utils/jwtutil"
	"krm-backend/utils/logs"

	"github.com/gin-gonic/gin"
)

// CORS中间件
func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Header("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

// Logger中间件，打印请求和响应信息
func LoggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		method := c.Request.Method
		path := c.Request.URL.Path
		query := c.Request.URL.RawQuery

		// 读取body内容
		var bodyBytes []byte
		if c.Request.Body != nil {
			bodyBytes, _ = ioutil.ReadAll(c.Request.Body)
			c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))
		}

		logs.Info(map[string]interface{}{
			"method": method,
			"path": path,
			"query": query,
			"body": string(bodyBytes),
		}, "请求信息")

		// 捕获响应内容
		blw := &bodyLogWriter{body: bytes.NewBufferString(""), ResponseWriter: c.Writer}
		c.Writer = blw

		c.Next()

		latency := time.Since(start)
		status := c.Writer.Status()
		logs.Info(map[string]interface{}{
			"status": status,
			"latency": latency.String(),
			"response": blw.body.String(),
		}, "响应信息")
	}
}

type bodyLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w *bodyLogWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

func JWTAuth(r *gin.Context) {
	// 1. 除了login和logout之外的所有的接口，都要验证请求是否携带token，并且token是否合法
	requestUrl := r.FullPath()
	logs.Debug(map[string]interface{}{"请求路径": requestUrl}, "")
	if requestUrl == "/api/auth/login" || requestUrl == "/api/auth/logout" {
		logs.Debug(map[string]interface{}{"请求路径": requestUrl}, "登录和退出不需要验证token")
		r.Next()
		return
	}
	returnData := config.NewReturnData()
	// token
	// 其他接口需要验证token
	// 获取是否携带token
	tokenString := r.Request.Header.Get("Authorization")
	if tokenString == "" {
		// 说明请求没有携带token
		returnData.Status = 401
		returnData.Message = "请求未携带Token, 请登录后尝试"
		r.JSON(200, returnData)
		r.Abort()
		return
	}
	// token不为空，要去验证token是否合法
	claims, err := jwtutil.ParseToken(tokenString)
	if err != nil {
		returnData.Status = 401
		returnData.Message = "Token验证未通过"
		r.JSON(200, returnData)
		r.Abort()
		return
	}
	// 验证成功
	r.Set("claims", claims)
	r.Next()
}
