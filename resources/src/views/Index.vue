<template>
	<el-container class="page-container">
		<el-header class="layout-header" height="56px">
			<div class="layout-header-left">
				<div class="layout-header-handle">
					<div class="header-left-btn" @click="handleClickBtn"><i class="el-icon-menu"></i></div>
					<div class="height-left-logo">
						<img class="header-logo" src="https://chv.oss-cn-shanghai.aliyuncs.com/aegicare/logo_aegicare.png">
					</div>
				</div>
				<el-drawer class="header-left-menu" title="" v-model="drawer" :direction="direction" :before-close="handleClose" destroy-on-close :size="asidesize"  :modal="true" :show-close="true" >
					<div class="asider-menu">
						<el-asider-menu :is-collapse="isCollapse" @select="handleSelect"></el-asider-menu>
						<div class="aside-bottom-btn" @click="showAside">
							<span v-if="isCollapse" class="el-icon-s-unfold"></span>
							<span v-else class="el-icon-s-fold"></span>
						</div>
					</div>
				</el-drawer>
			</div>
			<div class="layout-header-middle">
				
			</div>
			<div class="layout-header-right">
				<router-link to="/" ><i class="el-icon-s-home"></i></router-link>
			</div>
		</el-header>
		<el-main class="layout-main">
			<el-card class="dashboard-card calendar-card">
				<el-calendar v-model="today"></el-calendar>
			</el-card>
			<el-card class="dashboard-card">
				<router-link to="/hospital" >医院管理</router-link>
			</el-card>
			<el-card class="dashboard-card">
				<router-link to="/hospital" >医院管理</router-link>
			</el-card>
			
		</el-main>
	</el-container>
</template>
<script>
	export default {
		name: 'Layout',
		data() {
			return {
				today: new Date(),
				count: 0,
				drawer: false,
				direction: 'ltr',
				aside_size:'mini',
				isCollapse:false
			}
		},
		props: {
			msg: String
		},
		computed:{
			asidesize:function(){
				if (this.aside_size=='mini') {
					this.isCollapse=true
					return "64px"
				}else{
					this.isCollapse=false
					return "240px"
				}
			},
		},
		watch:{
			$router:{
				handler(val,oldval){
					console.log('val',val)
					console.log('oldval',oldval)
				},
			},
			isCollapse(val,oldval) {
				this.showCollapse = val;
			},
		},
		
		methods: {
			showAside(){
				if (this.aside_size=='mini') {
					this.aside_size='normal'
				}else{
					this.aside_size='mini' 
				}
				this.isCollapse = !this.isCollapse;
			},
			handleClickBtn(){
				this.drawer = !this.drawer;
				console.log('this.drawer',this.drawer)
			},
			handleClose(done) {
				console.log('done',done)
				this.drawer = false;
			},
			handleSelect(val){
				console.log('handleSelect',val)
				this.drawer = false
			}
		}
	}
</script>
<style>
.page-container{
	height: 100%;
	width: 100%;
	margin:0;
	padding: 0;
	position: relative;
}
.layout-header{
	height: 56px;
	background: #fff;
	display: flex;
	padding: 0;
	border-bottom: 1px solid #d6d6d6;
}
.layout-header-left{
	width: 256px;
}
.layout-header-handle{
	width: 100%;
	height: 100%;
	display: flex;
}
.height-left-logo{
	padding: 4px 10px;
}
.header-logo{
	max-height: 48px;
}
.header-left-btn {
	min-width: 64px;
	width: 64px;
	height: 100%;
	background-color: #001529;
	color: #fff;
}
.layout-header-left .el-drawer{
	height: calc(100% - 56px);
	top: 56px;
	background-color: #001529;
}
.layout-header-left .el-drawer > .el-drawer__header{
	margin-bottom: 0px;
}
.layout-header-left .el-drawer > .el-drawer__body{
	height: calc(100% - 43px);
}
.layout-header-left .el-drawer .el-menu{
	border-right:0;
}
.header-left-btn i {
	color: #fff;
	font-size: 32px;
	line-height: 56px;
	padding-left: 14px;
}
.layout-header-middle{
	flex: 1;
}
.layout-header-right{
	width: auto;
	min-width: 160px;
	display: flex;
}

.layout-main{
	padding: 6px 6px 6px 12px;
	background-color: rgb(240, 249, 235);
}

.asider-menu{
	max-height: calc(100% - 50px);
	overflow-y: auto;
	overflow-x: hidden;
	position: relative;
}
.aside-bottom-btn{
	color: #fff;
	position: fixed;
	bottom: 12px;
	left: 20px;
	font-size: 18px;
}
.dashboard-card{
	max-width: 360px;
	width: 360px;
	margin-bottom: 12px;
	float: left;
}
.dashboard-card + .dashboard-card{
	margin-left: 12px;
}
.layout-header-right {
	line-height: 50px;
	font-size: 32px;
}
.calendar-card{
	width: 480px;
	height: 400px;
}
.calendar-card > .el-card__body{
	padding: 2px;
}
.calendar-card .el-calendar-table .el-calendar-day{
	height: 40px;
}
</style>
