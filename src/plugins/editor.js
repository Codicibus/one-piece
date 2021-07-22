import VueMarkdownEditor from '@kangc/v-md-editor'
import '@kangc/v-md-editor/lib/style/base-editor.css'
import vuepressTheme from '@kangc/v-md-editor/lib/theme/vuepress.js'
import '@kangc/v-md-editor/lib/theme/style/vuepress.css'
import createCopyCodePlugin from '@kangc/v-md-editor/lib/plugins/copy-code/index'
import '@kangc/v-md-editor/lib/plugins/copy-code/copy-code.css'

// 引入所有语言包
import Prism from 'prismjs'
// 定义编辑器代码主题 注:必须要在创建编辑器之前定义
VueMarkdownEditor.use(vuepressTheme, { Prism })
// 创建编辑器
VueMarkdownEditor.use(createCopyCodePlugin())
// 导入
export default app => {
	app.use(VueMarkdownEditor)
}
