import"./index.af140293.js";import{r as e,o as l,c as a,w as t,g as s}from"./vendor.3de699ef.js";const i={name:"HelloWorld",data:()=>({showProductDistribution:!1,showStatusDistribution:!1,list:[],page:{page:1,size:50,total:1,last:1},query:{page:1,size:50,order:"seq_no",sort:1,is_antiepileptic:0,is_clinic_research:0,is_import:0,is_vip:0,reanalysis:0},form:{},status:[{label:"样本开始检测",value:"2"},{label:"检测完成",value:"3"},{label:"开始生信分析",value:"4"},{label:"生信分析完成",value:"5"},{label:"生成临床报告",value:"6"},{label:"临床报告发送完成",value:"7"},{label:"Sanger验证",value:"8"},{label:"验证完成",value:"9"},{label:"生成验证报告",value:"10"},{label:"核酸提取",value:"11"},{label:"文库构建",value:"12"},{label:"外送测序",value:"13"},{label:"生成原始报告",value:"14"},{label:"生成初步报告",value:"15"},{label:"自建库重测",value:"20"},{label:"外送补测",value:"21"},{label:"外送重测",value:"22"},{label:"完毕",value:"32"},{label:"追加验证结果完成",value:"35"},{label:"报告初审",value:"40"},{label:"报告终审",value:"41"},{label:"验证报告初审",value:"42"},{label:"验证报告终审",value:"43"},{label:"推广审核完成",value:"44"},{label:"推广审核驳回",value:"45"},{label:"销售审核完成",value:"46"},{label:"销售审核驳回",value:"47"},{label:"报告审核驳回",value:"48"},{label:"验证报告审核驳回",value:"49"},{label:"报告待打包",value:"50"},{label:"上机测序",value:"51"},{label:"自测序补测",value:"52"},{label:"医学可出报告",value:"53"},{label:"纸质报告邮寄",value:"54"},{label:"数据已发送",value:"55"},{label:"问题样本",value:"56"},{label:"筛点完成",value:"57"},{label:"异常周期样本",value:"504"}]}),props:{msg:String},watch:{$router:{handler(e,l){console.log("val",e),console.log("oldval",l)}},isCollapse(e,l){this.showCollapse=e}},mounted(){},methods:{handleShowProductDistribution(){this.form=this.query,"page"in this.form&&delete this.form.page,"size"in this.form&&delete this.form.size,"order"in this.form&&delete this.form.order,"sort"in this.form&&delete this.form.sort,this.showProductDistribution=!0,this.$refs&&this.$refs.product&&this.$refs.product.getData()},handleShowStatusDistribution(){this.form=this.query,"page"in this.form&&delete this.form.page,"size"in this.form&&delete this.form.size,"order"in this.form&&delete this.form.order,"sort"in this.form&&delete this.form.sort,this.showStatusDistribution=!0,this.$refs&&this.$refs.status&&this.$refs.status.getData()},copy(e){var l=document.createElement("input");l.value=e,document.body.appendChild(l),l.select(),document.execCommand("Copy"),document.body.removeChild(l),this.$message({showClose:!0,message:"复制成功",type:"success"})}}};i.render=function(i,o,r,u,n,b){const h=e("el-monitor-medicine-statistics"),d=e("el-card");return l(),a(d,{class:"medicine_page page_wrapper"},{default:t((()=>[s(h,{ref:"status",query:n.form},null,8,["query"])])),_:1})};export default i;