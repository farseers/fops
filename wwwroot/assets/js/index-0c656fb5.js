import{_ as T}from"./index-a508280f.js";import{f as A}from"./index-43fd20bc.js";import{k as g,aN as P,e as U,$ as G,i as B,n as L,ah as r,ar as E,o as s,c as H,W as e,Q as n,a as Q,U as u,R as W,P as m,T as _,u as $}from"./@vue-d9027515.js";import{_ as j}from"./_plugin-vue_export-helper-c27b6911.js";import"./pinia-be7e992d.js";import"./vue-demi-71ba0ef2.js";import"./vue-router-a2fac16f.js";import"./vue-i18n-e0b9a81d.js";import"./@intlify-9e8a497c.js";import"./source-map-7d7e1c08.js";import"./vue-b46b26c7.js";import"./js-cookie-cf83ad76.js";import"./@element-plus-96a31696.js";import"./nprogress-8d2808ea.js";import"./axios-707ed124.js";import"./qs-bbfcbdf3.js";import"./side-channel-eafc5c70.js";import"./get-intrinsic-b9397c9a.js";import"./has-symbols-e8f3ca0e.js";import"./function-bind-22e7ee79.js";import"./has-26d28e02.js";import"./call-bind-9ceb8f5b.js";import"./object-inspect-7b08a46e.js";import"./element-plus-7da03efa.js";import"./lodash-es-36eb724a.js";import"./@vueuse-2d9216d7.js";import"./@ctrl-f8748455.js";import"./dayjs-6aac7dfa.js";import"./async-validator-dee29e8b.js";import"./memoize-one-297ddbcb.js";import"./escape-html-1d60d822.js";import"./normalize-wheel-es-ed76fb12.js";import"./@floating-ui-463e90e0.js";import"./mitt-f7ef348c.js";import"./vue-grid-layout-0ee01be3.js";import"./requestFSGet-738ae044.js";const M={class:"system-user-container layout-padding"},O={class:"system-user-search mb15"},q=g({name:"fopsTask"}),F=g({...q,setup(J){const f=A(),b=P(()=>T(()=>import("./editGroupDialog-d0a6666d.js"),["assets/js/editGroupDialog-d0a6666d.js","assets/js/index-43fd20bc.js","assets/js/index-a508280f.js","assets/js/@vue-d9027515.js","assets/js/pinia-be7e992d.js","assets/js/vue-demi-71ba0ef2.js","assets/js/vue-router-a2fac16f.js","assets/js/vue-i18n-e0b9a81d.js","assets/js/@intlify-9e8a497c.js","assets/js/source-map-7d7e1c08.js","assets/js/vue-b46b26c7.js","assets/js/js-cookie-cf83ad76.js","assets/js/@element-plus-96a31696.js","assets/js/nprogress-8d2808ea.js","assets/css/nprogress-8b89e2e0.css","assets/js/axios-707ed124.js","assets/js/qs-bbfcbdf3.js","assets/js/side-channel-eafc5c70.js","assets/js/get-intrinsic-b9397c9a.js","assets/js/has-symbols-e8f3ca0e.js","assets/js/function-bind-22e7ee79.js","assets/js/has-26d28e02.js","assets/js/call-bind-9ceb8f5b.js","assets/js/object-inspect-7b08a46e.js","assets/js/element-plus-7da03efa.js","assets/js/lodash-es-36eb724a.js","assets/js/@vueuse-2d9216d7.js","assets/js/@ctrl-f8748455.js","assets/js/dayjs-6aac7dfa.js","assets/js/async-validator-dee29e8b.js","assets/js/memoize-one-297ddbcb.js","assets/js/escape-html-1d60d822.js","assets/js/normalize-wheel-es-ed76fb12.js","assets/js/@floating-ui-463e90e0.js","assets/js/mitt-f7ef348c.js","assets/js/vue-grid-layout-0ee01be3.js","assets/css/index-c2520709.css","assets/js/requestFSGet-738ae044.js"])),S=U(),t=G({keyWord:"",enable:-1,taskStatus:-1,taskGroupId:"",tableData:{data:[],total:0,loading:!1,param:{pageNum:1,pageSize:10}}}),c=()=>{t.tableData.loading=!0;const l=new URLSearchParams;l.append("taskStatus",t.taskStatus.toString()),l.append("taskGroupId",t.taskGroupId),l.append("pageSize",t.tableData.param.pageSize.toString()),l.append("pageIndex",t.tableData.param.pageNum.toString()),f.taskList(l.toString()).then(function(o){o.Status?(t.tableData.data=o.Data.List,t.tableData.total=o.Data.RecordCount,setTimeout(()=>{t.tableData.loading=!1},500)):(t.tableData.data=[],setTimeout(()=>{t.tableData.loading=!1},500))})},v=()=>{c()},k=l=>{t.tableData.param.pageSize=l,c()},D=l=>{t.tableData.param.pageNum=l,c()},w=l=>{t.taskStatus=l};return B(()=>{L(()=>{c()})}),(l,o)=>{const h=r("el-input"),i=r("el-option"),y=r("el-select"),C=r("ele-Search"),z=r("el-icon"),x=r("el-button"),p=r("el-table-column"),d=r("el-tag"),N=r("el-table"),V=r("el-pagination"),I=r("el-card"),R=E("loading");return s(),H("div",M,[e(I,{shadow:"hover",class:"layout-padding-auto"},{default:n(()=>[Q("div",O,[e(h,{size:"default",modelValue:t.taskGroupId,"onUpdate:modelValue":o[0]||(o[0]=a=>t.taskGroupId=a),placeholder:"请输入任务组ID",style:{"max-width":"180px"},class:"ml10"},null,8,["modelValue"]),e(y,{modelValue:t.taskStatus,"onUpdate:modelValue":o[1]||(o[1]=a=>t.taskStatus=a),placeholder:"请选择调度状态",class:"ml10",onChange:w},{default:n(()=>[e(i,{label:"全部",value:-1}),e(i,{label:"未开始",value:0}),e(i,{label:"调度中",value:1}),e(i,{label:"调度失败",value:2}),e(i,{label:"执行中",value:3}),e(i,{label:"失败",value:4}),e(i,{label:"成功",value:5})]),_:1},8,["modelValue"]),e(x,{size:"default",type:"primary",class:"ml10",onClick:v},{default:n(()=>[e(z,null,{default:n(()=>[e(C)]),_:1}),u(" 查询 ")]),_:1})]),W((s(),m(N,{data:t.tableData.data,style:{width:"100%"}},{default:n(()=>[e(p,{prop:"Id",label:"序号",width:"60"}),e(p,{prop:"Name",label:"任务组名称","show-overflow-tooltip":""}),e(p,{prop:"Caption",label:"任务组标题","show-overflow-tooltip":""}),e(p,{prop:"StartAt",label:"开始时间","show-overflow-tooltip":""}),e(p,{prop:"RunAt",label:"运行时间","show-overflow-tooltip":""}),e(p,{prop:"RunSpeed",label:"运行耗时","show-overflow-tooltip":""}),e(p,{prop:"Progress",label:"进度","show-overflow-tooltip":""}),e(p,{label:"任务状态","show-overflow-tooltip":""},{default:n(a=>[a.row.Status==0?(s(),m(d,{key:0},{default:n(()=>[u("未开始")]),_:1})):_("",!0),a.row.Status==1?(s(),m(d,{key:1},{default:n(()=>[u("调度中")]),_:1})):_("",!0),a.row.Status==2?(s(),m(d,{key:2},{default:n(()=>[u("调度失败")]),_:1})):_("",!0),a.row.Status==3?(s(),m(d,{key:3},{default:n(()=>[u("执行中")]),_:1})):_("",!0),a.row.Status==4?(s(),m(d,{key:4},{default:n(()=>[u("失败")]),_:1})):_("",!0),a.row.Status==5?(s(),m(d,{key:5},{default:n(()=>[u("成功")]),_:1})):_("",!0)]),_:1})]),_:1},8,["data"])),[[R,t.tableData.loading]]),e(V,{onSizeChange:k,onCurrentChange:D,class:"mt15","pager-count":5,"page-sizes":[10,20,30],"current-page":t.tableData.param.pageNum,"onUpdate:currentPage":o[2]||(o[2]=a=>t.tableData.param.pageNum=a),background:"","page-size":t.tableData.param.pageSize,"onUpdate:pageSize":o[3]||(o[3]=a=>t.tableData.param.pageSize=a),layout:"total, sizes, prev, pager, next, jumper",total:t.tableData.total},null,8,["current-page","page-size","total"])]),_:1}),e($(b),{ref_key:"editDialogRef",ref:S,onRefresh:o[4]||(o[4]=a=>c())},null,512)])}}});const Tt=j(F,[["__scopeId","data-v-1963b33f"]]);export{Tt as default};
