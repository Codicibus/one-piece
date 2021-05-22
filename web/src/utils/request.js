import axios from 'axios'
import qs from 'qs'
import { getToken } from './auth'
import { message } from 'ant-design-vue'
import NProgress from 'nprogress'

// 请求配置
let request = axios.create({
	timeout: 1000 * 30, // 请求30s超时
	baseURL: import.meta.VITE_BASE_URL, // 基础路径
	responseType: 'json', // 请求响应类型
	responseEncoding: 'utf8', // 编码格式
	// withCredentials: true, // 跨域请求时要使用凭证
	/**
	 * 允许在向服务器发送前，修改请求数据
	 * 只能用在 'PUT', 'POST' , 'PATCH' 请求方法
	 */
	transformRequest: [data => qs.stringify(data)]
})

request.interceptors.request.use(
	config => {
		NProgress.start()
		config.headers['Authorization'] = getToken()
		return config
	},
	error => {
		console.log(error)
		return Promise.reject(error)
	}
)

// axios结束拦截
request.interceptors.response.use(
	response => {
		NProgress.done()
		const msg = response.data.meta.msg
		const status = response.data.meta.status
		const res = response.data.data
		switch (status) {
			case 200:
				message.success(msg)
				return res
			case 201:
				message.warning(msg)
				return res
			case 400:
				message.error(msg)
				break
			default:
				message({
					message: msg || '错误',
					type: 'error',
					duration: 2000
				})
				return Promise.reject(new Error(msg || '错误'))
		}
	},
	error => {
		NProgress.done()
		console.log('err' + error) // for debug
		ElMessage({
			message: error || '请求超时',
			type: 'error',
			duration: 2000
		})
		return Promise.reject(error)
	}
)

/**
 * @param {String} url // 请求路径
 * @param {object} params // 请求参数
 * @param {object} config // 请求配置
 */

let get = (url, params = {}, config = {}) => {
	return new Promise((resolve, reject) => {
		res.get(url, params, config)
			.then(res => {
				if (res.status === 200) return resolve(res)
			})
			.catch(err => reject(err))
	})
}

let post = (url, params = {}, config = {}) => {
	return new Promise((resolve, reject) => {
		res.post(url, params, config)
			.then(res => {
				if (res.status === 200) return resolve(res)
			})
			.catch(err => reject(err))
	})
}

export default {
	get,
	post
}
