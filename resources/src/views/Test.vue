<template>
	<div>
		<el-area-select></el-area-select>
	</div>
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
.distribution-chart{
	float: left;

}
.distribution-chart +.distribution-chart{
	margin-left: 8px;
}
</style>
