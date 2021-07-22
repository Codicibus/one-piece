<template>
	<a-layout-header class="header">
		<div class="personal">
			<a-avatar><UserOutlined /></a-avatar>
			<a-dropdown>
				<a-button type="link" @click.prevent>
					Link
					<DownOutlined />
				</a-button>
				<template #overlay>
					<a-menu>
						<a-menu-item>
							<a href="javascript:;">个人设置</a>
						</a-menu-item>
						<a-menu-divider />
						<a-menu-item @click="signOut"> 退出登录 </a-menu-item>
					</a-menu>
				</template>
			</a-dropdown>
		</div>
	</a-layout-header>
</template>

<script>
import { defineComponent } from 'vue'
import { DownOutlined, UserOutlined } from '@ant-design/icons-vue'
import { removeToken } from '@/utils/auth'
import { useRouter } from 'vue-router'
// import socket from '@/utils/socket'

export default defineComponent({
	components: {
		DownOutlined,
		UserOutlined
	},
	setup() {
		const router = useRouter()
		const signOut = () => {
			removeToken()
			// socket.close(1000)
			router.push('/login')
		}
		return { signOut }
	}
})
</script>

<style lang="less" scoped>
.header {
	background: #fff;
	padding: 0 16px;
}
.personal {
	float: right;
}
</style>
