import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'
import usePluginImport from 'vite-plugin-importer'
import { resolve } from 'path'
import proxy from './config/proxy'

export default defineConfig({
	// 在生产中服务时的基本公共路径
	base: '/',
	// 服务端渲染
	ssr: false,
	server: {
		proxy
	},
	strictPort: true, // 若端口已被占用则会直接退出
	resolve: {
		alias: {
			'@': resolve(__dirname, 'src'),
			'@components': resolve(__dirname, 'src/components'),
			'@article': resolve(__dirname, 'src/pages/admin/article')
		}
	},
	plugins: [
		vue(),
		usePluginImport({
			libraryName: 'ant-design-vue',
			libraryDirectory: 'es',
			style: 'css'
		})
	]
})
