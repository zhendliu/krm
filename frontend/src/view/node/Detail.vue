<script setup>
import { computed } from 'vue';

// 接收参数
const props = defineProps({
  item: {
    type: Object
  }
})
const getNodeRole = computed(() => {
  console.log("labels:", props.item.metadata.labels)
  // 获取labels所有的key
  const keyList = Object.keys(props.item.metadata.labels)
  if (keyList.length == 0) {
    return "工作节点"
  }
  // node-role.kubernetes.io/control-plane / node-role.kubernetes.io/master / node-role.kubernetes.io/controlplane
  if (keyList.includes('node-role.kubernetes.io/control-plane') || keyList.includes('node-role.kubernetes.io/controlplane') || keyList.includes('node-role.kubernetes.io/master')) {
    return "控制节点"
  }
  return "工作节点"
})
</script>

<template>
  <div>
    <el-descriptions 
    size="large"
    :column="3" 
    border>
      <el-descriptions-item label="主机名">
        <span>{{ props.item.metadata.name }}</span>
      </el-descriptions-item>
      <el-descriptions-item label="IP">
        <span>{{ props.item.status.addresses[0].address }}</span>
      </el-descriptions-item>
      <el-descriptions-item label="角色">
        <el-tag size="large">{{ getNodeRole }}</el-tag>
      </el-descriptions-item>
      <!-- 第二行数据展示 -->
      <el-descriptions-item label="节点状态">
        <el-tag size="large" 
        :type="props.item.status.conditions[props.item.status.conditions.length-1].status?'success':'danger'"
        >{{ props.item.status.conditions[props.item.status.conditions.length-1].reason }}</el-tag>
      </el-descriptions-item>
      <el-descriptions-item label="Pod">
        <el-button link type="primary">查看</el-button>
      </el-descriptions-item>
      <el-descriptions-item label="操作系统">
        {{ props.item.status.nodeInfo.osImage }}
      </el-descriptions-item>
      <el-descriptions-item label="Runtime">
        {{ props.item.status.nodeInfo.containerRuntimeVersion }}
      </el-descriptions-item>
      <el-descriptions-item label="CPU架构">
        {{ props.item.status.nodeInfo.architecture }}
      </el-descriptions-item>
      <el-descriptions-item label="组件版本">
        <span>
          Kubelet: {{ props.item.status.nodeInfo.kubeletVersion }}
        </span>
        <span style="margin-left: 10px;">
          KubeProxy:{{ props.item.status.nodeInfo.kubeProxyVersion }}
        </span>
      </el-descriptions-item>
      <el-descriptions-item label="系统内核">
        {{ props.item.status.nodeInfo.kernelVersion }}
      </el-descriptions-item>
    </el-descriptions>
  </div>
</template>

<style scoped>

</style>
