import axios from 'axios'
import qs from 'qs'
import md5 from 'md5'

let baseURL
// 判断环境变量
if (import.meta.env.MODE === 'production') {
	baseURL = import.meta.env.VITE_TRANSLATE_URL
}

// 请求配置
const http = axios.create({
	timeout: 1000 * 3, // 请求超时
	baseURL, // 基础路径
	transformRequest: [data => qs.stringify(data)]
})
const path = '/api/trans/vip/translate'
const from = 'en'
const to = 'zh'
const appid = '20200731000530374'
const translate = async q => {
	const salt = Date.now()
	let sign = md5(appid + q + salt + 'm0sg8AlMb3JxAEJ5uCZY')
	const { data } = await http.post(path, { q, from, to, appid, salt, sign })
	const error_msg = data.error_msg === 'Invalid Access Limit' ? '请降低翻译频率，3s后再试' : false
	const msg = data.trans_result ? data.trans_result[0].dst : error_msg
	return msg
}
export default translate
