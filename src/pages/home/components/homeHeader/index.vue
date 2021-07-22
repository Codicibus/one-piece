<template>
	<header class="homeHeader">
		<span class="logo"></span>
		<span class="classify">
			<a-button type="link">Link Button</a-button>
			<a-button type="link">Link Button</a-button>
			<a-button type="link">Link Button</a-button>
			<a-button type="link">Link Button</a-button>
			<a-button type="link">Link Button</a-button>
		</span>
	</header>
</template>
<script>
import { defineComponent } from 'vue'
export default defineComponent({
	setup() {}
})
</script>

<style lang="less" scope>
.homeHeader {
	width: 100%;
	height: 64px;
	line-height: 64px;
	padding: 0 20px;
	position: relative;
	z-index: 10;
	border-radius: 0 15px 0 0;
	border-bottom: 1px solid #5353534d;
	box-shadow: 0 0 5px #5353534d;
	.logo {
		padding: 10px 60px;
		background: #ccc;
		border-radius: 8px;
	}
	.classify {
		float: right;
	}
}
</style>
