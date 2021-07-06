<template>
	<div>404</div>
</template>
<script>
	export default {
		name: 'Error',
		data() {
			return {
				count: 0,
				drawer: true,
				direction: 'ltr',
				aside_size:'mini',
				isCollapse:false
			}
		},
		props: {
			msg: String
		},
		methods: {

		}
	}
</script>