import Layout from '@/layouts/index.vue'

const routes = [
	{ path: '/' },
	{
		path: '/login',
		component: () => import('@/pages/login/index.vue')
	},
	{
		path: '/admin',
		component: Layout,
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
				name: 'Setting',
				path: 'setting',
				// hidden: true, // 是否展示该路由
				component: () => import('@/pages/admin/setting/index.vue'),
				meta: {
					title: '基础设置',
					icon: 'el-icon-s-home'
				}
			},
			{
				name: 'Article',
				path: 'article',
				// hidden: true, // 是否展示该路由
				component: () => import('@/pages/admin/article/index.vue'),
				meta: {
					title: '文章管理',
					icon: 'el-icon-s-home'
				}
			}
		]
	}
]

export default routes
