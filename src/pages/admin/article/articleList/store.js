import request from '@/utils/request'
import { setToken } from '@/utils/auth'
import { defineStore } from 'pinia'

export default defineStore({
	id: 'articleList',
	state: () => ({
		visible: false,
		loading: false,
		selectedRows: [],
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
