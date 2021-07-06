import request from '@/utils/request'
import { columns } from './mock'
// import { setToken } from '@/utils/auth'
import { defineStore } from 'pinia'

export default defineStore({
	id: 'articleList',
	state: () => ({
		// 抽屉状态
		visible: false,
		loading: false,
		// 选择后的信息
		selectedRows: [],
		// 表单数据
		formData: {
			title: '', // 文章标题
			content_hash: '', // 文章唯一标识
			content: '', // 文章内容
			category: [], // 文章分类
			background_visible: false, // 背景图状态
			background_random: false, // 随机背景图状态
			background_pic: [] // 文章背景图地址
		},
		// 分页设置
		pagingSetup: { page_num: 1, page_size: 5 },
		// 文章
		articlesInfo: {},
		columns
	}),
	getters: {
		// 重置表单
		resetForm: state => {
			state.formData = {
				title: '', // 文章标题
				content_hash: '', // 文章唯一标识
				content: '', // 文章内容
				category: [], // 文章分类
				background_visible: false, // 背景图状态
				background_random: false, // 随机背景图状态
				background_pic: [] // 文章背景图地址
			}
		}
	},
	actions: {
		// 获取文章
		async getArticle() {
			try {
				this.loading = true
				const data = await request.get('/v1/article/get', {
					params: this.pagingSetup
				})
				this.articlesInfo = data
				this.loading = false
			} catch (error) {
				return console.error(error.message)
			}
		}
	}
})
