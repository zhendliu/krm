// 取出来调用后端接口的地址
import { API_CONFIG } from "../config/index.js";
import request from '../api/index.js'
// 在很多位置，都可能用到同一个接口
export const getClusterListHandler = () => {
    return request(API_CONFIG.getClusterListApi, {}, "get", "2000")
}
export const getClusterHandler = (id) => {
    return request(API_CONFIG.getClusterApi, {id}, "get", "2000")
}
export const deleteClusterHandler = (id) => {
    console.log("删除集群:", id)
    return request(API_CONFIG.deleteClusterApi, {id}, "get", "2000")
}

// 直接接收一个对象
export const addClusterHandler = (data) => {
    console.log("添加集群:", data)
    return request(API_CONFIG.addClusterApi, data, "post", "2000")
}

export const updateClusterHandler = (data) => {
    console.log("更新集群:", data)
    return request(API_CONFIG.updateClusterApi, data, "post", "2000")
}

// 节点
export const getNodeListHandler = (clusterId) => {
    console.log("查看节点:", )
    return request(API_CONFIG.getNodeListApi, {clusterId}, "get", "2000")
}
