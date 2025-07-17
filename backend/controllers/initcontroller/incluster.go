package initcontroller

import (
	"context"
	"krm-backend/config"
	"krm-backend/utils/logs"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func metadataInit() {
	logs.Debug(nil, "初始化元数据命名空间")
	// 1. 初始化config实例
	restConfig, err := clientcmd.BuildConfigFromFlags("", "config/meta.kubeconfig")
	if err != nil {
		logs.Error(map[string]interface{}{"msg": err.Error()}, "inCluster kubeconfig加载失败")
		panic(err.Error())
	}
	// 2. 创建客户端工具，clientset
	clientset, err := kubernetes.NewForConfig(restConfig)
	if err != nil {
		logs.Error(map[string]interface{}{"msg": err.Error()}, "inCluster客户端创建失败")
		panic(err.Error())
	}
	config.InClusterClientSet = clientset
	inClusterVersion, _ := clientset.Discovery().ServerVersion()
	// 1. 检查元数据命名空间是否存在
	_, err = clientset.CoreV1().Namespaces().Get(context.TODO(), config.MetadataNamespace, metav1.GetOptions{})
	if err != nil {
		// 不存在元数据命名空间
		// 创建元数据命名空间
		var metadataNamespace corev1.Namespace
		metadataNamespace.Name = config.MetadataNamespace
		_, err = clientset.CoreV1().Namespaces().Create(context.TODO(), &metadataNamespace, metav1.CreateOptions{})
		if err != nil {
			logs.Error(map[string]interface{}{"msg": err.Error()}, "元数据命名空间创建失败")
			panic(err.Error())
		}

		logs.Info(map[string]interface{}{"Namespace": config.MetadataNamespace, "inCluster版本": inClusterVersion.String()}, "元数据命名空间创建成功")
	} else {
		// 已经存在namespace
		logs.Info(map[string]interface{}{"Namespace": config.MetadataNamespace, "inCluster版本": inClusterVersion.String()}, "元数据命名空间已存在")
	}
	// 初始化ClusterKubeconfig
	config.ClusterKubeconfig = make(map[string]string)
	// 查询当前已经存在集群配置
	listOptions := metav1.ListOptions{
		LabelSelector: config.ClusterConfigSecretLabelKey + "=" + config.ClusterConfigSecretLabelValue,
	}
	secretList, _ := config.InClusterClientSet.CoreV1().Secrets(config.MetadataNamespace).List(context.TODO(), listOptions)
	for _, secret := range secretList.Items {
		// clusterId
		clusterId := secret.Name
		kubeconfig := secret.Data["kubeconfig"]
		config.ClusterKubeconfig[clusterId] = string(kubeconfig)
	}
	// 打印一个debug
	// fmt.Println("当前集群配置:", config.ClusterKubeconfig)
}
