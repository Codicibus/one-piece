import Layout from '@/layouts/index.vue'
import home from '@/pages/home/index.vue'
import NotFound from '@/pages/error/404.vue'

const routes = [
	{ path: '/:pathMatch(.*)*', name: 'NotFound', component: NotFound },
	{ path: '/', component: home },
	{
		name: 'Login',
		path: '/login',
		component: () => import('@/pages/login/index.vue')
	},
	{
		name: 'Admin',
		path: '/admin',
		component: Layout,
		redirect: '/admin/dashboard',
		children: [
			{
				name: 'Dashboard',
				path: 'dashboard',
				// hidden: true, // 是否展示该路由
				component: () => import('@/pages/admin/dashboard/index.vue'),
				meta: {
					title: '仪表盘',
					icon: 'el-icon-s-home'
					// activeMenu: '/admin/dashboard' // 控制默认激活展示的路由
				}
			},
			{
				name: 'ArticleList',
				path: 'articleList',
				// hidden: true, // 是否展示该路由
				component: () => import('@article/articleList/index.vue'),
				meta: {
					title: '文章列表',
					icon: 'el-icon-s-home'
					// activeMenu: '/admin/dashboard' // 控制默认激活展示的路由
				}
			},
			{
				name: 'ArticleClassify',
				path: 'articleClassify',
				// hidden: true, // 是否展示该路由
				component: () => import('@article/articleClassify/index.vue'),
				meta: {
					title: '文章分类',
					icon: 'el-icon-s-home'
				}
			},
			{
				name: 'ImagePool',
				path: 'imagePool',
				// hidden: true, // 是否展示该路由
				component: () => import('@article/imagePool/index.vue'),
				meta: {
					title: '图片池',
					icon: 'el-icon-s-home'
				}
			}
		]
	}
]

export default routes
