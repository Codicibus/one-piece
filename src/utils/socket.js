import { getToken } from '@/utils/auth'
const wsUrl = 'wss://api.amujun.com/v1/dashboard/ws/systat'
const token = getToken()
const socket = new WebSocket(wsUrl, [token])
export default socket
