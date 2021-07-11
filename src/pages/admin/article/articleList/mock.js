export const columns = [
	{
		title: '标题',
		dataIndex: 'name',
		slots: {
			customRender: 'name'
		},
		ellipsis: true
	},
	{
		title: '分类',
		dataIndex: 'category',
		ellipsis: true, // 自动省略
		// 筛选
		// filters: [
		// 	{
		// 		text: 'one',
		// 		value: 'one'
		// 	},
		// 	{
		// 		text: 'two',
		// 		value: 'two'
		// 	}
		// ],
		slots: {
			customRender: 'category'
		}
	},
	{
		title: '发布时间',
		dataIndex: 'created_at',
		sorter: true, // 排序
		slots: {
			customRender: 'created_at'
		}
	},
	{
		title: '更新时间',
		dataIndex: 'updated_at',
		sorter: true, // 排序
		slots: {
			customRender: 'updated_at'
		}
	},
	{
		title: '操作',
		dataIndex: 'action',
		width: '20%',
		slots: {
			customRender: 'action'
		}
	}
]
