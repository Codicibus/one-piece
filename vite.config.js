import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'
import usePluginImport from 'vite-plugin-importer'
import prismjs from 'vite-plugin-prismjs'
import babel from 'rollup-plugin-babel'
import { resolve } from 'path'
import proxy from './config/proxy'

import components from 'prismjs/components'
const allLanguages = Object.keys(components.languages).filter(item => item !== 'meta')

export default defineConfig({
	// 基础访问路径
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
		}),
		babel(),
		prismjs({
			languages: allLanguages
		})
	]
})
