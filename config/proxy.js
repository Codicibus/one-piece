export default {
	'/api': {
		target: 'https://fanyi-api.baidu.com',
		changeOrigin: true
	},
	'/v1/': {
		target: 'http://127.0.0.1:8080',
		changeOrigin: true
	}
}
