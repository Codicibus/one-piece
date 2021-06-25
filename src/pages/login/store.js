import request from '@/utils/request'
import { setToken } from '@/utils/auth'
import { defineStore } from 'pinia'

export default defineStore({
	id: 'login',
	state: () => ({
		loading: false
	}),
	getters: {},
	actions: {
		async userLogin(formData) {
			try {
				const { token } = await request.post('/v1/user/login', formData)
				setToken(token)
			} catch (error) {
				return
			}
		}
	}
})
