<template>
	<a-row :gutter="16" class="serverState">
		<a-col :span="8">
			<a-card title="CPU" bodyStyle="text-align: center;">
				<a-progress
					type="dashboard"
					:percent="dashboardStore.cpuPercent"
				/>
			</a-card>
		</a-col>
		<a-col :span="8">
			<a-card
				title="内存"
				bodyStyle="display: flex; justify-content: space-evenly;"
			>
				<a-progress
					type="dashboard"
					:percent="dashboardStore.RAMPercent"
				/>
				<div class="detail">
					<div>
						{{ dashboardStore.RAMStat('used') }}
						<a-divider type="vertical" /> 已使用内容
					</div>
					<div>
						{{ dashboardStore.RAMStat('free') }}
						<a-divider type="vertical" /> 剩余内容
					</div>
					<div>
						{{ dashboardStore.RAMStat('available') }}
						<a-divider type="vertical" /> 可用总量
					</div>
					<div>
						{{ dashboardStore.RAMStat('total') }}
						<a-divider type="vertical" /> 内存总量
					</div>
				</div>
			</a-card>
		</a-col>
		<a-col :span="8">
			<a-card title="网络" bodyStyle="text-align: center;">
				<a-progress type="dashboard" :percent="75" />
			</a-card>
		</a-col>
	</a-row>
</template>

<script>
import { defineComponent, onMounted } from 'vue'
import socket from '@/utils/socket'
import useStore from './store'
let dashboardStore = useStore()
const getSocket = () => {
	socket.onmessage = evn => {
		let serveDate = JSON.parse(evn.data)
		dashboardStore.$state = { ...serveDate }
	}
}
export default defineComponent({
	name: 'Dashboard',
	setup() {
		onMounted(() => getSocket())
		return { getSocket, dashboardStore }
	}
})
</script>

<style lang="less" scoped>
.serverState {
	margin-bottom: 16px;
	.detail {
		line-height: 200%;
	}
}
</style>
