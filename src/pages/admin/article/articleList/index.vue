<template>
	<div class="articleList">
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
				showTotal: total => `共 ${total} 条数据`,
				pageSizeOptions: ['5', '10', '15', '20'],
				showSizeChanger: true,
				showQuickJumper: true
			}"
			@change="changePaging"
			:scroll="{ y: 'calc(100vh - 305px)' }"
		>
			<!-- 标题 -->
			<template #name="{ text }">{{ text }}</template>
			<!-- 分类 -->
			<template #category="{ text }">
				<a-tag color="blue">{{ text }}</a-tag>
			</template>
			<!-- 发布时间 -->
			<template #created_at="{ text }">{{ text }}</template>
			<!-- 更新时间 -->
			<template #updated_at="{ text }">{{ text }}</template>
			<!-- 操作 start -->
			<template #action="{ record }">
				<a-space :size="8">
					<a-button type="primary" @click="onEdit(record)">
						<EditOutlined />编辑
					</a-button>
					<a-popconfirm
						title="确定要删除?"
						@confirm="onDelete(record)"
					>
						<a-button type="primary" danger>
							<DeleteOutlined />删除
						</a-button>
					</a-popconfirm>
				</a-space>
			</template>
			<!-- 操作 end -->
		</a-table>
		<!-- 抽屉 -->
		<Modal />
	</div>
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
		const onSelectChange = selectedRowKeys =>
			articleListStore.$patch({ selectedRowKeys })

		// 编辑文章
		const onEdit = record =>
			articleListStore.$patch({
				visible: true,
				formData: record,
				editingMode: 'edit'
			})
		// 删除文章
		const onDelete = record => {
			articleListStore.deleteArticle(toRaw(record.content_hash))
			// 重新获取文章
			articleListStore.getArticle()
		}
		// 修改分页
		const changePaging = pag => {
			articleListStore.$patch({
				pagingSetup: {
					page_num: pag.current,
					page_size: pag.pageSize
				}
			})
			// 重新获取文章
			articleListStore.getArticle()
		}

		return {
			articleListStore,
			onSelectChange,
			onEdit,
			onDelete,
			changePaging
		}
	}
})
</script>

<style lang="less" scoped>
.articleList {
	height: calc(100%);
	// overflow: auto;
}
</style>
