<script setup>
import { MoonNight } from '@element-plus/icons-vue'
import { MENU_CONFIG } from '../../../config/menu.js';
import {useIsCollapse} from '../../../store/index.js'
import {storeToRefs} from 'pinia'
import '../../../assets/iconfont/iconfont.css'
import '../../../assets/iconfont/customIconfont.css'
import { ref } from 'vue';
// const isCollapse = ref(true)
const {isCollapse} = storeToRefs(useIsCollapse())
</script>

<template>
    <el-aside class="el-aside" :style="{width: isCollapse?'65px':'240px'}" style="border-right: 1px solid #cccccc;">
      <!-- logo -->
      <div class="aside-logo">
        <router-link to="/">
          <el-button text style="font-size: 25px;">
            <el-icon style="margin-right: 10px; padding-left: 14px;"><MoonNight /></el-icon>
            <span v-show="!isCollapse">KRM</span>
          </el-button>
        </router-link>
      </div>
      <!-- menu -->
      <div id="menu">
        <el-menu
          :default-active="$route.path"
          class="el-menu-vertical-demo"
          router
          :collapse="isCollapse"
          style="border: none;"
          :collapse-transition="false"
        >
          <el-sub-menu v-for="menu in MENU_CONFIG" :key="menu.index" :index="menu.index">
            <template #title>
              <el-icon><span style="font-size: 20px;" :class="menu.icon"></span></el-icon>
              <span>{{menu.title}}</span>
            </template>
            <!-- 判断是否有子菜单 -->
            <template v-if="menu.subMenu">
              <el-sub-menu v-for="subMenu in menu.subMenu" :key="subMenu.index" :index="subMenu.index">
                <template #title>
                  <el-icon><span  style="font-size: 20px;"  :class="subMenu.icon"></span></el-icon>
                  <span>{{subMenu.title}}</span>
                </template>
                <el-menu-item v-for="item in subMenu.items" :key="item.index" :index="item.index">
                <template #title>{{ item.title }}</template>
              </el-menu-item>
              </el-sub-menu>
            </template>
            <template v-else>
              <el-menu-item v-for="item in menu.items" :key="item.index" :index="item.index">
                <template #title>{{ item.title }}</template>
              </el-menu-item>
            </template>
          </el-sub-menu>
          <!-- <el-sub-menu index="/user">
            <template #title>
              <el-icon><location /></el-icon>
              <span>用户管理</span>
            </template>
            <el-menu-item index="/user/add">
              <el-icon><icon-menu /></el-icon>
              <template #title>添加用户</template>
            </el-menu-item>
          </el-sub-menu>
          <el-menu-item index="3">
            <el-icon><icon-menu /></el-icon>
            <template #title>Navigator Three</template>
          </el-menu-item> -->
        </el-menu>
      </div>
    </el-aside>
</template>

<style lang="less" scoped>
.aside-logo{
  height: 50px;
  button {
    width: 100%;
    height: 100%;
  }
}
// 使用aside实现折叠效果
.el-aside {
  transition: width 0.15s;
  -webkit-transition: width 0.15s;
  -moz-transition: width 0.15s;
  -webkit-transition: width 0.15s;
  -o-transition: width 0.15s;
}
</style>
