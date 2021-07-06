<template>
	<a-drawer
		title="编辑文章"
		width="60%"
		height="100vh"
		v-model:visible="articleListStore.visible"
		:body-style="{ paddingBottom: '80px' }"
		@onClose="closeDrawer"
	>
		<a-form
			v-model:model="articleListStore.formData"
			:rules="rules"
			layout="vertical"
			ref="formRef"
		>
			<a-row :gutter="16">
				<a-col :span="12">
					<a-form-item label="文章标题" name="title">
						<a-input
							v-model:value="articleListStore.formData.title"
							:placeholder="rules.title.message"
							allow-clear
						/>
					</a-form-item>
				</a-col>
				<a-col :span="12">
					<a-form-item label="文章分类" name="category">
						<a-select
							v-model:value="
								articleListStore.formData.category.value
							"
							mode="multiple"
							style="width: 100%"
							:placeholder="rules.category.message"
							:options="
								articleListStore.formData.category.options
							"
							allowClear
						/>
					</a-form-item>
				</a-col>
			</a-row>
			<a-row :gutter="16">
				<a-col :span="12">
					<a-form-item label="文章背景图">
						<a-switch
							v-model:checked="
								articleListStore.formData.background_visible
							"
							@change="
								!articleListStore.formData.background_visible &&
									(articleListStore.formData.background_random = false)
							"
						/>
					</a-form-item>
				</a-col>
				<a-col :span="12">
					<a-form-item label="文章随机图">
						<a-switch
							v-model:checked="
								articleListStore.formData.background_random
							"
							:disabled="
								!articleListStore.formData.background_visible
							"
						/>
					</a-form-item>
				</a-col>
			</a-row>
			<a-row
				:gutter="16"
				v-if="
					articleListStore.formData.background_visible &&
					!articleListStore.formData.background_random
				"
			>
				<a-col :span="24">
					<a-form-item name="articleBackgroundImage">
						<a-upload
							action="/v1/upload/pic"
							list-type="picture"
							name="local_file"
							:headers="{ Authorization: getToken() }"
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
					<a-form-item label="文章内容" name="content">
						<Editor :rules="rules" />
					</a-form-item>
				</a-col>
			</a-row>
		</a-form>
		<div class="footer">
			<a-button style="margin-right: 8px" @click="closeDrawer"
				>返回</a-button
			>
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
import Editor from '../editor/index.vue'
import useStore from '../../store'
import { getToken } from '@/utils/auth'
const articleListStore = useStore()
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
		LoadingOutlined,
		Editor
	},
	setup() {
		const rules = {
			title: { required: true, message: '请输入文章标题' },
			category: {
				required: true,
				type: 'object',
				message: '请输入或选择文章分类'
			},
			content: { required: true, message: '请输入文章内容' }
		}

		articleListStore.$patch(state => {
			state.formData.category = {
				options,
				value: ['a10', 'c12']
			}
		})

		console.log(toRaw(articleListStore.formData.category))

		const closeDrawer = () => articleListStore.$patch({ visible: false })

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

		const submit = () => {
			formRef.value
				.validate()
				.then(() => {
					console.log('values', toRaw(articleListStore.formData))
					articleListStore.closeDrawer
				})
				.catch(error => {
					return console.error('error', error)
				})
		}

		return {
			articleListStore,
			formRef,
			rules,
			getToken,
			closeDrawer,
			submit,
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
