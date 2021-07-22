<template>
	<div class="article">col</div>
</template>

<script>
import { defineComponent } from 'vue'
export default defineComponent({
	setup() {}
})
</script>

<style lang="less" scoped>
.article {
	width: 100%;
	height: 150px;
	padding: 16px;
	border: 1px solid rgba(255, 255, 255, 0.3);
	border-radius: 8px;
	background: rgba(255, 255, 255, 0.3);
	box-shadow: 2px 2px 3px rgba(170, 170, 170, 0.3);
	backdrop-filter: blur(20px);
}
</style>
