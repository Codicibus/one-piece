<template>
	<v-md-editor
		height="calc(100vh - 385px)"
		v-model="articleListStore.formData.content"
		:disabled-menus="[]"
		@upload-image="handleUploadImage"
		@copy-code-success="handleCopyCodeSuccess"
	/>
</template>

<script>
import { defineComponent } from 'vue'
import useStore from '../../store'
const articleListStore = useStore()

export default defineComponent({
	name: 'editor',
	props: ['rules'],
	setup() {
		const handleUploadImage = (event, insertImage, files) => {
			// 拿到 files 之后上传到文件服务器，然后向编辑框中插入对应的内容
			console.log(event, files)
			// 此处只做示例
			insertImage({
				url: 'https://ss0.bdstatic.com/70cFvHSh_Q1YnxGkpoWK1HF6hhy/it/u=1269952892,3525182336&fm=26&gp=0.jpg',
				desc: files[0].name
			})
		}

		const handleCopyCodeSuccess = code => {
			console.log(code)
		}
		return {
			articleListStore,
			handleUploadImage,
			handleCopyCodeSuccess
		}
	}
})
</script>

<style lang="less" scoped>
.batchProcessing {
	margin-bottom: 16px;
	height: calc(100%);
}
</style>
