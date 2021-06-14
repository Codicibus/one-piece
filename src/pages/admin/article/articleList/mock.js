export const columns = [
	{
		title: '标题',
		dataIndex: 'title',
		slots: {
			customRender: 'title'
		},
		ellipsis: true
	},
	{
		title: '分类',
		dataIndex: 'classify',
		width: '10%',
		// 筛选名称
		filters: [
			{
				text: 'one',
				value: 'one'
			},
			{
				text: 'two',
				value: 'two'
			}
		],
		slots: {
			customRender: 'classify'
		}
	},
	{
		title: '发布时间',
		dataIndex: 'time',
		sorter: true,
		ellipsis: true,
		slots: {
			customRender: 'time'
		}
	},
	{
		title: '操作',
		dataIndex: 'action',
		width: '30%',
		slots: {
			customRender: 'action'
		}
	}
]
export const dataSource = [
	{
		id: 1,
		title: '1',
		classify: 'John Brown',
		time: 18
	},
	{
		id: 2,
		title: '2',
		classify: 'John Brown',
		time: 32
	},
	{
		id: 3,
		title: '3',
		classify: 'John Brown',
		time: 32
	}
]
