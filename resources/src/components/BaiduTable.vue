<template>
  <el-card class="input_table page_wrapper">
    <div class="page_header">
      <el-form :inline="true" :model="query" class="monitot-form-inline">
        <div class="el-form-item el-form-item--small">
          <div class="el-input el-input-group el-input--small el-input-group--prepend">
            <input class="el-input__inner" type="text" autocomplete="off" placeholder="请输入名称" v-model="query.name" @change="getList" @keyup="getList">
            <div class="el-input-group__append">
              <button class="el-button el-button--default" type="button" @click="getList">
                <i class="el-icon-refresh"></i>
              </button>
            </div>
          </div>
        </div>
        <div class="el-form-item el-form-item--small">
          <div class="el-input el-input-group el-input--small el-input-group--prepend">
            <input class="el-input__inner" type="text" autocomplete="off" placeholder="请输入名称" v-model="name" @change="getList" @keyup="getList">
          </div>
        </div>
      </el-form>
    </div>
    <div class="page_body">
      <el-table :data="list" style="width: 100%" @row-click="handleRowClick" :highlight-current-row="false">
        <el-table-column prop="name" label="名称"></el-table-column>
        <el-table-column prop="address" label="名称"></el-table-column>
        <el-table-column label="操作" width="120">
          <template #default="scope">
            <button class="el-button el-button--default el-button--mini is-round" type="button" @click="saveTest(scope.row)"><i class="el-icon-edit"></i></button>
            
          </template>
        </el-table-column>
      </el-table>
    </div>
    <div class="page_footer">
      <div class="batch-handle">
        总记录数 {{count}}
      </div>
    </div>
  </el-card>
</template>
<script>
  import { watch } from 'vue'
  import api from '../api';
  export default {
    name: 'ElBaiduTable',
    data() {
      return {
        name:'',
        list:[],
        query:{
          name:"",
        },
        form:{

        },
      }
    },
    props: {
      keyword: String
    },
    computed:{
      count(){
        return this.list.length
      }
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
      this.getList()
      watch(
        () => _this.keyword,
        (toParams, previousParams) => {

          _this.query.name = toParams
          _this.name = toParams
          console.log('toParams',toParams)
          console.log('previousParams',previousParams)
          console.log('watch_keyword',_this.query)
          _this.getList()
        }
        )
    },
    methods: {
      getList() {
        const _this = this;
        api.getBaiduPages.send({'name':this.query.name})
        .then(result => {
          if (result.status==0) {
            var exp = result.results; 
            if (!exp && typeof(exp)!="undefined" && exp!=0){
              _this.list = []
            }else{
              _this.list = result.results
            }
          }
        })
        .catch(e => {
          console.log(e)
        })
      },
      saveTest(item){
        const _this = this;
        console.log('item',item)
        var postdata = {'name':this.name,'alias':item['name'],'province':item['province'],'city':item['city'],'area':item['area'],'lat':item['location']['lat'],'lng':item['location']['lng']}
        if (postdata['name'] == postdata['alias']){
          postdata['alias'] = ''
        }
        api.saveTest.send(postdata)
        .then(result => {
          console.log("SaveTest的结果",result)
          if (result.state==2000) {
            _this.$emit('save',{'name':this.query.name,'alias':item['name'],'province':item['province'],'city':item['city'],'area':item['area'],'lat':item['location']['lat'],'lng':item['location']['lng']})
          }
        })
        .catch(e => {
          console.log(e)
        })
      },
      newTest(item){
        const _this = this;
        console.log('item',item)
        var postdata = {'name':item['name'],'alias':'','province':item['province'],'city':item['city'],'area':item['area'],'lat':item['location']['lat'],'lng':item['location']['lng']}
        if (postdata['name'] == postdata['alias']){
          postdata['alias'] = ''
        }
        api.saveTest.send(postdata)
        .then(result => {
          console.log("SaveTest的结果",result)
          if (result.state==2000) {
            _this.$emit('save',{'name':item['name'],'alias':'','province':item['province'],'city':item['city'],'area':item['area'],'lat':item['location']['lat'],'lng':item['location']['lng']})
          }
        })
        .catch(e => {
          console.log(e)
        })
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

.input_table .el-table__body tr.bg-blue:hover > td{
  background-color: #6fb6ff;
  color:#fff;
}

</style>
