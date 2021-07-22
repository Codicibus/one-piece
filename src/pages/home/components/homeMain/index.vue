<template>
	<main class="homeMain">
		<div>
			<a-row class="card" :gutter="[0, 32]">
				<!-- 文章 -->
				<a-col :span="24">
					<home-article />
				</a-col>
				<a-col :span="24">
					<home-article />
				</a-col>
				<a-col :span="24">
					<home-article />
				</a-col>
				<!-- 分页 -->
				<a-col :span="24" class="pagination">
					<a-pagination
						show-quick-jumper
						v-model:current="current1"
						:total="500"
						@change="onChange"
					/>
				</a-col>
			</a-row>
		</div>
	</main>
</template>
<script>
import { defineComponent, ref } from 'vue'
import HomeArticle from './Article.vue'
export default defineComponent({
	components: { HomeArticle },
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
	.card {
		width: 100%;
		height: 100%;
		.pagination {
			text-align: center;
		}
	}
}
</style>
