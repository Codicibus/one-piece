import request from '@/utils/request'
import { defineStore } from 'pinia'
import dayjs from 'dayjs'
import relativeTime from 'dayjs/plugin/relativeTime'
import 'dayjs/locale/zh-cn'

export default defineStore({
	id: 'homeMain',
	state: () => ({
		loading: false,
		// 分页设置
		pagingSetup: { page_num: 1, page_size: 5 },
		// 文章
		articles: [],
		// 文章信息
		articleInfo: {}
	}),
	getters: {},
	actions: {
		// 获取文章
		async getArticle() {
			try {
				this.loading = true
				const data = await request.get('/v1/article/get', {
					params: this.pagingSetup
				})
				this.articleInfo = data
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
				this.articles = res
				this.loading = false
			} catch (error) {
				return console.error(error.message)
			}
		}
	}
})
