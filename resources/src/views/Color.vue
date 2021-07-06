<template>
	<div class="block">
		<span class="demonstration">总样本量</span>
		<el-color-picker v-model="color01"></el-color-picker>
	</div>
	<div class="block">
		<span class="demonstration">正常样本量</span>
		<el-color-picker v-model="color02"></el-color-picker>
	</div>
	<div class="block">
		<span class="demonstration">正常达标样本量</span>
		<el-color-picker v-model="color03"></el-color-picker>
	</div>
	<div class="block">
		<span class="demonstration">正常达标率</span>
		<el-color-picker v-model="color04"></el-color-picker>
	</div>

	<div class="block">
		<span class="demonstration">异常样本量</span>
		<el-color-picker v-model="color05"></el-color-picker>
	</div>
	<div class="block">
		<span class="demonstration">异常达标样本量</span>
		<el-color-picker v-model="color06"></el-color-picker>
	</div>
	<div class="block">
		<span class="demonstration">异常样本比例</span>
		<el-color-picker v-model="color07"></el-color-picker>
	</div>
	<div class="block">
		<span class="demonstration">异常达标率</span>
		<el-color-picker v-model="color08"></el-color-picker>
	</div>
	<div class="block">
		<span class="demonstration">达标数</span>
		<el-color-picker v-model="color09"></el-color-picker>
	</div>
	<div class="block">
		<span class="demonstration">超期数</span>
		<el-color-picker v-model="color10"></el-color-picker>
	</div>
	<div class="block">
		<span class="demonstration">达标比例</span>
		<el-color-picker v-model="color11"></el-color-picker>
	</div>

	<el-input v-model=all></el-input>
</template>

<script>
	export default {
		name:"Test",
		data() {
			return {
				color01: '#409EFF',
				color02: '#409EFF',
				color03: '#409EFF',
				color04: '#409EFF',
				color05: '#409EFF',
				color06: '#409EFF',
				color07: '#409EFF',
				color08: '#409EFF',
				color09: '#409EFF',
				color10: '#409EFF',
				color11: '#409EFF',
			}
		},
		computed:{
			all:function(){
				return this.color01 + ',' +this.color02 + ',' +this.color03 + ',' +this.color04 + ',' +this.color05 + ',' +this.color06 + ',' +this.color07 + ',' +this.color08 + ',' +this.color09 + ',' +this.color10 + ',' +this.color11 
			},
		},
	};
</script>

<style>

</style>
