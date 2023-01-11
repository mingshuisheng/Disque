import {createApp} from 'vue'
import App from './App.vue'
import {router} from "./router";
import 'uno.css'
import 'element-plus/theme-chalk/dark/css-vars.css'
import "./index.scss"
import { autoAnimatePlugin } from '@formkit/auto-animate/vue'
// import 'element-plus/dist/index.css'

createApp(App).use(router).use(autoAnimatePlugin).mount('#app')
useTitle('Disque')
