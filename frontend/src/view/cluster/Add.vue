<script setup>
import { onBeforeMount, reactive, toRefs,ref, onMounted } from 'vue';
// import {addUserHandler, updateUserHandler} from '../../api/user.js'
import {addClusterHandler as addHandler, updateClusterHandler as updateHandler} from '../../api/cluster.js'
import { ElMessage  } from 'element-plus'
const data = reactive({
  itemForm: {
  }
})
const props = defineProps({
  // 判断当前是创建还是更新
  method: {
    type: String,
    default: "Create"
  },
  itemForm: {
    type: Object
  }
})
const emits = defineEmits(["rollback"])
const itemFormRef = ref()
let loading = ref(false)
const {itemForm} = toRefs(data)
const clear = ()=> {
  // 实现清空表单
  itemFormRef.value.resetFields()
}
const rules = reactive({
    id: [
        { required: true, message: '请输入集群ID', trigger: 'blur' },
    ],
    displayName: [
        { required: true, message: '请输入集群别名', trigger: 'blur' },
    ],
    city: [
        { required: true, message: '请输入集群位置', trigger: 'blur' },
    ],
    kubeconfig: [
        { required: true, message: '请输入集群的kubeconfig配置', trigger: 'blur' },
    ],
})
const submit = ()=> {
  itemFormRef.value.validate((valid)=>{
    // valid => true
    if(valid){
      // loading.value = true
      // console.log(loading.value)
      if (props.method == "Create") {
          // 创建用户，异步的
          addHandler(data.itemForm).then((response)=>{
            // console.log("response:", response)
            ElMessage({
                    message: response.data.message,
                    type: 'success',
            })
            // loading.value = false
            // emits('rollback')
          })
      } else {
        // 创建用户，异步的
        updateHandler(data.itemForm).then((response)=>{
            // console.log("response:", response)
            ElMessage({
                    message: response.data.message,
                    type: 'success',
            })
            // loading.value = false
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
  // data.itemForm = props.itemForm
  // 把对象转成一个字符串
  // const jsonString = JSON.stringify(props.itemForm)
  // 再把字符串转成JSON，最后赋值给data.itemForm
  // data.itemForm = JSON.parse(jsonString)
  data.itemForm = JSON.parse(JSON.stringify(props.itemForm))
})
</script>

<template>
  <el-form 
  :rules="rules" 
  ref="itemFormRef" 
  :model="itemForm" 
  label-width="140px"
  >
    <div class="cluster-info">
      <!-- 集群配置基本信息 -->
      <div>
        <el-form-item class="form-item" prop="id" label="集群ID" required>
          <el-input v-model="itemForm.id" autocomplete="off" placeholder="请输入集群ID" :disabled="props.method=='Edit'" />
        </el-form-item>
        <el-form-item class="form-item" prop="displayName" label="集群名称" required>
          <el-input v-model="itemForm.displayName" autocomplete="off" placeholder="请输入集群别名" />
        </el-form-item>
        <el-form-item class="form-item" label="位置" required>
          <el-col :span="10">
            <el-form-item prop="city" required>
              <el-input placeholder="所在城市" v-model="itemForm.city" autocomplete="off" />
            </el-form-item>
          </el-col>
          <el-col :span="4">-</el-col>
          <el-col :span="10">
            <el-form-item  prop="district" >
              <el-input placeholder="所在区/县" v-model="itemForm.district" autocomplete="off" />
            </el-form-item>
          </el-col>
        </el-form-item>
      </div>
      <!-- kubeconfig -->
      <div>
        <el-form-item prop="kubeconfig" label-width="120px" style="width: 600px;" required>
          <el-input type="textarea" :rows="6" v-model="itemForm.kubeconfig" autocomplete="off" placeholder="请输入集群kubeconfig配置" />
        </el-form-item>
      </div>
    </div>

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
  width: 420px;
  margin: 0 auto 20px 0;
}
.cluster-info{
  display: flex;
  flex-direction: row;
}
</style>
