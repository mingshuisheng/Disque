<script setup lang='ts'>

import { isDark } from './composables'
import type { MenuItem } from './types/MenuItem'
import { Sunny, Moon } from '@element-plus/icons-vue'

let router = useRouter()
let route = useRoute()

const items = ref<MenuItem[]>([
  {
    text: '文件',
    onClick: () => {
      if(!route.path.startsWith('/folder')){
        router.push('/folder/1')
      }
    }
  },
  {
    text: '传输',
    // onClick: () => router.push("/about")
    subItems: [
      {
        text: '上传',
        onClick: () => router.push('/transmission/uploadPage')
      },
      {
        text: '下载',
        onClick: () => router.push('/transmission/downloadPage')
      },
      {
        text: '已完成',
        onClick: () => router.push('/transmission/transmissionCompleted')
      }
    ]
  },
  {
    text: '回收站'
    // onClick: null
  }
])

const theme = ref(unref(isDark))
const onThemeChange = toggleDark


</script>

<template>
  <el-container class='common-layout'>
    <el-aside class='aside-menu'>
      <el-row class='tac'>
        <el-col class='menu-container'>
          <div class='aside-head'>
            <div class='aside-title'>网盘</div>
            <el-switch
              size="large"
              class='theme-switch'
              $='theme'
              @change='onThemeChange'
              inline-prompt
              :active-icon='Moon'
              :inactive-icon='Sunny' />
            <!--            <img class="theme-icon" :src="themIcon" alt="" @click="toggleDark()">-->
          </div>
          <el-menu
            class='el-menu-vertical menu-list'
            default-active='0'
          >
            <template v-for='(item, i) in items'>
              <el-menu-item v-if='!item.subItems' :key='i' @click='item.onClick' :index='i.toString()'>
                <span>{{ item.text }}</span>
              </el-menu-item>
              <el-sub-menu v-if='!!item.subItems' :key='i' @click='item.onClick' :index='i.toString()'>
                <template #title>
                  <span>{{ item.text }}</span>
                </template>
                <el-menu-item v-for='(subItem, subIndex) in item.subItems' :key='subIndex' @click='subItem.onClick'
                              :index='`${i}-${subIndex}`'>
                  <span>{{ subItem.text }}</span>
                </el-menu-item>
              </el-sub-menu>
            </template>
          </el-menu>
        </el-col>
      </el-row>
    </el-aside>
    <el-main v-auto-animate>
      <router-view />
    </el-main>
  </el-container>
</template>

<style lang='scss' scoped>

.common-layout {
  width: 100%;
  height: 100%;

  .aside-menu {
    width: 200px;

    .tac {
      padding-left: 20px;
      height: 100%;

      .menu-container {
        display: flex;
        flex-direction: column;

        .menu-list {
          flex: 1
        }
      }
    }

    .aside-head {
      display: flex;
      justify-content: space-between;
      align-items: center;
      height: 40px;
      padding-right: 20px;

      .theme-switch {
        --el-switch-off-color: #2c2c2c;
        --el-switch-on-color: #f2f2f2;
        --el-switch-border-color: var(--el-border-color-lighter);
        --el-color-white: var(--el-bg-color);
      }

      .aside-title {
        font-size: var(--el-font-size-extra-large);
      }

      .theme-icon {
        width: 25px;
        height: 25px;
        margin-right: 20px;

        &:hover {
          cursor: pointer;
        }
      }
    }

  }
}

</style>
