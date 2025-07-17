<script setup>
import { onBeforeMount, reactive, toRefs,ref } from 'vue';
import { getUserListHandler, deleteUserHandler } from '../../api/user.js';
import { getClusterListHandler as getListHandler, deleteClusterHandler as deleteHandler, getClusterHandler as getDetail  } from '../../api/cluster.js';
import { ElMessage,ElMessageBox  } from 'element-plus'
import { Close,Check } from '@element-plus/icons-vue'
import Add from './Add.vue';
const data = reactive({
  items: [],
  itemForm: {
    id: "",
    displayName: "",
    city: "",
    district: "",
    kubeconfig: ""
  }
})
const loading = ref(true)
const addUserDialog = ref(false)
const defaultMethod = ref("Create")

// 获取列表
const getList = () => {
  loading.value = true
  getListHandler().then((response)=>{
    console.log("response:", response.data.data.items)
    data.items = response.data.data.items
    loading.value = false
    // addUserDialog.value = false
  })
}
// 删除集群
const deleteItem = (row) => {
  console.log("当前行的数据:", row)
  ElMessageBox.confirm(
    '你确定要删除集群: '+ row.id,
    {
      confirmButtonText: '确认',
      cancelButtonText: '取消',
      type: 'warning',
    }
  )
    .then(() => {
      loading.value = true
      deleteHandler(row.id).then((response)=>{
        ElMessage({
            message: response.data.message,
            type: 'success',
        })
        // 用户删除后，自动更新列表
        getList()
        // loading.value = false
      })
    })
    .catch(() => {
      return
    })
}
// 添加
const add = () => {
  data.itemForm = {}
  defaultMethod.value = "Create"
  addUserDialog.value = true
}
const edit = (id) => {
  loading.value = true
  // 1. 获取编辑的集群ID
  console.log("编辑集群:", id)
  // 2. 获取当前集群的配置详情
  getDetail(id).then((response)=>{
    // 3. 查询到数据以后，把数据赋值给我们的itemForm
    console.log("获取到集群详情:", response.data.data.item)
    data.itemForm = response.data.data.item
    // 编辑的dialog给打开
    defaultMethod.value = "Edit"
    addUserDialog.value = true
    loading.value = false
  })
}
onBeforeMount(()=>{
  getList()
})
const closeDialog = () => {
  console.log("close dialog")
  getList()
}
const rollback = () => {
  console.log("close dialog")
  addUserDialog.value = false
  getList()
}
const {items, itemForm} = toRefs(data)
</script>

<template>
    <el-card element-loading-text="玩命加载中..." class="box-card">
      <template #header>
        <div class="card-header">
          <span>集群列表</span>
          <el-button @click="add()" class="button" type="primary">添加</el-button>
        </div>
      </template>
      <!-- 放真实数据的地址，card body -->
      <el-table
      border 
      height="480" 
      :data="items" 
      style="width: 100%"
      >
        <el-table-column fixed prop="" align="center" label="集群ID" width="180">
          <template #default="scope">
            <router-link v-if="scope.row.status == 'Active'" :to="{path: '/api/cluster/detail', query: {'clusterId': scope.row.id}}">{{ scope.row.id }}</router-link>
            <span v-else style="color: red;">{{ scope.row.id }}</span>
          </template>
        </el-table-column>
        <el-table-column prop="displayName" align="center" label="集群别名" width="180" />
        <el-table-column prop="version" align="center" label="版本" width="100" sortable />
        <el-table-column prop="" align="center" label="位置" width="180">
          <template #default="scope">
            {{ scope.row.city }} - {{ scope.row.district }}
          </template>
        </el-table-column>
        <el-table-column prop="" align="center" label="状态" width="80" sortable>
          <template #default="scope">
            <el-icon v-if="scope.row.status == 'Active'" color="green"><Check /></el-icon>
            <el-icon v-else  color="red"><Close /></el-icon>
          </template>
        </el-table-column>
        <el-table-column prop="version" align="center" label="集群规模" width="100" />
        <el-table-column fixed="right" align="center" label="操作">
          <template #default="scope">
            <el-button :disabled="scope.row.status != 'Active'" link type="primary" @click="edit(scope.row.id)">编辑</el-button>
            <el-button link type="warning"  @click="deleteItem(scope.row)">删除</el-button>
          </template>
        </el-table-column>
      </el-table>
      <!-- 尾部的布局 -->
      <div class="auto-space-between" style="margin-top: 10px;">
        <!-- 前面批量操作 -->
        <div></div>
        <!-- 总数的统计 -->
        <div>总数: {{ items.length }}</div>
      </div>
    </el-card>
    <el-dialog
    destroy-on-close 
    @close="closeDialog" 
    v-model="addUserDialog" 
    :title="defaultMethod=='Create'?'添加集群':'更新集群'"
    width="70%">
      <!-- <Add @rollback="getUserList"></Add> -->
      <Add :item-form="itemForm" :method="defaultMethod" @rollback="rollback()"></Add>
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
