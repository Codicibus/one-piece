<template>
	<div class="home">
		<div class="container">
			<div class="leftContainer">
				<home-aside />
			</div>
			<div class="rightContainer">
				<home-header />
				<home-main />
				<home-footer />
			</div>
		</div>
	</div>
</template>
<script>
import { defineComponent } from 'vue'
import HomeHeader from './components/homeHeader/index.vue'
import HomeAside from './components/homeAside/index.vue'
import HomeMain from './components/homeMain/index.vue'
import HomeFooter from './components/homeFooter/index.vue'
export default defineComponent({
	components: { HomeHeader, HomeAside, HomeMain, HomeFooter },
	setup() {}
})
</script>

<style lang="less" scope>
.home {
	width: 100%;
	height: 100vh;
	background: url('@/assets/background.jpg') no-repeat;
	// background: -webkit-linear-gradient(45deg, rgb(175, 238, 255) 37%, rgb(165, 213, 255) 73%);
	.container {
		width: 1100px;
		height: 85%;
		display: flex;
		position: absolute;
		top: 50%;
		left: 50%;
		transform: translate(-50%, -50%);
		border-radius: 15px;
		background: rgba(255, 255, 255, 0.3);
		box-shadow: 0 0 5px rgba(170, 170, 170, 0.3);
		.leftContainer {
			width: 200px;
			height: 100%;
		}
		.rightContainer {
			width: 900px;
			height: 100%;
		}
	}
}
</style>
