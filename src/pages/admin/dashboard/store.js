import { unitTransform } from '@/utils/util'
import { defineStore } from 'pinia'

export default defineStore({
	id: 'dashboard',
	state: () => ({
		loading: false,
		cpu_info: [],
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
		cpuInfo(tpye) {
			this.cpu_info.forEach(item => {
				return item[tpye]
			})
		},
		RAMStat(tpye) {
			return unitTransform(this.mem_stat[tpye]).toFixed(0) + ' MB'
		}
	}
})
