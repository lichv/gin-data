import{r as e,o as l,c as a,w as s,g as o,h as i}from"./vendor.2906b451.js";const d={name:"Layout",data:()=>({today:new Date,count:0,drawer:!1,direction:"ltr",aside_size:"mini",isCollapse:!1}),props:{msg:String},computed:{asidesize:function(){return"mini"==this.aside_size?(this.isCollapse=!0,"64px"):(this.isCollapse=!1,"240px")}},watch:{$router:{handler(e,l){console.log("val",e),console.log("oldval",l)}},isCollapse(e,l){this.showCollapse=e}},methods:{showAside(){"mini"==this.aside_size?this.aside_size="normal":this.aside_size="mini",this.isCollapse=!this.isCollapse},handleClickBtn(){this.drawer=!this.drawer,console.log("this.drawer",this.drawer)},handleClose(e){console.log("done",e),this.drawer=!1},handleSelect(e){console.log("handleSelect",e),this.drawer=!1}}},t={class:"layout-header-left"},n={class:"layout-header-handle"},r=o("i",{class:"el-icon-menu"},null,-1),c=o("div",{class:"height-left-logo"},[o("img",{class:"header-logo",src:"https://chv.oss-cn-shanghai.aliyuncs.com/aegicare/logo_aegicare.png"})],-1),h={class:"asider-menu"},u={key:0,class:"el-icon-s-unfold"},m={key:1,class:"el-icon-s-fold"},p=o("div",{class:"layout-header-middle"},null,-1),f={class:"layout-header-right"},g=o("i",{class:"el-icon-s-home"},null,-1),w=i("医院管理"),C=i("医院管理");d.render=function(i,d,_,y,v,b){const z=e("el-asider-menu"),k=e("el-drawer"),S=e("router-link"),V=e("el-header"),x=e("el-calendar"),A=e("el-card"),B=e("el-main"),U=e("el-container");return l(),a(U,{class:"page-container"},{default:s((()=>[o(V,{class:"layout-header",height:"56px"},{default:s((()=>[o("div",t,[o("div",n,[o("div",{class:"header-left-btn",onClick:d[1]||(d[1]=(...e)=>b.handleClickBtn&&b.handleClickBtn(...e))},[r]),c]),o(k,{class:"header-left-menu",title:"",modelValue:v.drawer,"onUpdate:modelValue":d[3]||(d[3]=e=>v.drawer=e),direction:v.direction,"before-close":b.handleClose,"destroy-on-close":"",size:b.asidesize,modal:!0,"show-close":!0},{default:s((()=>[o("div",h,[o(z,{"is-collapse":v.isCollapse,onSelect:b.handleSelect},null,8,["is-collapse","onSelect"]),o("div",{class:"aside-bottom-btn",onClick:d[2]||(d[2]=(...e)=>b.showAside&&b.showAside(...e))},[v.isCollapse?(l(),a("span",u)):(l(),a("span",m))])])])),_:1},8,["modelValue","direction","before-close","size"])]),p,o("div",f,[o(S,{to:"/"},{default:s((()=>[g])),_:1})])])),_:1}),o(B,{class:"layout-main"},{default:s((()=>[o(A,{class:"dashboard-card calendar-card"},{default:s((()=>[o(x,{modelValue:v.today,"onUpdate:modelValue":d[4]||(d[4]=e=>v.today=e)},null,8,["modelValue"])])),_:1}),o(A,{class:"dashboard-card"},{default:s((()=>[o(S,{to:"/hospital"},{default:s((()=>[w])),_:1})])),_:1}),o(A,{class:"dashboard-card"},{default:s((()=>[o(S,{to:"/hospital"},{default:s((()=>[C])),_:1})])),_:1})])),_:1})])),_:1})};export default d;
