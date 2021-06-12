<template>
	<div class="login">
		<a-typography-title :level="2" class="title">
			ONE PIECE
		</a-typography-title>
		<a-form :model="formState" layout="vertical">
			<a-form-item>
				<a-input
					v-model:value="formState.username"
					placeholder="请输入用户名"
				>
					<template #prefix>
						<UserOutlined />
					</template>
				</a-input>
			</a-form-item>
			<a-form-item>
				<a-input-password
					v-model:value="formState.password"
					placeholder="请输入密码"
					@pressEnter="onSubmit"
				>
					<template #prefix>
						<LockOutlined />
					</template>
				</a-input-password>
			</a-form-item>
			<a-form-item>
				<a-button
					type="primary"
					@click="onSubmit"
					block
					v-model:loading="store.loading"
				>
					登录
				</a-button>
			</a-form-item>
		</a-form>
	</div>
</template>
<script>
import { defineComponent, reactive, toRaw } from 'vue'
import { useRouter } from 'vue-router'
import { UserOutlined, LockOutlined } from '@ant-design/icons-vue'
import useStore from '@/store'
import useLoginStore from './store'

export default defineComponent({
	components: { UserOutlined, LockOutlined },
	setup() {
		const loginStore = useLoginStore()
		const store = useStore()
		const router = useRouter()
		const formState = reactive({ username: '', password: '' })
		const formData = toRaw(formState)
		const onSubmit = async () => {
			store.$state = { loading: true }
			await loginStore.userLogin(formData)
			await router.push('/admin/dashboard')
			store.$state = { loading: false }
		}
		return { formState, onSubmit, store, loginStore }
	}
})
</script>

<style lang="less" scope>
.login {
	.title {
		text-align: center;
		margin-bottom: 35px;
	}
	width: 300px;
	position: absolute;
	top: 50%;
	left: 50%;
	transform: translate(-50%, -50%);
}
</style>
