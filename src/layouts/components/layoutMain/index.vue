<template>
	<a-layout-content class="content">
		<router-view />
	</a-layout-content>
</template>

<style lang="less" scoped>
.content {
	margin: 16px 16px 0;
	height: calc(100vh - 170px);
	overflow: hidden;
}
</style>
