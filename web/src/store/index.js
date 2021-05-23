import { createStore } from 'vuex'
<<<<<<< HEAD
=======
import sidebar from './modules/sidebar'
>>>>>>> 6c89e1c72d17c3c65ff3abcf8517c479a0e5ab16
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
<<<<<<< HEAD
=======
	sidebar,
>>>>>>> 6c89e1c72d17c3c65ff3abcf8517c479a0e5ab16
	user
}

export default createStore({
<<<<<<< HEAD
	strict: import.meta.env.MODE !== 'production',
=======
>>>>>>> 6c89e1c72d17c3c65ff3abcf8517c479a0e5ab16
	state,
	getters,
	mutations,
	modules
})
