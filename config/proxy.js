export default {
	proxy: {
		'/v1/': {
			target: 'http://127.0.0.1:8081',
			changeOrigin: true
		}
	}
}
