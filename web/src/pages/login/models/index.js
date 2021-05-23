import { reactive } from 'vue'
import { getUsers } from '../request'
// 列表数据
export const state = reactive({
	loading: true, // 加载状态
	tableList: [], // 列表数据
	total: 0, // 数据总条数
	queryInfo: {
		// 列表数据请求参数
		query: '',
		pagenum: 1,
		pagesize: 10
	},
	addUserDialogVisible: false, // 新增弹框
	updateUserDialogVisible: false, // 修改弹框
	updateForm: {} // 修改弹框组件数据
})

// 获得用户列表
export const getList = async () => {
	state.loading = true
	let res = await getUsers(state.queryInfo)
	if (res) {
		state.tableList = res.users
		state.total = res.total
	}
	state.loading = false
}

// 表头单元格样式
export const headerCellStyle = {
	backgroundColor: '#1283F8',
	color: '#fff',
	textAlign: 'center'
}
