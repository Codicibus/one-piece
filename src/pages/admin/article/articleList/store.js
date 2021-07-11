import request from '@/utils/request'
import { columns } from './mock'
// import { setToken } from '@/utils/auth'
import { defineStore } from 'pinia'
import dayjs from 'dayjs'
import relativeTime from 'dayjs/plugin/relativeTime'
import 'dayjs/locale/zh-cn'

export default defineStore({
	id: 'articleList',
	state: () => ({
		// 抽屉状态
		visible: false,
		loading: false,
		// 选择后的文章hash
		selectedRowKeys: [],
		// 编辑模式
		editingMode: '',
		// 表单数据
		formData: {
			title: '', // 文章标题
			content_hash: '', // 文章内容唯一标识
			content: '', // 文章内容
			category: '', // 文章分类
			background_visible: false, // 背景图状态
			background_random: false, // 随机背景图状态
			background_pic: '', // 文章背景图地址
			created_at: '', // 文章创建时间
			updated_at: '' // 文章创建时间
		},
		// 分页设置
		pagingSetup: { page_num: 1, page_size: 10 },
		// 文章
		articlesInfo: {},
		columns,
		fileList: [] // 导入文件列表
	}),
	getters: {
		getImageList() {
			let image = this.formData.background_pic
			if (image) {
				return [
					{
						uid: 0,
						url: 'https://api.amujun.com/' + this.formData.background_pic
					}
				]
			}
			return []
		},
		getFileList() {}
	},
	actions: {
		// 重置表单
		resetForm() {
			this.formData = {
				title: '', // 文章标题
				content_hash: '', // 文章唯一标识
				content: '', // 文章内容
				category: '', // 文章分类
				background_visible: false, // 背景图状态
				background_random: false, // 随机背景图状态
				background_pic: '' // 文章背景图地址
			}
		},
		// 获取文章
		async getArticle() {
			try {
				this.loading = true
				const data = await request.get('/v1/article/get', {
					params: this.pagingSetup
				})
				const res = data.articles.map(item => {
					dayjs.extend(relativeTime).locale('zh-cn')
					let nowCreateTime = dayjs(item.updated_at).format('YYYY/MM/DD')
					let nowUpdateTime = dayjs(item.updated_at).fromNow()
					return {
						...item,
						name: item.title,
						created_at: nowCreateTime,
						updated_at: nowUpdateTime
					}
				})
				data.articles = res
				this.articlesInfo = data
				this.loading = false
			} catch (error) {
				return console.error(error.message)
			}
		},
		// 新增文章
		async addArticle(formData) {
			try {
				this.loading = true
				await request.post('/v1/article/post', formData)
				this.loading = false
			} catch (error) {
				return console.error(error.message)
			}
		},
		// 编辑文章
		async editArticle(formData) {
			try {
				this.loading = true
				await request.post('/v1/article/update', {
					...formData,
					old_article_hash: formData.content_hash
				})
				this.loading = false
			} catch (error) {
				return console.error(error.message)
			}
		},
		// 删除文章
		async deleteArticle(content_hash) {
			try {
				this.loading = true
				await request.post('/v1/article/delete', { content_hash })
				this.loading = false
			} catch (error) {
				return console.error(error.message)
			}
		}
	}
})
