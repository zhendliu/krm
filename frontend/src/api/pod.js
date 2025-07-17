// 取出来调用后端接口的地址
import { API_CONFIG } from "../config/index.js";
import request from '../api/index.js'
// 在很多位置，都可能用到同一个接口
export const getPodListHandler = (clusterId, namespace) => {
    return request(API_CONFIG.getPodListApi, {clusterId, namespace}, "get", "2000")
}

export const deletePodHandler = (clusterId, namespace, name, force) => {
    return request(API_CONFIG.deletePodApi, {clusterId, namespace, name, force}, "post", "2000")
}



