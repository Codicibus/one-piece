import request from '@/utils/request'
import { setToken } from '@/utils/auth'
// import translate from '@/utils/translate'
import { defineStore } from 'pinia'
export default defineStore({
	id: 'login',
	state: () => ({}),
	getters: {},
	actions: {
		async userLogin(formData) {
			try {
				const { token } = await request.post('/v1/user/login', formData)
				setToken(token)
			} catch (error) {
				console.error(error.message)
				// const data = await translate(error.message)
				// console.log(data)
			}
		}
	}
})
