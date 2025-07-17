// 封装axios
/*
    1. 封装可以简化我们接口调用的代码，可以去掉一些重复的代码
    2. 换掉Axios也是比较简单的
    3. 
*/
import axios from 'axios'
import { CONFIG } from '../config/index.js';
import {ref} from 'vue';
import router from '../router/index.js'
import { ElMessage, ElLoading } from 'element-plus'
let loadingInstance = ref()
// 添加请求拦截器
axios.interceptors.request.use(function (config) {
    loadingInstance.value = ElLoading.service({
        text: '玩命加载中...',
        background: 'rgba(0,0,0,0.5)'
    })
    // 在发送请求之前做些什么
    console.log("请求拦截器:", config)
    // 添加时间戳在请求当中
    if (config.method == "get") {
        // 解决缓存
        let timeStamp = (new Date()).getTime()
        console.log("timeStamp:", timeStamp)
        // 有灭有params
        if (config.params) {
            // 说明本身是有params配置的
            config.params.timeStamp = timeStamp
        } else {
            config.params = {
                timeStamp: timeStamp
            }
        }
    }
    // 取出来token的值
    let TokenValue = ""
    try {
        TokenValue = window.localStorage.getItem(CONFIG.TOKEN_NAME)
    } catch (error) {
        TokenValue = ""
    }
    if (TokenValue == "" || TokenValue == null) {
        config.headers[CONFIG.TOKEN_NAME] = ""
    } else {
        config.headers[CONFIG.TOKEN_NAME] = TokenValue
    }
    return config;
  }, function (error) {
    // 对请求错误做些什么
    ElMessage({
        type: "error",
        message: "请求错误: " + error.message
    })
    return Promise.reject(error);
});
// 添加响应拦截器
axios.interceptors.response.use(function (response) {
    // 2xx 范围内的状态码都会触发该函数。
    // 对响应数据做点什么
    // 判断相应状态码是否正常
    console.log("response:", response)
    if (response.data.status === 200) {
        loadingInstance.value.close()
        return Promise.resolve(response)
    } else if (response.data.status === 401) {
        // 说明我们的token已经失效了
        ElMessage({
            type: "warning",
            message: response.data.message
        })
        // 删除token
        window.localStorage.removeItem(CONFIG.TOKEN_NAME)
        // 调整到登录页
        // 如果取到当前的path
        // route
        router.currentRoute.path != "/login" && router.replace("/login")
        loadingInstance.value.close()
    }

    return response;
  }, function (error) {
    // 超出 2xx 范围的状态码都会触发该函数。
    // 对响应错误做点什么
    ElMessage({
        type: "error",
        message: "请求错误: " + error.message
    })
    loadingInstance.value.close()
    return Promise.reject(error);
  });
const request = (url = '', data = {}, method = "get", timeout = 5000) => {
    console.log("使用封装函数去处理请求")
    return new Promise((resolve, reject)=>{
        console.log("使用axios请求接口")
        // GET POST 
        const methodLower = method.toLowerCase() 
        if (methodLower === 'get') {
            axios({
                method: methodLower,
                params: data,
                timeout: timeout,
                url: url,
            }).then((response)=>{
                // 能正常拿到数据
                resolve(response)
            }).catch((error)=>{
                reject(error)
            })
        } else if (methodLower === "post") {
            axios({
                method: methodLower,
                data: data,
                timeout: timeout,
                url: url,
            }).then((response)=>{
                // 能正常拿到数据
                resolve(response)
            }).catch((error)=>{
                reject(error)
            })
        }
    })
}

export default request