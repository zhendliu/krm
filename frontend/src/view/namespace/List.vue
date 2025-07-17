<script setup>
import { onBeforeMount, reactive, toRefs,ref,computed  } from 'vue';
import { deleteNamespaceHandler as deleteHandler, getNamespaceListHandler as getListHandler, createNamespaceHandler as createHandler  } from '../../api/namespace.js';
import { getClusterListHandler  } from '../../api/cluster.js';
import {useRoute} from 'vue-router'
import { ElMessage  } from 'element-plus'
import Edit from './Edit.vue'
import Detail from './Detail.vue'
const route = useRoute()
const data = reactive({
  items: [],
  clusterList: [],
  clusterId: "",
  editItem: {},
  editName: "",
  detailItem: {},
  detailName: "",
  createName: "",
  deleteName: "",
  deleteNameConfirm: ""
})
const loading = ref(true)
const editDialog = ref(false)
const detailDialog = ref(false)
const createNamespaceDialog = ref(false)
const deleteNamespaceDialog = ref(false)
// 获取列表
const getList = (clusterId) => {
  console.log("查询集群:", clusterId)
  loading.value = true
  editDialog.value = false
  getListHandler(clusterId).then((response)=>{
    console.log("response:", response.data.data.items)
    data.items = response.data.data.items
    loading.value = false
    // editDialog.value = false
  })
}
// 编辑
const edit = (row) => {
  loading.value = true
  console.log("编辑Namespace:", row)
  data.editItem = row
  data.editName = row.metadata.name
  editDialog.value = true
  loading.value = false
}
// 删除namespace
const deleteNamespace = (row) => {
  console.log("删除namespace:", row.metadata.name)
  data.deleteName = row.metadata.name
  deleteNamespaceDialog.value = true
}
const submitDeleteNamespace = () => {
  console.log("删除Namespace:", data.deleteName)
  if (data.deleteName != data.deleteNameConfirm) {
    ElMessage({
        message: "删除失败: 确定删除的名称不正确",
        type: 'warning',
    })
    return
  }
  // loading.value = true
  // 调用创建namespace的接口
  deleteHandler(data.clusterId, data.deleteName).then((response)=>{
    ElMessage({
            message: response.data.message,
            type: 'success',
    })
    // 关闭loading
    // loading.value = false
    // 查询新列表
    getList(data.clusterId)
    // 关闭弹窗
    deleteNamespaceDialog.value = false
    data.deleteNameConfirm = ""
  })
}
// 展示数据的方法
const detail = (row)=> {
  data.detailItem = row
  data.detailName = row.metadata.name
  detailDialog.value = true
}
const getClusterList = async () => {
  await getClusterListHandler().then(
    (res) => {
      data.clusterList = res.data.data.items
    }
  )
}
const createNamespace = ()=> {
  createNamespaceDialog.value = true
}
const submitCreateNamespace = () => {
  console.log("创建namespace:", data.createName)
  // loading.value = true
  // 调用创建namespace的接口
  createHandler(data.clusterId, data.createName).then((response)=>{
    ElMessage({
            message: response.data.message,
            type: 'success',
    })
    // 关闭loading
    // loading.value = false
    // 查询新列表
    getList(data.clusterId)
    // 关闭弹窗
    createNamespaceDialog.value = false
    data.createName = ""
  })
}

onBeforeMount(async ()=>{
  // 查询列表之前，要知道查询哪个集群的节点列表
  // 1. 查询到集群列表
  // 2. 如果没有指定集群，则给一个默认值，进行获取集群列表
  // 3. 如果指定了集群，就直接查询当前集群的节点列表
  // 4. 切换了集群之后，应该重新获取节点列表
  // 5. 集群不可用，则不可选
  // 1. 获取集群列表
  await getClusterList()
  // 2. 定义一个默认的集群
  const defaultClusterSelected = data.clusterList[0].id
  console.log("默认的集群是:",defaultClusterSelected)
  // 3. 获取当前请求是否携带了集群ID的参数
  const curClusterId = route.query.clusterId // curClusterId有可能是空的
  data.clusterId = curClusterId?curClusterId:defaultClusterSelected
  getList(data.clusterId)
})
const closeDialog = () => {
  console.log("close dialog")
  getList(data.clusterId)
}
const rollback = () => {
  console.log("close dialog")
  editDialog.value = false
  getList()
}
const filterTableData = computed(() =>
  data.items.filter(
    (item) =>
      !search.value ||
      item.metadata.name.toLowerCase().includes(search.value.toLowerCase())
  )
)
const search = ref('')
const { clusterId,
        clusterList, 
        editItem, 
        editName,
        detailItem, 
        detailName,
        createName,
        deleteName,
        deleteNameConfirm
        } = toRefs(data)
</script>

<template>
    <el-card element-loading-text="玩命加载中..." class="box-card">
      <template #header>
        <div class="card-header">
          <div>
            <span>命名空间列表</span>
          </div>
          <div class="card-header">
            <el-select
            filterable
            v-model="clusterId" 
            placeholder="选择集群" 
            @change="getList(clusterId)"
            style="width: 180px;"
            >
              <el-option
                v-for="item in clusterList"
                :key="item.id"
                :label="item.displayName"
                :value="item.id"
                :disabled="item.status == 'InActive'"
              />
            </el-select>
            <el-input v-model="search" size="small" placeholder="搜索" style="width: 140px; margin-left: 10px; height: 32px;" />
            <el-button type="primary" style="margin-left: 10px;" @click="createNamespace">创建</el-button>
          </div>
          
        </div>
      </template>
      <!-- 放真实数据的地址，card body -->
      <el-table
      border 
      height="480" 
      :data="filterTableData" 
      style="width: 100%"
      >
        <el-table-column fixed prop="" align="center" label="名称" width="480">
          <template #default="scope">
            <el-button link type="primary"  @click="detail(scope.row)">{{ scope.row.metadata.name }}</el-button>
          </template>
        </el-table-column>
        <el-table-column prop="metadata.creationTimestamp" align="center" label="创建时间" width="240" sortable />
        <el-table-column prop="" align="center" label="状态" width="120" sortable>
          <template #default="scope">
            <span v-if="scope.row.status.phase == 'Active'" style="color: green;">Active</span>
            <span v-else  style="color: red;">{{ scope.row.status.phase }}</span>
          </template>
        </el-table-column>
        <el-table-column fixed="right" align="center" label="操作" >
          <template #default="scope">
            <el-button link type="primary" @click="edit(scope.row)">资源复制</el-button>
            <el-button link type="primary" @click="edit(scope.row)">编辑</el-button>
            <el-button link type="warning" @click="deleteNamespace(scope.row)">删除</el-button>
          </template>
        </el-table-column>
      </el-table>
      <!-- 尾部的布局 -->
      <div class="auto-space-between" style="margin-top: 10px;">
        <!-- 总数的统计 -->
        <div></div>
        <div>总数: {{ filterTableData.length }}</div>
      </div>
    </el-card>
    <!-- 编辑namespace -->
    <el-dialog
    destroy-on-close 
    @close="closeDialog" 
    v-model="editDialog" 
    :title="'正在编辑集群:' + clusterId + '  命名空间:' + editName"
    width="70%">
      <Edit @rollback="getList" :cluster-id="data.clusterId" :item-form="editItem"></Edit>
    </el-dialog>
    <!-- 详情 -->
    <el-dialog
    destroy-on-close 
    v-model="detailDialog" 
    :title="'集群:' + clusterId + '  节点:' + detailName"
    width="70%">
      <Detail :item="detailItem"></Detail>
    </el-dialog>
    <!-- 创建namespace -->
    <el-dialog
    destroy-on-close 
    v-model="createNamespaceDialog" 
    title="创建命名空间"
    width="30%">
      <el-input v-model.trim="createName" placeholder="请输入Namespace名称" />
      <el-button style="margin-top: 20px;" type="primary" @click="submitCreateNamespace">创建</el-button>
    </el-dialog>
    <!-- 删除namespace -->
    <el-dialog
    destroy-on-close 
    v-model="deleteNamespaceDialog" 
    :title="'删除命名空间:' + deleteName"
    width="30%">
      <el-input v-model.trim="deleteNameConfirm" placeholder="请确认删除的名称" />
      <el-button style="margin-top: 20px;" type="warning" @click="submitDeleteNamespace">删除</el-button>
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
