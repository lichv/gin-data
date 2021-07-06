<template>
	<el-wechat-login></el-wechat-login>
</template>
<script>
	export default {
		name: 'Login',
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