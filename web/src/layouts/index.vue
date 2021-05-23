<template>
	<a-layout style="min-height: 100vh">
		<layoutSidebar />
		<a-layout>
			<layoutHeader />
			<layoutMain />
			<layoutFooter />
		</a-layout>
	</a-layout>
</template>

<script setup>
import layoutHeader from './components/layoutHeader/index.vue'
import layoutSidebar from './components/layoutSidebar/index.vue'
import layoutMain from './components/layoutMain/index.vue'
import layoutFooter from './components/layoutFooter/index.vue'
</script>
