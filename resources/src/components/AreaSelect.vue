<template>
	<el-card class="area-select-card">
		<el-input placeholder="输入关键字进行过滤" v-model="filterText"></el-input>
		<div class="area-select-content">
			<el-tree ref="tree" :data="regions" node-key="code" :expand-on-click-node="false" :render-content="renderContent" :filter-node-method="filterNode" @node-click="handleNodeClick"> </el-tree>
		</div>
	</el-card>
</template>
<script>
	import pca from '../lib/pca-code.json'
	let id = 1000;

	export default {
		name: 'ElAreaSelect',
		data(){
			const data = [{
				id: 1,
				label: '一级 1',
				children: [{
					id: 4,
					label: '二级 1-1',
					children: [{
						id: 9,
						label: '三级 1-1-1'
					}, {
						id: 10,
						label: '三级 1-1-2'
					}]
				}]
			}, {
				id: 2,
				label: '一级 2',
				children: [{
					id: 5,
					label: '二级 2-1'
				}, {
					id: 6,
					label: '二级 2-2'
				}]
			}, {
				id: 3,
				label: '一级 3',
				children: [{
					id: 7,
					label: '二级 3-1'
				}, {
					id: 8,
					label: '二级 3-2'
				}]
			}];
			return {
				regions:[],
				filterText: '',
				formInline: {
					user: '',
					region: ''
				},
				form: {
					name: '',
					region: '',
					date1: '',
					date2: '',
					delivery: false,
					type: [],
					resource: '',
					desc: ''
				},
				data: JSON.parse(JSON.stringify(data)),
			}
		},
		props: {

		},
		watch: {
			filterText(val) {
				this.$refs.tree.filter(val);
			}
		},
		mounted() {
			this.regions = pca;
		},
		methods:{
			handleNodeClick(data,node,el){
				console.log('select',{"code":data.code,"name":data.name})
				this.$emit('select',{"code":data.code,"name":data.name})
			},
			onSubmit() {
				console.log('submit!');
			},
			renderContent(h, { node, data, store }) {
				return h("span", {
					class: "custom-tree-node"
				}, h("span", {
					class: "tree-item-text"
				}, data.name));
			},
			filterNode(value, data) {
				if (!value) return true;
				return data.name.indexOf(value) !== -1;
			}
		},
		computed:{

		},
		created(){

		}
	};
</script>
<style >
.area-select-card{
	height: 100%;
}
.el-card.area-select-card > .el-card__body{
	height: 100%;
	overflow-y: scroll;
	padding: 12px;
}
</style>