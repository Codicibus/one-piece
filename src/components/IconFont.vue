<template>
	<icon-font class="icon" :type="name" />
</template>
<script>
import { createFromIconfontCN } from '@ant-design/icons-vue'
import { defineComponent } from 'vue'
const IconFont = createFromIconfontCN({
	scriptUrl: '//at.alicdn.com/t/font_2269916_7sel28rjojq.js'
})
export default defineComponent({
	components: {
		IconFont
	},
	props: {
		name: String
	},
	setup({ name }) {
		return { name }
	}
})
</script>
