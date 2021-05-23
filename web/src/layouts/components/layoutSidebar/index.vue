<template>
	<a-layout-sider v-model:inline-collapsed="collapsed" collapsible>
		<a-menu theme="dark" v-model:selectedKeys="selectedKeys" mode="inline">
			<div class="logo" />
			<a-menu-item key="1">
				<span>
					<pie-chart-outlined />
					<router-link to="/admin" active-class="routerActive">仪表盘</router-link>
				</span>
			</a-menu-item>
			<a-menu-item key="2">
				<span>
					<pie-chart-outlined />
					<router-link to="/admin/setting" active-class="routerActive">基础设置</router-link>
				</span>
			</a-menu-item>
			<a-sub-menu key="sub1">
				<template #title>
					<user-outlined />
					<span>文章管理</span>
				</template>
				<a-menu-item key="sub1-1">文章列表</a-menu-item>
				<a-menu-item key="sub1-2">文章分类</a-menu-item>
				<a-menu-item key="sub1-3">分页设置</a-menu-item>
				<a-menu-item key="sub1-4">图片池</a-menu-item>
			</a-sub-menu>
		</a-menu>
	</a-layout-sider>
</template>
<script setup>
import { PieChartOutlined, UserOutlined } from '@ant-design/icons-vue'
import { defineComponent, ref } from 'vue'

const selectedKeys = ref(['1'])
const collapsed = ref(false)

defineComponent({
	components: {
		PieChartOutlined,
		UserOutlined
	}
})
</script>
<style lang="less" scope>
.logo {
	height: 32px;
	margin: 16px;
	background: rgba(255, 255, 255, 0.3);
}
.routerActive {
	color: white;
}
</style>
