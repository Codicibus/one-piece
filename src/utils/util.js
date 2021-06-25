export const bytesToGB = bytes => {
	if (bytes === 0) return '0 B'
	const k = 1024
	let GB = bytes / k
	return GB
	//toPrecision(3) 后面保留一位小数，如1.0GB                                                                                                                  //return (bytes / Math.pow(k, i)).toPrecision(3) + ' ' + sizes[i];
}
