import{_ as F}from"./index-1fb9eb7d.js";import{f as B}from"./index-3be0d155.js";import{k as V,aN as K,e as L,$,i as H,ah as n,ar as Q,o as f,c as y,W as t,Q as o,a as p,U as d,R as Y,P as w,V as m,T as D,u as j,bb as O,b9 as W}from"./@vue-d9027515.js";import{_ as q}from"./_plugin-vue_export-helper-c27b6911.js";import"./pinia-be7e992d.js";import"./vue-demi-71ba0ef2.js";import"./vue-router-a2fac16f.js";import"./vue-i18n-e0b9a81d.js";import"./@intlify-9e8a497c.js";import"./source-map-7d7e1c08.js";import"./vue-b46b26c7.js";import"./js-cookie-cf83ad76.js";import"./@element-plus-96a31696.js";import"./nprogress-8d2808ea.js";import"./axios-707ed124.js";import"./qs-bbfcbdf3.js";import"./side-channel-eafc5c70.js";import"./get-intrinsic-b9397c9a.js";import"./has-symbols-e8f3ca0e.js";import"./function-bind-22e7ee79.js";import"./has-26d28e02.js";import"./call-bind-9ceb8f5b.js";import"./object-inspect-7b08a46e.js";import"./element-plus-7da03efa.js";import"./lodash-es-36eb724a.js";import"./@vueuse-2d9216d7.js";import"./@ctrl-f8748455.js";import"./dayjs-6aac7dfa.js";import"./async-validator-dee29e8b.js";import"./memoize-one-297ddbcb.js";import"./escape-html-1d60d822.js";import"./normalize-wheel-es-ed76fb12.js";import"./@floating-ui-463e90e0.js";import"./mitt-f7ef348c.js";import"./vue-grid-layout-0ee01be3.js";import"./requestFSGet-d236e20b.js";const u=g=>(O("data-v-26207a12"),g=g(),W(),g),G={class:"system-user-container layout-padding"},J={class:"system-user-search mb15"},X=u(()=>p("label",null,"应用名称",-1)),Z=u(()=>p("label",{class:"ml10"},"执行端IP",-1)),ee=u(()=>p("label",{class:"ml10"},"KEY",-1)),ae=u(()=>p("label",{class:"ml10"},"Field",-1)),te=u(()=>p("label",{class:"ml10"},"往前推N分钟的数据",-1)),le=u(()=>p("label",{class:"ml10"},"执行时间大于x毫秒的记录",-1)),oe=["onClick"],ne=u(()=>p("br",null,null,-1)),pe=u(()=>p("br",null,null,-1)),ie={key:1},se=V({name:"fopsTaskRunning"}),re=V({...se,setup(g){const S=B(),k=K(()=>F(()=>import("./detailV2Dialog-cc846fa7.js"),["assets/js/detailV2Dialog-cc846fa7.js","assets/js/index-3be0d155.js","assets/js/index-1fb9eb7d.js","assets/js/@vue-d9027515.js","assets/js/pinia-be7e992d.js","assets/js/vue-demi-71ba0ef2.js","assets/js/vue-router-a2fac16f.js","assets/js/vue-i18n-e0b9a81d.js","assets/js/@intlify-9e8a497c.js","assets/js/source-map-7d7e1c08.js","assets/js/vue-b46b26c7.js","assets/js/js-cookie-cf83ad76.js","assets/js/@element-plus-96a31696.js","assets/js/nprogress-8d2808ea.js","assets/css/nprogress-8b89e2e0.css","assets/js/axios-707ed124.js","assets/js/qs-bbfcbdf3.js","assets/js/side-channel-eafc5c70.js","assets/js/get-intrinsic-b9397c9a.js","assets/js/has-symbols-e8f3ca0e.js","assets/js/function-bind-22e7ee79.js","assets/js/has-26d28e02.js","assets/js/call-bind-9ceb8f5b.js","assets/js/object-inspect-7b08a46e.js","assets/js/element-plus-7da03efa.js","assets/js/lodash-es-36eb724a.js","assets/js/@vueuse-2d9216d7.js","assets/js/@ctrl-f8748455.js","assets/js/dayjs-6aac7dfa.js","assets/js/async-validator-dee29e8b.js","assets/js/memoize-one-297ddbcb.js","assets/js/escape-html-1d60d822.js","assets/js/normalize-wheel-es-ed76fb12.js","assets/js/@floating-ui-463e90e0.js","assets/js/mitt-f7ef348c.js","assets/js/vue-grid-layout-0ee01be3.js","assets/css/index-c2520709.css","assets/js/requestFSGet-d236e20b.js","assets/css/detailV2Dialog-abc41127.css"])),h=L(),e=$({appName:"",appIp:"",key:"",field:"",startMin:0,searchUseTs:0,tableData:{data:[],total:0,loading:!1,param:{pageNum:1,pageSize:10}}}),b=()=>{e.tableData.loading=!0;var s={appName:e.appName,appIp:e.appIp,key:e.key,field:e.field,startMin:e.startMin.toString(),searchUseTs:e.searchUseTs.toString(),pageSize:e.tableData.param.pageSize.toString(),pageIndex:e.tableData.param.pageNum.toString()};const l=new URLSearchParams(s).toString();S.slowRedis(l).then(function(i){i.Status?(e.tableData.data=i.Data.List,e.tableData.total=i.Data.RecordCount,e.tableData.loading=!1):(e.tableData.data=[],e.tableData.loading=!1)})},x=s=>{h.value.openDialog(s)},C=s=>{e.tableData.param.pageSize=s,b()},z=s=>{e.tableData.param.pageNum=s,b()},N=()=>{b()};return H(()=>{b()}),(s,l)=>{const i=n("el-input"),c=n("el-option"),I=n("el-select"),E=n("ele-Search"),U=n("el-icon"),v=n("el-button"),r=n("el-table-column"),_=n("el-tag"),T=n("el-table"),R=n("el-pagination"),A=n("el-card"),M=Q("loading");return f(),y("div",G,[t(A,{shadow:"hover",class:"layout-padding-auto"},{default:o(()=>[p("div",J,[X,t(i,{size:"default",modelValue:e.appName,"onUpdate:modelValue":l[0]||(l[0]=a=>e.appName=a),placeholder:"应用名称",style:{"max-width":"180px"}},null,8,["modelValue"]),Z,t(i,{size:"default",modelValue:e.appIp,"onUpdate:modelValue":l[1]||(l[1]=a=>e.appIp=a),placeholder:"执行端IP",style:{"max-width":"180px"}},null,8,["modelValue"]),ee,t(i,{size:"default",modelValue:e.key,"onUpdate:modelValue":l[2]||(l[2]=a=>e.key=a),placeholder:"KEY",style:{"max-width":"180px"}},null,8,["modelValue"]),ae,t(i,{size:"default",modelValue:e.field,"onUpdate:modelValue":l[3]||(l[3]=a=>e.field=a),placeholder:"Field",style:{"max-width":"180px"}},null,8,["modelValue"]),te,t(I,{modelValue:e.startMin,"onUpdate:modelValue":l[4]||(l[4]=a=>e.startMin=a),placeholder:"往前推N分钟的数据",clearable:"",class:"ml10"},{default:o(()=>[t(c,{label:"全部",value:0}),t(c,{label:"1小时耗时最高",value:60}),t(c,{label:"30分钟耗时最高",value:30}),t(c,{label:"10分钟耗时最高",value:10}),t(c,{label:"5分钟耗时最高",value:5}),t(c,{label:"1分钟耗时最高",value:1})]),_:1},8,["modelValue"]),le,t(i,{size:"default",modelValue:e.searchUseTs,"onUpdate:modelValue":l[5]||(l[5]=a=>e.searchUseTs=a),placeholder:"执行时间大于x毫秒的记录",style:{"max-width":"180px"}},null,8,["modelValue"]),t(v,{size:"default",type:"primary",class:"ml10",onClick:N},{default:o(()=>[t(U,null,{default:o(()=>[t(E)]),_:1}),d(" 查询 ")]),_:1})]),Y((f(),w(T,{data:e.tableData.data,style:{width:"100%"}},{default:o(()=>[t(r,{width:"180px",label:"TraceID","show-overflow-tooltip":""},{default:o(a=>[p("span",{onClick:P=>x(a.row)},m(a.row.TraceId),9,oe)]),_:1}),t(r,{width:"250px",label:"应用","show-overflow-tooltip":""},{default:o(a=>[t(_,null,{default:o(()=>[d(m(a.row.AppName),1)]),_:2},1024),ne,t(_,null,{default:o(()=>[d(m(a.row.AppIp),1)]),_:2},1024),pe,t(_,null,{default:o(()=>[d(m(a.row.AppId),1)]),_:2},1024)]),_:1}),t(r,{width:"120px",prop:"UseDesc",label:"执行耗时","show-overflow-tooltip":""}),t(r,{prop:"Key",label:"Key","show-overflow-tooltip":""}),t(r,{prop:"Field",label:"Field","show-overflow-tooltip":""}),t(r,{width:"300px",label:"异常","show-overflow-tooltip":""},{default:o(a=>[a.row.Exception!=null?(f(),w(_,{key:0},{default:o(()=>[d(m(a.row.Exception.ExceptionCallFile)+":"+m(a.row.Exception.ExceptionCallLine)+" "+m(a.row.Exception.ExceptionCallFuncName),1)]),_:2},1024)):D("",!0),a.row.Exception!=null?(f(),y("br",ie)):D("",!0),a.row.Exception!=null?(f(),w(_,{key:2},{default:o(()=>[d(m(a.row.Exception.ExceptionMessage),1)]),_:2},1024)):(f(),w(_,{key:3},{default:o(()=>[d("无")]),_:1}))]),_:1}),t(r,{width:"180px",prop:"CreateAt",label:"执行时间","show-overflow-tooltip":""}),t(r,{label:"操作",width:"100"},{default:o(a=>[t(v,{size:"small",text:"",type:"primary",onClick:P=>x(a.row)},{default:o(()=>[d("追踪")]),_:2},1032,["onClick"])]),_:1})]),_:1},8,["data"])),[[M,e.tableData.loading]]),t(R,{onSizeChange:C,onCurrentChange:z,class:"mt15","pager-count":5,"page-sizes":[10,20,30],"current-page":e.tableData.param.pageNum,"onUpdate:currentPage":l[6]||(l[6]=a=>e.tableData.param.pageNum=a),background:"","page-size":e.tableData.param.pageSize,"onUpdate:pageSize":l[7]||(l[7]=a=>e.tableData.param.pageSize=a),layout:"total, sizes, prev, pager, next, jumper",total:e.tableData.total},null,8,["current-page","page-size","total"])]),_:1}),t(j(k),{ref_key:"detailDialogRef",ref:h,onRefresh:l[8]||(l[8]=a=>b())},null,512)])}}});const je=q(re,[["__scopeId","data-v-26207a12"]]);export{je as default};
