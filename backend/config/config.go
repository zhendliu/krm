// 配置层 存放程序的配置信息
package config

import (
	"fmt"
	"krm-backend/utils/logs"
	"path"
	"runtime"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"k8s.io/client-go/kubernetes"
)

const (
	TimeFormat                    string = "2006-01-02 15:04:05"
	ClusterConfigSecretLabelKey   string = "kubeasy.com/cluster.metadata"
	ClusterConfigSecretLabelValue string = "true"
)

var (
	Port       string
	JwtSignKey string
	JwtExpTime int64 // jwt token过期时间，单位：分钟
	Username   string
	Password   string
	// inCluster相关配置
	MetadataNamespace string // 元数据存储的namespace
	// inCluster 客户端
	InClusterClientSet *kubernetes.Clientset
	ClusterKubeconfig  map[string]string
)

type ReturnData struct {
	Status  int                    `json:"status"`
	Message string                 `json:"message"`
	Data    map[string]interface{} `json:"data"`
}

// 构造函数
func NewReturnData() ReturnData {
	returnData := ReturnData{}
	returnData.Status = 200
	data := make(map[string]interface{})
	returnData.Data = data
	return returnData
}

func initLogConfig(logLevel string) {
	// 配置程序的日志输出级别
	if logLevel == "debug" {
		logrus.SetLevel(logrus.DebugLevel)
	} else {
		logrus.SetLevel(logrus.InfoLevel)
	}
	// 文件名和行号加进去
	logrus.SetReportCaller(true)
	// 日志格式改成json
	logrus.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: TimeFormat,
		// runtime.Frame: 帧，可用于获取调用者返回的PC值的函数、文件或者是行信息
		CallerPrettyfier: func(f *runtime.Frame) (function string, file string) {
			//
			fileName := path.Base(f.File)
			return f.Function, fileName
		},
	})
}

func init() {
	// 环境变量加载我们的程序配置
	viper.SetDefault("LOG_LEVEL", "debug")
	// 获取程序启动端口号的配置
	viper.SetDefault("PORT", ":8080")
	// 获取jwt加密的secret
	viper.SetDefault("JWT_SIGN_KEY", "dukuan")
	// 获取jwt过期时间的配置
	viper.SetDefault("JWT_EXPIRE_TIME", 120)
	// 配置用户名密码的默认值
	// 加密用户名和密码 md5
	// 默认值：kunyu，kunyu123，部署时，请修改密码
	viper.SetDefault("USERNAME", "3e324c158f5ff42ffc8e60864ec32a50")
	viper.SetDefault("PASSWORD", "32037208992070c9ccadd14f01610673")
	// inCluster相关配置
	viper.SetDefault("METADATA_NAMESPACE", "krm")
	viper.AutomaticEnv()
	logLevel := viper.GetString("LOG_LEVEL") // 获取程序的配置
	// 加载日志输出格式
	initLogConfig(logLevel)
	logs.Debug(nil, "开始加载程序配置")
	Port = viper.GetString("PORT")
	JwtSignKey = viper.GetString("JWT_SIGN_KEY")
	JwtExpTime = viper.GetInt64("JWT_EXPIRE_TIME")
	// 获取用户名和密码
	Username = viper.GetString("USERNAME")
	Password = viper.GetString("PASSWORD")
	MetadataNamespace = viper.GetString("METADATA_NAMESPACE")
	// inCluster 相关配置
	fmt.Println("开始获取InClusterClientSet")
	fmt.Printf("Username:%+v,Password:%+v", Username, Password)
	logs.Debug(map[string]interface{}{"用户名": Username, "密码": Password}, "获取到程序用户名密码配置")
}
