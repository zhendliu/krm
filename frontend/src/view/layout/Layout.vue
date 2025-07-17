<script setup>
import Aside from './components/Aside.vue';
import Footer from './components/Footer.vue';
import Header from './components/Header.vue';
import Main from './components/Main.vue';
</script>

<template>
  <div class="common-layout">
    <el-container class="el-container">
      <Aside></Aside>
      <el-container class="el-container__el-container">
        <Header></Header>
        <Main></Main>
        <Footer></Footer>
      </el-container>
    </el-container>
  </div>
</template>

<style lang="less" scoped>
.el-container{
  /* 100%: 只是铺满当前的元素，也就是它会受限于父元素 */
  /* vh: 不会受限于父元素 */
  height: 100vh;
  width: 100vw;
  .el-container__el-container{
    /* 网页变成弹性布局 */
    display: flex; 
    flex-direction: column;
  }
}

</style>
