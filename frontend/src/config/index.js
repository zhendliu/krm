// 用来放置项目的配置信息
//
console.log("获取环境变量:", import.meta.env)
const BASE_URL = import.meta.env.VITE_BASE_URL
export const API_CONFIG = {
    loginApi: `${BASE_URL}/auth/login`,
    logoutApi: `${BASE_URL}/auth/logout`,
    getUserListApi: `${BASE_URL}/user/list`,
    deleteUserApi: `${BASE_URL}/user/delete`,
    addUserApi: `${BASE_URL}/user/add`,
    updateUserApi: `${BASE_URL}/user/update`,
    // 定义集群资源操作的接口
    getClusterListApi: `${BASE_URL}/cluster/list`,
    addClusterApi: `${BASE_URL}/cluster/add`,
    deleteClusterApi: `${BASE_URL}/cluster/delete`,
    getClusterApi: `${BASE_URL}/cluster/get`,
    updateClusterApi: `${BASE_URL}/cluster/update`,
    // 节点接口
    getNodeListApi: `${BASE_URL}/node/list`,
    updateNodeApi: `${BASE_URL}/node/update`,
    // namespace接口
    getNamespaceListApi: `${BASE_URL}/namespace/list`,
    createNamespaceApi: `${BASE_URL}/namespace/create`,
    deleteNamespaceApi: `${BASE_URL}/namespace/delete`,
    updateNamespaceApi: `${BASE_URL}/namespace/update`,
    // Pod接口
    getPodListApi: `${BASE_URL}/pod/list`,
    deletePodApi: `${BASE_URL}/pod/delete`,
    getPodApi: `${BASE_URL}/pod/get`,
}
export const CONFIG = {
    TOKEN_NAME: "Authorization"
}

