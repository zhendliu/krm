<script setup>
import { ArrowLeft,SwitchButton,ArrowRight } from '@element-plus/icons-vue'
import {logoutHandler} from '../../../api/login.js'
import {CONFIG} from '../../../config/index.js'
import { ElMessage,ElMessageBox  } from 'element-plus'
import {useRouter} from 'vue-router'
import {useIsCollapse} from '../../../store/index.js'
import {storeToRefs} from 'pinia'
const router = useRouter()
const logout = ()=>{
    // 1. 后端没有存储token，直接删除local storage的Authorization
    // 2. 后端存储了token，需要先调用后端的退出接口，然后正常处理后，再去删除Authorization
    ElMessageBox.confirm(
    '你确定要退出吗？',
    {
      confirmButtonText: '确认',
      cancelButtonText: '取消',
      type: 'warning',
    }
  )
    .then(() => {
      logoutHandler().then((response)=>{
        ElMessage({
            message: response.data.message,
            type: 'success',
        })
        // 删除本地token
        window.localStorage.removeItem(CONFIG.TOKEN_NAME)
        router.replace("/login")
      })
    })
    .catch(() => {
      return
    })
}
const collapseChange = ()=> {
    const isCollapse = useIsCollapse()
    isCollapse.changeIsCollapse()
}
const {isCollapse} = storeToRefs(useIsCollapse())
</script>

<template>

    <el-header class="el-header">
        <el-button @click="collapseChange()">
            <el-icon v-show="!isCollapse" :size="20"><ArrowLeft /></el-icon>
            <el-icon v-show="isCollapse" :size="20"><ArrowRight /></el-icon>
            <!-- <el-icon  :size="25"><Expand /></el-icon> -->
        </el-button>
        <el-button @click="logout()">
            <el-icon :size="20"><SwitchButton /></el-icon>
        </el-button>
    </el-header>

</template>

<style lang="less" scoped>
.el-header{
    height: 50px;
    display: flex;
    /* 均匀分布 */
    justify-content: space-between; 
    align-items: center;
    button {
        border: 0;
    }
}

</style>
