import request from '@/utils/request'
import { setToken } from '@/utils/auth'
// import translate from '@/utils/translate'
import { defineStore } from 'pinia'

export default defineStore({
	id: 'articleList',
	state: () => ({
		visible: false,
		loading: false,
		form: {
			articleTitle: '',
			articleClassify: {},
			fixedState: false,
			randomState: false,
			fileList: [],
			articleContent: ''
		}
	}),
	getters: {},
	actions: {}
})
