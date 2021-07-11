<template>
	<a-layout-sider v-model:collapsed="collapsed" collapsible inlineCollapsed class="sider">
		<!-- logo -->
		<a-typography-title :level="4" class="logo"> ONE PIECE </a-typography-title>
		<a-menu
			theme="dark"
			v-model:selectedKeys="selectedKeys"
			v-model:openKeys="openKeys"
			mode="inline"
		>
			<!-- 仪表盘 -->
			<a-menu-item key="Dashboard">
				<router-link :to="{ name: 'Dashboard' }">
					<span>
						<pie-chart-outlined />
						<span>仪表盘</span>
					</span>
				</router-link>
			</a-menu-item>
			<!-- 文章管理 -->
			<a-sub-menu key="sub1">
				<template #title>
					<user-outlined />
					<span>文章管理</span>
				</template>
				<a-menu-item key="ArticleList">
					<router-link :to="{ name: 'ArticleList' }"> 文章列表 </router-link>
				</a-menu-item>
				<a-menu-item key="ArticleClassify">
					<router-link :to="{ name: 'ArticleClassify' }"> 文章分类 </router-link>
				</a-menu-item>
				<a-menu-item key="ImagePool">
					<router-link :to="{ name: 'ImagePool' }"> 图片池 </router-link>
				</a-menu-item>
			</a-sub-menu>
		</a-menu>
	</a-layout-sider>
</template>

<script>
import { defineComponent, ref } from 'vue'
import { PieChartOutlined, UserOutlined } from '@ant-design/icons-vue'
import { useRoute } from 'vue-router'
export default defineComponent({
	components: {
		PieChartOutlined,
		UserOutlined
	},
	setup() {
		const route = useRoute()
		const selectedKeys = ref([route.name])
		const collapsed = ref(false)
		const openKeys = ref(['sub1'])
		return { selectedKeys, collapsed, openKeys, route }
	}
})
</script>

<style lang="less" scope>
.sider {
	.logo {
		text-align: center;
		margin: 10px auto !important;
		color: white !important;
	}
}
</style>
