import router from './index'
import { getToken } from '@/utils/auth'
// 进度条
import NProgress from 'nprogress'
// NProgress配置右上角loading样式,默认为true
NProgress.configure({ showSpinner: false })

const isExt = path => /https|www/.test(path)

router.beforeEach((to, from, next) => {
	// 开始进度条
	NProgress.start()
	// 阻止外链跳转在路由中跳转
	if (isExt(to.path)) return next(from.path)
	// 提示NotFound
	if (to.name === 'NotFound') return next()
	const token = getToken()
	if (token) {
		if (to.path === '/login') return next(from.path)
		next()
	} else {
		if (to.path.indexOf('/admin') !== -1) return next('/login')
		next()
	}
})

router.afterEach(() => NProgress.done())
