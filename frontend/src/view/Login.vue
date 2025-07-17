<script setup>
import { reactive,ref, watch } from 'vue';
import { User,Lock } from '@element-plus/icons-vue'
import {login} from '../api/login.js'
import { ElMessage } from 'element-plus'
import {CONFIG} from '../config/index.js'
import {useRouter} from 'vue-router'
const loginInfo = reactive({
    username: "",
    password: ""
})
const router = useRouter()
const loginRef = ref()
const rules = reactive({
    username: [
        { required: true, message: '请输入用户名', trigger: 'blur' },
    ],
    password: [
        { required: true, message: '请输入密码', trigger: 'blur' },
    ],
})
let loginButtonDisabled = ref(true)
// 监听username和password输入的情况
watch([()=> loginInfo.username,()=>loginInfo.password], ()=>{
  loginRef.value.validate((valid)=>{
    // valid => true
    if(valid){
      // 表单校验成功
      loginButtonDisabled.value = false 
    } else {
      loginButtonDisabled.value = true 
    }
  })
})
// 实现一个登录的接口
const submitForm = () => {
  console.log("loginInfo:", loginInfo)
  login(loginInfo.username, loginInfo.password).then((response)=>{
    console.log("登录response:", response)
    if (response.data.status == 200) {
      // 先拿到token
      const token = response.data.data.token
      window.localStorage.setItem(CONFIG.TOKEN_NAME, token)
      // 登录成功
      ElMessage({
        message: response.data.message,
        type: 'success',
      })
      // 登录成功跳转到 首页
      router.replace("/")
    }
  })
}
</script>

<template>
  <!-- 当我们无法控制某个元素时，可以加一层div去控制 -->
  <div id="login" style="width: 100vw;">
    <el-card class="box-card">
      <h2>后台管理系统</h2>
      <el-form
          ref="loginRef"
          :model="loginInfo"
          :rules="rules"
      >
          <el-form-item prop="username" class="form-item">
              <el-input 
              placeholder="请输入用户名" 
              clearable 
              :prefix-icon="User" 
              v-model="loginInfo.username" 
              />
          </el-form-item>
          <el-form-item prop="password" class="form-item">
              <el-input 
              placeholder="请输入密码"  
              :prefix-icon="Lock" 
              show-password 
              v-model="loginInfo.password" 
              type="password" 
              autocomplete="off" />
          </el-form-item>

          <el-form-item>
          <el-button style="margin: 10px auto 10px auto;" :disabled="loginButtonDisabled"  type="primary" @click="submitForm()"
              >登录</el-button
          >
          </el-form-item>
      </el-form>
    </el-card>
  </div>

</template>

<style scoped>
.text {
  font-size: 14px;
}

.item {
  padding: 18px 0;
}

.box-card {
  width: 480px;
  margin: 0 auto;
}
.form-item{
  width: 240px; 
  margin: 0 auto 20px auto;
}
</style>
