<template>
	<a-layout-sider v-model:inline-collapsed="collapsed" collapsible>
		<a-menu theme="dark" v-model:selectedKeys="selectedKeys" mode="inline">
			<div class="logo" />
			<!-- 仪表盘 -->
			<a-menu-item key="1">
				<router-link :to="{ name: 'Dashboard' }">
					<pie-chart-outlined />
					<span>仪表盘</span>
				</router-link>
			</a-menu-item>
			<!-- 基础设置 -->
			<a-menu-item key="2">
				<router-link :to="{ name: 'Setting' }">
					<pie-chart-outlined />
					<span>基础设置</span>
				</router-link>
			</a-menu-item>
			<!-- 文章管理 -->
			<a-sub-menu key="sub1">
				<template #title>
					<span>
						<user-outlined />
						<span>文章管理</span>
					</span>
				</template>
				<router-link :to="{ name: 'ArticleList' }">
					<a-menu-item key="sub1-1">文章列表</a-menu-item>
				</router-link>
				<router-link :to="{ name: 'ArticleClassify' }">
					<a-menu-item key="sub1-2">文章分类</a-menu-item>
				</router-link>
				<router-link :to="{ name: 'PagingSetting' }">
					<a-menu-item key="sub1-3">分页设置</a-menu-item>
				</router-link>
				<router-link :to="{ name: 'ImagePool' }">
					<a-menu-item key="sub1-4">图片池</a-menu-item>
				</router-link>
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
</style>
