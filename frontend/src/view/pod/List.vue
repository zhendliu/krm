<script setup>
import List from '../components/List.vue';
import ClusterNamespaceSelector from '../components/ClusterNamespaceSelector.vue'
import { getPodListHandler as getListHandler, deletePodHandler as deleteHandler  } from '../../api/pod.js';
import {reactive, computed, ref, toRefs} from 'vue'
import { ElMessage,ElMessageBox  } from 'element-plus'
import CodeMirror  from '../components/CodeMirror.vue'
import { obj2Yaml } from '../../utils/utils';
const data = reactive({
  items: [],
  clusterId: "",
  namespace: "",
  yamlContent: ""
})
// 定义一个查询Pod列表的函数
const getList = () => {
  getListHandler(data.clusterId, data.namespace).then(
    (res) => {
      data.items = res.data.data.items
      console.log("items:",data.items)
    }
  )
}
// 回调函数
const rollback = (clusterId, namespace)=> {
  data.clusterId = clusterId
  data.namespace = namespace
  getList()
}
const search = ref('')
const filterTableData = computed(() =>
  data.items.filter(
    (item) =>
      !search.value ||
      item.metadata.name.toLowerCase().includes(search.value.toLowerCase())
  )
)
// 删除的方法
const deleteItem = (row, force) => {
  console.log("删除:", row)
  ElMessageBox.confirm(
    `你确定要删除命名空间: ${data.namespace} Pod: ${row.metadata.name}吗？`,
    {
      confirmButtonText: '确认',
      cancelButtonText: '取消',
      type: 'warning',
    }
  )
    .then(() => {
      deleteHandler(data.clusterId, data.namespace, row.metadata.name, force).then(
        (res) => {
          ElMessage({
              message: res.data.message,
              type: 'warning',
          })
          getList()
        }
      )
    })
    .catch(() => {
      return
    })

  
}
// 判断pod的重启次数和容器状态
const getPodRestartCount = computed(() => (row) => {
  // console.log("row:", row)
  let restartCount = 0
  let readyCount = 0
  row.status.containerStatuses.forEach(item => {
    restartCount += item.restartCount
    if (item.ready){
      // 表明当前容器已经启动完成
      readyCount += 1
    }
  })
  // 获取当前Pod一共有多少个容器
  const containerTotal = row.status.containerStatuses.length
  // 1/2 2/2
  const readyStatus = `${readyCount}/${containerTotal}`
  return [restartCount, readyStatus]
})
// 展示Pod详情
const showDetail = ref(false)
const detail = (row) => {
  const item = JSON.parse(JSON.stringify(row))
  // 删除managedFields
  delete item.metadata.managedFields
  // 把apiVersion kind补上去
  const itemtemp = {
    "apiVersion": "v1",
    "kind": "Pod",
    "metadata": item.metadata,
    "spec": item.spec
  }
  // 把当前数据转成yaml
  const yamlData = obj2Yaml(itemtemp)
  data.yamlContent = yamlData
  showDetail.value = true
}
const {yamlContent} = toRefs(data)
</script>

<template>
  <List title="Pod列表">
    <template #header-options>
      <ClusterNamespaceSelector @namespaceChangedRollback="rollback">

      </ClusterNamespaceSelector>
      <el-input v-model="search" size="small" placeholder="搜索" style="width: 140px; margin-left: 10px; height: 32px;" />
    </template>
    <template #main>
      <el-table
      border 
      height="480" 
      :data="filterTableData" 
      style="width: 100%"
      :default-sort="{prop: 'metadata.creationTimestamp', order: 'descending'}"
      >
        <el-table-column fixed prop="" align="center" label="名称" width="320">
          <template #default="scope">
            <el-button link type="primary"  @click="detail(scope.row)">{{ scope.row.metadata.name }}</el-button>
          </template>
        </el-table-column>
        <el-table-column prop="metadata.creationTimestamp" align="center" label="创建时间" width="240" sortable />
        <el-table-column prop="metadata.namespace" align="center" label="命名空间" width="240" />
        <el-table-column prop="" align="center" label="状态" width="120" sortable>
          <template #default="scope">
            <span v-if="scope.row.status.phase == 'Running'" style="color: green;">Running</span>
            <span v-else  style="color: red;">{{ scope.row.status.phase }}</span>
          </template>
        </el-table-column>
        <el-table-column prop="" align="center" label="重启次数" width="120">
          <template #default="scope">
            <span v-if="!scope.row.metadata.deletionTimestamp">{{ getPodRestartCount(scope.row)[0] }}</span>
          </template>
        </el-table-column>
        <el-table-column prop="" align="center" label="容器状态" width="120">
          <template #default="scope">
            <span v-if="!scope.row.metadata.deletionTimestamp">{{ getPodRestartCount(scope.row)[1] }}</span>
          </template>
        </el-table-column>
        <el-table-column prop="status.podIP" align="center" label="Pod IP" width="240" />
        <el-table-column prop="spec.nodeName" align="center" label="宿主机主机名" width="240" />
        <el-table-column prop="status.hostIP" align="center" label="宿主机IP" width="240" />
        <el-table-column fixed="right" align="center" label="操作" width="320" >
          <template #default="scope">
            <el-button v-if="!scope.row.metadata.deletionTimestamp" link type="warning" @click="deleteItem(scope.row, false)">删除</el-button>
            <el-button v-else link type="warning" @click="deleteItem(scope.row, true)">强制删除</el-button>
            <el-button link v-if="!scope.row.metadata.deletionTimestamp" type="primary">执行命令</el-button>
            <el-button link v-if="!scope.row.metadata.deletionTimestamp" type="primary">查看日志</el-button>
          </template>
        </el-table-column>
      </el-table>
    </template>
  </List>
  <el-dialog
  v-model="showDetail" 
  width="70%">
  <template #header="{ titleId, titleClass }">
    <div class="my-header">
      <div :id="titleId" :class="titleClass" style="margin: 10px auto 0 auto;">
        <el-button type="primary" title="资源所在的集群" >
          <i class="iconfont icon-kubernetes" style="font-size: 20px;" />
          <span style="font-size: 20px;margin-left: 10px;">{{ data.clusterId }}</span>
        </el-button>
        <el-button type="primary" title="资源所在的命名空间" >
          <i class="iconfont icon-Namespace" style="font-size: 20px;" />
          <span style="font-size: 20px;margin-left: 10px;">{{ data.namespace }}</span>
        </el-button>
      </div>
    </div>
  </template>
    <CodeMirror
    v-model="yamlContent"
    ></CodeMirror>
  </el-dialog>
</template>

<style scoped>
.my-header {
  display: flex;
  flex-direction: row;
  justify-content: space-between;
}
</style>
