import { createStore } from 'vuex'
import sidebar from './modules/sidebar'
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
	sidebar,
	user
}

export default createStore({
	state,
	getters,
	mutations,
	modules
})
