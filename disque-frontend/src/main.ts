import {createApp} from 'vue'
import App from './App.vue'
import {router} from "./router";
import 'uno.css'
import 'element-plus/theme-chalk/dark/css-vars.css'
import "./index.scss"
// import 'element-plus/dist/index.css'

createApp(App).use(router).mount('#app')

