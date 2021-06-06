// import request from '@/utils/request'
import { defineStore } from 'pinia'
import loginStore from '@/pages/login/store'
import dashboardStore from '@/pages/admin/dashboard/store'

export default defineStore(loginStore, dashboardStore)
