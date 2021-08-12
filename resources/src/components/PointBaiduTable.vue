<template>
	<el-card class="point_baidu_table page_wrapper">
		<div class="page_header">
			<el-form :inline="true" :model="query" class="monitot-form-inline">
				<div class="el-form-item el-form-item--small">
					<div class="el-input el-input-group el-input--small el-input-group--prepend">
							<input class="el-input__inner" type="text" autocomplete="off" placeholder="请输入名称" v-model="query.adcode" @change="getList" @keyup="getList">
							<div class="el-input-group__append">
								<button class="el-button el-button--default" type="button" @click="handlerReset">
									<i class="el-icon-refresh"></i>
								</button>
							</div>
						</div>
				</div>
			</el-form>
		</div>
		<div class="page_body">
			<el-table :data="list" style="width: 100%" @header-click="headerClick" @row-click="handleRowClick" :highlight-current-row="false">
				<el-table-column prop="name" label="名称"></el-table-column>
				<el-table-column label="操作" width="64">
					<template #default="scope">
						<button class="el-button el-button--default el-button--mini is-round" type="button" @click="handleDel(scope.row)">
							<i class="el-icon-delete"></i>
						</button>
					</template>
				</el-table-column>
			</el-table>
		</div>
		<div class="page_footer">
			<div class="batch-handle">
			</div>
			<div class="pagination">
				<el-pagination small layout="total, prev, pager, next" :page-size="page.size" :total="page.total" @next-click="handleNext" @prev-click="handlePrev" @current-change="handleChangePage" @size-change="handleChangeSize" :pager-count="2"></el-pagination>
			</div>
		</div>
	</el-card>
</template>
<script>
	import { watch } from 'vue'
	import dayjs from 'dayjs';
	import api from '../api';
	export default {
		name: 'ElPointBaiduTable',
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
					adcode:"",
				},
				form:{

				},
			}
		},
		props: {
			keyword: String
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
			const _this = this;
			console.log('come in')
			this.getPage(1)
			watch(
	        () => _this.keyword,
	        (toParams, previousParams) => {
	          _this.query.adcode = toParams
	          console.log('watch_keyword',_this.query)
	          _this.getList()
	        }
	        )
		},
		methods: {
			getPage(page){
				const _this = this;
				this.query.page = page
				console.log(this.query)
				api.getPointBaiduPage.send(this.query)
				.then(result => {
					console.log("getPage的结果",result)
					if (result.state==2000) {
						var exp = result.data; 
						if (!exp && typeof(exp)!="undefined" && exp!=0){
							_this.list = []
						}else{
							_this.list = result.data
						}
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
			handlerReset(){
				this.query.name = ''
				this.getList()
			},
			handleDel(item){
				const _this = this;
				console.log('item',item)
				api.delPointBaidu.send({'id':item.id})
				.then(result => {
					console.log("getPage的结果",result)
					if (result.state==2000) {
						_this.getPage(1)
					}
				})
				.catch(e => {
					console.log(e)
				})
			},
			handleRowClick(row,column,event){
				console.log('row',row)
				console.log('column',column)
				console.log('event',event)
				this.$emit('select',row)
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

.point_baidu_table .el-table__body tr.bg-blue:hover > td{
  background-color: #6fb6ff;
  color:#fff;
}

</style>
