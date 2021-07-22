<template>
	<main class="homeMain">
		<div>
			<home-card />
		</div>
	</main>
</template>
<script>
import { defineComponent, ref } from 'vue'
import HomeCard from './Card.vue'
export default defineComponent({
	components: { HomeCard },
	setup() {
		const current1 = ref(1)
		const onChange = pageNumber => {
			console.log('Page: ', pageNumber)
		}
		return {
			current1,
			onChange
		}
	}
})
</script>

<style lang="less" scope>
.homeMain {
	width: 100%;
	height: calc(100% - 130px);
	overflow-x: hidden;
	overflow-y: auto;
	padding: 32px;
}
</style>
