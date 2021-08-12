<template>
  <el-card class="input_table page_wrapper">
    <div class="page_body">
      <el-table :data="list" style="width: 100%" @row-click="handleRowClick" :highlight-current-row="false">
        <el-table-column prop="name" label="名称"></el-table-column>
        <el-table-column label="操作" width="64">
          <template #default="scope">
            <button class="el-button el-button--default el-button--mini is-round" type="button" @click="copy(scope.row.name)">
              <i class="el-icon-document-copy"></i>
            </button>
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
  import dayjs from 'dayjs'
  import api from '../api';
  export default {
    name: 'ElInputTable',
    data() {
      return {
        list:[],
        query:{
          order:'id',
          sort :1,
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
    mounted() {
      const _this = this;
      console.log('come in')
      this.getList()
      watch(
          () => _this.keyword,
          (toParams, previousParams) => {
            _this.query.name = toParams
            console.log('watch_keyword',_this.query)
            console.log('toParams',toParams)
            console.log('previousParams',previousParams)
            _this.getList()
          }
          )
    },
    methods: {
      getList() {
        const _this = this;
        api.getInputList.send({})
        .then(result => {
          if (result.state==2000) {
            var exp = result.data; 
            if (!exp && typeof(exp)!="undefined" && exp!=0){
              _this.list = []
            }else{
              _this.list = result.data
            }
          }
        })
        .catch(e => {
          console.log(e)
        })
      },
      handleDel(item){
        const _this = this;
        console.log('item',item)
        api.delTest.send({'id':item.id})
        .then(result => {
          console.log("delInput的结果",result)
          if (result.state==2000) {
            _this.getList()
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

.input_table .el-table__body tr.bg-blue:hover > td{
  background-color: #6fb6ff;
  color:#fff;
}

</style>
