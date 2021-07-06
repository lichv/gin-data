<template>
	<div>
		<div>
			<el-form :inline="true" :model="query" class="demo-form-inline">
				<el-form-item label="时间范围">
					<el-date-picker v-model="datepicker" type="daterange" align="right" unlink-panels range-separator="至" start-placeholder="开始日期" end-placeholder="结束日期" :shortcuts="shortcuts" @change="handleChange"     >     </el-date-picker>
				</el-form-item>
				<el-form-item label="">
					<el-button size="mini" round  icon="el-icon-search" @click="getData" type="primary"></el-button>
				</el-form-item>
			</el-form>
		</div>
		<div>
			<div :id="grade_id" :style="{ width:'1440px',height:'300px' }"></div>
			<div :id="wgs_id" :style="{ width:'1440px',height:'300px' }"></div>
			<div :id="wes_id" :style="{ width:'1440px',height:'300px' }"></div>
			<div :id="plastosome_id" :style="{ width:'1440px',height:'300px' }"></div>
			<div :id="sanger_id" :style="{ width:'1440px',height:'300px' }"></div>
			<div :id="anti_epileptic_id" :style="{ width:'1440px',height:'300px' }"></div>
			<div :id="epileptic_gene_1_0_id" :style="{ width:'1440px',height:'300px' }"></div>
			<div :id="besides_id" :style="{ width:'1440px',height:'300px' }"></div>
		</div>
	</div>
</template>
<script>
	import { watch } from 'vue'
	import dayjs from 'dayjs';
	import api from '../api';
	export default {
		name: 'ElMonitorMedicineStatistics',
		data(){
			return{
				shortcuts: [{
					text: '最近一周',
					value: (() => {
						const end = new Date()
						const start = new Date()
						start.setTime(start.getTime() - 3600 * 1000 * 24 * 7)
						return [start, end]
					})(),
				}, {
					text: '最近一个月',
					value: (() => {
						const end = new Date()
						const start = new Date()
						start.setTime(start.getTime() - 3600 * 1000 * 24 * 30)
						return [start, end]
					})(),
				}, {
					text: '最近三个月',
					value: (() => {
						const end = new Date()
						const start = new Date()
						start.setTime(start.getTime() - 3600 * 1000 * 24 * 90)
						return [start, end]
					})(),
				}, {
					text: '最近六个月',
					value: (() => {
						const end = new Date()
						const start = new Date()
						start.setTime(start.getTime() - 3600 * 1000 * 24 * 180)
						return [start, end]
					})(),
				}, {
					text: '最近一年',
					value: (() => {
						const end = new Date()
						const start = new Date()
						start.setTime(start.getTime() - 3600 * 1000 * 24 * 365)
						return [start, end]
					})(),
				}],
				datepicker:[],
				id:Math.random().toString(36).substr(2),
				wgs_id:Math.random().toString(36).substr(2),
				wes_id:Math.random().toString(36).substr(2),
				plastosome_id:Math.random().toString(36).substr(2),
				sanger_id:Math.random().toString(36).substr(2),
				anti_epileptic_id:Math.random().toString(36).substr(2),
				epileptic_gene_1_0_id:Math.random().toString(36).substr(2),
				besides_id:Math.random().toString(36).substr(2),
				grade_id:Math.random().toString(36).substr(2),
				data:[],
				medicines:[],
				form:{},
				chartPie: {},
				wgs:{},
				wes:{},
				plastosome:{},
				sanger:{},
				anti_epileptic:{},
				epileptic_gene_1_0:{},
				besides:{},
				grade:{},
				dates:[],
				query:{},
				begin:dayjs().add(-3, 'month').format('YYYY-MM-DD'),
				end:dayjs().format('YYYY-MM-DD'),
			};
		},
		props: {
			query:{
				type:Object,
				default:function(){
					return {"timely": "720"}
				},
			}
		},

		mounted() {
			let _this = this;
			this.form = this.query
			if ('page' in this.form) {
				delete this.form['page']
			}
			if ('size' in this.form) {
				delete this.form['size']
			}
			if ('order' in this.form) {
				delete this.form['order']
			}
			if ('sort' in this.form) {
				delete this.form['sort']
			}
			var today = dayjs().format('YYYY-MM-DD')
			var begin = dayjs().add(-3, 'month').format('YYYY-MM-DD')
			this.datepicker = [dayjs().add(-3,'month'),dayjs()]
			this.getData(begin,today)
		},
		methods:{
			initChartLine(domId,title, xData, yData,legend) {
				let getchart = this.$echarts.init(
					document.getElementById(domId)
					);

				var option = {
					color:['#5470c6', '#91cc75', '#fac858', '#ee6666', '#73c0de', '#3ba272', '#fc8452', '#9a60b4', '#ea7ccc'],
					title: {
						text: title
					},
					legend: {
						data: legend
					},
					tooltip: {
						trigger: "axis",
						axisPointer: {
							type: "none",
						},
					},
					toolbox: {
						show: true,
						orient: 'vertical',
						left: 'right',
						top: 'center',
						feature: {
							mark: {show: true},
							dataView: {show: true, readOnly: false},
							magicType: {show: true, type: ['line', 'bar', 'stack', 'tiled']},
							restore: {show: true},
							saveAsImage: {show: true}
						}
					},
					xAxis: {
						type: "category",
						data: xData
					},
					yAxis: {
						type: "value",
					},
					series: yData,
				};
				getchart.setOption(option);
				window.addEventListener("resize", () => {
					getchart.resize();
				});
				getchart.resize();

			},
			handlerReset(){
				this.getData()
			},
			getData:function(){
				let _this = this;
				api.getMedicineStatistics.send({'begin':this.begin,'end':this.end})
				.then(result => {
					
					if (result.state==2000) {
						_this.data = result.data
						_this.medicines = result.data.medicines
						_this.wgs = result.data.wgs
						_this.wes = result.data.wes
						_this.plastosome = result.data.plastosome
						_this.sanger = result.data.sanger
						_this.anti_epileptic = result.data.anti_epileptic
						_this.epileptic_gene_1_0 = result.data.epileptic_gene_1_0
						_this.besides = result.data.besides
						_this.grade = result.data.grade
						_this.dates = result.data.dates

						_this.initChartLine(_this.grade_id,"分数",result.data.dates,result.data.grade,result.data.medicines)
						_this.initChartLine(_this.wgs_id,"WGS",result.data.dates,result.data.wgs,result.data.medicines)
						_this.initChartLine(_this.wes_id,"WES",result.data.dates,result.data.wes,result.data.medicines)
						_this.initChartLine(_this.plastosome_id,"线粒体",result.data.dates,result.data.plastosome,result.data.medicines)
						_this.initChartLine(_this.sanger_id,"Sanger",result.data.dates,result.data.sanger,result.data.medicines)
						_this.initChartLine(_this.anti_epileptic_id,"抗癫痫药物",result.data.dates,result.data.anti_epileptic,result.data.medicines)
						_this.initChartLine(_this.epileptic_gene_1_0_id,"癫痫基因1.0",result.data.dates,result.data.epileptic_gene_1_0,result.data.medicines)
						_this.initChartLine(_this.besides_id,"其他项目",result.data.dates,result.data.besides,result.data.medicines)

						console.log('result',result.data)

					}
				})
				.catch(e => {
					console.log(e)
				})
			},
			handleChange(val){
				var begin = dayjs(val[0]).format('YYYY-MM-DD')
				var end = dayjs(val[1]).format('YYYY-MM-DD')
				console.log('begin',begin)
				console.log('end',end)
				this.begin = begin
				this.end = end
				this.getData()

			}
		},
	};
</script>
