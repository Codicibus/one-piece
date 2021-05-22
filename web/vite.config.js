import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'
import { resolve } from 'path'

// https://vitejs.dev/config/
export default defineConfig({
	base: '/',
	resolve: {
		alias: {
			'@': resolve(__dirname, 'src'),
			com: resolve(__dirname, 'src/components')
		}
	},
	plugins: [vue()]
})
