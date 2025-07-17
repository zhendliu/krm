// 取出来调用后端接口的地址
import { API_CONFIG } from "../config/index.js";
import request from '../api/index.js'
// 在很多位置，都可能用到同一个接口
export const getUserListHandler = () => {
    return request(API_CONFIG.getUserListApi, {}, "get", "2000")
}
export const deleteUserHandler = (id) => {
    console.log("删除用户:", id)
    return request(API_CONFIG.deleteUserApi, {id}, "get", "1000")
}

// 直接接收一个对象
export const addUserHandler = (userForm) => {
    console.log("添加用户:", userForm)
    return request(API_CONFIG.addUserApi, userForm, "post", "2000")
}

export const updateUserHandler = (userForm) => {
    console.log("更新用户:", userForm)
    return request(API_CONFIG.updateUserApi, userForm, "post", "2000")
}