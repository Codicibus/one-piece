const state = {
	user: {},
	token: '',
	username: ''
}

const mutations = {
	M_SET_USER(state, user) {
		state.user = { ...user }
	},
	M_SET_TOKEN(state, token) {
		state.token = token
	},
	M_SET_USERNAME(state, username) {
		state.username = username
	}
}

export default {
	namespaced: true,
	state,
	mutations
}
