<template>
	<div :id="id" :style="{ width: '100%', height: '100%',minWidth:'200px',minHeight:'200px' }"></div>
</template>
<script>
	import api from '../api';
	export default {
		name: 'ElWechatLogin',
		data(){
			return{
				id:Math.random().toString(36).substr(2),
				wx: null,
			};
		},
		props: {
			wechat_code: {
				type: String,
				default: "8oaj0ph"
			},
		},

		mounted() {
			this.getData()
		},
		methods:{
			getData:function(){
				let _this = this;
				console.log('getData',_this)
				api.getWechatConfig.send({"code":this.wechat_code})
				.then(result => {
					console.log('result',result)
					if (result.state==2000) {
						_this.data = result.data
						_this.wx = new WxLogin({
							self_redirect:true,
							id:_this.id, 
							appid: result.data.appid, 
							scope: result.data.scope, 
							redirect_uri: result.data.auth_redirect_url,
							state: _this.wechat_code+'@'+window.location.href,
							style: 'black'
						});
					}
				})
				.catch(e => {
					console.log(e)
				})
			},
		},
	};
</script>
<style>
iframe {
	margin: 0 auto;
	width: 100%;
	top: 20px;
	position: absolute;
}
</style>
