import { defineStore } from 'pinia'

// 创建一个全局的状态、容器。
// 容器接收两个参数，第一个参数：容器的ID，第二个：容器的内容
// defineStore返回的是一个函数，这个函数按照useXXX去命名
// counter--> useCounter  storeDemo --> useStoreDemo
export const useStoreDemo = defineStore('storeDemo',{
    // 容器的内容
    // state：用来存储全局状态/数据，可以理解为数据配置的位置
    // data
    state: ()=>{
        return {
            msg: "Hello, Pinia"
        }
    },
    // 相当于计算属性
    getters: {},
    // 定义修改数据的方法
    // 相当于methods
    actions: {
        // 接收参数
        changeStoreDemo(value) {
            this.msg = value
        }
    }
})
// 怎么被其他组件使用呢？

export const useIsCollapse = defineStore('isCollapse',{
    // 容器的内容
    // state：用来存储全局状态/数据，可以理解为数据配置的位置
    // data
    state: ()=>{
        return {
            isCollapse: false
        }
    },
    // 相当于计算属性
    getters: {},
    // 定义修改数据的方法
    // 相当于methods
    actions: {
        // 接收参数
        changeIsCollapse() {
            this.isCollapse = !this.isCollapse
        }
    }
})