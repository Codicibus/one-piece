import axios from 'axios'
import qs from 'qs'
import { getToken } from './auth'
import { message } from 'ant-design-vue'
import NProgress from 'nprogress'

// 请求配置
const request = axios.create({
	timeout: 1000 * 5, // 请求超时
	baseURL: import.meta.env.VITE_BASE_URL, // 基础路径
	responseType: 'json', // 请求响应类型
	responseEncoding: 'utf8', // 编码格式
	/**
	 * 允许在向服务器发送前，修改请求数据
	 * 只能用在 'PUT', 'POST' , 'PATCH' 请求方法
	 */
	transformRequest: [data => qs.stringify(data)]
})

request.interceptors.request.use(
	config => {
		NProgress.start()
		config.url !== '/v1/user/login' &&
			(config.headers['Authorization'] = getToken())
		return config
	},
	error => {
		console.error(error)
		return Promise.reject(error)
	}
)

// axios结束拦截
request.interceptors.response.use(
	response => {
		NProgress.done()
		const msg = response.data.msg
		const status = response.data.code
		const res = response.data.data
		switch (status) {
			case 0:
				message.success(msg || '错误')
				return res
			case 7:
				message.error(msg || '错误')
				return res
			default:
				message.error(msg || '错误')
				return Promise.reject(new Error(msg || '错误'))
		}
	},
	error => {
		NProgress.done()
		console.error('err' + error) // for debug
		message.error(error || '错误')
		return Promise.reject(error)
	}
)

export default request
