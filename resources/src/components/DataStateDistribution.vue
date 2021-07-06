<template>
	<el-card class="data-flag-distribution-card" >
		<template #header>
			<div class="card-header">
				<span>临时表 state状态</span>
			</div>
		</template>
		<div :id="id" :style="{ width:'400px',height:'320px' }"></div>
	</el-card>
</template>
<script>
	import { watch } from 'vue'
	import api from '../api';
	export default {
		name: 'ElDataStateDistribution',
		data(){
			return{
				id:Math.random().toString(36).substr(2),
				data:[],
				form:{},
				chartPie: {},
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
			this.getData()
			setInterval(function () {
				_this.getData()
			},5000)
		},
		methods:{
			initChartBar(domId, echartData) {
				var option = {

					dataZoom: [
					{
						type: "inside",
						throttle: "50",
						minValueSpan: 10,
						start: 100,
						end: 100,
						zoomLock: true,
					},
					],
					title: {
						text: echartData.title,
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
					tooltip: {
						trigger: "axis",
						axisPointer: {
							type: "none",
						},
					},
					xAxis: {
						type: "category",
						data: echartData.x,
						axisTick: {
							show: false,
						},
					},
					yAxis: {
						type: "value",
					},
					series: [
					{
						type: "bar",
						data: echartData.y,
					},
					],
				};

				let getchart = this.$echarts.init(document.getElementById(domId));

				getchart.setOption(option);
				getchart.resize();
				window.addEventListener("resize", () => {
					getchart.resize();
				});
			},
			initChartPie(domId, pData) {
				if (!this.chartPie[domId]) {
					let getchart = this.$echarts.init(document.getElementById(domId));

					var option = {
						tooltip: {
							trigger: "item",
						},
						toolbox: {
							show: true,
							orient: 'vertical',
							left: 'right',
							top: 'center',
							feature: {
								mark: {show: true},
								dataView: {show: true, readOnly: false},
								saveAsImage: {show: true}
							}
						},
						legend: {
							orient: 'vertical',
							left: 'left',
						},
						series: [
						{
							name: pData.name,
							type: "pie",
							radius: "50%",
							data: pData.data,
							emphasis: {
								itemStyle: {
									shadowBlur: 10,
									shadowOffsetX: 0,
									shadowColor: "rgba(0, 0, 0, 0.5)",
								},
							},
						},
						],
					};
					getchart.setOption(option);
					getchart.resize();
					window.addEventListener("resize", () => {
						getchart.resize();
					});
					console.log('创建数据')
					this.chartPie[domId] = getchart;
				} else {
					console.log('更新数据')
					var option = this.chartPie[domId].getOption();
					// option.series[0].data = [];
					// this.chartPie[domId].setOption(option);
					option.series[0].data = pData.data;
					this.chartPie[domId].setOption(option);
				}
			},
			initChartLine(xData, yData) {
				if (!this.chartLine) {
					let getchart = this.$echarts.init(
						document.getElementById("echart-line")
						);

					var option = {
						dataZoom: [
						{
							type: "inside",
							throttle: "50",
							minValueSpan: 5,
							start: 100,
							end: 100,
							zoomLock: true,
						},
						],
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
							data: xData,
							axisTick: {
								show: false,
							},
						},
						yAxis: {
							type: "value",
						},
						series: [
						{
							type: "line",
							data: yData,
						},
						],
					};
					getchart.setOption(option);
					window.addEventListener("resize", () => {
						getchart.resize();
					});
					getchart.resize();

					this.chartLine = getchart;
				} else {
					var option = this.chartLine.getOption();
					option.xAxis[0].data = [];
					option.series[0].data = [];
					this.chartLine.setOption(option);
					option.xAxis[0].data = xData;
					option.series[0].data = yData;
					option.dataZoom[0].start = 100;
					option.dataZoom[0].end = 100;
					this.chartLine.setOption(option);
				}
			},
			getData:function(){
				let _this = this;
				api.getDataStateDistribution.send(this.form)
				.then(result => {
					console.log('result',result)
					if (result.state==2000) {
						var data = []
						result.data.forEach(function(item){
							data.push({'name':item.name,'value':item.count})
						})
						_this.data = data
						_this.initChartPie(_this.id,{title:"状态分布",name:"状态分布",data:data})
					}
				})
				.catch(e => {
					console.log(e)
				})
			},
		},
	};
</script>
<style >
.data-state-distribution-card{
	width: 400px;
	height: 360px;
}
</style>