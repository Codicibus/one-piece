// 引入andt
import {
	Typography,
	Input,
	BackTop,
	Form,
	Button,
	Layout,
	Menu,
	Card,
	Empty,
	Statistic,
	Row,
	Col,
	Progress,
	Breadcrumb
} from 'ant-design-vue'

// 导入
export default app => {
	app.use(Typography)
	app.use(Input)
	app.use(BackTop)
	app.use(Form)
	app.use(Button)
	app.use(Layout)
	app.use(Menu)
	app.use(Card)
	app.use(Empty)
	app.use(Statistic)
	app.use(Row)
	app.use(Col)
	app.use(Progress)
	app.use(Breadcrumb)
}
