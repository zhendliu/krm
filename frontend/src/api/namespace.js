// 取出来调用后端接口的地址
import { API_CONFIG } from "../config/index.js";
import request from '../api/index.js'
// 在很多位置，都可能用到同一个接口
export const getNamespaceListHandler = (clusterId) => {
    return request(API_CONFIG.getNamespaceListApi, {clusterId}, "get", "2000")
}

export const createNamespaceHandler = (clusterId, name) => {
    return request(API_CONFIG.createNamespaceApi, {clusterId, name}, "post", "2000")
}

export const deleteNamespaceHandler = (clusterId, name) => {
    return request(API_CONFIG.deleteNamespaceApi, {clusterId, name}, "get", "2000")
}

export const updateNamespaceHandler = (clusterId, name, item) => {
    console.log("更新Namespace:", clusterId, name, item)
    return request(API_CONFIG.updateNamespaceApi, {clusterId, name, item}, "post", "2000")
}

