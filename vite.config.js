import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'
import usePluginImport from 'vite-plugin-importer'
import { resolve } from 'path'

export default defineConfig({
	base: '/',
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
