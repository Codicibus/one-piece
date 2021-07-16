<template>
	<main>
		<home-aside />
		<div class="content">
			<home-card />
			<footer>MIT Licensed | Copyright © 2021 Codicibus</footer>
		</div>
	</main>
</template>
<script>
import { defineComponent } from 'vue'
import HomeCard from './Card.vue'
import HomeAside from './Aside.vue'
export default defineComponent({
	components: { HomeCard, HomeAside },
	setup() {}
})
</script>

<style lang="less" scope>
main {
	width: 100%;
	height: 100%;
	display: flex;
	padding: 16px;
	.content {
		width: 100%;
	}
}
</style>
