import{r as e,o as s,c as l,w as a,g as i}from"./vendor.2906b451.js";const o={name:"Layout",data:()=>({count:0,drawer:!1,direction:"ltr",aside_size:"mini",isCollapse:!1}),props:{msg:String},computed:{asidesize:function(){return"mini"==this.aside_size?(this.isCollapse=!0,"64px"):(this.isCollapse=!1,"240px")}},watch:{$router:{handler(e,s){console.log("val",e),console.log("oldval",s)}},isCollapse(e,s){this.showCollapse=e}},methods:{showAside(){"mini"==this.aside_size?this.aside_size="normal":this.aside_size="mini",this.isCollapse=!this.isCollapse},handleClickBtn(){this.drawer=!this.drawer,console.log("this.drawer",this.drawer)},handleClose(e){console.log("done",e),this.drawer=!1},handleSelect(e){console.log("handleSelect",e),this.drawer=!1}}},t={class:"layout-header-left"},d={class:"layout-header-handle"},n=i("i",{class:"el-icon-menu"},null,-1),r=i("div",{class:"height-left-logo"},[i("img",{class:"header-logo",src:"https://chv.oss-cn-shanghai.aliyuncs.com/aegicare/logo_aegicare.png"})],-1),c={class:"asider-menu"},h={key:0,class:"el-icon-s-unfold"},u={key:1,class:"el-icon-s-fold"},m=i("div",{class:"layout-header-middle"},null,-1),p={class:"layout-header-right"},g=i("i",{class:"el-icon-s-home"},null,-1);o.render=function(o,w,f,C,v,y){const _=e("el-asider-menu"),z=e("el-drawer"),k=e("router-link"),b=e("el-header"),S=e("router-view"),x=e("el-main"),A=e("el-container");return s(),l(A,{class:"page-container"},{default:a((()=>[i(b,{class:"layout-header",height:"56px"},{default:a((()=>[i("div",t,[i("div",d,[i("div",{class:"header-left-btn",onClick:w[1]||(w[1]=(...e)=>y.handleClickBtn&&y.handleClickBtn(...e))},[n]),r]),i(z,{class:"header-left-menu",title:"",modelValue:v.drawer,"onUpdate:modelValue":w[3]||(w[3]=e=>v.drawer=e),direction:v.direction,"before-close":y.handleClose,"destroy-on-close":"",size:y.asidesize,modal:!0,"show-close":!0},{default:a((()=>[i("div",c,[i(_,{"is-collapse":v.isCollapse,onSelect:y.handleSelect},null,8,["is-collapse","onSelect"]),i("div",{class:"aside-bottom-btn",onClick:w[2]||(w[2]=(...e)=>y.showAside&&y.showAside(...e))},[v.isCollapse?(s(),l("span",h)):(s(),l("span",u))])])])),_:1},8,["modelValue","direction","before-close","size"])]),m,i("div",p,[i(k,{to:"/"},{default:a((()=>[g])),_:1})])])),_:1}),i(x,{class:"layout-main"},{default:a((()=>[i(S)])),_:1})])),_:1})};export default o;