export const unitTransform = (bytes, unit) => {
	if (bytes === 0) return '0 B'
	let MB = bytes / 1024 / 1024
	let GB = bytes / 1024 / 1024 / 1024
	switch (unit) {
		case 'MB':
			return MB
		case 'GB':
			return GB
		default:
			return MB
	}
}
