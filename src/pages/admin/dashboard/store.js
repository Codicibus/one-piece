import request from '@/utils/request'
import { setToken } from '@/utils/auth'
// import translate from '@/utils/translate'
import { defineStore } from 'pinia'

export default defineStore({
	id: 'dashboard',
	state: () => ({
		loading: false
	}),
	getters: {},
	actions: {}
})
