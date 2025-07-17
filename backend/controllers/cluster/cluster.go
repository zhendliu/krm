package cluster

import (
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

type ClusterInfo struct {
	Id          string `json:"id"`
	DisplayName string `json:"displayName"`
	City        string `json:"city"`
	District    string `json:"district"`
}

// 定义一个结构体，用于描述集群的状态
type ClusterStatus struct {
	ClusterInfo
	Version string `json:"version"`
	Status  string `json:"status"`
}

// 定义一个结构体，用于描述创建集群所用的配置信息
type ClusterConfig struct {
	ClusterInfo
	Kubeconfig string `json:"kubeconfig"`
}

// 结构体的方法，用于判断集群的状态
func (c *ClusterConfig) getClusterStatus() (ClusterStatus, error) {
	// 判断集群是否是正常
	clusterStatus := ClusterStatus{}
	clusterStatus.ClusterInfo = c.ClusterInfo
	// 创建一个clientset
	restConfig, err := clientcmd.RESTConfigFromKubeConfig([]byte(c.Kubeconfig))
	if err != nil {
		return clusterStatus, err
	}
	clientset, err := kubernetes.NewForConfig(restConfig)
	if err != nil {
		return clusterStatus, err
	}
	serverVersion, err := clientset.Discovery().ServerVersion()
	if err != nil {
		return clusterStatus, err
	}
	clusterVersion := serverVersion.String()
	clusterStatus.Version = clusterVersion
	clusterStatus.Status = "Active"
	return clusterStatus, nil
}
