import router from './index'
import { getToken } from '@/utils/auth'
import { message } from 'ant-design-vue'

// 进度条
import NProgress from 'nprogress'
// NProgress配置右上角loading样式,默认为true
NProgress.configure({ showSpinner: true })

const isExt = path => /https|www/.test(path)

router.beforeEach((to, from, next) => {
	// 开始进度条
	NProgress.start()
	// 阻止外链跳转在路由中跳转
	if (isExt(to.path)) return next(from.path)
	// 提示NotFound
	if (to.name === 'NotFound') return next()
	const token = getToken()
	if (to.path === '/login' || to.path === '/') {
		NProgress.done()
		next()
	} else {
		if (token) return next()
		message.warning('请先登录')
		next('/login')
	}
})

router.afterEach(() => {
	// 完成进度条
	NProgress.done()
})
