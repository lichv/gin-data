<template>
  <div class="hospital-page">
    <el-input-table ref="input" class="input-table" @select="handleSelect"></el-input-table>
    <el-baidu-table ref="baidu" class="baidu-table" @save="handleSave" :keyword="keywords"></el-baidu-table>
    <el-google-table ref="baidu" class="baidu-table" @save="handleSave" :keyword="keywords"></el-google-table>
  </div>
</template>

<script>
  import { watch } from 'vue'
  export default {
    name: 'App',
    data(){
      return{
        keywords:'',
        selected :{},
      };
    },
    mounted() {
      const _this = this;
      watch(
        () => _this.selected,
        (toParams, previousParams) => {
          _this.query.name = toParams
          console.log('watch_keyword',_this.query)
          console.log('toParams',toParams)
          console.log('previousParams',previousParams)          
        }
        )
    },
    methods: {
      handleSelect(item){
        console.log('handleSelect',item)
        this.keywords = item.name
      },
      handleSave(item){
        console.log('handleSave',item)
        this.$refs.input.getList()
      }
    }
  }
</script>
<style>
.hospital-page{
  height: 100%;
  background-color: #f0f9eb;
  display: flex;
}
.input-table{
  width: 360px;
}
.baidu-table{
  margin-left: 6px;
  width: 640px;
}
</style>
