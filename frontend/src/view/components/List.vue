<!-- 封装通用的列表组件 -->
<script setup>
const props = defineProps({
  // 判断当前是创建还是更新
  title: {
    type: String,
    default: ""
  },
})
</script>

<template>
    <el-card class="box-card">
      <template #header>
        <div class="card-header">
          <div>
            <span>{{props.title}}</span>
          </div>
          <div class="card-header">
            <!-- 这是改成插槽，放置header的操作组件 -->
            <slot name="header-options"></slot>
          </div>
          
        </div>
      </template>
      <!-- 放真实数据的地址，card body -->
      <!-- 这里改成插槽，放置主要内容 -->
      <slot name="main"></slot>
      <!-- 尾部的布局 -->
      <div class="auto-space-between" style="margin-top: 10px;">
        <!-- 总数的统计 -->
        <!-- 改成插槽，放置左侧的功能 -->
        <div>
            <slot name="footer-left"></slot>
        </div>
        <!-- 改成插槽，放置右侧的功能 -->
        <div>
            <slot name="footer-right"></slot>
        </div>
      </div>
    </el-card>
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
