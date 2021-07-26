<template>
	<div class="article">
		<a-typography-title :level="3" class="title">{{ articleInfo.title }}</a-typography-title>
		<a-typography-paragraph
			class="content"
			:ellipsis="{ rows: 2 }"
			:content="articleInfo.content"
		/>
		<div class="info">
			<a-space>
				<span><IconFont name="iconlishi" /> {{ articleInfo.created_at }}</span>
				<span><IconFont name="iconshuqian" /> {{ articleInfo.category }}</span>
			</a-space>
		</div>
	</div>
</template>

<script>
import { defineComponent } from 'vue'
import IconFont from '@/components/IconFont.vue'
export default defineComponent({
	components: { IconFont },
	props: {
		articleInfo: Object
	},
	setup({ articleInfo }) {
		return {
			articleInfo
		}
	}
})
</script>

<style lang="less" scoped>
.article {
	width: 100%;
	padding: 16px;
	border: 1px solid rgba(255, 255, 255, 0.3);
	border-radius: 8px;
	background: rgba(255, 255, 255, 0.3);
	box-shadow: 2px 2px 3px rgba(170, 170, 170, 0.3);
	backdrop-filter: blur(20px);
	.title {
		cursor: pointer;
		font-weight: bold;
		color: #5c5c5c;
	}
	.content {
		line-height: 2;
		color: #585858;
	}
	.info {
		color: #585858;
	}
}
</style>
