<script setup>
import { onMounted,reactive,toRefs  } from 'vue';
import { getNamespaceListHandler } from '../../api/namespace.js';
import { getClusterListHandler  } from '../../api/cluster.js';
import {useRoute} from 'vue-router'
const route = useRoute()
// 接收参数
const props = defineProps({
  // 判断当前是创建还是更新
  showNamespace: {
    type: Boolean,
    default: true
  },
})
const data = reactive({
  clusterList: [],
  namespaceList: [],
  clusterId: "",
  namespace: ""
})
// 定义两个事件，用来处理当集群和namespace产生变化后，应该做哪些处理
const emits = defineEmits(['clusterChangedRollback', 'namespaceChangedRollback'])
// 定义namespace发生变化需要做的事情
const namespaceChangedRollbackHandler = () => {
    emits('namespaceChangedRollback', data.clusterId, data.namespace)
}
// 定义集群发生变化需要做的事情
const clusterChangedRollbackHandler = () => {
    emits('clusterChangedRollback', data.clusterId)
}

// 查询集群列表
const getNamespaceList = async () => {
    console.log("查询集群:", data.clusterId)
    await getNamespaceListHandler(data.clusterId).then((response)=>{
        console.log("response:", response.data.data.items)
        data.namespaceList = response.data.data.items
    })
    // 获取当前请求是否携带了namespace的参数
    const curNamespace = route.query.namespace // curClusterId有可能是空的
    data.namespace = curNamespace?curNamespace:"default"
    // 回调
    namespaceChangedRollbackHandler()
}

const getClusterList = async () => {
  await getClusterListHandler().then(
    (res) => {
      data.clusterList = res.data.data.items
    }
  )
}
const clusterChanged = () => {
    if (props.showNamespace) {
        getNamespaceList()
    } else {
        clusterChangedRollbackHandler()
    }
}
const namespaceChanged = () => {
    namespaceChangedRollbackHandler()
}
onMounted(async ()=>{
    await getClusterList()
    // 2. 定义一个默认的集群
    const defaultClusterSelected = data.clusterList[0].id
    console.log("默认的集群是:",defaultClusterSelected)
    // 3. 获取当前请求是否携带了集群ID的参数
    const curClusterId = route.query.clusterId // curClusterId有可能是空的
    data.clusterId = curClusterId?curClusterId:defaultClusterSelected
    // 查询namespace
    clusterChanged()
})
const { clusterId,
        clusterList, 
        namespace,
        namespaceList
        } = toRefs(data)
</script>

<template>
    <el-select
    filterable
    v-model="clusterId" 
    placeholder="选择集群" 
    @change="clusterChanged()"
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
    <el-select
    filterable
    v-if="props.showNamespace"
    v-model="namespace" 
    placeholder="选择Namespace" 
    @change="namespaceChanged()"
    style="width: 180px;margin-left: 10px;"
    >
        <el-option
        v-for="item in namespaceList"
        :key="item.metadata.name"
        :label="item.metadata.name"
        :value="item.metadata.name"
        />
    </el-select>
</template>

<style scoped>

</style>
