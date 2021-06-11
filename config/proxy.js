export default {
	'/api': {
		target: 'https://fanyi-api.baidu.com',
		changeOrigin: true,
		rewrite: path => path.replace(/^\/api/, '')
	}
}
