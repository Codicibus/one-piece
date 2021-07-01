export const bytesToGB = bytes => {
	if (bytes === 0) return '0 B'
	let GB = bytes / 1024 / 1024 / 1024
	return GB
}
