export default {
	'/api': {
		target: 'https://fanyi-api.baidu.com',
		changeOrigin: true
	},
	// '^/v1/dashboard': {
	// 	target: 'wss://api.amujun.com',
	// 	changeOrigin: true,
	// 	ws: true
	// 	// rewrite: path => path.replace(/^\/v1/, '')
	// },
	'/v1': {
		target: 'https://api.amujun.com',
		changeOrigin: true
	}
}
