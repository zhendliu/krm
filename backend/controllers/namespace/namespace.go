package namespace

// 如果对集群中的namespace进行管理
// 1. 首先要知道对哪个集群中的namespace进行管理
// 2. kubeconfig，secret，clusterId --> secret --> data --> kubeconfig --> clientset
// 		kubeconfig 存储在一个变量当中，map[string]string map[string]ClientSet
// 3. 创建一个clientset
// 4. 操作namespace
// {clusterId:xxx, name: xxxx, namespace:xxx}
