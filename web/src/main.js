import { createApp } from 'vue'
import App from './App.vue'

import router from './router'
import '@/router/intercept' // 路由拦截

// import store from './store'

import { createPinia } from 'pinia'

// 引入ant-design-vue
import Antd from 'ant-design-vue'
import 'ant-design-vue/dist/antd.css'

// css全局样式
import './styles/index.less'
import './styles/antd.less'
import 'nprogress/nprogress.css'

const app = createApp(App)
app.config.productionTip = false

app.use(router)
// app.use(store)
app.use(createPinia())
app.use(Antd)
app.mount('#app')

if (import.meta.env.MODE === 'production') {
	console.log('当前是生产环境')
} else {
	console.log('当前是开发环境')
}
