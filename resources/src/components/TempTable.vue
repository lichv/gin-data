<template>
	<el-card class="hospital_table page_wrapper">
		<div class="page_header">
			<el-form :inline="true" :model="query" class="monitot-form-inline">
				<div class="el-form-item el-form-item--small">
					<div class="el-form-item__content">
						<div class="el-input el-input-group el-input--small el-input-group--prepend">
							<input class="el-input__inner" type="text" autocomplete="off" placeholder="请输入名称" v-model="query.name" @change="getList" @keyup="getList">
							<div class="el-input-group__append">
								<button class="el-button el-button--default" type="button" @click="handlerReset">
									<i class="el-icon-refresh"></i>
								</button>
							</div>
						</div>
					</div>
				</div>
			</el-form>
		</div>
		<div class="page_body">
			<el-table :data="list" style="width: 100%" @header-click="headerClick" @row-click="handleRowClick" :row-class-name="tableRowClassName">
				<el-table-column prop="name" label="名称"></el-table-column>
				<el-table-column label="操作" width="120">
					<template #default="scope">
						<button class="el-button el-button--default el-button--mini is-round" type="button" @click="handleDel(scope.row)">
							<i class="el-icon-delete"></i>
						</button>
						<button class="el-button el-button--default el-button--mini is-round" type="button" @click="handleMove(scope.row)">
							<i class="el-icon-right"></i>
						</button>
					</template>
				</el-table-column>
			</el-table>
		</div>
		<div class="page_footer">
			<div class="batch-handle">
			</div>
			<div class="pagination">
				<el-pagination small layout="total, prev, pager, next" :page-size="page.size" :total="page.total" @next-click="handleNext" @prev-click="handlePrev" @current-change="handleChangePage" @size-change="handleChangeSize" :pager-count="3"></el-pagination>
			</div>
		</div>
	</el-card>
</template>
<script>
	import dayjs from 'dayjs';
	import api from '../api';
	export default {
		name: 'ElTempTable',
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
					order:'id',
					sort :1,
					name:"",
				},
				form:{

				},
				selected:{},
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

		},
		mounted() {
			console.log('come in')
			this.getPage(1)
		},
		methods: {
			getPage(page){
				const _this = this;
				this.query.page = page
				console.log(this.query)
				api.getTempPage.send(this.query)
				.then(result => {
					console.log("getPage的结果",result)
					if (result.state==2000) {
						_this.list = result.data
						_this.page = result.page
						_this.query.page = _this.page.page + 1

					}
				})
				.catch(e => {
					console.log(e)
				})
			},
			reset(){
				this.query={
					page:1,
					size:20,
				}
				this.getPage(1)
			},
			getList() {
				this.getPage(1)
			},			
			handleNext(){
				var page = 1;
				if (this.page.page < this.page.last) {
					page = this.page.page + 1;
				}
				this.getPage(page)
			},
			handlePrev(){
				var page = 1;
				if (this.page.page > 2 ) {
					page = this.page.page - 1;
				}
				this.getPage(page)
			},
			handleChangePage(page){
				this.getPage(page)
			},
			handleChangeSize(size){
				this.query.size = size
				this.getPage(1)
			},
			headerClick(column, event){
				console.log('column',column)
				if (column.property=='id' || column.property=='status' || column.property=='timely' || column.property=='user' || column.property=='time') {
					if (column.property == 'id') {
						this.query.order = 'id'
					}else if (column.property == 'name') {
						this.query.order = 'name'
					}

					if (!this.query.sort) {
						this.query.sort = 0
					}

					if (this.query.sort == 2) {
						this.query.sort = 0
					}else{
						this.query.sort = this.query.sort + 1
					}

					this.getPage(1)
				}
			},
			handleSelect(item){

			},
			handleDel(item){
				const _this = this;
				console.log('item',item)
				api.delTemp.send({'id':item.id})
				.then(result => {
					console.log("getPage的结果",result)
					if (result.state==2000) {
						_this.getPage(_this.query.page)
					}
				})
				.catch(e => {
					console.log(e)
				})
			},
			handleMove(item){
				const _this = this;
				console.log('item',item)
				api.moveTemp2Hospital.send({'id':item.id})
				.then(result => {
					console.log("handleMove的结果",result)
					if (result.state==2000) {
						_this.getPage(_this.query.page)
						_this.$emit('move',item)
					}
				})
				.catch(e => {
					console.log(e)
				})
			},
			handleRowClick(row,column,event){
				this.selected = row;
				this.$emit('select',row)
			},
			tableRowClassName({row, rowIndex}) {
				if (row.id === this.selected.id) {
					return 'bg-blue';
				} else if (rowIndex === 3) {
					return '';
				}
				return '';
			},
			handlerReset(){
				this.query.name = ''
				this.getList()
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

			foramtDatetime(val){
				if (val) {
					return dayjs(val).format('YYYY-MM-DD HH:mm:ss')
				}
				return ''
			}
		}
	}
</script>
<style>

.page_body{
	max-height: calc(100% - 80px);
	overflow-y: auto;
}
.page_footer{
	padding: 6px;
	border-top: 1px solid #d6d6d6;
}
.page_footer .batch-handle{
	width: 336px;
	min-width: 336px;
	float: left;
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
	float: right;
	text-align: center;
	width: 256px;
	min-width: 256px;
}
.monitot-form-inline .el-radio-button__orig-radio:checked + .el-radio-button__inner{
	color: rgb(255, 255, 255);
	background-color: rgb(182 217 253);
	border-color: rgb(182 217 253);
}
.hospital_table .el-radio-button--mini .el-radio-button__inner{
	padding: 4px 6px;
}
.el-form-item__label{
	font-size: 10px;
}
.el-table tr.bg-blue{
	background-color: #6fb6ff;
	color:#fff;
}
.el-table tr.bg-blue:hover{
	background-color: #6fb6ff;
	color:#fff;
}
.el-table__body tr.hover-row.bg-blue>td{
	background-color: #6fb6ff;
	color:#fff;
}
</style>
