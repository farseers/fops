import{_ as at}from"./index-1f38180f.js";import{f as et}from"./index-d1a984d2.js";import{a as V}from"./element-plus-7da03efa.js";import{k as T,aN as ot,e as lt,$ as nt,f as st,i as it,_ as pt,ah as i,ar as rt,o as p,c as b,W as n,Q as s,a as d,F as L,a9 as G,U as g,P as c,R as dt,T as v,u as ct,V as m,bb as ut,b9 as gt}from"./@vue-d9027515.js";import{_ as mt}from"./_plugin-vue_export-helper-c27b6911.js";import"./pinia-be7e992d.js";import"./vue-demi-71ba0ef2.js";import"./vue-router-a2fac16f.js";import"./vue-i18n-e0b9a81d.js";import"./@intlify-9e8a497c.js";import"./source-map-7d7e1c08.js";import"./vue-b46b26c7.js";import"./js-cookie-cf83ad76.js";import"./@element-plus-96a31696.js";import"./nprogress-8d2808ea.js";import"./axios-707ed124.js";import"./qs-bbfcbdf3.js";import"./side-channel-eafc5c70.js";import"./get-intrinsic-b9397c9a.js";import"./has-symbols-e8f3ca0e.js";import"./function-bind-22e7ee79.js";import"./has-26d28e02.js";import"./call-bind-9ceb8f5b.js";import"./object-inspect-7b08a46e.js";import"./mitt-f7ef348c.js";import"./vue-grid-layout-0ee01be3.js";import"./lodash-es-36eb724a.js";import"./@vueuse-2d9216d7.js";import"./@ctrl-f8748455.js";import"./dayjs-6aac7dfa.js";import"./async-validator-dee29e8b.js";import"./memoize-one-297ddbcb.js";import"./escape-html-1d60d822.js";import"./normalize-wheel-es-ed76fb12.js";import"./@floating-ui-463e90e0.js";import"./requestFSGet-6eec0a90.js";const _t=D=>(ut("data-v-0421a01c"),D=D(),gt(),D),ft={class:"system-user-container layout-padding"},bt={class:"system-user-search mb15"},Dt={class:"flex-warp",style:{background:"#e0e0e0"}},ht=["onClick"],It={class:"appItem"},vt={class:"appItem"},yt={class:"appItem"},Lt={class:"appItem"},wt={class:"appItem"},St={class:"appItem"},kt={class:"appItem"},Ct=_t(()=>d("h3",{style:{padding:"5px"}},"构建队列",-1)),Nt=T({name:"fopsApp"}),xt=T({...Nt,setup(D){const u=et(),F=ot(()=>at(()=>import("./dialog-ae86d902.js"),["assets/js/dialog-ae86d902.js","assets/js/index-d1a984d2.js","assets/js/index-1f38180f.js","assets/js/@vue-d9027515.js","assets/js/pinia-be7e992d.js","assets/js/vue-demi-71ba0ef2.js","assets/js/vue-router-a2fac16f.js","assets/js/vue-i18n-e0b9a81d.js","assets/js/@intlify-9e8a497c.js","assets/js/source-map-7d7e1c08.js","assets/js/vue-b46b26c7.js","assets/js/js-cookie-cf83ad76.js","assets/js/@element-plus-96a31696.js","assets/js/nprogress-8d2808ea.js","assets/css/nprogress-8b89e2e0.css","assets/js/axios-707ed124.js","assets/js/qs-bbfcbdf3.js","assets/js/side-channel-eafc5c70.js","assets/js/get-intrinsic-b9397c9a.js","assets/js/has-symbols-e8f3ca0e.js","assets/js/function-bind-22e7ee79.js","assets/js/has-26d28e02.js","assets/js/call-bind-9ceb8f5b.js","assets/js/object-inspect-7b08a46e.js","assets/js/element-plus-7da03efa.js","assets/js/lodash-es-36eb724a.js","assets/js/@vueuse-2d9216d7.js","assets/js/@ctrl-f8748455.js","assets/js/dayjs-6aac7dfa.js","assets/js/async-validator-dee29e8b.js","assets/js/memoize-one-297ddbcb.js","assets/js/escape-html-1d60d822.js","assets/js/normalize-wheel-es-ed76fb12.js","assets/js/@floating-ui-463e90e0.js","assets/js/mitt-f7ef348c.js","assets/js/vue-grid-layout-0ee01be3.js","assets/css/index-c2520709.css","assets/js/requestFSGet-6eec0a90.js","assets/css/detailDialog-2d838ab0.css"])),y=lt(),t=nt({logDialogIsShow:!1,logContent:"",tableData:{data:[],total:0,loading:!1,param:{pageNum:1,pageSize:10}},tableLogData:{data:[],total:0,loading:!1,param:{pageNum:1,pageSize:10}},appName:"",logId:0,clusterId:0,clusterData:[]}),w=()=>{t.tableData.loading=!0;const a=[];u.appsList({}).then(function(e){if(e.Status){for(let r=0;r<e.Data.length;r++){var l=e.Data[r];l.FrameworkGitsStr=Q(l.FrameworkGits),l.AppGitStr=W(l.AppGit),a.push(l)}setTimeout(()=>{t.tableData.data=a,t.tableData.total=a.length,t.tableData.loading=!1},500)}else t.tableData.data=[],setTimeout(()=>{t.tableData.loading=!1},500)})},_=()=>{t.tableLogData.loading=!0;const a={appName:"",pageIndex:t.tableLogData.param.pageNum,pageSize:t.tableLogData.param.pageSize};u.buildList(a).then(function(e){e.Status?(t.tableLogData.data=e.Data.List,t.tableLogData.total=e.Data.RecordCount,t.tableLogData.loading=!1):(t.tableLogData.data=[],t.tableLogData.loading=!1)})},$=()=>{t.tableData.loading=!0,u.clusterList({}).then(function(a){if(a.Status){var e=[];for(let r=0;r<a.Data.length;r++){var l=a.Data[r];r==0&&(t.clusterId=l.Id),l.Name=l.Name+" - "+l.DockerName,e.push(l)}t.clusterData=e}else t.tableData.data=[]})},B=a=>{t.clusterId=a},R=a=>{y.value.openDialog(a)},U=(a,e)=>{y.value.openDialog(a,e)},E=a=>{t.tableLogData.param.pageSize=a,_()},P=a=>{t.appName=a.AppName,t.tableLogData.param.pageNum=1,t.tableLogData.param.pageSize=10,_()},M=a=>{t.tableLogData.param.pageNum=a,_()};let S=null;st(()=>t.logDialogIsShow,(a,e)=>{a?S=setInterval(H,1e3):clearInterval(S)});const O=a=>{t.logId=a.Id,u.buildLog({buildId:t.logId}).then(function(e){t.logContent=e,t.logDialogIsShow=!0})},H=()=>{u.buildLog({buildId:t.logId}).then(function(a){console.log(a),t.logContent=a})},j=a=>{var e={AppName:a.AppName,ClusterId:t.clusterId};u.buildAdd(e).then(async function(l){l.Status?(V.success("添加成功"),_()):V.error(l.StatusMessage)})},Q=a=>{var e=[];for(let l=0;l<a.length;l++)u.gitInfo({gitId:a[l]}).then(function(r){r.Status&&e.push(r.Data.Name)});return e},W=a=>{var e=[];return u.gitInfo({gitId:a}).then(function(l){l.Status&&e.push(l.Data.Name)}),e};let k=null;return it(()=>{w(),_(),$(),k=setInterval(_,1e3)}),pt(()=>{clearInterval(k)}),(a,e)=>{const l=i("el-option"),r=i("el-select"),q=i("ele-FolderAdd"),J=i("el-icon"),h=i("el-button"),I=i("el-tag"),C=i("el-col"),N=i("el-empty"),f=i("el-table-column"),K=i("el-table"),X=i("el-pagination"),Y=i("el-row"),x=i("el-card"),Z=i("el-dialog"),tt=rt("loading");return p(),b("div",ft,[n(x,{shadow:"hover",class:"layout-padding-auto"},{default:s(()=>[d("div",bt,[n(r,{modelValue:t.clusterId,"onUpdate:modelValue":e[0]||(e[0]=o=>t.clusterId=o),placeholder:"请选择集群",class:"ml10",onChange:B},{default:s(()=>[(p(!0),b(L,null,G(t.clusterData,o=>(p(),c(l,{key:o.Id,label:o.Name,value:o.Id},null,8,["label","value"]))),128))]),_:1},8,["modelValue"]),n(h,{size:"default",type:"success",class:"ml10",onClick:e[1]||(e[1]=o=>R("add"))},{default:s(()=>[n(J,null,{default:s(()=>[n(q)]),_:1}),g(" 新增应用 ")]),_:1})]),d("div",Dt,[n(Y,null,{default:s(()=>[t.tableData.data.length>0?(p(),c(C,{key:0,style:{float:"left",background:"#ffffff",margin:"10px"},xs:16,sm:16,md:16,lg:16,xl:16,class:"mb15"},{default:s(()=>[(p(!0),b(L,null,G(t.tableData.data,(o,A)=>(p(),b("div",{class:"flex-warp-item",key:A},[d("div",{class:"flex-warp-item-box",onClick:z=>P(o)},[d("div",It,"应用ID："+m(o.AppId),1),d("div",vt,[n(I,null,{default:s(()=>[g(m(o.AppName),1)]),_:2},1024)]),d("div",yt,"容器版本："+m(o.DockerVer),1),d("div",Lt,"集群版本："+m(o.ClusterVer),1),d("div",wt,"AppGit："+m(o.AppGitName)+" ",1),d("div",St,"容器文件路径："+m(o.DockerfilePath),1),d("div",kt,[n(h,{size:"default",onClick:z=>U("edit",o),type:"warning"},{default:s(()=>[g("修改")]),_:2},1032,["onClick"]),n(h,{onClick:z=>j(o),size:"default",type:"danger"},{default:s(()=>[g("构建")]),_:2},1032,["onClick"])])],8,ht)]))),128))]),_:1})):(p(),c(N,{key:1,description:"暂无数据"})),t.tableLogData.data.length>0?(p(),c(C,{key:2,style:{float:"left",background:"#ffffff",margin:"10px","margin-left":"0",padding:"5px"},xs:8,sm:8,md:8,lg:8,xl:8,class:"mb15"},{default:s(()=>[Ct,t.tableLogData.data.length>0?(p(),b(L,{key:0},[dt((p(),c(K,{data:t.tableLogData.data,style:{"min-width":"576px"}},{default:s(()=>[n(f,{prop:"Id",label:"编号",width:"60"}),n(f,{prop:"AppName",label:"应用名称"}),n(f,{label:"状态",width:"100","show-overflow-tooltip":""},{default:s(o=>[o.row.Status==0?(p(),c(I,{key:0,style:{color:"#7a7a7a"}},{default:s(()=>[g("未开始")]),_:1})):v("",!0),o.row.Status==1?(p(),c(I,{key:1},{default:s(()=>[g("构建中")]),_:1})):v("",!0),o.row.Status==2?(p(),c(I,{key:2,style:{color:"green"}},{default:s(()=>[g("完成")]),_:1})):v("",!0)]),_:1}),n(f,{prop:"FinishAt",width:"170",label:"完成时间"}),n(f,{label:"操作",width:"100"},{default:s(o=>[n(h,{size:"small",text:"",type:"primary",onClick:A=>O(o.row)},{default:s(()=>[g("构建日志")]),_:2},1032,["onClick"])]),_:1})]),_:1},8,["data"])),[[tt,t.tableLogData.loading]]),n(X,{onSizeChange:E,onCurrentChange:M,class:"mt15","pager-count":5,"page-sizes":[10,20,30],"current-page":t.tableLogData.param.pageNum,"onUpdate:currentPage":e[2]||(e[2]=o=>t.tableLogData.param.pageNum=o),background:"","page-size":t.tableLogData.param.pageSize,"onUpdate:pageSize":e[3]||(e[3]=o=>t.tableLogData.param.pageSize=o),layout:"total, sizes, prev, pager, next, jumper",total:t.tableLogData.total},null,8,["current-page","page-size","total"])],64)):v("",!0)]),_:1})):(p(),c(N,{key:3,description:"暂无数据"}))]),_:1})])]),_:1}),n(ct(F),{ref_key:"appDialogRef",ref:y,onRefresh:e[4]||(e[4]=o=>w())},null,512),n(Z,{title:"构建日志",modelValue:t.logDialogIsShow,"onUpdate:modelValue":e[5]||(e[5]=o=>t.logDialogIsShow=o),width:"700px;",height:"300px;"},{default:s(()=>[n(x,{shadow:"hover",class:"layout-padding-auto"},{default:s(()=>[d("div",null,m(t.logContent),1)]),_:1})]),_:1},8,["modelValue"])])}}});const ua=mt(xt,[["__scopeId","data-v-0421a01c"]]);export{ua as default};
