// 取出来调用后端接口的地址
import { API_CONFIG } from "../config/index.js";
import request from '../api/index.js'
// 在很多位置，都可能用到同一个接口
export const login = (username, password) => {
    /*
        {
            username: username,
            password: password
        }
        ==> {username, password}
    */ 
    return request(API_CONFIG.loginApi, {username, password}, "post", "2000")
}

export const logoutHandler = () => {
    /*
        {
            username: username,
            password: password
        }
        ==> {username, password}
    */ 
    return request(API_CONFIG.logoutApi, {}, "get", "2000")
}