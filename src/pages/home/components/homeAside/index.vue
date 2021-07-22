<template>
	<aside class="aside">
		<header class="header">
			<a-space>
				<span class="red"></span>
				<span class="yellow"></span>
				<span class="green"></span>
			</a-space>
		</header>
		<main class="main">
			<div>
				<div class="classifyName">iCloud</div>
				<div>classifyName</div>
				<div>classifyName</div>
				<div>classifyName</div>
			</div>
			<div>classifyName</div>
			<div>classifyName</div>
			<div>classifyName</div>
			<div>classifyName</div>
			<div>classifyName</div>
			<div>classifyName</div>
			<div>classifyName</div>
		</main>
	</aside>
</template>
<script>
import { defineComponent } from 'vue'
export default defineComponent({
	setup() {}
})
</script>

<style lang="less" scoped>
.aside {
	width: 200px;
	height: 100%;
	border-radius: 15px 0 0 15px;
	// border-right: 2px solid #ffffff4d;
	background: rgba(0, 0, 0, 0.5);
	> .header {
		padding: 20px 25px;
		span {
			display: inline-block;
			width: 13px;
			height: 13px;
			border-radius: 50%;
			box-shadow: 1px 1px 2px #8585854d;
		}
		.red {
			background: #f56c6c;
		}
		.yellow {
			background: #e6a23c;
		}
		.green {
			background: #67c23a;
		}
	}
	> .main {
		color: white;
		font-size: 15px;
		height: calc(100% - 62px);
		padding: 0 25px 25px;
		overflow-y: auto;
		overflow-x: hidden;
		.classifyName {
			color: #515970;
		}
	}
}
</style>
