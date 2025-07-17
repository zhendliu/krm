// 控制器层 实现路由的处理逻辑
package controllers

import (
	"errors"
	"krm-backend/config"
	"krm-backend/utils/logs"
	"net/http"

	"github.com/dotbalo/kubeutils/kubeutils"
	"github.com/gin-gonic/gin"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

// 定义全局的数据结构
type BasicInfo struct {
	ClusterId  string      `json:"clusterId" form:"clusterId"`
	Namespace  string      `json:"namespace" form:"namespace"`
	Name       string      `json:"name" form:"name"`
	Item       interface{} `json:"item"`
	DeleteList []string    `json:"deleteList"`
}
type Info struct {
	BasicInfo
	ReturnData    config.ReturnData
	LabelSelector string `json:"labelSelector" form:"labelSelector"`
	FieldSelector string `json:"fieldSelector" form:"fieldSelector"`
	// 判断是否是强制删除
	Force bool `json:"force" form:"force"`
}

func NewInfo(r *gin.Context, info *Info, returnDataMsg string) (kubeconfig string) {
	// 首先获取请求的类型
	requestMethod := r.Request.Method
	var err error
	// var returnData config.ReturnData
	info.ReturnData.Message = returnDataMsg
	info.ReturnData.Status = 200
	if requestMethod == "GET" {
		err = r.ShouldBindQuery(&info)
	} else if requestMethod == "POST" {
		err = r.ShouldBindJSON(&info)
	} else {
		err = errors.New("不支持的请求类型")
	}
	logs.Debug(map[string]interface{}{"info": info}, "数据绑定结果")
	if err != nil {
		msg := "请求出错: " + err.Error()
		info.ReturnData.Message = msg
		info.ReturnData.Status = 400
		logs.Error(nil, msg)
		r.JSON(http.StatusOK, info.ReturnData)
		return
	}
	// 获取kubeconfig
	kubeconfig = config.ClusterKubeconfig[info.ClusterId]

	return kubeconfig
}

func (c *Info) Create(r *gin.Context, kubeUtilsInterface kubeutils.KubeUtilser) {
	err := kubeUtilsInterface.Create(c.Namespace)
	if err != nil {
		msg := "创建失败: " + err.Error()
		c.ReturnData.Message = msg
		c.ReturnData.Status = 400
		logs.Error(nil, msg)
	}
	r.JSON(200, c.ReturnData)
}

func (c *Info) Update(r *gin.Context, kubeUtilsInterface kubeutils.KubeUtilser) {
	err := kubeUtilsInterface.Update(c.Namespace)
	if err != nil {
		msg := "更新失败: " + err.Error()
		c.ReturnData.Message = msg
		c.ReturnData.Status = 400
		logs.Error(nil, msg)
	}
	r.JSON(200, c.ReturnData)
}
func (c *Info) List(r *gin.Context, kubeUtilsInterface kubeutils.KubeUtilser) {
	items, err := kubeUtilsInterface.List(c.Namespace, c.LabelSelector, c.FieldSelector)
	if err != nil {
		msg := "查询失败: " + err.Error()
		c.ReturnData.Message = msg
		c.ReturnData.Status = 400
		logs.Error(nil, msg)
	} else {
		c.ReturnData.Data = make(map[string]interface{})
		c.ReturnData.Data["items"] = items
	}
	r.JSON(200, c.ReturnData)
}

func (c *Info) Get(r *gin.Context, kubeUtilsInterface kubeutils.KubeUtilser) {
	item, err := kubeUtilsInterface.Get(c.Namespace, c.Name)
	if err != nil {
		msg := "查询失败: " + err.Error()
		c.ReturnData.Message = msg
		c.ReturnData.Status = 400
		logs.Error(nil, msg)
	} else {
		c.ReturnData.Data = make(map[string]interface{})
		c.ReturnData.Data["item"] = item
	}
	r.JSON(200, c.ReturnData)
}

func (c *Info) Delete(r *gin.Context, kubeUtilsInterface kubeutils.KubeUtilser) {
	var gracePeriodSeconds int64
	if c.Force {
		// 强制删除
		var s int64 = 0
		gracePeriodSeconds = s
	}
	err := kubeUtilsInterface.Delete(c.Namespace, c.Name, &gracePeriodSeconds)
	if err != nil {
		msg := "删除失败: " + err.Error()
		c.ReturnData.Message = msg
		c.ReturnData.Status = 400
		logs.Error(nil, msg)
	}
	r.JSON(200, c.ReturnData)
}

func (c *Info) DeleteList(r *gin.Context, kubeUtilsInterface kubeutils.KubeUtilser) {
	var gracePeriodSeconds int64
	if c.Force {
		// 强制删除
		var s int64 = 0
		gracePeriodSeconds = s
	}
	err := kubeUtilsInterface.DeleteList(c.Namespace, c.BasicInfo.DeleteList, &gracePeriodSeconds)
	if err != nil {
		msg := "删除失败: " + err.Error()
		c.ReturnData.Message = msg
		c.ReturnData.Status = 400
		logs.Error(nil, msg)
	}
	r.JSON(200, c.ReturnData)
}

func BasicInit(r *gin.Context, item interface{}) (clientset *kubernetes.Clientset, basicInfo BasicInfo, err error) {
	basicInfo = BasicInfo{}
	basicInfo.Item = item
	// 首先获取请求的类型
	requestMethod := r.Request.Method
	if requestMethod == "GET" {
		err = r.ShouldBindQuery(&basicInfo)
	} else if requestMethod == "POST" {
		err = r.ShouldBindJSON(&basicInfo)
	} else {
		err = errors.New("不支持的请求类型")
	}
	logs.Debug(map[string]interface{}{"basicInfo": basicInfo}, "数据绑定结果")
	if err != nil {
		msg := "请求出错: " + err.Error()
		return nil, basicInfo, errors.New(msg)
	}
	if basicInfo.Namespace == "" {
		basicInfo.Namespace = "default"
	}
	// 获取kubeconfig
	kubeconfig := config.ClusterKubeconfig[basicInfo.ClusterId]
	restConfig, err := clientcmd.RESTConfigFromKubeConfig([]byte(kubeconfig))
	if err != nil {
		msg := "解析kubeconfig错误: " + err.Error()
		return nil, basicInfo, errors.New(msg)
	}
	clientset, err = kubernetes.NewForConfig(restConfig)
	if err != nil {
		msg := "创建clientset失败: " + err.Error()
		return nil, basicInfo, errors.New(msg)
	}
	return clientset, basicInfo, nil
}

// create delete get list update --> 接口
// kubeutils
//
