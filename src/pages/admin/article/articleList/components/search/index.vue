<template
	#filterDropdown="{
		setSelectedKeys,
		selectedKeys,
		confirm,
		clearFilters,
		column
	}"
>
	<div class="search">
		<a-input
			ref="searchInput"
			:placeholder="`Search ${column.dataIndex}`"
			:value="selectedKeys[0]"
			style="width: 188px; margin-bottom: 8px; display: block"
			@change="
				e => setSelectedKeys(e.target.value ? [e.target.value] : [])
			"
			@pressEnter="handleSearch(selectedKeys, confirm, column.dataIndex)"
		/>
		<a-button
			type="primary"
			size="small"
			style="width: 90px; margin-right: 8px"
			@click="handleSearch(selectedKeys, confirm, column.dataIndex)"
		>
			<template #icon><SearchOutlined /></template>
			Search
		</a-button>
		<a-button
			size="small"
			style="width: 90px"
			@click="handleReset(clearFilters)"
		>
			Reset
		</a-button>
	</div>
</template>

<script>
import { defineComponent, toRaw } from 'vue'
import useStore from '../../store'
import { SearchOutlined } from '@ant-design/icons-vue'

const articleListStore = useStore()

export default defineComponent({
	name: 'search',
	components: { SearchOutlined },
	setup() {
		const state = reactive({
			searchText: '',
			searchedColumn: ''
		})
		const searchInput = ref()
		const handleSearch = (selectedKeys, confirm, dataIndex) => {
			confirm()
			state.searchText = selectedKeys[0]
			state.searchedColumn = dataIndex
		}
		const handleReset = clearFilters => {
			clearFilters()
			state.searchText = ''
		}
		return {
			articleListStore,
			handleSearch,
			handleReset,
			searchText: '',
			searchInput,
			searchedColumn: ''
		}
	}
})
</script>

<style lang="less" scoped>
.search {
	padding: 8px;
}
</style>
