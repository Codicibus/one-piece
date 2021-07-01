import { bytesToGB } from '@/utils/util'
import { defineStore } from 'pinia'

export default defineStore({
	id: 'dashboard',
	state: () => ({
		loading: false,
		cpu_percent: null,
		mem_stat: {},
		net_stat: []
	}),
	getters: {
		cpuPercent() {
			let percent = this.cpu_percent
			if (percent) {
				return Number(percent.toFixed(0))
			} else {
				return 0
			}
		},
		RAMPercent() {
			let percent = this.mem_stat.usedPercent
			if (percent) {
				return Number(percent.toFixed(0))
			} else {
				return 0
			}
		}
	},
	actions: {
		RAMStat(tpye) {
			switch (tpye) {
				case 'total':
					return bytesToGB(this.mem_stat.total).toFixed(2) + ' GB'
				case 'available':
					return bytesToGB(this.mem_stat.available).toFixed(2) + ' GB'
				case 'used':
					return bytesToGB(this.mem_stat.used).toFixed(2) + ' GB'
				case 'free':
					return bytesToGB(this.mem_stat.free).toFixed(2) + ' GB'
			}
		}
	}
})
