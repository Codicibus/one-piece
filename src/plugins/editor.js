import VueMarkdownEditor from '@kangc/v-md-editor'
import '@kangc/v-md-editor/lib/style/base-editor.css'

// import vuepressTheme from '@kangc/v-md-editor/lib/theme/vuepress.js'
// import '@kangc/v-md-editor/lib/theme/style/vuepress.css'

import githubTheme from '@kangc/v-md-editor/lib/theme/github.js'
import '@kangc/v-md-editor/lib/theme/style/github.css'

import createCopyCodePlugin from '@kangc/v-md-editor/lib/plugins/copy-code/index'
import '@kangc/v-md-editor/lib/plugins/copy-code/copy-code.css'

// 引入所有语言包
import hljs from 'highlight.js'
// import Prism from 'prismjs'
VueMarkdownEditor.use(githubTheme, {
	// Prism
	Hljs: hljs
})

VueMarkdownEditor.use(createCopyCodePlugin())

// 导入
export default app => {
	app.use(VueMarkdownEditor)
}
