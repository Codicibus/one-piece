import { createApp } from 'vue'
import App from './App.vue'

import router from './router'
import store from './store'

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
app.use(store)
app.use(Antd)
app.mount('#app')

if (import.meta.env.MODE === 'production') {
	console.log('当前是生产环境')
} else {
	console.log('当前是开发环境')
}
