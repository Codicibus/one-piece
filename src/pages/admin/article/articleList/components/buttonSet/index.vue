<template>
	<a-space :size="10" class="batchProcessing">
		<a-button type="primary" @click="addArticle">
			<PlusOutlined />
			添加文章
		</a-button>
		<a-popconfirm title="确定要删除?" @confirm="onDelete">
			<a-button type="primary" danger>
				<DeleteOutlined />
				批量删除
			</a-button>
		</a-popconfirm>
		<a-upload
			action="https://www.mocky.io/v2/5cc8019d300000980a055e76"
			v-model:file-list="articleListStore.fileList"
		>
			<a-button type="primary">
				<ImportOutlined />
				批量导入
			</a-button>
		</a-upload>
		<a-upload
			action="https://www.mocky.io/v2/5cc8019d300000980a055e76"
			v-model:file-list="articleListStore.fileList"
		>
			<a-button type="primary">
				<UploadOutlined />
				批量导出
			</a-button>
		</a-upload>
	</a-space>
</template>

<script>
import { defineComponent, toRaw } from 'vue'
import useStore from '../../store'
import {
	PlusOutlined,
	DeleteOutlined,
	ImportOutlined,
	UploadOutlined
} from '@ant-design/icons-vue'

const articleListStore = useStore()

export default defineComponent({
	name: 'articleList',
	components: {
		PlusOutlined,
		DeleteOutlined,
		ImportOutlined,
		UploadOutlined
	},
	setup() {
		// 添加文章
		const addArticle = () => {
			// 打开窗口,修改编辑模式为添加
			articleListStore.$patch({ visible: true, editingMode: 'add' })
			// 重置表单
			articleListStore.resetForm()
		}
		// 批量删除文章
		const onDelete = async () => {
			const selectedRowKeys = toRaw(articleListStore.selectedRowKeys)
			const content_hash = selectedRowKeys.join(',')
			await articleListStore.deleteArticle(content_hash)
			// 重新获取文章
			await articleListStore.getArticle()
		}
		return {
			articleListStore,
			addArticle,
			onDelete
		}
	}
})
</script>

<style lang="less" scoped>
.batchProcessing {
	margin-bottom: 16px;
}
</style>
