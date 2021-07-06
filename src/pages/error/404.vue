<template>
	<a-empty :image="simpleImage" description="NotFound" class="notFound" />
</template>
<script>
import { defineComponent } from 'vue'
import { Empty } from 'ant-design-vue'
export default defineComponent({
	name: 'NotFound',
	setup() {
		return {
			simpleImage: Empty.PRESENTED_IMAGE_SIMPLE
		}
	}
})
</script>

<style lang="less">
.notFound {
	margin: 0;
	position: absolute;
	left: 50%;
	top: 50%;
	transform: translate(-50%, -50%);
}
</style>
