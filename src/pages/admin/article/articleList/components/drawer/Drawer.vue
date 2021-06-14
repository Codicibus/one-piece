<template>
	<a-drawer
		title="编辑文章"
		width="50%"
		v-model:visible="articleListStore.visible"
		:body-style="{ paddingBottom: '80px' }"
		@close="onClose"
	>
		<a-form
			v-model:model="form"
			:rules="rules"
			layout="vertical"
			ref="formRef"
		>
			<a-row :gutter="16">
				<a-col :span="12">
					<a-form-item label="文章标题" name="articleTitle">
						<a-input
							v-model:value="form.articleTitle"
							:placeholder="rules.articleTitle.message"
							allow-clear
						/>
					</a-form-item>
				</a-col>
				<a-col :span="12">
					<a-form-item label="文章分类" name="articleClassify">
						<a-select
							v-model:value="form.articleClassify.value"
							mode="multiple"
							style="width: 100%"
							:placeholder="rules.articleClassify.message"
							:options="form.articleClassify.options"
							allowClear
						/>
					</a-form-item>
				</a-col>
			</a-row>
			<a-row :gutter="16">
				<a-col :span="12">
					<a-form-item label="文章背景图">
						<a-switch
							v-model:checked="form.fixedState"
							@change="
								!form.fixedState && (form.randomState = false)
							"
						/>
					</a-form-item>
				</a-col>
				<a-col :span="12">
					<a-form-item label="文章随机图">
						<a-switch
							v-model:checked="form.randomState"
							:disabled="!form.fixedState"
						/>
					</a-form-item>
				</a-col>
			</a-row>
			<a-row :gutter="16" v-if="form.fixedState && !form.randomState">
				<a-col :span="24">
					<a-form-item name="articleBackgroundImage">
						<a-upload
							:action="uploadURL"
							list-type="picture"
							name="local_file"
							:headers="{ Authorization: getToken() }"
							v-model:file-list="form.fileList"
							@change="handleChange"
						>
							<a-button>
								<upload-outlined></upload-outlined>
								上传文章背景图
							</a-button>
						</a-upload>
					</a-form-item>
				</a-col>
			</a-row>
			<a-row :gutter="16">
				<a-col :span="24">
					<a-form-item label="文章内容" name="articleContent">
						<a-textarea
							v-model:value="form.articleContent"
							:rows="4"
							:placeholder="rules.articleContent.message"
						/>
					</a-form-item>
				</a-col>
			</a-row>
		</a-form>
		<div class="footer">
			<a-button style="margin-right: 8px" @click="onClose">返回</a-button>
			<a-button type="primary" @click="submit">完成</a-button>
		</div>
	</a-drawer>
</template>
<script>
import {
	PlusOutlined,
	UploadOutlined,
	LoadingOutlined
} from '@ant-design/icons-vue'
import { defineComponent, toRaw, ref } from 'vue'
import { message } from 'ant-design-vue'
import useStore from '../../store'
import { getToken } from '@/utils/auth'
const articleListStore = useStore()
const { form } = articleListStore
const formRef = ref()
const options = []
for (let i = 0; i < 1000; i++) {
	const value = `${i.toString(36)}${i}`
	options.push({
		value,
		disabled: i === 10
	})
}

export default defineComponent({
	components: {
		PlusOutlined,
		UploadOutlined,
		LoadingOutlined
	},
	setup() {
		const rules = {
			articleTitle: { required: true, message: '请输入文章标题' },
			articleClassify: {
				required: true,
				type: 'object',
				message: '请输入或选择文章分类'
			},
			articleContent: { required: true, message: '请输入文章内容' }
		}

		articleListStore.$patch(state => {
			state.form.articleClassify = {
				options,
				value: ['a10', 'c12']
			}
		})

		const handleChange = info => {
			if (info.file.status !== 'uploading') {
				console.log(info.file, info.fileList)
			}
			if (info.file.status === 'done') {
				message.success(`${info.file.name} 图片上传成功`)
			} else if (info.file.status === 'error') {
				message.error(`${info.file.name} 图片上传失败`)
			}
		}

		const onClose = () => articleListStore.$patch({ visible: false })

		const submit = () => {
			formRef.value
				.validate()
				.then(() => {
					console.log('values', toRaw(form))
					articleListStore.$patch({ visible: false })
				})
				.catch(error => {
					return console.error('error', error)
				})
		}

		return {
			articleListStore,
			form,
			formRef,
			rules,
			onClose,
			submit,
			uploadURL: import.meta.env.VITE_UPLOAD_URL,
			getToken,
			handleChange
		}
	}
})
</script>

<style lang="less" scoped>
.footer {
	position: absolute;
	right: 0;
	bottom: 0;
	width: 100%;
	border-top: 1px solid #e9e9e9;
	padding: 10px 16px;
	background: #fff;
	text-align: right;
	z-index: 1;
}
</style>
