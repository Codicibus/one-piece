import request from '@/utils/request'
import { setToken } from '@/utils/auth'

const loginStore = {
	id: 'login',
	state: () => ({
		loading: false
	}),
	getters: {},
	actions: {
		userLogin: async formData => {
			const { token } = await request.post('/v1/user/login', formData)
			setToken(token)
		}
	}
}
export default loginStore
