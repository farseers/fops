import{_ as H}from"./index-74aeefad.js";import{f as Q}from"./index-2149a754.js";import{k as L,aN as j,e as M,$ as O,f as q,i as F,ah as n,ar as G,o as d,c as I,W as a,Q as l,a as m,U as p,R as J,P as f,V as r,T as v,u as K,bb as X,b9 as Y}from"./@vue-d9027515.js";import{_ as Z}from"./_plugin-vue_export-helper-c27b6911.js";import"./pinia-be7e992d.js";import"./vue-demi-71ba0ef2.js";import"./vue-router-a2fac16f.js";import"./vue-i18n-e0b9a81d.js";import"./@intlify-9e8a497c.js";import"./source-map-7d7e1c08.js";import"./vue-b46b26c7.js";import"./js-cookie-cf83ad76.js";import"./@element-plus-96a31696.js";import"./nprogress-8d2808ea.js";import"./axios-707ed124.js";import"./qs-bbfcbdf3.js";import"./side-channel-eafc5c70.js";import"./get-intrinsic-b9397c9a.js";import"./has-symbols-e8f3ca0e.js";import"./function-bind-22e7ee79.js";import"./has-26d28e02.js";import"./call-bind-9ceb8f5b.js";import"./object-inspect-7b08a46e.js";import"./element-plus-7da03efa.js";import"./lodash-es-36eb724a.js";import"./@vueuse-2d9216d7.js";import"./@ctrl-f8748455.js";import"./dayjs-6aac7dfa.js";import"./async-validator-dee29e8b.js";import"./memoize-one-297ddbcb.js";import"./escape-html-1d60d822.js";import"./normalize-wheel-es-ed76fb12.js";import"./@floating-ui-463e90e0.js";import"./mitt-f7ef348c.js";import"./vue-grid-layout-0ee01be3.js";import"./requestFSGet-61715c66.js";const b=w=>(X("data-v-3b45e0a2"),w=w(),Y(),w),ee={class:"system-user-container layout-padding"},te={class:"system-user-search mb15"},ae=b(()=>m("label",null,"TraceId",-1)),oe=b(()=>m("label",{class:"ml10"},"应用名称",-1)),le=b(()=>m("label",{class:"ml10"},"执行端IP",-1)),ie=b(()=>m("label",{class:"ml10"},"日志内容",-1)),ne=b(()=>m("label",{class:"ml10"},"日志类型",-1)),pe=["onClick"],re=b(()=>m("br",null,null,-1)),se={key:1,type:"info",size:"mini"},de=L({name:"fopsLogList"}),me=L({...de,setup(w,{expose:x}){const z=Q(),C=j(()=>H(()=>import("./detailDialog-f5365baf.js"),["assets/js/detailDialog-f5365baf.js","assets/js/index-2149a754.js","assets/js/index-74aeefad.js","assets/js/@vue-d9027515.js","assets/js/pinia-be7e992d.js","assets/js/vue-demi-71ba0ef2.js","assets/js/vue-router-a2fac16f.js","assets/js/vue-i18n-e0b9a81d.js","assets/js/@intlify-9e8a497c.js","assets/js/source-map-7d7e1c08.js","assets/js/vue-b46b26c7.js","assets/js/js-cookie-cf83ad76.js","assets/js/@element-plus-96a31696.js","assets/js/nprogress-8d2808ea.js","assets/css/nprogress-8b89e2e0.css","assets/js/axios-707ed124.js","assets/js/qs-bbfcbdf3.js","assets/js/side-channel-eafc5c70.js","assets/js/get-intrinsic-b9397c9a.js","assets/js/has-symbols-e8f3ca0e.js","assets/js/function-bind-22e7ee79.js","assets/js/has-26d28e02.js","assets/js/call-bind-9ceb8f5b.js","assets/js/object-inspect-7b08a46e.js","assets/js/element-plus-7da03efa.js","assets/js/lodash-es-36eb724a.js","assets/js/@vueuse-2d9216d7.js","assets/js/@ctrl-f8748455.js","assets/js/dayjs-6aac7dfa.js","assets/js/async-validator-dee29e8b.js","assets/js/memoize-one-297ddbcb.js","assets/js/escape-html-1d60d822.js","assets/js/normalize-wheel-es-ed76fb12.js","assets/js/@floating-ui-463e90e0.js","assets/js/mitt-f7ef348c.js","assets/js/vue-grid-layout-0ee01be3.js","assets/css/index-c2520709.css","assets/js/requestFSGet-61715c66.js","assets/css/detailDialog-2d838ab0.css"])),D=M(),t=O({traceId:"",appName:"",appIp:"",logContent:"",logLevel:-1,tableData:{data:[],total:0,loading:!1,param:{pageNum:1,pageSize:10}},dialog:{isShowDialog:!1,type:"",title:"",submitTxt:""}});q(()=>t.logLevel,(i,o)=>{console.log(`count 从 ${o} 变为 ${i}`),u()});const u=()=>{t.tableData.loading=!0;var i={appName:t.appName,appIp:t.appIp,traceId:t.traceId,logContent:t.logContent,logLevel:t.logLevel.toString(),pageSize:t.tableData.param.pageSize.toString(),pageIndex:t.tableData.param.pageNum.toString()};const o=new URLSearchParams(i).toString();z.logList(o).then(function(s){s.Status?(t.tableData.data=s.Data.List,t.tableData.total=s.Data.RecordCount,t.tableData.loading=!1):(t.tableData.data=[],t.tableData.loading=!1)})},h=i=>{D.value.openDialog(i)},V=i=>{t.dialog.isShowDialog=!0,t.traceId=i.TraceIdN,u()},S=i=>{t.tableData.param.pageSize=i,u()},N=i=>{t.tableData.param.pageNum=i,u()},k=()=>{u()};return F(()=>{u()}),x({openDialog:V}),(i,o)=>{const s=n("el-input"),g=n("el-option"),T=n("el-select"),U=n("ele-Search"),A=n("el-icon"),y=n("el-button"),_=n("el-table-column"),c=n("el-tag"),P=n("el-table"),R=n("el-pagination"),$=n("el-card"),E=n("el-dialog"),B=G("loading");return d(),I("div",ee,[a(E,{title:t.dialog.title,modelValue:t.dialog.isShowDialog,"onUpdate:modelValue":o[8]||(o[8]=e=>t.dialog.isShowDialog=e),width:"80%"},{default:l(()=>[a($,{shadow:"hover",class:"layout-padding-auto"},{default:l(()=>[m("div",te,[ae,a(s,{size:"default",modelValue:t.traceId,"onUpdate:modelValue":o[0]||(o[0]=e=>t.traceId=e),placeholder:"链路ID",style:{"max-width":"180px"}},null,8,["modelValue"]),oe,a(s,{size:"default",modelValue:t.appName,"onUpdate:modelValue":o[1]||(o[1]=e=>t.appName=e),placeholder:"应用名称",style:{"max-width":"150px","padding-left":"5px"}},null,8,["modelValue"]),le,a(s,{size:"default",modelValue:t.appIp,"onUpdate:modelValue":o[2]||(o[2]=e=>t.appIp=e),placeholder:"执行端IP",style:{"max-width":"120px","padding-left":"5px"}},null,8,["modelValue"]),ie,a(s,{size:"default",modelValue:t.logContent,"onUpdate:modelValue":o[3]||(o[3]=e=>t.logContent=e),placeholder:"日志内容",style:{"max-width":"250px","padding-left":"5px"}},null,8,["modelValue"]),ne,a(T,{modelValue:t.logLevel,"onUpdate:modelValue":o[4]||(o[4]=e=>t.logLevel=e),placeholder:"日志类型",class:"ml10",style:{"max-width":"110px"},size:"mini"},{default:l(()=>[a(g,{label:"全部",value:-1}),a(g,{label:"Trace",value:0}),a(g,{label:"Debug",value:1}),a(g,{label:"Info",value:2}),a(g,{label:"Warning",value:3}),a(g,{label:"Error",value:4}),a(g,{label:"Critical",value:5})]),_:1},8,["modelValue"]),a(y,{size:"default",type:"primary",class:"ml10",onClick:k},{default:l(()=>[a(A,null,{default:l(()=>[a(U)]),_:1}),p(" 查询 ")]),_:1})]),J((d(),f(P,{data:t.tableData.data,style:{width:"100%"}},{default:l(()=>[a(_,{width:"180px",label:"LogID","show-overflow-tooltip":""},{default:l(e=>[m("span",{onClick:W=>h(e.row)},r(e.row.LogIdN),9,pe)]),_:1}),a(_,{width:"200px",label:"应用","show-overflow-tooltip":""},{default:l(e=>[a(c,{size:"mini"},{default:l(()=>[p(r(e.row.AppName)+" "+r(e.row.AppIp),1)]),_:2},1024),re,p(" "+r(e.row.AppId),1)]),_:1}),a(_,{width:"100px",prop:"LogLevel",label:"日志类型","show-overflow-tooltip":""},{default:l(e=>[e.row.LogLevel=="Info"?(d(),f(c,{key:0,size:"mini"},{default:l(()=>[p(r(e.row.LogLevel),1)]),_:2},1024)):e.row.LogLevel=="Debug"?(d(),f(c,{key:1,type:"info",size:"mini"},{default:l(()=>[p(r(e.row.LogLevel),1)]),_:2},1024)):e.row.LogLevel=="Warn"?(d(),f(c,{key:2,type:"warning",size:"mini"},{default:l(()=>[p(r(e.row.LogLevel),1)]),_:2},1024)):e.row.LogLevel=="Error"?(d(),f(c,{key:3,type:"danger",size:"mini"},{default:l(()=>[p(r(e.row.LogLevel),1)]),_:2},1024)):v("",!0)]),_:1}),a(_,{label:"内容","show-overflow-tooltip":""},{default:l(e=>[e.row.TraceId>0?(d(),f(c,{key:0,type:"info",size:"mini"},{default:l(()=>[p("TraceID "+r(e.row.TraceId),1)]),_:2},1024)):v("",!0),e.row.TraceId>0?(d(),I("br",se)):v("",!0),p(" "+r(e.row.Content),1)]),_:1}),a(_,{width:"180px",prop:"Component",label:"组件","show-overflow-tooltip":""}),a(_,{width:"180px",prop:"CreateAt",label:"请求时间","show-overflow-tooltip":""}),a(_,{label:"操作",width:"100"},{default:l(e=>[a(y,{size:"small",text:"",type:"primary",onClick:W=>h(e.row)},{default:l(()=>[p("详情")]),_:2},1032,["onClick"])]),_:1})]),_:1},8,["data"])),[[B,t.tableData.loading]]),a(R,{onSizeChange:S,onCurrentChange:N,class:"mt15","pager-count":5,"page-sizes":[10,20,30],"current-page":t.tableData.param.pageNum,"onUpdate:currentPage":o[5]||(o[5]=e=>t.tableData.param.pageNum=e),background:"","page-size":t.tableData.param.pageSize,"onUpdate:pageSize":o[6]||(o[6]=e=>t.tableData.param.pageSize=e),layout:"total, sizes, prev, pager, next, jumper",total:t.tableData.total},null,8,["current-page","page-size","total"])]),_:1}),a(K(C),{ref_key:"detailDialogRef",ref:D,onRefresh:o[7]||(o[7]=e=>u())},null,512)]),_:1},8,["title","modelValue"])])}}});const Ge=Z(me,[["__scopeId","data-v-3b45e0a2"]]);export{Ge as default};