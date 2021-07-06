<template>
	<a-space :size="10" class="batchProcessing">
		<a-button type="primary" @click="addArticle">
			<template #icon>
				<PlusOutlined />
			</template>
			添加文章
		</a-button>
		<a-popconfirm title="确定要删除?" @confirm="onDelete(record)">
			<a-button type="danger">
				<template #icon>
					<DeleteOutlined />
				</template>
				批量删除
			</a-button>
		</a-popconfirm>
		<a-button type="primary">
			<template #icon>
				<EditOutlined />
			</template>
			导入文章
		</a-button>
		<a-button type="primary">
			<template #icon>
				<EditOutlined />
			</template>
			批量导出
		</a-button>
	</a-space>
</template>

<script>
import { defineComponent, toRaw } from 'vue'
import useStore from '../../store'
import {
	EditOutlined,
	DeleteOutlined,
	PlusOutlined
} from '@ant-design/icons-vue'

const articleListStore = useStore()

export default defineComponent({
	name: 'articleList',
	components: { EditOutlined, DeleteOutlined, PlusOutlined },
	setup() {
		const openDrawer = () => articleListStore.$patch({ visible: true })
		const addArticle = () => {
			articleListStore.resetForm
			openDrawer()
		}
		const onDelete = record => {
			console.log(record)
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
