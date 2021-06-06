import request from '@/utils/request'

const loginStore = {
	id: 'login',
	state: () => ({
		username: '',
		password: ''
	}),
	computed: {},
	actions: {
		userLogin: () => {
			// request.
		}
	}
}
export default loginStore
