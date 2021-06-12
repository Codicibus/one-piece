import axios from 'axios'
import qs from 'qs'
import md5 from 'md5'

// 请求配置
const http = axios.create({
	timeout: 1000 * 3, // 请求超时
	baseURL: import.meta.env.VITE_TRANSLATE_URL, // 基础路径
	transformRequest: [data => qs.stringify(data)]
})

const path = '/api/trans/vip/translate'
const from = 'en'
const to = 'zh'
const appid = '20200731000530374'
const translate = q => {
	const salt = Date.now()
	const sign = md5(appid + q + salt + 'm0sg8AlMb3JxAEJ5uCZY')
	http.post(path, { q, from, to, appid, salt, sign })
}

export default translate
