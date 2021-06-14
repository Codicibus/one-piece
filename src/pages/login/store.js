import request from '@/utils/request'
import { setToken } from '@/utils/auth'
// import translate from '@/utils/translate'
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
				// console.error(error.message)
				// const data = await translate(error.message)
				// console.log(data)
			}
		}
	}
})
