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
		<a-col :span="8">
			<a-card title="CPU" :bordered="false">
				<a-progress type="dashboard" :percent="cpu_percent" />
			</a-card>
		</a-col>
		<a-col :span="8">
			<a-card title="内存" :bordered="false">
				<a-progress type="dashboard" :percent="75" />
				<div>
					<!-- <span>{{ totalMemory }}</span> -->
					/
					<span></span>
				</div>
			</a-card>
		</a-col>
		<a-col :span="8">
			<a-card title="网络" :bordered="false">
				<a-progress type="dashboard" :percent="75" />
			</a-card>
		</a-col>
	</a-row>
	{{ dashboardStore.serveDate }}
</template>

<script>
// import { message } from 'ant-design-vue'
import { computed, defineComponent, onMounted } from 'vue'
import socket from '@/utils/socket'
import useStore from './store'
let dashboardStore = useStore()
const getSocket = () => {
	socket.onopen = () => console.log('socket:onopen')
	socket.onmessage = evn => {
		let data = JSON.parse(evn.data)
		dashboardStore.$state = { serveDate: data }
	}
	socket.onclose = () => console.log('socket.onclose')
}
export default defineComponent({
	name: 'Dashboard',
	setup() {
		const cpu_percent = computed(() => {
			let percent = dashboardStore.serveDate.cpu_percent
			if (percent) {
				return percent.toFixed(0)
			} else {
				return 0
			}
		})

		onMounted(() => getSocket())
		return { getSocket, dashboardStore, cpu_percent }
	}
})
</script>

<style lang="less" scoped>
.articleState {
	margin-bottom: 16px;
}
// .serverState;
</style>
