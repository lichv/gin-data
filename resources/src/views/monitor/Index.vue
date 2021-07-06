<template>
	<el-card class="monitor_page page_wrapper">
		<div class="page_header">
			<el-form :inline="true" :model="query" class="monitot-form-inline">
				<el-form-item label="产品名称：" label-width="84px">
					<el-input v-model="query.seq_type" placeholder="产品名称" clearable></el-input>
				</el-form-item>
				<el-form-item label="状态：" label-width="84px">
					<el-select v-model="query.status" placeholder="所有状态" @change="search">
						<el-option label="所有状态" value="0"></el-option>
						<el-option v-for="item in status" :label="item.label" :value="item.value"></el-option>
					</el-select>
				</el-form-item>
				<el-form-item label="超出时长：" label-width="84px">
					<el-select v-model="query.timely" placeholder="区域" @change="search">
						<el-option label="14自然日" value="336"></el-option>
						<el-option label="15自然日" value="360"></el-option>
						<el-option label="20自然日" value="480"></el-option>
						<el-option label="25自然日" value="600"></el-option>
						<el-option label="28自然日" value="672"></el-option>
						<el-option label="1个月" value="720"></el-option>
						<el-option label="1000小时" value="1000"></el-option>
						<el-option label="2个月" value="1440"></el-option>
					</el-select>
				</el-form-item>
				<el-form-item label="癫痫1.0：" label-width="84px">
					<el-radio-group size="mini" v-model="query.is_antiepileptic" @change="search">
						<el-radio-button label="0">所有</el-radio-button>
						<el-radio-button label="1" >是</el-radio-button>
						<el-radio-button label="-1">否</el-radio-button>
					</el-radio-group>
				</el-form-item>
				<el-form-item label="是否科研：" label-width="84px">
					<el-radio-group size="mini" v-model="query.is_clinic_research" @change="search">
						<el-radio-button label="0">所有</el-radio-button>
						<el-radio-button label="1" >是</el-radio-button>
						<el-radio-button label="-1">否</el-radio-button>
					</el-radio-group>
				</el-form-item>
				<el-form-item label="重要标志位：" label-width="96px">
					<el-radio-group size="mini" v-model="query.is_import" @change="search">
						<el-radio-button label="0">所有</el-radio-button>
						<el-radio-button label="1" >是</el-radio-button>
						<el-radio-button label="-1">否</el-radio-button>
					</el-radio-group>
				</el-form-item>
				<el-form-item label="是否VIP：" label-width="96px">
					<el-radio-group size="mini" v-model="query.is_vip" @change="search">
						<el-radio-button label="0">所有</el-radio-button>
						<el-radio-button label="1" >是</el-radio-button>
						<el-radio-button label="-1">否</el-radio-button>
					</el-radio-group>
				</el-form-item>
				<el-form-item label="是否重分析：" label-width="96px">
					<el-radio-group size="mini" v-model="query.reanalysis" @change="search">
						<el-radio-button label="0">所有</el-radio-button>
						<el-radio-button label="1" >是</el-radio-button>
						<el-radio-button label="-1">否</el-radio-button>
					</el-radio-group>
				</el-form-item>
				<el-form-item>
					<el-button type="primary" size="mini" @click="search">查询</el-button>
					<el-button  size="mini" @click="reset">重置</el-button>
					<el-button size="mini" @click="handleShowProductDistribution">产品分布</el-button>
					<el-button size="mini" @click="handleShowStatusDistribution">状态分布</el-button>
					<el-button size="mini" @click="handleCopyAS">批量复制AS</el-button>
				</el-form-item>
			</el-form>
		</div>
		<div class="page_body">
			<el-table :data="list" style="width: 100%" @header-click="headerClick">
				<el-table-column fixed prop="seq_no" label="检测编号⧫" width="80">
					<template #default="scope">
						<p @click="copy(scope.row.seq_no)">{{scope.row.seq_no}}</p>
					</template>
				</el-table-column>
				<el-table-column prop="user" label="用户⧫" width="110">
					<template #default="scope">
						<p>用户：{{scope.row.name}}</p>
						<p>性别：{{scope.row.gender}}</p>
						<p @click="copy(scope.row.patient)">{{scope.row.patient}}</p>
						<p v-if="scope.row.is_proband">先证者</p><p v-else>亲属</p>
						<p @click="copy(scope.row.family_no)">{{scope.row.family_no}}</p>
					</template>
				</el-table-column>
				<el-table-column prop="seq_type" label="产品名称" >
					<template #default="scope">
						<p @click="copy(scope.row.seq_type)">{{scope.row.seq_type}}</p>
					</template>
				</el-table-column>
				<el-table-column prop="seq_company" label="送检公司" width="100">
					<template #default="scope">
						<p @click="copy(scope.row.seq_company)">{{scope.row.seq_company}}</p>
					</template>
				</el-table-column>
				<el-table-column prop="urgent_level" label="普通/加急/首次送检" width="140">
					<template #default="scope">
						<span v-if="scope.row.urgent_level==1">普通</span><span v-else-if="scope.row.urgent_level==2">加急</span><span v-else>首次送检</span>
					</template>
				</el-table-column>
				<el-table-column prop="status" label="样本状态⧫" width="120">
					<template #default="scope">
						<p >{{ $filters.findStatusByCode(scope.row.status,status) }}</p>
					</template>
				</el-table-column>
				
				<el-table-column prop="lab" label="实验室" width="160">
					<template #default="scope">
						<p @click="copy(scope.row.sample_lab_no)">样本编号：{{scope.row.sample_lab_no}}</p>
						<p @click="copy(scope.row.sample_type)">样本类型：{{scope.row.sample_type}}</p>
						<p @click="copy(scope.row.lab_seq_type)">检测类型：{{scope.row.lab_seq_type}}</p>
					</template>
				</el-table-column>				
				<el-table-column prop="people" label="负责人" width="120">
					<template #default="scope">
						<p @click="copy(scope.row.auditor_ad)">医学顾问：{{scope.row.auditor_ad}}</p>
						<p @click="copy(scope.row.sale)">销售员：{{scope.row.sale}}</p>
						<p @click="copy(scope.row.department)">处理部门：{{scope.row.department}}</p>
					</template>
				</el-table-column>	
				<el-table-column prop="timely" label="花费时长⧫" width="80"></el-table-column>
				<el-table-column prop="remark" label="备注" width="80">
					<template #default="scope">
						<el-popover placement="top-start" :width="200" trigger="hover" v-if="scope.row.trial_remark!=''" >
							<template #reference>
								<p>报告备注</p>
							</template>
							<p v-html="scope.row.trial_remark"></p>
						</el-popover>
						<el-popover placement="top-start" :width="200" trigger="hover"  v-if="scope.row.remark!=''" >
							<template #reference>
								<p>其他备注</p>
							</template>
							<p v-html="scope.row.remark"></p>
						</el-popover>
					</template>
				</el-table-column>
				<el-table-column label="标记" width="150">
					<template #default="scope">
						<p>是否癫痫1.0：{{scope.row.is_antiepileptic}}</p>
						<p>是否科研样本：{{scope.row.is_clinic_research}}</p>
						<p>是否重要标志位：{{scope.row.is_import}}</p>
						<p>是否VIP：{{scope.row.is_vip}}</p>
						<p>是否重分析：{{scope.row.reanalysis}}</p>
					</template>
				</el-table-column>
				<el-table-column prop="time" label="时间⧫" width="180">
					<template #default="scope">
						<p>采样：{{foramtDatetime(scope.row.collection_date)}}</p>
						<p>录单：{{foramtDatetime(scope.row.initial_dt)}}</p>
						<p>修改：{{foramtDatetime(scope.row.latest_time)}}</p>
					</template>
				</el-table-column>
			</el-table>
		</div>
		<div class="page_footer">
			<div class="batch-handle">
				<div class="selected-title">已选中 <span class="selected—count">0</span> 项</div>
				<div class="vux-flexbox selection-items-box vux-flex-row">
					<el-button size="mini" disabled>批量转移</el-button>
					<el-button size="mini" disabled>批量审核</el-button>
					<el-button size="mini" disabled>批量删除</el-button>
				</div>
			</div>
			<div class="pagination">
				<el-pagination small layout="total, sizes, prev, pager, next" :page-size="page.size" :total="page.total" @next-click="handleNext" @prev-click="handlePrev" @current-change="handleChangePage" @size-change="handleChangeSize"></el-pagination>
			</div>
		</div>
		<el-dialog title="产品分布" v-model="showProductDistribution" width="1600px">
			<el-monitor-product-distribution ref="product" :query="form"></el-monitor-product-distribution>
		</el-dialog>
		<el-dialog title="状态分布" v-model="showStatusDistribution" width="1080px">
			<el-monitor-status-distribution ref="status" :query="form"></el-monitor-status-distribution>
		</el-dialog>
	</el-card>
</template>
<script>
	import dayjs from 'dayjs';
	import api from '../../api';
	export default {
		name: 'Monitor',
		data() {
			return {
				showProductDistribution:false,
				showStatusDistribution:false,
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
					is_antiepileptic:0,
					is_clinic_research:0,
					is_import:0,
					is_vip:0,
					reanalysis:0,
				},
				form:{

				},
				status:[
				{"label":"样本开始检测","value":"2"},
				{"label":"检测完成","value":"3"},
				{"label":"开始生信分析","value":"4"},
				{"label":"生信分析完成","value":"5"},
				{"label":"生成临床报告","value":"6"},
				{"label":"临床报告发送完成","value":"7"},
				{"label":"Sanger验证","value":"8"},
				{"label":"验证完成","value":"9"},
				{"label":"生成验证报告","value":"10"},
				{"label":"核酸提取","value":"11"},
				{"label":"文库构建","value":"12"},
				{"label":"外送测序","value":"13"},
				{"label":"生成原始报告","value":"14"},
				{"label":"生成初步报告","value":"15"},
				{"label":"自建库重测","value":"20"},
				{"label":"外送补测","value":"21"},
				{"label":"外送重测","value":"22"},
				{"label":"完毕","value":"32"},
				{"label":"追加验证结果完成","value":"35"},
				{"label":"报告初审","value":"40"},
				{"label":"报告终审","value":"41"},
				{"label":"验证报告初审","value":"42"},
				{"label":"验证报告终审","value":"43"},
				{"label":"推广审核完成","value":"44"},
				{"label":"推广审核驳回","value":"45"},
				{"label":"销售审核完成","value":"46"},
				{"label":"销售审核驳回","value":"47"},
				{"label":"报告审核驳回","value":"48"},
				{"label":"验证报告审核驳回","value":"49"},
				{"label":"报告待打包","value":"50"},
				{"label":"上机测序","value":"51"},
				{"label":"自测序补测","value":"52"},
				{"label":"医学可出报告","value":"53"},
				{"label":"纸质报告邮寄","value":"54"},
				{"label":"数据已发送","value":"55"},
				{"label":"问题样本","value":"56"},
				{"label":"筛点完成","value":"57"},
				{"label":"异常周期样本","value":"504"},
				],
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
		mounted() {
			console.log('come in')
			this.getPage(1)
		},
		methods: {
			handleShowProductDistribution(){
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
				this.showProductDistribution = true;
				if (this.$refs && this.$refs.product) {
					this.$refs.product.getData()
				}
			},
			handleShowStatusDistribution(){
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
				this.showStatusDistribution = true;
				if (this.$refs && this.$refs.status) {
					this.$refs.status.getData()
				}
			},
			reset(){
				this.query={
					page:1,
					size:20,
					is_antiepileptic:false,
					is_clinic_research:false,
					is_import:false,
					is_vip:false,
					reanalysis:false,
				}
				this.getPage(1)
			},
			search() {
				this.getPage(1)
			},
			getPage(page){
				const _this = this;
				this.query.page = page
				console.log(this.query)
				api.getMonitorPages.send(this.query)
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
				if (column.property=='seq_no' || column.property=='status' || column.property=='timely' || column.property=='user' || column.property=='time') {
					if (column.property=='seq_no' || column.property=='status' || column.property=='timely') {
						this.query.order = column.property
					}else if (column.property == 'user') {
						this.query.order = 'patient'
					}else if (column.property == 'time') {
						this.query.order = 'initial_dt'
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
			handleCopyAS(){
				const _this = this;
				api.getSeqNos.send(this.query)
				.then(result => {
					console.log(result)
					if (result.state==2000) {
						var list = result.data
						var text = "'"+list.join("','")+"'"
						_this.copy(text)
					}
				})
				.catch(e => {
					console.log(e)
				})
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

.page_wrapper{
	height: 100%;
	min-height: 100px;
}
.el-card.page_wrapper > .el-card__body{
	padding: 12px;
	height: 100%;
	overflow-y: auto;
}
.page_header{
	border-bottom: 1px solid #d6d6d6;
}
@media(min-width: 669px){
	.page_body{
		max-height: calc(100% - 256px);
		overflow-y: auto;
	}
}
@media(min-width: 911px){
	.page_body{
		max-height: calc(100% - 178px);
		overflow-y: auto;
	}
}
@media(min-width: 1226px){
	.page_body{
		max-height: calc(100% - 128px);
		overflow-y: auto;
	}
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
	width: 420px;
	min-width: 420px;
}
.monitot-form-inline .el-radio-button__orig-radio:checked + .el-radio-button__inner{
	color: rgb(255, 255, 255);
    background-color: rgb(182 217 253);
    border-color: rgb(182 217 253);
}
.monitor_page .el-radio-button--mini .el-radio-button__inner{
	padding: 4px 6px;
}
</style>
