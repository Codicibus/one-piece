import { createStore } from 'vuex'
import user from './modules/user'
import getters from './getters'

const state = {
	couter: 0
}

const mutations = {
	add(state) {
		state.couter++
	}
}

const modules = {
	user
}

export default createStore({
	strict: import.meta.env.MODE !== 'production',
	state,
	getters,
	mutations,
	modules
})
