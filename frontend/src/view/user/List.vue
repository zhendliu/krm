<script setup>
import { onBeforeMount, reactive, toRefs,ref } from 'vue';
import { getUserListHandler, deleteUserHandler } from '../../api/user.js';
import { ElMessage,ElMessageBox  } from 'element-plus'
import Add from './Add.vue';
const data = reactive({
  items: [],
  userForm: {
    username: "123",
    qq: "",
    address: ""
  }
})
const loading = ref(true)
const addUserDialog = ref(false)
const defaultMethod = ref("Create")

// 获取用户列表的方法
const getUserList = () => {
  loading.value = true
  getUserListHandler().then((response)=>{
    console.log("response:", response.data.data.items)
    data.items = response.data.data.items
    loading.value = false
    // addUserDialog.value = false
  })
}
// 删除用户的方法
const deleteUser = (row) => {
  console.log("当前行的数据:", row)
  ElMessageBox.confirm(
    '你确定要删除用户: '+ row.username,
    {
      confirmButtonText: '确认',
      cancelButtonText: '取消',
      type: 'warning',
    }
  )
    .then(() => {
      loading.value = true
      deleteUserHandler(row.id).then((response)=>{
        ElMessage({
            message: response.data.message,
            type: 'success',
        })
        // 用户删除后，自动更新列表
        getUserList()
        // loading.value = false
      })
    })
    .catch(() => {
      return
    })
}
// 添加用户的方法
const addUser = () => {
  data.userForm = {}
  defaultMethod.value = "Create"
  addUserDialog.value = true
}
const editUser = (row) => {
  console.log("row:", row)
  defaultMethod.value = "Edit"
  addUserDialog.value = true
  data.userForm = row
}
onBeforeMount(()=>{
  getUserList()
})
const closeDialog = () => {
  console.log("close dialog")
  getUserList()
}
const rollback = () => {
  console.log("close dialog")
  addUserDialog.value = false
  getUserList()
}
const {items, userForm} = toRefs(data)
</script>

<template>
    <el-card element-loading-text="玩命加载中..." class="box-card">
      <template #header>
        <div class="card-header">
          <span>用户列表</span>
          <el-button @click="addUser()" class="button" text>添加</el-button>
        </div>
      </template>
      <!-- 放真实数据的地址，card body -->
      <el-table
      border 
      height="480" 
      :data="items" 
      style="width: 100%"
      >
        <el-table-column sortable fixed prop="username" label="用户名" width="180" />
        <el-table-column sortable    prop="id" label="ID" width="180" />
        <el-table-column prop="qq" label="QQ号" width="180" />
        <el-table-column prop="address" label="地址" width="180" />
        <el-table-column fixed="right" label="操作">
          <template #default="scope">
            <el-button link type="primary" size="small" @click="editUser(scope.row)">编辑</el-button>
            <el-button link type="warning" size="small"  @click="deleteUser(scope.row)">删除</el-button>
          </template>
        </el-table-column>
      </el-table>
    </el-card>
    <el-dialog
    destroy-on-close 
    @close="closeDialog" 
    v-model="addUserDialog" 
    :title="defaultMethod=='Create'?'添加用户':'更新用户'" 
    width="40%">
      <!-- <Add @rollback="getUserList"></Add> -->
      <Add :user-form="userForm" :method="defaultMethod" @rollback="rollback()"></Add>
    </el-dialog>

</template>

<style scoped>
.card-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.text {
  font-size: 14px;
}

.item {
  margin-bottom: 18px;
}
</style>
