import{_ as V}from"./index-32e628b7.js";import{f as R}from"./index-a5919b88.js";import{k as g,aN as T,e as U,$ as A,i as E,ah as r,ar as G,o as p,c as P,W as t,Q as l,a as B,U as s,R as W,P as m,T as _,u as H}from"./@vue-d9027515.js";import{_ as Q}from"./_plugin-vue_export-helper-c27b6911.js";import"./pinia-be7e992d.js";import"./vue-demi-71ba0ef2.js";import"./vue-router-a2fac16f.js";import"./vue-i18n-e0b9a81d.js";import"./@intlify-9e8a497c.js";import"./source-map-7d7e1c08.js";import"./vue-b46b26c7.js";import"./js-cookie-cf83ad76.js";import"./@element-plus-96a31696.js";import"./nprogress-8d2808ea.js";import"./axios-707ed124.js";import"./qs-bbfcbdf3.js";import"./side-channel-eafc5c70.js";import"./get-intrinsic-b9397c9a.js";import"./has-symbols-e8f3ca0e.js";import"./function-bind-22e7ee79.js";import"./has-26d28e02.js";import"./call-bind-9ceb8f5b.js";import"./object-inspect-7b08a46e.js";import"./element-plus-7da03efa.js";import"./lodash-es-36eb724a.js";import"./@vueuse-2d9216d7.js";import"./@ctrl-f8748455.js";import"./dayjs-6aac7dfa.js";import"./async-validator-dee29e8b.js";import"./memoize-one-297ddbcb.js";import"./escape-html-1d60d822.js";import"./normalize-wheel-es-ed76fb12.js";import"./@floating-ui-463e90e0.js";import"./mitt-f7ef348c.js";import"./vue-grid-layout-0ee01be3.js";const $={class:"system-user-container layout-padding"},j={class:"system-user-search mb15"},J=g({name:"fopsTaskRunning"}),M=g({...J,setup(O){const f=R(),b=T(()=>V(()=>import("./editGroupDialog-f3bb8bf2.js"),["assets/js/editGroupDialog-f3bb8bf2.js","assets/js/index-a5919b88.js","assets/js/index-32e628b7.js","assets/js/@vue-d9027515.js","assets/js/pinia-be7e992d.js","assets/js/vue-demi-71ba0ef2.js","assets/js/vue-router-a2fac16f.js","assets/js/vue-i18n-e0b9a81d.js","assets/js/@intlify-9e8a497c.js","assets/js/source-map-7d7e1c08.js","assets/js/vue-b46b26c7.js","assets/js/js-cookie-cf83ad76.js","assets/js/@element-plus-96a31696.js","assets/js/nprogress-8d2808ea.js","assets/css/nprogress-8b89e2e0.css","assets/js/axios-707ed124.js","assets/js/qs-bbfcbdf3.js","assets/js/side-channel-eafc5c70.js","assets/js/get-intrinsic-b9397c9a.js","assets/js/has-symbols-e8f3ca0e.js","assets/js/function-bind-22e7ee79.js","assets/js/has-26d28e02.js","assets/js/call-bind-9ceb8f5b.js","assets/js/object-inspect-7b08a46e.js","assets/js/element-plus-7da03efa.js","assets/js/lodash-es-36eb724a.js","assets/js/@vueuse-2d9216d7.js","assets/js/@ctrl-f8748455.js","assets/js/dayjs-6aac7dfa.js","assets/js/async-validator-dee29e8b.js","assets/js/memoize-one-297ddbcb.js","assets/js/escape-html-1d60d822.js","assets/js/normalize-wheel-es-ed76fb12.js","assets/js/@floating-ui-463e90e0.js","assets/js/mitt-f7ef348c.js","assets/js/vue-grid-layout-0ee01be3.js","assets/css/index-c2520709.css"])),v=U(),e=A({keyWord:"",taskGroupId:"",logLevel:-1,tableData:{data:[],total:0,loading:!1,param:{pageNum:1,pageSize:10}}}),c=()=>{e.tableData.loading=!0;const n=new URLSearchParams;n.append("logLevel",e.logLevel.toString()),n.append("taskGroupId",e.taskGroupId.toString()),n.append("pageSize",e.tableData.param.pageSize.toString()),n.append("pageIndex",e.tableData.param.pageNum.toString()),f.taskLogList(n.toString()).then(function(o){o.Status?(e.tableData.data=o.Data.List,e.tableData.total=o.Data.RecordCount,e.tableData.loading=!1):(e.tableData.data=[],e.tableData.loading=!1)})},D=n=>{e.tableData.param.pageSize=n,c()},L=n=>{e.tableData.param.pageNum=n,c()},w=()=>{c()};return E(()=>{c()}),(n,o)=>{const k=r("el-input"),d=r("el-option"),y=r("el-select"),h=r("ele-Search"),C=r("el-icon"),S=r("el-button"),i=r("el-table-column"),u=r("el-tag"),z=r("el-table"),I=r("el-pagination"),x=r("el-card"),N=G("loading");return p(),P("div",$,[t(x,{shadow:"hover",class:"layout-padding-auto"},{default:l(()=>[B("div",j,[t(k,{size:"default",modelValue:e.taskGroupId,"onUpdate:modelValue":o[0]||(o[0]=a=>e.taskGroupId=a),placeholder:"请输入任务组ID",style:{"max-width":"180px"}},null,8,["modelValue"]),t(y,{modelValue:e.logLevel,"onUpdate:modelValue":o[1]||(o[1]=a=>e.logLevel=a),placeholder:"请选择日志等级",clearable:"",class:"ml10"},{default:l(()=>[t(d,{label:"全部",value:-1}),t(d,{label:"Trace",value:0}),t(d,{label:"Debug",value:1}),t(d,{label:"Info",value:2}),t(d,{label:"Warning",value:3}),t(d,{label:"Error",value:4}),t(d,{label:"Critical",value:5})]),_:1},8,["modelValue"]),t(S,{size:"default",type:"primary",class:"ml10",onClick:w},{default:l(()=>[t(C,null,{default:l(()=>[t(h)]),_:1}),s(" 查询 ")]),_:1})]),W((p(),m(z,{data:e.tableData.data,style:{width:"100%"}},{default:l(()=>[t(i,{prop:"Id",label:"序号",width:"180"}),t(i,{prop:"Name",label:"实现Job的特性名称","show-overflow-tooltip":""}),t(i,{prop:"Caption",label:"任务组标题","show-overflow-tooltip":""}),t(i,{prop:"Ver",label:"版本","show-overflow-tooltip":""}),t(i,{prop:"TaskId",label:"任务ID","show-overflow-tooltip":""}),t(i,{prop:"LogLevel",label:"日志级别","show-overflow-tooltip":""},{default:l(a=>[a.row.LogLevel==0?(p(),m(u,{key:0},{default:l(()=>[s("Trace")]),_:1})):_("",!0),a.row.LogLevel==1?(p(),m(u,{key:1},{default:l(()=>[s("Debug")]),_:1})):_("",!0),a.row.LogLevel==2?(p(),m(u,{key:2},{default:l(()=>[s("Information")]),_:1})):_("",!0),a.row.LogLevel==3?(p(),m(u,{key:3},{default:l(()=>[s("Warning")]),_:1})):_("",!0),a.row.LogLevel==4?(p(),m(u,{key:4},{default:l(()=>[s("Error")]),_:1})):_("",!0),a.row.LogLevel==5?(p(),m(u,{key:5},{default:l(()=>[s("Critical")]),_:1})):_("",!0),a.row.LogLevel==6?(p(),m(u,{key:6},{default:l(()=>[s("NoneLevel")]),_:1})):_("",!0)]),_:1}),t(i,{prop:"Content",label:"日志内容","show-overflow-tooltip":""}),t(i,{prop:"CreateAt",label:"日志时间","show-overflow-tooltip":""})]),_:1},8,["data"])),[[N,e.tableData.loading]]),t(I,{onSizeChange:D,onCurrentChange:L,class:"mt15","pager-count":5,"page-sizes":[10,20,30],"current-page":e.tableData.param.pageNum,"onUpdate:currentPage":o[2]||(o[2]=a=>e.tableData.param.pageNum=a),background:"","page-size":e.tableData.param.pageSize,"onUpdate:pageSize":o[3]||(o[3]=a=>e.tableData.param.pageSize=a),layout:"total, sizes, prev, pager, next, jumper",total:e.tableData.total},null,8,["current-page","page-size","total"])]),_:1}),t(H(b),{ref_key:"editDialogRef",ref:v,onRefresh:o[4]||(o[4]=a=>c())},null,512)])}}});const xe=Q(M,[["__scopeId","data-v-36f033c6"]]);export{xe as default};
