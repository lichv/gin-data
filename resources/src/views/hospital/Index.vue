<template>
	<div class="data-dashboard-page page_wrapper">
		<div class="column-card region-card">
			<el-area-select @select="handleSelectArea" ></el-area-select>
		</div>
		<div class="column-card">
			<el-point-hospital-table class="one-row-card-table" :adcode="baiduQuery" :keyword="tempQuery"></el-point-hospital-table>
		</div>
		<div class="column-card">
			<el-temp-table class="temp-table two-row-card-table" @select="handleSelectTemp" @move="handleMoveTemp" :keyword="tempQuery"></el-temp-table>
			<el-hospital-table class="hospital-table two-row-card-table" @select="handleSelect" :keyword="hospitalQuery"></el-hospital-table>
		</div>
		<div class="other-card">
			<el-hospital-form class="hospital-form" :item="selected"></el-hospital-form>
		</div>
	</div>
</template>
<script>
	import api from '../../api';
	export default {
		name: 'HelloWorld',
		data() {
			return {
				baiduQuery:'',
				gaodeQuery:'',
				tencentQuery:'',
				tempQuery:'',
				hospitalQuery:'',
				list:[],
				page:{
					page:1,
					size:50,
					total:1,
					last:1,
				},
				query:{
					page:1,
					size:50,
					order:'seq_no',
					sort :1,
				},
				selected:{

				},
				selectedTemp:{

				},
				query:'',
			}
		},
		props: {
			
		},
		watch:{
			$router:{
				handler(val,oldval){
					console.log('val',val)
					console.log('oldval',oldval)
				},
			},
		},
		mounted() {
			
		},
		methods: {
			handleSelectArea(item){
				console.log('handleSelectArea',item)
				this.baiduQuery = item.code
				this.gaodeQuery = item.code
				this.tencentQuery = item.code
				this.tempQuery = item.name
				this.hospitalQuery = item.name


			},
			handleSelect(row){
				console.log('handleSelect',row)
				this.selected = row
			},
			handleSelectTemp(temp){
				this.selectedTemp = temp
				this.query = temp.name
			},
			handleMoveTemp(temp){
				console.log('handleMoveTemp',temp)
				this.query = temp.name
			},
			copy(text){
				var _input = document.createElement("input");
				_input.value = text;
				document.body.appendChild(_input);
				_input.select();
				document.execCommand("Copy");
				document.body.removeChild(_input);
				this.$message({
					showClose: true,
					message: "复制成功",
					type: "success"
				});
			},
			
		}
	}
</script>
<style>
.data-dashboard-page{
	display: flex;
	min-width: 1460px;
}
.column-card{
	width: 300px;
	height: 100%;
}
.column-card + .column-card{
	margin-left: 6px;
}
.column-card.region-card{
	width: 200px;
}
.two-row-card-table{
	height: calc((100% - 6px)/2);
}
.two-row-card-table + .two-row-card-table{
	margin-top: 6px;
}
.three-row-card-table{
	height: calc((100% - 12px)/3);	
}
.three-row-card-table + .three-row-card-table{
	margin-top: 6px;
}
.other-card{
	flex: 1;
	margin-left: 8px;
}

.hospital-form{
	flex: 1;
	margin-left: 6px;
}
</style>
