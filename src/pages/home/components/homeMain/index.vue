<template>
	<!-- <a-spin size="large" wrapperClassName="spin" :spinning="homeMainStore.articles"> -->
	<main class="homeMain">
		<!-- <a-back-top :visibilityHeight="10" /> -->
		<div>
			<a-row class="card" :gutter="[0, 32]">
				<!-- 文章 -->
				<a-col :span="24" v-for="item in homeMainStore.articles" :key="item.content_hash">
					<home-article :articleInfo="item" />
				</a-col>
				<!-- 分页 -->
				<a-col :span="24" class="pagination">
					<a-pagination
						show-quick-jumper
						v-model:current="homeMainStore.articleInfo.page_num"
						:total="homeMainStore.articleInfo.total"
						@change="onChange"
					/>
				</a-col>
			</a-row>
		</div>
	</main>
	<!-- </a-spin> -->
</template>
<script>
import { defineComponent, ref } from 'vue'
import HomeArticle from './Article.vue'
import useStore from './store'
export default defineComponent({
	components: { HomeArticle },
	setup() {
		const homeMainStore = useStore()
		// 获取文章
		homeMainStore.getArticle()
		const current1 = ref(1)
		const onChange = pageNumber => {
			console.log('Page: ', pageNumber)
			// homeMainStore.$patch({
			// 	pagingSetup: {
			// 		page_num: pageNumber
			// 	}
			// })
			// 获取文章
			homeMainStore.getArticle()
		}
		return {
			homeMainStore,
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
	.spin {
		height: 80%;
	}
	.card {
		width: 100%;
		height: 100%;
		.pagination {
			text-align: center;
		}
	}
}
</style>
