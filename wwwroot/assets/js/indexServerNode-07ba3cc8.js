import{_ as N}from"./index-a79ea617.js";import{f as T}from"./index-5e41c337.js";import{E as z,a as v}from"./element-plus-7da03efa.js";import{k as D,aN as B,e as E,$ as A,i as R,ah as n,ar as $,o as y,c as P,W as e,Q as s,a as U,U as d,R as L,P as M,u as W}from"./@vue-d9027515.js";import{_ as O}from"./_plugin-vue_export-helper-c27b6911.js";import"./pinia-be7e992d.js";import"./vue-demi-71ba0ef2.js";import"./vue-router-a2fac16f.js";import"./vue-i18n-e0b9a81d.js";import"./@intlify-9e8a497c.js";import"./source-map-7d7e1c08.js";import"./vue-b46b26c7.js";import"./js-cookie-cf83ad76.js";import"./@element-plus-96a31696.js";import"./nprogress-8d2808ea.js";import"./axios-707ed124.js";import"./qs-bbfcbdf3.js";import"./side-channel-eafc5c70.js";import"./get-intrinsic-b9397c9a.js";import"./has-symbols-e8f3ca0e.js";import"./function-bind-22e7ee79.js";import"./has-26d28e02.js";import"./call-bind-9ceb8f5b.js";import"./object-inspect-7b08a46e.js";import"./mitt-f7ef348c.js";import"./vue-grid-layout-0ee01be3.js";import"./lodash-es-36eb724a.js";import"./@vueuse-2d9216d7.js";import"./@ctrl-f8748455.js";import"./dayjs-6aac7dfa.js";import"./async-validator-dee29e8b.js";import"./memoize-one-297ddbcb.js";import"./escape-html-1d60d822.js";import"./normalize-wheel-es-ed76fb12.js";import"./@floating-ui-463e90e0.js";const G={class:"system-user-container layout-padding"},Q={class:"system-user-search mb15",style:{display:"none"}},j=D({name:"fopsTaskTimeOut"}),q=D({...j,setup(F){const u=T(),w=B(()=>N(()=>import("./editGroupDialog-4ed7ba57.js"),["assets/js/editGroupDialog-4ed7ba57.js","assets/js/index-5e41c337.js","assets/js/index-a79ea617.js","assets/js/@vue-d9027515.js","assets/js/pinia-be7e992d.js","assets/js/vue-demi-71ba0ef2.js","assets/js/vue-router-a2fac16f.js","assets/js/vue-i18n-e0b9a81d.js","assets/js/@intlify-9e8a497c.js","assets/js/source-map-7d7e1c08.js","assets/js/vue-b46b26c7.js","assets/js/js-cookie-cf83ad76.js","assets/js/@element-plus-96a31696.js","assets/js/nprogress-8d2808ea.js","assets/css/nprogress-8b89e2e0.css","assets/js/axios-707ed124.js","assets/js/qs-bbfcbdf3.js","assets/js/side-channel-eafc5c70.js","assets/js/get-intrinsic-b9397c9a.js","assets/js/has-symbols-e8f3ca0e.js","assets/js/function-bind-22e7ee79.js","assets/js/has-26d28e02.js","assets/js/call-bind-9ceb8f5b.js","assets/js/object-inspect-7b08a46e.js","assets/js/element-plus-7da03efa.js","assets/js/lodash-es-36eb724a.js","assets/js/@vueuse-2d9216d7.js","assets/js/@ctrl-f8748455.js","assets/js/dayjs-6aac7dfa.js","assets/js/async-validator-dee29e8b.js","assets/js/memoize-one-297ddbcb.js","assets/js/escape-html-1d60d822.js","assets/js/normalize-wheel-es-ed76fb12.js","assets/js/@floating-ui-463e90e0.js","assets/js/mitt-f7ef348c.js","assets/js/vue-grid-layout-0ee01be3.js","assets/css/index-c2520709.css"])),c=E(),a=A({keyWord:"",enable:-1,taskStatus:-1,clientId:"",tableData:{data:[],total:0,loading:!1,param:{pageNum:1,pageSize:10}}}),m=()=>{a.tableData.loading=!0;var i={};u.serverNodeList(i).then(function(t){t.Status?(a.tableData.data=t.Data.List,a.tableData.total=t.Data.RecordCount,setTimeout(()=>{a.tableData.loading=!1},500)):(a.tableData.data=[],setTimeout(()=>{a.tableData.loading=!1},500))})},g=i=>{},h=i=>{c.value.openDialog(type,i)},k=i=>{z.confirm(`此操作将永久删除：“${i.Name}”，是否继续?`,"提示",{confirmButtonText:"确认",cancelButtonText:"取消",type:"warning"}).then(()=>{u.taskDel({TaskGroupId:i.Id}).then(function(t){t.Status?(m(),v.success("删除成功")):v.error(t.StatusMessage)})}).catch(()=>{})};return R(()=>{m()}),(i,t)=>{const _=n("el-input"),l=n("el-option"),f=n("el-select"),x=n("ele-Search"),V=n("el-icon"),p=n("el-button"),r=n("el-table-column"),I=n("el-table"),C=n("el-card"),S=$("loading");return y(),P("div",G,[e(C,{shadow:"hover",class:"layout-padding-auto"},{default:s(()=>[U("div",Q,[e(_,{size:"default",modelValue:a.keyWord,"onUpdate:modelValue":t[0]||(t[0]=o=>a.keyWord=o),placeholder:"请输入任务名称",style:{"max-width":"180px"}},null,8,["modelValue"]),e(f,{modelValue:a.enable,"onUpdate:modelValue":t[1]||(t[1]=o=>a.enable=o),placeholder:"请选择运行状态",clearable:"",class:"ml10"},{default:s(()=>[e(l,{label:"全部",value:"-1"}),e(l,{label:"停止",value:"0"}),e(l,{label:"运行中",value:"1"})]),_:1},8,["modelValue"]),e(f,{modelValue:a.taskStatus,"onUpdate:modelValue":t[2]||(t[2]=o=>a.taskStatus=o),placeholder:"请选择调度状态",clearable:"",class:"ml10"},{default:s(()=>[e(l,{label:"全部",value:"-1"}),e(l,{label:"未开始",value:"0"}),e(l,{label:"调度中",value:"1"}),e(l,{label:"调度失败",value:"2"}),e(l,{label:"执行中",value:"3"}),e(l,{label:"失败",value:"4"}),e(l,{label:"成功",value:"5"})]),_:1},8,["modelValue"]),e(_,{size:"default",modelValue:a.clientId,"onUpdate:modelValue":t[3]||(t[3]=o=>a.clientId=o),placeholder:"请输入客户端ID",style:{"max-width":"180px"}},null,8,["modelValue"]),e(p,{size:"default",type:"primary",class:"ml10"},{default:s(()=>[e(V,null,{default:s(()=>[e(x)]),_:1}),d(" 查询 ")]),_:1})]),L((y(),M(I,{data:a.tableData.data,style:{width:"100%"}},{default:s(()=>[e(r,{prop:"Id",label:"序号",width:"60"}),e(r,{prop:"Name",label:"服务器节点名称","show-overflow-tooltip":""}),e(r,{prop:"Ip",label:"IP","show-overflow-tooltip":""}),e(r,{prop:"Port",label:"端口","show-overflow-tooltip":""}),e(r,{prop:"IsLeader",label:"是否主节点","show-overflow-tooltip":""}),e(r,{prop:"ActivateAt",label:"激活时间","show-overflow-tooltip":""}),e(r,{label:"操作",width:"100",style:{display:"none"}},{default:s(o=>[e(p,{size:"small",text:"",type:"primary",onClick:b=>g(o.row)},{default:s(()=>[d("详情信息")]),_:2},1032,["onClick"]),e(p,{size:"small",text:"",type:"primary",onClick:b=>h("edit",o.row)},{default:s(()=>[d("修改")]),_:2},1032,["onClick"]),e(p,{size:"small",text:"",type:"primary",onClick:b=>k(o.row)},{default:s(()=>[d("删除")]),_:2},1032,["onClick"])]),_:1})]),_:1},8,["data"])),[[S,a.tableData.loading]])]),_:1}),e(W(w),{ref_key:"editDialogRef",ref:c,onRefresh:t[4]||(t[4]=o=>m())},null,512)])}}});const Ne=O(q,[["__scopeId","data-v-20883b09"]]);export{Ne as default};
