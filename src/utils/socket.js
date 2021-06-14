import { getToken } from '@/utils/auth'
const token = getToken()
const wsUrl = `${import.meta.env.VITE_WS_URL}/v1/dashboard/systat`
const socket = new WebSocket(wsUrl, [token])
socket.onconnecting = evn => {
	console.log('socket:onconnecting:', evn)
}
socket.onopen = evn => {
	console.log('socket:onopen:', evn)
}
socket.onmessage = evn => {
	const data = JSON.parse(evn.data)
	console.log(data)
}
socket.onclose = evn => {
	console.log('socket.onclose:', evn)
}
socket.onerror = evn => {
	console.error(evn)
}

export default socket
