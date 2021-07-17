<template>
	<a-card title="状态">
		<a-row :gutter="16" class="serverState">
			<a-col :span="6" class="state">
				<div class="title">CPU使用情况</div>
				<a-progress type="dashboard" :percent="dashboardStore.cpuPercent" />
				<div class="detail">
					{{ dashboardStore.cpuInfo('cores') + '个核心' }}
					<a-divider type="vertical" />
					{{ dashboardStore.cpuInfo('mhz') }}
				</div>
			</a-col>
			<a-col :span="6" class="state">
				<div class="title">内存使用情况</div>
				<a-progress type="dashboard" :percent="dashboardStore.RAMPercent" />
				<div class="detail">
					{{ dashboardStore.RAMStat('used') }}
					<a-divider type="vertical" />
					{{ dashboardStore.RAMStat('total') }}
				</div>
			</a-col>
			<a-col :span="6" class="state">
				<a-progress type="dashboard" :percent="75" />
			</a-col>
			<a-col :span="6" class="state">
				<a-progress type="dashboard" :percent="75" />
			</a-col>
		</a-row>
	</a-card>
</template>

<script>
import { defineComponent } from 'vue'
// import socket from '@/utils/socket'
import useStore from './store'
let dashboardStore = useStore()
// const getSocket = () => {
// 	socket.onmessage = evn => {
// 		let serveDate = JSON.parse(evn.data)
// 		dashboardStore.$state = { ...serveDate }
// 	}
// }
export default defineComponent({
	name: 'Dashboard',
	setup() {
		// onMounted(() => getSocket())
		return { dashboardStore }
	}
})
</script>

<style lang="less" scoped>
.serverState {
	margin-bottom: 16px;
	.state {
		display: flex;
		flex-flow: column;
		justify-content: space-evenly;
		align-items: center;
	}
}
</style>
