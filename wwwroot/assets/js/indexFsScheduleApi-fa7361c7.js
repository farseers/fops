import{_ as P}from"./index-858e7551.js";import{f as $}from"./index-0358cb65.js";import{k as I,aN as F,e as L,$ as B,f as H,i as Q,ah as p,ar as W,o as n,c as k,W as t,Q as o,a as c,U as s,R as j,P as u,V as i,T as b,u as O,bb as q,b9 as J}from"./@vue-d9027515.js";import{_ as K}from"./_plugin-vue_export-helper-c27b6911.js";import"./pinia-be7e992d.js";import"./vue-demi-71ba0ef2.js";import"./vue-router-a2fac16f.js";import"./vue-i18n-e0b9a81d.js";import"./@intlify-9e8a497c.js";import"./source-map-7d7e1c08.js";import"./vue-b46b26c7.js";import"./js-cookie-cf83ad76.js";import"./@element-plus-96a31696.js";import"./nprogress-8d2808ea.js";import"./axios-707ed124.js";import"./qs-bbfcbdf3.js";import"./side-channel-eafc5c70.js";import"./get-intrinsic-b9397c9a.js";import"./has-symbols-e8f3ca0e.js";import"./function-bind-22e7ee79.js";import"./has-26d28e02.js";import"./call-bind-9ceb8f5b.js";import"./object-inspect-7b08a46e.js";import"./element-plus-7da03efa.js";import"./lodash-es-36eb724a.js";import"./@vueuse-2d9216d7.js";import"./@ctrl-f8748455.js";import"./dayjs-6aac7dfa.js";import"./async-validator-dee29e8b.js";import"./memoize-one-297ddbcb.js";import"./escape-html-1d60d822.js";import"./normalize-wheel-es-ed76fb12.js";import"./@floating-ui-463e90e0.js";import"./mitt-f7ef348c.js";import"./vue-grid-layout-0ee01be3.js";import"./requestFSGet-6e3c6110.js";const _=h=>(q("data-v-e1c3fcb4"),h=h(),J(),h),X={class:"system-user-container layout-padding"},Y={class:"system-user-search mb15"},Z=_(()=>c("label",null,"应用名称",-1)),ee=_(()=>c("label",{class:"ml10"},"执行端IP",-1)),ae=_(()=>c("label",{class:"ml10"},"任务名称",-1)),te=_(()=>c("label",{class:"ml10"},"任务组ID",-1)),le=_(()=>c("label",{class:"ml10"},"任务ID",-1)),oe=_(()=>c("label",{class:"ml10"},"往前推N分钟的数据",-1)),se=_(()=>c("label",{class:"ml10"},"执行时间",-1)),ie=["onClick"],ne=_(()=>c("br",null,null,-1)),pe={key:2},re={key:1},de=I({name:"fopsTaskRunning"}),me=I({...de,setup(h){const v=$(),z=F(()=>P(()=>import("./detailV2Dialog-6e18abaf.js"),["assets/js/detailV2Dialog-6e18abaf.js","assets/js/index-858e7551.js","assets/js/@vue-d9027515.js","assets/js/pinia-be7e992d.js","assets/js/vue-demi-71ba0ef2.js","assets/js/vue-router-a2fac16f.js","assets/js/vue-i18n-e0b9a81d.js","assets/js/@intlify-9e8a497c.js","assets/js/source-map-7d7e1c08.js","assets/js/vue-b46b26c7.js","assets/js/js-cookie-cf83ad76.js","assets/js/@element-plus-96a31696.js","assets/js/nprogress-8d2808ea.js","assets/css/nprogress-8b89e2e0.css","assets/js/axios-707ed124.js","assets/js/qs-bbfcbdf3.js","assets/js/side-channel-eafc5c70.js","assets/js/get-intrinsic-b9397c9a.js","assets/js/has-symbols-e8f3ca0e.js","assets/js/function-bind-22e7ee79.js","assets/js/has-26d28e02.js","assets/js/call-bind-9ceb8f5b.js","assets/js/object-inspect-7b08a46e.js","assets/js/element-plus-7da03efa.js","assets/js/lodash-es-36eb724a.js","assets/js/@vueuse-2d9216d7.js","assets/js/@ctrl-f8748455.js","assets/js/dayjs-6aac7dfa.js","assets/js/async-validator-dee29e8b.js","assets/js/memoize-one-297ddbcb.js","assets/js/escape-html-1d60d822.js","assets/js/normalize-wheel-es-ed76fb12.js","assets/js/@floating-ui-463e90e0.js","assets/js/mitt-f7ef348c.js","assets/js/vue-grid-layout-0ee01be3.js","assets/css/index-c2520709.css","assets/js/index-0358cb65.js","assets/js/requestFSGet-6e3c6110.js","assets/css/detailV2Dialog-9a9be94f.css"])),x=L(),a=B({keyWord:"",appName:"",appIp:"",taskName:"",taskGroupId:"",searchUseTs:0,taskId:"",startMin:0,tableData:{data:[],total:0,loading:!1,param:{pageNum:1,pageSize:10}}});H(()=>a.startMin,(r,l)=>{console.log(`count 从 ${l} 变为 ${r}`),g()});const g=()=>{a.tableData.loading=!0;var r={appName:a.appName,appIp:a.appIp,taskName:a.taskName,taskGroupId:a.taskGroupId.toString(),startMin:a.startMin.toString(),searchUseTs:a.searchUseTs.toString(),taskId:a.taskId.toString(),pageSize:a.tableData.param.pageSize.toString(),pageIndex:a.tableData.param.pageNum.toString()};const l=new URLSearchParams(r).toString();v.linkTraceFScheduleList(l).then(function(d){d.Status?(a.tableData.data=d.Data.List,a.tableData.total=d.Data.RecordCount,a.tableData.loading=!1):(a.tableData.data=[],a.tableData.loading=!1)})},y=r=>{x.value.openDialog(r)},S=r=>{a.tableData.param.pageSize=r,g()},V=r=>{a.tableData.param.pageNum=r,g()},N=()=>{g()};return Q(()=>{g()}),(r,l)=>{const d=p("el-input"),w=p("el-option"),U=p("el-select"),T=p("ele-Search"),C=p("el-icon"),D=p("el-button"),f=p("el-table-column"),m=p("el-tag"),E=p("el-table"),A=p("el-pagination"),G=p("el-card"),M=W("loading");return n(),k("div",X,[t(G,{shadow:"hover",class:"layout-padding-auto"},{default:o(()=>[c("div",Y,[Z,t(d,{size:"default",modelValue:a.appName,"onUpdate:modelValue":l[0]||(l[0]=e=>a.appName=e),placeholder:"应用名称",style:{"max-width":"150px","padding-left":"5px"}},null,8,["modelValue"]),ee,t(d,{size:"default",modelValue:a.appIp,"onUpdate:modelValue":l[1]||(l[1]=e=>a.appIp=e),placeholder:"执行端IP",style:{"max-width":"120px","padding-left":"5px"}},null,8,["modelValue"]),ae,t(d,{size:"default",modelValue:a.taskName,"onUpdate:modelValue":l[2]||(l[2]=e=>a.taskName=e),placeholder:"任务名称",style:{"max-width":"180px","padding-left":"5px"}},null,8,["modelValue"]),te,t(d,{size:"default",modelValue:a.taskGroupId,"onUpdate:modelValue":l[3]||(l[3]=e=>a.taskGroupId=e),placeholder:"任务组ID",style:{"max-width":"90px","padding-left":"5px"}},null,8,["modelValue"]),le,t(d,{size:"default",modelValue:a.taskId,"onUpdate:modelValue":l[4]||(l[4]=e=>a.taskId=e),placeholder:"任务ID",style:{"max-width":"180px","padding-left":"5px"}},null,8,["modelValue"]),oe,t(U,{modelValue:a.startMin,"onUpdate:modelValue":l[5]||(l[5]=e=>a.startMin=e),placeholder:"往前推N分钟的数据",class:"ml10",style:{"max-width":"150px"},size:"mini"},{default:o(()=>[t(w,{label:"全部",value:0}),t(w,{label:"1小时耗时最高",value:60}),t(w,{label:"30分钟耗时最高",value:30}),t(w,{label:"10分钟耗时最高",value:10}),t(w,{label:"5分钟耗时最高",value:5}),t(w,{label:"1分钟耗时最高",value:1})]),_:1},8,["modelValue"]),se,t(d,{size:"default",modelValue:a.searchUseTs,"onUpdate:modelValue":l[6]||(l[6]=e=>a.searchUseTs=e),placeholder:"执行时间大于毫秒的记录",style:{"max-width":"80px","padding-left":"5px"}},null,8,["modelValue"]),s(" ms "),t(D,{size:"default",type:"primary",class:"ml10",onClick:N},{default:o(()=>[t(C,null,{default:o(()=>[t(T)]),_:1}),s(" 查询 ")]),_:1})]),j((n(),u(E,{data:a.tableData.data,style:{width:"100%"}},{default:o(()=>[t(f,{width:"180px",label:"TraceID","show-overflow-tooltip":""},{default:o(e=>[c("span",{onClick:R=>y(e.row)},i(e.row.TraceIdN),9,ie)]),_:1}),t(f,{width:"200px",label:"应用","show-overflow-tooltip":""},{default:o(e=>[t(m,{size:"mini"},{default:o(()=>[s(i(e.row.AppName)+" "+i(e.row.AppIp),1)]),_:2},1024),ne,s(" "+i(e.row.AppId),1)]),_:1}),t(f,{width:"120px",prop:"UseDesc",label:"执行耗时","show-overflow-tooltip":""},{default:o(e=>[e.row.UseTs>1e8?(n(),u(m,{key:0,size:"mini",type:"danger"},{default:o(()=>[s(i(e.row.UseDesc),1)]),_:2},1024)):e.row.UseTs>5e7?(n(),u(m,{key:1,size:"mini",type:"warning"},{default:o(()=>[s(i(e.row.UseDesc),1)]),_:2},1024)):e.row.UseTs>1e6?(n(),u(m,{key:2,size:"mini"},{default:o(()=>[s(i(e.row.UseDesc),1)]),_:2},1024)):(n(),u(m,{key:3,size:"mini",type:"success"},{default:o(()=>[s(i(e.row.UseDesc),1)]),_:2},1024))]),_:1}),t(f,{label:"任务名称","show-overflow-tooltip":""},{default:o(e=>[e.row.TaskGroupId>0?(n(),u(m,{key:0,size:"mini"},{default:o(()=>[s("任务组Id："+i(e.row.TaskGroupId),1)]),_:2},1024)):b("",!0),e.row.TaskId>0?(n(),u(m,{key:1,size:"mini",type:"success"},{default:o(()=>[s("任务Id："+i(e.row.TaskId),1)]),_:2},1024)):b("",!0),e.row.TaskGroupId>0||e.row.TaskId>0?(n(),k("br",pe)):b("",!0),s(" "+i(e.row.TaskName),1)]),_:1}),t(f,{width:"300px",label:"异常","show-overflow-tooltip":""},{default:o(e=>[e.row.Exception!=null?(n(),u(m,{key:0,size:"mini"},{default:o(()=>[s(i(e.row.Exception.ExceptionCallFile)+":"+i(e.row.Exception.ExceptionCallLine)+" "+i(e.row.Exception.ExceptionCallFuncName),1)]),_:2},1024)):b("",!0),e.row.Exception!=null?(n(),k("br",re)):b("",!0),e.row.Exception!=null?(n(),u(m,{key:2,size:"mini"},{default:o(()=>[s(i(e.row.Exception.ExceptionMessage),1)]),_:2},1024)):(n(),u(m,{key:3,size:"mini",type:"info"},{default:o(()=>[s("无")]),_:1}))]),_:1}),t(f,{width:"180px",prop:"CreateAt",label:"请求时间","show-overflow-tooltip":""}),t(f,{label:"操作",width:"100"},{default:o(e=>[t(D,{size:"small",text:"",type:"primary",onClick:R=>y(e.row)},{default:o(()=>[s("追踪")]),_:2},1032,["onClick"])]),_:1})]),_:1},8,["data"])),[[M,a.tableData.loading]]),t(A,{onSizeChange:S,onCurrentChange:V,class:"mt15","pager-count":5,"page-sizes":[10,20,30],"current-page":a.tableData.param.pageNum,"onUpdate:currentPage":l[7]||(l[7]=e=>a.tableData.param.pageNum=e),background:"","page-size":a.tableData.param.pageSize,"onUpdate:pageSize":l[8]||(l[8]=e=>a.tableData.param.pageSize=e),layout:"total, sizes, prev, pager, next, jumper",total:a.tableData.total},null,8,["current-page","page-size","total"])]),_:1}),t(O(z),{ref_key:"detailDialogRef",ref:x,onRefresh:l[9]||(l[9]=e=>g())},null,512)])}}});const qe=K(me,[["__scopeId","data-v-e1c3fcb4"]]);export{qe as default};
