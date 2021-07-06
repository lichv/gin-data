<template>
	<el-card class="hospital-form page_wrapper">
		<template #header>
			<div class="card-header">
				<div class="card-header-left">
					<span>医院信息</span>
					<span v-if="'id' in form && form.id != '' && form.id != 0">-{{form.id}}</span>
				</div>
				<div class="card-header-right">
					<el-button type="primary" @click="handlerSubmit" size="mini">提交</el-button>
					<el-button  @click="handlerReset" size="mini">重置</el-button>
				</div>
			</div>
		</template>
		<el-form ref="form" :inline="true" :model="form" label-width="80px">
			<el-form-item label="医院编号" >
				<el-input v-model="form.code"></el-input>
			</el-form-item>
			<el-form-item label="医院名称" >
				<el-input v-model="form.name"></el-input>
			</el-form-item>
			<el-form-item label="医院别称" >
				<el-input v-model="form.alias"></el-input>
			</el-form-item>
			<el-form-item label="医院logo" >
				<el-input v-model="form.logo"></el-input>
			</el-form-item>
			<el-form-item label="医院等级" >
				<el-select v-model="form.grade" filterable allow-create default-first-option placeholder="请选择医院等级">
					<el-option v-for="item in grade_options" :key="item.value" :label="item.label" :value="item.value"></el-option>
				</el-select>
			</el-form-item>
			<el-form-item label="医院性质" >
				<el-select v-model="form.nature" filterable allow-create default-first-option placeholder="请选择医院性质">
					<el-option key="公立医院" label="公立医院" value="公立医院"></el-option>
					<el-option key="民营医院" label="民营医院" value="民营医院"></el-option>
				</el-select>
			</el-form-item>
			<el-form-item label="医院类型" >
				<el-select v-model="form.type" filterable allow-create default-first-option placeholder="请选择医院类型">
					<el-option key="综合医院" label="综合医院" value="综合医院"></el-option>
					<el-option key="专科医院" label="专科医院" value="专科医院"></el-option>
				</el-select>
			</el-form-item>
			<el-form-item label="成立年份" >
				<el-input v-model="form.regyear" placeholder="例：2016年"></el-input>
			</el-form-item>
			<el-form-item label="机构规模" >
				<el-input v-model="form.size" placeholder="例：1000 人"></el-input>
			</el-form-item>
			<el-form-item label="信用代码" >
				<el-input v-model="form.creditcode"></el-input>
			</el-form-item>
			<el-form-item label="法定代表人" >
				<el-input v-model="form.corporations"></el-input>
			</el-form-item>
			<el-form-item label="公司类型" >
				<el-input v-model="form.typename"></el-input>
			</el-form-item>
			<el-form-item label="网址" >
				<el-input v-model="form.url"></el-input>
			</el-form-item>
			<el-form-item label="电话" >
				<el-input v-model="form.phone" placeholder="例：13012345678"></el-input>
			</el-form-item>
			<el-form-item label="邮箱" >
				<el-input v-model="form.email" placeholder="例：88888@qq.com"></el-input>
			</el-form-item>
			<el-form-item label="所在省" >
				<el-input v-model="form.province" placeholder="例：xx省"></el-input>
			</el-form-item>
			<el-form-item label="所在市" >
				<el-input v-model="form.city" placeholder="例：xx市"></el-input>
			</el-form-item>
			<el-form-item label="所在县区" >
				<el-input v-model="form.town" placeholder="例：xx县"></el-input>
			</el-form-item>
			<el-form-item label="详细地址" >
				<el-input v-model="form.address" placeholder="例：xx街道xx小区xx楼xx房"></el-input>
			</el-form-item>
			<el-form-item label="经纬度" >
				<el-input v-model="form.location" placeholder="例：116.659893,39.9013"></el-input>
			</el-form-item>
			<el-form-item label="在职员工" >
				<el-input v-model="form.members" placeholder="例：100 人"></el-input>
			</el-form-item>
			<el-form-item label="核定床位数" >
				<el-input v-model="form.beds" placeholder="例：100 张"></el-input>
			</el-form-item>
			<el-form-item label="年急诊量" >
				<el-input v-model="form.emergency" placeholder="例：10000 次"></el-input>
			</el-form-item>
			<el-form-item label="年住院人次" >
				<el-input v-model="form.inpatients" placeholder="例：1000 人"></el-input>
			</el-form-item>
			<el-form-item label="来源" >
				<el-input v-model="form.source" placeholder="网址，例：https://www.tianyancha.com/company/3127473242"></el-input>
			</el-form-item>
			<el-form-item label="天眼查" >
				<el-input v-model="form.tianyancha" placeholder="网址，例：https://www.tianyancha.com/company/3127473242"></el-input>
			</el-form-item>
			<el-form-item label="医学百科" >
				<el-input v-model="form.yixue" placeholder="网址，例：https://www.tianyancha.com/company/3127473242"></el-input>
			</el-form-item>
			<el-form-item label="时间">
				<el-date-picker v-model="form.addtime" type="datetime" placeholder="选择日期时间"> </el-date-picker>
			</el-form-item>
			<el-form-item label="预留标记">
				<el-checkbox-group v-model="form.flag" size="mini" >
					<el-checkbox-button label="有效" :key="1">有效</el-checkbox-button>
					<el-checkbox-button label="无效" :key="2">无效</el-checkbox-button>
				</el-checkbox-group>
			</el-form-item>
			<el-form-item label="有效标记">
				<el-checkbox-group v-model="form.state" size="mini" >
					<el-checkbox-button :label="1" :key="1">有效</el-checkbox-button>
					<el-checkbox-button :label="2" :key="2">无效</el-checkbox-button>
				</el-checkbox-group>
			</el-form-item>
			<el-form-item label="通过标记">
				<el-checkbox-group v-model="form.status" size="mini" >
					<el-checkbox-button :label="1" :key="1">通过</el-checkbox-button>
					<el-checkbox-button :label="2" :key="2">未通过</el-checkbox-button>
				</el-checkbox-group>
			</el-form-item>
		</el-form>
		<el-form ref="form" :model="form" label-width="80px">
			<el-form-item label="医院介绍" >
				<el-input type="textarea" v-model="form.introduction" :rows="6"></el-input>
			</el-form-item>
			<el-form-item label="其他信息" >
				<el-input type="textarea" v-model="form.others" :rows="4"></el-input>
			</el-form-item>
		</el-form>
	</el-card>
</template>
<script>
	import { watch } from 'vue'
	import api from '../api';
	export default {
		name: 'ElHospitalForm',
		data(){
			return{
				form: {
					"id":0,
					"code":"",
					"name":"",
					"alias":"",
					"logo":"",
					"grade":"",
					"nature":"",
					"type":"",
					"regyear":"",
					"size":"",
					"province":"",
					"city":"",
					"town":"",
					"address":"",
					"location":"",
					"members":"",
					"beds":"",
					"emergency":"",
					"inpatients":"",
					"creditcode":"",
					"corporations":"",
					"typename":"",
					"introduction":" ",
					"others":"",
					"source":"",
					"url":"",
					"phone":"",
					"email":"",
					"addtime":"",
					"flag":1,
					"state":1,
					"status":0
				},
				grade_options: [
				{
					value: '三级甲等',
					label: '三级甲等'
				}, 
				{
					value: '三级乙等',
					label: '三级乙等'
				}, 
				{
					value: '三级丙等',
					label: '三级丙等'
				},
				{
					value: '三级医院',
					label: '三级医院'
				}, 
				{
					value: '二级甲等',
					label: '二级甲等'
				}, 
				{
					value: '二级乙等',
					label: '二级乙等'
				},
				{
					value: '二级丙等',
					label: '二级丙等'
				},
				{
					value: '二级医院',
					label: '二级医院'
				},
				{
					value: '一级甲等',
					label: '一级甲等'
				}, 
				{
					value: '一级乙等',
					label: '一级乙等'
				},
				{
					value: '一级丙等',
					label: '一级丙等'
				},
				{
					value: '一级医院',
					label: '一级医院'
				},
				],

			};
		},
		props: {
			item:{
				type: Object,
				default:function(){
					return {
						name: '',
						region: '',
						date1: '',
						date2: '',
						delivery: false,
						type: [],
						resource: '',
						desc: '',
						flag:1,
						state:1,
						status:1,
					};
				}
			},
		},

		mounted() {
			let _this= this
			this.data = this.form
			watch(
				() => this.item,
				(toValue, previousValue) => {
					_this.form = toValue
					console.log('watch_value',_this.form)
				}
				)
		},
		methods:{
			handlerReset(){
				this.form = {
					"id":0,
					"code":"",
					"name":"",
					"alias":"",
					"logo":"",
					"grade":"",
					"nature":"",
					"type":"",
					"regyear":"",
					"size":"",
					"province":"",
					"city":"",
					"town":"",
					"address":"",
					"location":"",
					"members":"",
					"beds":"",
					"emergency":"",
					"inpatients":"",
					"creditcode":"",
					"corporations":"",
					"typename":"",
					"introduction":" ",
					"others":"",
					"source":"",
					"url":"",
					"phone":"",
					"email":"",
					"addtime":"",
					"flag":1,
					"state":1,
					"status":0
				}
			},
			handlerSubmit(){
				const _this = this;
				api.saveHospital.send(this.form)
				.then(result => {
					console.log("save的结果",result)
					if (result.state==2000) {
						_this.$emit('input',_this.form)
						_this.$message({
							showClose: true,
							message: "修改成功",
							type: "success"
						});
					}
				})
				.catch(e => {
					console.log(e)
				})
			},
		}
	};
</script>
<style>
.hospital-form .card-header{
	display: flex;
}
.hospital-form .card-header .card-header-left{
	flex :1;
}
.hospital-form .card-header .card-header-right{
	width: 128px;
}
.hospital-form .el-form-item__label{
	font-size: 10px;
}
.hospital-form .el-form--inline .el-form-item--small .el-form-item__content{
	width: 200px;
}
.hospital-form .el-form--inline .el-date-editor.el-input{
	width: 200px;
}
</style>
