import { bytesToGB } from '@/utils/util'
import { defineStore } from 'pinia'

export default defineStore({
	id: 'dashboard',
	state: () => ({
		loading: false,
		serveDate: {}
	}),
	getters: {},
	actions: {}
})
