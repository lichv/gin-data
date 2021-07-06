<template>
	<div class="hospital-page page_wrapper">
		<el-temp-table class="temp-table" @select="handleSelectTemp" @move="handleMoveTemp"></el-temp-table>
		<el-hospital-table class="hospital-table" @select="handleSelect" :keyword="query"></el-hospital-table>
		<el-hospital-form class="hospital-form" :item="selected"></el-hospital-form>
	</div>
</template>
<script>
	import api from '../../api';
	export default {
		name: 'HelloWorld',
		data() {
			return {
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
.hospital-page{
	display: flex;
}
.temp-table{
	width: 300px;
	min-width: 300px;
}
.hospital-table{
	width: 300px;
	min-width: 300px;
	margin-left: 6px;
}
.hospital-form{
	flex: 1;
	margin-left: 6px;
}
</style>
