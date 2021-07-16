<template>
	<HomeHeader />
	<HomeMain />
</template>
<script>
import { defineComponent } from 'vue'
import HomeHeader from './components/homeHeader/index.vue'
import HomeMain from './components/homeMain/index.vue'
export default defineComponent({
	components: { HomeHeader, HomeMain },
	setup() {}
})
</script>

<style lang="less" scope>
footer {
	padding: 20px;
	text-align: center;
}
</style>
