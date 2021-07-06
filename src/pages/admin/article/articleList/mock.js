export const columns = [
	{
		title: '标题',
		dataIndex: 'title',
		slots: {
			customRender: 'title',
			filterDropdown: 'filterDropdown'
		},
		ellipsis: true
	},
	{
		title: '内容',
		dataIndex: 'content',
		width: '20%',
		ellipsis: true, // 自动省略
		slots: {
			customRender: 'content'
		}
	},
	{
		title: '分类',
		dataIndex: 'category',
		ellipsis: true, // 自动省略
		// 筛选
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
			customRender: 'category'
		}
	},
	{
		title: '发布时间',
		dataIndex: 'time',
		sorter: true, // 排序
		slots: {
			customRender: 'time'
		}
	},
	{
		title: '更新时间',
		dataIndex: 'time',
		sorter: true, // 排序
		slots: {
			customRender: 'time'
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
