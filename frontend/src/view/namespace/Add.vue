<script setup>
import { onBeforeMount, reactive, toRefs,ref, onMounted } from 'vue';
import {addUserHandler, updateUserHandler} from '../../api/user.js'
import { ElMessage  } from 'element-plus'
const data = reactive({
  userForm: {
    // username: "",
    // qq: "",
    // address: ""
  }
})
const props = defineProps({
  // 判断当前是创建还是更新
  method: {
    type: String,
    default: "Create"
  },
  userForm: {
    type: Object
  }
})
const emits = defineEmits(["rollback"])
const userFormRef = ref()
let loading = ref(false)
const {userForm} = toRefs(data)
const clear = ()=> {
  // 实现清空表单
  userFormRef.value.resetFields()
}
const rules = reactive({
    username: [
        { required: true, message: '请输入姓名', trigger: 'blur' },
    ],
    qq: [
        { required: true, message: '请输入QQ号', trigger: 'blur' },
    ],
})
const submit = ()=> {
  userFormRef.value.validate((valid)=>{
    // valid => true
    if(valid){
      loading.value = true
      console.log(loading.value)
      if (props.method == "Create") {
          // 创建用户，异步的
          addUserHandler(data.userForm).then((response)=>{
            // console.log("response:", response)
            ElMessage({
                    message: response.data.message,
                    type: 'success',
            })
            loading.value = false
            // emits('rollback')
          })
      } else {
        // 创建用户，异步的
        updateUserHandler(data.userForm).then((response)=>{
            // console.log("response:", response)
            ElMessage({
                    message: response.data.message,
                    type: 'success',
            })
            loading.value = false
            emits('rollback')
        })
      }

    } else {
      ElMessage({
                message: "请完善表单内容",
                type: 'warning',
            })
    }
  })
  
  
}
// 
onMounted(()=>{
  // 仍有问题，传递的仍然是内存地址
  // data.userForm = props.userForm
  // 把对象转成一个字符串
  // const jsonString = JSON.stringify(props.userForm)
  // 再把字符串转成JSON，最后赋值给data.userForm
  // data.userForm = JSON.parse(jsonString)
  data.userForm = JSON.parse(JSON.stringify(props.userForm))
})
</script>

<template>
  <el-form :rules="rules" ref="userFormRef" :model="userForm" >
      <el-form-item class="form-item" prop="username" label="姓名" label-width="64px">
        <el-input v-model="userForm.username" autocomplete="off" />
      </el-form-item>
      <el-form-item class="form-item" prop="qq" label="QQ" label-width="64px">
        <el-input v-model="userForm.qq" autocomplete="off" />
      </el-form-item>
      <el-form-item class="form-item" prop="address" label="地址" label-width="64px">
        <el-input v-model="userForm.address" autocomplete="off" />
      </el-form-item>
      <span style="display: block; margin-top: 30px;" class="span-button">
        <el-button @click="clear()">清空</el-button>
        <el-button @click="submit()" type="primary">
          {{ props.method=="Create"?"添加":"更新" }}
        </el-button>
      </span>
    </el-form>
</template>

<style scoped>
/* from */
.form-item{
  width: 480px;
  margin: 20px auto;
}
</style>
