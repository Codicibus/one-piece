<template>
	<a-row class="container" :gutter="[16, 16]">
		<a-col :span="12">
			<div class="article">Col</div>
		</a-col>
		<a-col :span="12">
			<div class="article">Col</div>
		</a-col>
		<a-col :span="12">
			<div class="article">Col</div>
		</a-col>
		<a-col :span="12">
			<div class="article">Col</div>
		</a-col>
		<a-col :span="12">
			<div class="article">Col</div>
		</a-col>
		<a-col :span="12">
			<div class="article">Col</div>
		</a-col>
	</a-row>
</template>

<script>
import { defineComponent } from 'vue'
export default defineComponent({
	setup() {}
})
</script>

<style lang="less" scoped>
.container {
	flex: 1;
	.article {
		width: 100%;
		height: 200px;
		padding: 15px;
		background: wheat;
		border-radius: 8px;
	}
}
</style>
