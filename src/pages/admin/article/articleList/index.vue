<template>
	<a-card class="batchCard">
		<a-space :size="10" class="batchProcessing">
			<a-button type="primary" @click="openDrawer">
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
				批量导出
			</a-button>
		</a-space>
	</a-card>
	<a-table
		:row-selection="{
			selectedRows: articleListStore.selectedRows,
			onChange: onSelectChange
		}"
		:columns="columns"
		:row-key="record => record.id"
		:data-source="dataSource"
		:pagination="{ current: 1, pageSize: 5, total: 10 }"
		:loading="articleListStore.loading"
		:scroll="{ y: '100%' }"
		style="height: 0"
	>
		<!-- 标题 -->
		<template #title="{ text }">{{ text }}</template>
		<!-- 分类 -->
		<template #classify="{ text }">
			<a-tag color="blue">{{ text }}</a-tag>
		</template>
		<!-- 发布时间 -->
		<template #time="{ text }">{{ text }}</template>
		<!-- 操作 start -->
		<template #action>
			<a-space :size="8">
				<a-button type="primary" @click="openDrawer">
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
	<Drawer />
</template>

<script>
import { defineComponent, toRaw } from 'vue'
import useStore from './store'
import Drawer from './components/Drawer/Drawer.vue'
import {
	SearchOutlined,
	EditOutlined,
	DeleteOutlined,
	PlusOutlined
} from '@ant-design/icons-vue'
import { columns, dataSource } from './mock'
const articleListStore = useStore()

export default defineComponent({
	name: 'articleList',
	components: {
		SearchOutlined,
		EditOutlined,
		DeleteOutlined,
		PlusOutlined,
		Drawer
	},
	setup() {
		const onSelectChange = (selectedRowKeys, selectedRows) => {
			articleListStore.$patch({ selectedRows })
			const keys = toRaw(articleListStore.selectedRows)
			console.log(keys, selectedRowKeys)
		}
		const openDrawer = () => articleListStore.$patch({ visible: true })
		const onDelete = record => {
			console.log(record)
		}
		return {
			articleListStore,
			columns,
			dataSource,
			openDrawer,
			onSelectChange,
			onDelete
		}
	}
})
</script>

<style lang="less" scoped>
.batchCard {
	margin-bottom: 16px;
	.batchProcessing {
		float: right;
	}
}
</style>
