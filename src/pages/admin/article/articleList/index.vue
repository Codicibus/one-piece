<template>
	<!-- 按钮组 -->
	<button-set />
	<!-- 表格 -->
	<a-table
		:row-selection="{
			selectedRows: articleListStore.selectedRows,
			onChange: onSelectChange
		}"
		:loading="articleListStore.loading"
		:columns="articleListStore.columns"
		:row-key="record => record.content_hash"
		:data-source="articleListStore.articlesInfo.articles"
		:pagination="{
			pageSize: articleListStore.pagingSetup.page_size,
			total: articleListStore.articlesInfo.total,
			showSizeChange: (page_num, page_size) =>
				articleListStore.$path({ pagingSetup: page_num, page_size })
		}"
		:scroll="{ y: 400 }"
	>
		<!-- 标题 -->
		<template #title="{ text }">{{ text }}</template>
		<!-- 内容 -->
		<template #content="{ text }">{{ text }}</template>
		<!-- 分类 -->
		<template #category="{ text }">
			<a-tag color="blue">{{ text }}</a-tag>
		</template>
		<!-- 发布时间 -->
		<template #time="{ text }">{{ text }}</template>
		<!-- 更新时间 -->
		<!-- <template #time="{ text }">{{ text }}</template> -->
		<!-- 操作 start -->
		<template #action="{ record }">
			<a-space :size="8">
				<a-button type="primary" @click="onEdit(record)">
					<template #icon>
						<EditOutlined />
					</template>
					编辑
				</a-button>
				<a-popconfirm title="确定要删除?" @confirm="onDelete(record)">
					<a-button type="danger">
						<template #icon>
							<DeleteOutlined />
						</template>
						删除
					</a-button>
				</a-popconfirm>
			</a-space>
		</template>
		<!-- 操作 end -->
	</a-table>
	<!-- 抽屉 -->
	<Modal />
</template>

<script>
import { EditOutlined, DeleteOutlined } from '@ant-design/icons-vue'
import { defineComponent, toRaw } from 'vue'
import useStore from './store'
import Modal from './components/modal/index.vue'
import ButtonSet from './components/buttonSet/index.vue'
const articleListStore = useStore()

export default defineComponent({
	name: 'articleList',
	components: { Modal, ButtonSet, EditOutlined, DeleteOutlined },
	setup() {
		// 获取文章
		articleListStore.getArticle()

		// 选择文章
		const onSelectChange = (selectedRowKeys, selectedRows) => {
			articleListStore.$patch({ selectedRows })
			console.log(selectedRowKeys, toRaw(selectedRows))
		}

		// 编辑文章
		const onEdit = record =>
			articleListStore.$patch({ visible: true, formData: record })

		// 删除文章
		const onDelete = record => {
			console.log(record)
			// 重新获取文章
			articleListStore.getArticle()
		}
		return {
			articleListStore,
			onSelectChange,
			onEdit,
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
