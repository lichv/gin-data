<template>
	<el-card class="monitor_page page_wrapper">
		<div class="page_header">
			<el-form :inline="true" :model="formInline" class="demo-form-inline">
				<el-form-item label="用户">
					<el-input v-model="formInline.user" placeholder="AP号 | AF号"></el-input>
				</el-form-item>
				<el-form-item label="区域">
					<el-select v-model="formInline.region" placeholder="区域">
						<el-option label="区域一" value="shanghai"></el-option>
						<el-option label="区域二" value="beijing"></el-option>
					</el-select>
				</el-form-item>
				<el-form-item>
					<el-button type="primary" size="mini" @click="onSubmit">查询</el-button>
					<el-button type="primary" size="mini" @click="onSubmit">导出</el-button>
				</el-form-item>
			</el-form>
		</div>
		<div class="page_body">
			<el-table :data="tableData" style="width: 100%">
				<el-table-column prop="date" label="日期" width="180"></el-table-column>
				<el-table-column prop="name" label="姓名" width="180"></el-table-column>
				<el-table-column prop="address" label="地址"></el-table-column>
			</el-table>
		</div>
		<div class="page_footer">
			<div class="batch-handle">
				<div class="selected-title">已选中 <span class="selected—count">5</span> 项</div>
				<div class="vux-flexbox selection-items-box vux-flex-row">
					<el-button size="mini">批量转移</el-button>
					<el-button size="mini">批量审核</el-button>
					<el-button size="mini">批量删除</el-button>
				</div>
			</div>
			<div class="pagination">
				<el-pagination @size-change="handleSizeChange" @current-change="handleCurrentChange" :current-page="currentPage4" :page-sizes="[100, 200, 300, 400]" :page-size="100" layout="total, sizes, prev, pager, next, jumper" :total="400">
				</el-pagination>
			</div>
		</div>
	</el-card>
</template>
<script>
	export default {
		name: 'Demo',
		data() {
			return {
				formInline: {
					user: '',
					region: ''
				},
				tableData: [{
					date: '2016-05-02',
					name: '王小虎',
					address: '上海市普陀区金沙江路 1518 弄'
				}, {
					date: '2016-05-04',
					name: '王小虎',
					address: '上海市普陀区金沙江路 1517 弄'
				}, {
					date: '2016-05-01',
					name: '王小虎',
					address: '上海市普陀区金沙江路 1519 弄'
				}, {
					date: '2016-05-03',
					name: '王小虎',
					address: '上海市普陀区金沙江路 1516 弄'
				}]
			}
		},
		props: {
			msg: String
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
			onSubmit() {
				console.log('submit!');
			},
		}
	}
</script>
<style>
.el-card.page_wrapper > .el-card__body{
	padding: 12px;
}
.page_footer{
	padding: 6px;
	display: flex;
}
.page_footer .batch-handle{
  width: 400px;
  display: flex;
  
}
.page_footer .batch-handle .selected-title{
 margin: 0 4px;
 line-height: 28px;
 padding-right: 20px;
}
.page_footer .batch-handle .el-button{
  padding: 4px 10px;
}
.page_footer .pagination{
  flex:1;
  text-align: center;
}
</style>
