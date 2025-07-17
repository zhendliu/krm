<script setup>
import { onBeforeMount, reactive, toRefs,ref,computed  } from 'vue';
import { getNodeListHandler as getListHandler, getClusterListHandler  } from '../../api/cluster.js';
import {useRoute} from 'vue-router'
import Edit from './Edit.vue'
import Detail from './Detail.vue'
const route = useRoute()
const data = reactive({
  items: [],
  clusterList: [],
  clusterId: "",
  editItem: {},
  editNodeName: "",
  detailItem: {},
  detailNodeName: "",
  createName: "",
})
const loading = ref(true)
const editDialog = ref(false)
const detailDialog = ref(false)
const value1 = ref(false)

// 获取列表
const getList = (clusterId) => {
  console.log("查询集群:", clusterId)
  loading.value = true
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
  console.log("编辑节点:", row)
  data.editItem = row
  data.editNodeName = row.metadata.name
  editDialog.value = true
  loading.value = false
}
// 展示数据的方法
const detail = (row)=> {
  data.detailItem = row
  data.detailNodeName = row.metadata.name
  detailDialog.value = true
}
const getClusterList = async () => {
  await getClusterListHandler().then(
    (res) => {
      data.clusterList = res.data.data.items
    }
  )
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
      item.metadata.name.toLowerCase().includes(search.value.toLowerCase()) ||
      item.status.addresses[0].address.toLowerCase().includes(search.value.toLowerCase())
  )
)
const search = ref('')
const {clusterId,clusterList, editItem, editNodeName,detailItem, detailNodeName} = toRefs(data)
</script>

<template>
    <el-card element-loading-text="玩命加载中..." class="box-card">
      <template #header>
        <div class="card-header">
          <div>
            <span>节点列表</span>
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
        <el-table-column fixed prop="" align="center" label="主机名" >
          <template #default="scope">
            <el-button link type="primary"  @click="detail(scope.row)">{{ scope.row.metadata.name }}</el-button>
          </template>
        </el-table-column>
        <el-table-column prop="status.addresses[0].address" align="center" label="IP" width="180" />
        <!-- 角色稍等处理 -->
        <el-table-column prop="metadata.name" align="center" label="角色" width="100" sortable />
        <!-- 节点状态稍等处理 -->
        <el-table-column prop="" align="center" label="状态" width="80" sortable>
          <template #default="scope">
            <span v-if="scope.row.status.conditions[scope.row.status.conditions.length-1].status == 'True'" style="color: green;">Ready</span>
            <span v-else  style="color: red;">NotReady</span>
          </template>
        </el-table-column>
        <el-table-column prop="" align="center" label="禁止调度" width="180">
          <template #default="scope">
            <el-switch v-model="value1" />
          </template>
        </el-table-column>
        <el-table-column  prop="" align="center" label="驱逐" width="180">
          <template #default="scope">
            <el-switch v-model="value1" />
          </template>
        </el-table-column>
        <el-table-column fixed="right" align="center" label="操作" width="180">
          <template #default="scope">
            <el-button link type="primary" @click="edit(scope.row)">编辑</el-button>
          </template>
        </el-table-column>
      </el-table>
      <!-- 尾部的布局 -->
      <div class="auto-space-between" style="margin-top: 10px;">
        <!-- 前面批量操作 -->
        <div>批量操作</div>
        <!-- 总数的统计 -->
        <div>总数: {{ filterTableData.length }}</div>
      </div>
    </el-card>
    <el-dialog
    destroy-on-close 
    @close="closeDialog" 
    v-model="editDialog" 
    :title="'正在编辑集群:' + clusterId + '  节点:' + editNodeName"
    width="70%">
      <Edit @rollback="getList" :cluster-id="data.clusterId" :item-form="editItem"></Edit>
    </el-dialog>
    <el-dialog
    destroy-on-close 
    v-model="detailDialog" 
    :title="'集群:' + clusterId + '  节点:' + detailNodeName"
    width="70%">
      <Detail :item="detailItem"></Detail>
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
