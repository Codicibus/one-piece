import { createApp } from 'vue'
import App from './App.vue'

import router from './router'
import '@/router/intercept' // 路由拦截
import pinia from './plugins/pinia'
import Antd from './plugins/antd'

// css全局样式
import './styles/index.less'
import 'nprogress/nprogress.css'

const app = createApp(App)
app.config.productionTip = false

app.use(router)
app.use(pinia)
app.use(Antd)
app.mount('#app')
if (import.meta.env.MODE === 'production') {
	console.log('当前是生产环境')
} else {
	console.log('当前是开发环境')
}
