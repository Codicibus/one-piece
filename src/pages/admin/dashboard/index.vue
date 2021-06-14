<template>
	<a-row :gutter="16">
		<a-col :span="8" class="articleState">
			<a-card title="Card title" :bordered="false">
				<p>card content</p>
			</a-card>
		</a-col>
		<a-col :span="8">
			<a-card title="Card title" :bordered="false">
				<p>card content</p>
			</a-card>
		</a-col>
		<a-col :span="8">
			<a-card title="Card title" :bordered="false">
				<p>card content</p>
			</a-card>
		</a-col>
	</a-row>
	<a-row :gutter="16" class="serverState">
		<a-col :span="6">
			<a-card title="Card title" :bordered="false">
				<a-progress type="dashboard" :percent="75" />
			</a-card>
		</a-col>
		<a-col :span="6">
			<a-card title="Card title" :bordered="false">
				<a-progress type="dashboard" :percent="75" />
			</a-card>
		</a-col>
		<a-col :span="6">
			<a-card title="Card title" :bordered="false">
				<a-progress type="dashboard" :percent="75" />
			</a-card>
		</a-col>
		<a-col :span="6">
			<a-card title="Card title" :bordered="false">
				<a-progress type="dashboard" :percent="75" />
			</a-card>
		</a-col>
	</a-row>
</template>

<script>
import { message } from 'ant-design-vue'
import { defineComponent, onUnmounted } from 'vue'
import socket from '@/utils/socket'
export default defineComponent({
	name: 'Dashboard',
	setup() {
		onUnmounted(() => {
			socket.close(1000)
			if (socket.readyState === 2) {
				message.success('websocket正在关闭...')
			} else if (socket.readyState === 3) {
				message.success('websocket已关闭')
			}
		})
	}
})
</script>

<style lang="less" scoped>
.articleState {
	margin-bottom: 16px;
}
// .serverState;
</style>
