<template>
	<header>
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
header {
	width: 100%;
	height: 64px;
	line-height: 64px;
	padding: 0 20px;
	border-bottom: 1px solid #ccc;
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
