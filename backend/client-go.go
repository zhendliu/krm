// 项目的总入口
package main

import (
	"context"
	"encoding/json"
	"fmt"
	"krm-backend/utils/logs"

	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func mainbak() {
	// 1. 初始化config实例
	config, err := clientcmd.BuildConfigFromFlags("", "./config/meta.kubeconfig")
	if err != nil {
		panic(err.Error())
	}
	// 2. 创建客户端工具，clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}
	// 3. 操作集群
	// kubectl api-resources | grep ingress
	// clientset.AppsV1()
	// clientset.NetworkingV1()
	// clientset.StorageV1()
	pods, err := clientset.CoreV1().Pods("kube-system").List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		logs.Error(nil, "查询Pod列表失败")
	} else {
		fmt.Printf("There are %d pods in the cluster\n", len(pods.Items))

	}
	// 查询deployment列表
	deployments, _ := clientset.AppsV1().Deployments("").List(context.TODO(), metav1.ListOptions{})
	deploymentItems := deployments.Items
	for _, deploy := range deploymentItems {
		fmt.Printf("当前资源的名字是: %s, namespace: %s\n", deploy.Name, deploy.Namespace)
	}
	// clientset.CoreV1().Namespaces().List(context.TODO(), metav1.ListOptions{})
	// Get 方法
	podDetail, err := clientset.CoreV1().Pods("default").Get(context.TODO(), "nginx-77b4fdf86c-vhrfn", metav1.GetOptions{})
	if err != nil {
		fmt.Println("查询详情失败")
	} else {
		fmt.Println("查询到详情:", podDetail)
		fmt.Println("第一个容器的镜像是:", podDetail.Spec.Containers[0].Image)
	}
	// 获取namespace详情
	nsDetail, _ := clientset.CoreV1().Namespaces().Get(context.TODO(), "kube-system", metav1.GetOptions{})
	fmt.Println("namespace详情:", nsDetail)
	// 更新操作
	// 获取deployment 并修改
	// deployDetail, _ := clientset.AppsV1().Deployments("default").Get(context.TODO(), "nginx", metav1.GetOptions{})
	// fmt.Println("查询到deployment,名字是: ", deployDetail.Name)
	// // 获取当前的label
	// // labels := deployDetail.Labels
	// // labels["newlabel"] = "newlabel-value"
	// deployDetail.Annotations["new-anno"] = "new-anno"
	// // 我通过创建的新的labels进行修改，还用赋值给deployDetail吗？
	// // 不需要的。
	// newReplicas := int32(3)
	// deployDetail.Spec.Replicas = &newReplicas
	// deployDetail.Spec.Template.Spec.Containers[0].Image = "nginx:1.22.1"
	// _, err = clientset.AppsV1().Deployments("default").Update(context.TODO(), deployDetail, metav1.UpdateOptions{})
	// if err != nil {
	// 	// 说明更新失败
	// 	fmt.Println("更新失败:", err.Error())
	// }
	// // 删除一个资源
	// err = clientset.CoreV1().Pods("default").Delete(context.TODO(), "nginx-85954cf5d7-xjffh", metav1.DeleteOptions{})
	// if err != nil {
	// 	// 删除失败
	// 	fmt.Println("删除失败:", err.Error())
	// }
	// err = clientset.AppsV1().Deployments("default").Delete(context.TODO(), "nginx", metav1.DeleteOptions{})
	// if err != nil {
	// 	// 删除失败
	// 	fmt.Println("删除Deploy失败:", err.Error())
	// }
	var newNamespace corev1.Namespace
	newNamespace.Name = "test1"
	_, err = clientset.CoreV1().Namespaces().Create(context.TODO(), &newNamespace, metav1.CreateOptions{})
	if err != nil {
		fmt.Println("创建namespace失败:", err.Error())
	}
	// 创建一个deployment
	var newDeployment appsv1.Deployment
	newDeployment.Name = "nginx"
	newDeployment.Namespace = "test1"
	label := make(map[string]string)
	label["app"] = "nginx"
	label["version"] = "v1"

	// 此处更改的是deployment的labels
	newDeployment.Labels = label
	// 此处更改的是selector下面的参数
	newDeployment.Spec.Selector = &metav1.LabelSelector{}
	newDeployment.Spec.Selector.MatchLabels = label
	// 此处更改的是pod的标签
	newDeployment.Spec.Template.Labels = label
	// 先去声明一个pod，然后赋值给newdeployment
	// 创建容器
	var containers []corev1.Container
	// newDeployment.Spec.Template.Spec.Containers[0].Name = "nginx"
	// newDeployment.Spec.Template.Spec.Containers[0].Image = "nginx"
	// 先声明一个容器
	var container corev1.Container
	container.Image = "redis"
	container.Name = "redis"
	containers = append(containers, container)
	container.Image = "nginx"
	container.Name = "nginx"
	containers = append(containers, container)
	newDeployment.Spec.Template.Spec.Containers = containers
	_, err = clientset.AppsV1().Deployments("test1").Create(context.TODO(), &newDeployment, metav1.CreateOptions{})
	if err != nil {
		fmt.Println("创建Deployment失败:", err.Error())
	}
	// 通过json串去创建k8s的资源
	deployJson := `{
		"kind": "Deployment",
		"apiVersion": "apps/v1",
		"metadata": {
			"name": "redis",
			"creationTimestamp": null,
			"labels": {
				"app": "redis"
			}
		},
		"spec": {
			"replicas": 1,
			"selector": {
				"matchLabels": {
					"app": "redis"
				}
			},
			"template": {
				"metadata": {
					"creationTimestamp": null,
					"labels": {
						"app": "redis"
					}
				},
				"spec": {
					"containers": [
						{
							"name": "redis",
							"image": "redis",
							"resources": {}
						}
					]
				}
			},
			"strategy": {}
		},
		"status": {}
	}`
	var newDeployment2 appsv1.Deployment
	// 把json串转成struct
	err = json.Unmarshal([]byte(deployJson), &newDeployment2)
	if err != nil {
		fmt.Println("json转struct失败:", err.Error())
	}
	fmt.Println("json转struct之后Deployment配置详情:", newDeployment2)
	_, err = clientset.AppsV1().Deployments("default").Create(context.TODO(), &newDeployment2, metav1.CreateOptions{})
	if err != nil {
		fmt.Println("创建Deployment失败:", err.Error())
	}
}
