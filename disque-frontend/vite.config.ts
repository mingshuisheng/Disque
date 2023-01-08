import path from 'path'
import {defineConfig} from 'vite'
import Vue from '@vitejs/plugin-vue'
import Icons from 'unplugin-icons/vite'
import IconsResolver from 'unplugin-icons/resolver'
import AutoImport from 'unplugin-auto-import/vite'
import Components from 'unplugin-vue-components/vite'
import {ElementPlusResolver, VueUseComponentsResolver} from 'unplugin-vue-components/resolvers'
import UnoCSS from 'unocss/vite'
import Inspect from 'vite-plugin-inspect'
import VueMacros from 'unplugin-vue-macros/vite'
import {transformShortVmodel} from '@vue-macros/short-vmodel'
import VueJsx from '@vitejs/plugin-vue-jsx'
import VueRouter from 'unplugin-vue-router/vite'
import {VueRouterAutoImports} from 'unplugin-vue-router'

const pathSrc = path.resolve(__dirname, 'src')


// https://vitejs.dev/config/
export default defineConfig({
  resolve: {
    alias: {
      '@': pathSrc,
    },
  },
  plugins: [
    VueRouter({
      routesFolder: "src/pages",
      extensions: ['.vue'],
      exclude: [],
      dts: path.resolve(pathSrc, 'types', 'typed-router.d.ts'),
      routeBlockLang: 'json5',
      importMode: 'async',
    }),
    AutoImport({
      include: [
        /\.[tj]sx?$/, // .ts, .tsx, .js, .jsx
        /\.vue$/, /\.vue\?vue/, // .vue
        /\.md$/, // .md
      ],
      imports: [
        'vue',
        VueRouterAutoImports,
        {'vue-router/auto': ['useLink']},
        {'@/assets': ['sunIcon', 'unKnownIcon', 'folderIcon', 'moonIcon']},
        '@vueuse/core',
      ],
      defaultExportByFilename: false,
      dirs: [path.resolve(pathSrc, 'composables')],
      resolvers: [
        ElementPlusResolver(),
        IconsResolver(),
      ],
      vueTemplate: true,
      dts: path.resolve(pathSrc, 'types', 'auto-imports.d.ts'),
    }),
    Components({
      resolvers: [
        IconsResolver(),
        ElementPlusResolver(),
        VueUseComponentsResolver()
      ],
      dts: path.resolve(pathSrc, 'types', 'components.d.ts'),
    }),
    Icons({
      autoInstall: true,
    }),
    UnoCSS(),
    Inspect(),
    // DefineOptions(),
    VueMacros({
      setupBlock: true,
      plugins: {
        vue: Vue({
          include: [/\.vue$/, /setup\.[cm]?[jt]sx?$/],
          reactivityTransform: true,
          template: {
            compilerOptions: {
              nodeTransforms: [
                transformShortVmodel({
                  prefix: '$',
                }),
              ],
            },
          },
        }),
        vueJsx: VueJsx(),
      },
      // defineModel: {
      //   unified: false
      // },
    }),
  ],
  server: {
    proxy: {
      "/api": {
        target: "http://localhost:9000",
        rewrite: path => path.replace(/^\/api/, ''),
      }
    }
  }
})
