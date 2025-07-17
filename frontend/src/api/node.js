// 取出来调用后端接口的地址
import { API_CONFIG } from "../config/index.js";
import request from '../api/index.js'
// 节点
export const updateNodeHandler = (clusterId, name, item) => {
    console.log("更新节点:", clusterId, name, item)
    return request(API_CONFIG.updateNodeApi, {clusterId, name, item}, "post", "2000")
}