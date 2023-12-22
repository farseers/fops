import{f as W}from"./index-824e04e8.js";import{k as v,e as A,$ as B,i as P,ah as i,ar as R,o as p,P as s,Q as o,a as g,W as t,V as f,U as n,R as E,T as u,bb as H,b9 as Q}from"./@vue-d9027515.js";import{_ as j}from"./_plugin-vue_export-helper-c27b6911.js";import"./index-46c45a09.js";import"./pinia-be7e992d.js";import"./vue-demi-71ba0ef2.js";import"./vue-router-a2fac16f.js";import"./vue-i18n-e0b9a81d.js";import"./@intlify-9e8a497c.js";import"./source-map-7d7e1c08.js";import"./vue-b46b26c7.js";import"./js-cookie-cf83ad76.js";import"./@element-plus-96a31696.js";import"./nprogress-8d2808ea.js";import"./axios-707ed124.js";import"./qs-bbfcbdf3.js";import"./side-channel-eafc5c70.js";import"./get-intrinsic-b9397c9a.js";import"./has-symbols-e8f3ca0e.js";import"./function-bind-22e7ee79.js";import"./has-26d28e02.js";import"./call-bind-9ceb8f5b.js";import"./object-inspect-7b08a46e.js";import"./element-plus-7da03efa.js";import"./lodash-es-36eb724a.js";import"./@vueuse-2d9216d7.js";import"./@ctrl-f8748455.js";import"./dayjs-6aac7dfa.js";import"./async-validator-dee29e8b.js";import"./memoize-one-297ddbcb.js";import"./escape-html-1d60d822.js";import"./normalize-wheel-es-ed76fb12.js";import"./@floating-ui-463e90e0.js";import"./mitt-f7ef348c.js";import"./vue-grid-layout-0ee01be3.js";const M=b=>(H("data-v-7b83bee1"),b=b(),Q(),b),$={class:"system-user-container layout-padding",style:{width:"100%"}},q={class:"system-user-search mb15"},F=M(()=>g("br",null,null,-1)),J={style:{color:"blue"}},K=v({name:"fopsTask"}),O=v({...K,setup(b,{expose:D}){const w=W();A();const e=B({keyWord:"",taskGroupId:"",logLevel:-1,tableData:{data:[],total:0,loading:!1,param:{pageNum:1,pageSize:10}},dialog:{isShowDialog:!1,type:"",title:"",submitTxt:""}}),c=()=>{e.tableData.loading=!0;const l=new URLSearchParams;l.append("logLevel",e.logLevel.toString()),l.append("taskGroupId",e.taskGroupId.toString()),l.append("pageSize",e.tableData.param.pageSize.toString()),l.append("pageIndex",e.tableData.param.pageNum.toString()),w.taskLogList(l.toString()).then(function(r){r.Status?(e.tableData.data=r.Data.List,e.tableData.total=r.Data.RecordCount,setTimeout(()=>{e.tableData.loading=!1},500)):(e.tableData.data=[],setTimeout(()=>{e.tableData.loading=!1},500))})},h=l=>{e.taskGroupId=l.Id,e.dialog.isShowDialog=!0,e.dialog.title=l.Caption,c()},S=()=>{e.dialog.isShowDialog=!1},L=l=>{e.tableData.param.pageSize=l,c()},y=l=>{e.tableData.param.pageNum=l,c()},k=()=>{c()};return P(()=>{c()}),D({openDialog:h,closeDialog:S}),(l,r)=>{const d=i("el-option"),C=i("el-select"),I=i("ele-Search"),z=i("el-icon"),V=i("el-button"),_=i("el-table-column"),m=i("el-tag"),N=i("el-table"),T=i("el-pagination"),x=i("el-card"),U=i("el-dialog"),G=R("loading");return p(),s(U,{modelValue:e.dialog.isShowDialog,"onUpdate:modelValue":r[3]||(r[3]=a=>e.dialog.isShowDialog=a),style:{height:"80%",width:"70%"}},{default:o(()=>[g("div",$,[t(x,{shadow:"hover",class:"layout-padding-auto"},{default:o(()=>[g("div",q,[g("span",null,"任务组名称："+f(e.dialog.title),1),t(C,{modelValue:e.logLevel,"onUpdate:modelValue":r[0]||(r[0]=a=>e.logLevel=a),placeholder:"请选择日志等级",clearable:"",class:"ml10"},{default:o(()=>[t(d,{label:"全部",value:-1}),t(d,{label:"Trace",value:0}),t(d,{label:"Debug",value:1}),t(d,{label:"Info",value:2}),t(d,{label:"Warning",value:3}),t(d,{label:"Error",value:4}),t(d,{label:"Critical",value:5})]),_:1},8,["modelValue"]),t(V,{size:"default",type:"primary",class:"ml10",onClick:k},{default:o(()=>[t(z,null,{default:o(()=>[t(I)]),_:1}),n(" 查询 ")]),_:1})]),E((p(),s(N,{data:e.tableData.data,style:{width:"100%"}},{default:o(()=>[t(_,{prop:"Id",label:"序号",width:"100"}),t(_,{label:"名称",width:"250"},{default:o(a=>[g("span",null,f(a.row.Caption),1),F,g("span",null,[n(f(a.row.Name)+"（",1),g("span",J,"Ver:"+f(a.row.Ver),1),n("）")])]),_:1}),t(_,{prop:"TaskId",width:"220",label:"任务ID","show-overflow-tooltip":""}),t(_,{prop:"LogLevel",width:"100",label:"日志级别","show-overflow-tooltip":""},{default:o(a=>[a.row.LogLevel==0?(p(),s(m,{key:0},{default:o(()=>[n("Trace")]),_:1})):u("",!0),a.row.LogLevel==1?(p(),s(m,{key:1},{default:o(()=>[n("Debug")]),_:1})):u("",!0),a.row.LogLevel==2?(p(),s(m,{key:2},{default:o(()=>[n("Information")]),_:1})):u("",!0),a.row.LogLevel==3?(p(),s(m,{key:3},{default:o(()=>[n("Warning")]),_:1})):u("",!0),a.row.LogLevel==4?(p(),s(m,{key:4},{default:o(()=>[n("Error")]),_:1})):u("",!0),a.row.LogLevel==5?(p(),s(m,{key:5},{default:o(()=>[n("Critical")]),_:1})):u("",!0),a.row.LogLevel==6?(p(),s(m,{key:6},{default:o(()=>[n("NoneLevel")]),_:1})):u("",!0)]),_:1}),t(_,{prop:"Content",label:"日志内容"}),t(_,{prop:"CreateAt",width:"200",label:"日志时间","show-overflow-tooltip":""})]),_:1},8,["data"])),[[G,e.tableData.loading]]),t(T,{onSizeChange:L,onCurrentChange:y,class:"mt15","pager-count":5,"page-sizes":[10,20,30],"current-page":e.tableData.param.pageNum,"onUpdate:currentPage":r[1]||(r[1]=a=>e.tableData.param.pageNum=a),background:"","page-size":e.tableData.param.pageSize,"onUpdate:pageSize":r[2]||(r[2]=a=>e.tableData.param.pageSize=a),layout:"total, sizes, prev, pager, next, jumper",total:e.tableData.total},null,8,["current-page","page-size","total"])]),_:1})])]),_:1},8,["modelValue"])}}});const xe=j(O,[["__scopeId","data-v-7b83bee1"]]);export{xe as default};
