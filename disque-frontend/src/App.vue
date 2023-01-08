<script setup lang="ts">

import {isDark} from "./composables";
import {moonIcon, sunIcon} from "./assets";

let router = useRouter();

const items = ref<{ text: string, onClick?(): void }[]>([
  {
    text: '文件',
    onClick: () => router.push("/folder/1")
  },
  {
    text: '传输',
    onClick: () => router.push("/about")
  },
  {
    text: '回收站',
    // onClick: null
  }
])

const themIcon = computed(() => isDark.value ? sunIcon : moonIcon)

</script>

<template>
  <el-container class="common-layout">
    <el-aside class="aside-menu">
      <el-row class="tac">
        <el-col class="menu-container">
          <div class="aside-head">
            <div class="aside-title">网盘</div>
            <img class="theme-icon" :src="themIcon" alt="" @click="toggleDark()">
          </div>
          <el-menu
            class="el-menu-vertical menu-list"
            default-active="0"
          >
            <el-menu-item :key="i" @click="item.onClick" :index="i.toString()" v-for="(item, i) in items">
              <span>{{ item.text }}</span>
            </el-menu-item>
          </el-menu>
        </el-col>
      </el-row>
    </el-aside>
    <el-main>
      <router-view/>
    </el-main>
  </el-container>
</template>

<style lang="scss" scoped>

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
