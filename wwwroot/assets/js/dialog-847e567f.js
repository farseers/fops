import{f as J}from"./index-0358cb65.js";import{s as K}from"./@intlify-9e8a497c.js";import{a as g}from"./element-plus-7da03efa.js";import{k as v,e as I,$ as Q,i as W,ah as s,ar as j,o as _,c as q,W as l,Q as o,a as V,U as m,V as h,P as k,u as X,R as Y}from"./@vue-d9027515.js";import"./index-858e7551.js";import"./pinia-be7e992d.js";import"./vue-demi-71ba0ef2.js";import"./vue-router-a2fac16f.js";import"./vue-i18n-e0b9a81d.js";import"./vue-b46b26c7.js";import"./js-cookie-cf83ad76.js";import"./@element-plus-96a31696.js";import"./nprogress-8d2808ea.js";import"./axios-707ed124.js";import"./qs-bbfcbdf3.js";import"./side-channel-eafc5c70.js";import"./get-intrinsic-b9397c9a.js";import"./has-symbols-e8f3ca0e.js";import"./function-bind-22e7ee79.js";import"./has-26d28e02.js";import"./call-bind-9ceb8f5b.js";import"./object-inspect-7b08a46e.js";import"./mitt-f7ef348c.js";import"./vue-grid-layout-0ee01be3.js";import"./source-map-7d7e1c08.js";import"./lodash-es-36eb724a.js";import"./@vueuse-2d9216d7.js";import"./@ctrl-f8748455.js";import"./dayjs-6aac7dfa.js";import"./async-validator-dee29e8b.js";import"./memoize-one-297ddbcb.js";import"./escape-html-1d60d822.js";import"./normalize-wheel-es-ed76fb12.js";import"./@floating-ui-463e90e0.js";import"./requestFSGet-6e3c6110.js";const Z={class:"system-user-dialog-container"},ee={class:"dialog-footer"},te={class:"system-user-search mb15"},le=v({name:"fopsAppDialog"}),Be=v({...le,emits:["refresh"],setup(oe,{expose:N,emit:S}){const d=J(),C=I(),F=I(),e=Q({ruleForm:{AppName:"",AppId:"",DockerVer:"",ShellScript:"",ClusterVer:"",AppGit:0,AppGitName:"",FrameworkGits:[],Dockerfile:"",DockerfilePath:"",IsHealth:!1},gitList:[],SelectItem:[],gitDialogIsShow:!1,dialog:{isShowDialog:!1,type:"",title:"",submitTxt:""},tableData:{data:[],total:0,loading:!1,param:{pageNum:1,pageSize:10}},gitType:1}),z=(a,t)=>{e.dialog.type=a,a==="edit"?(e.ruleForm=t,e.dialog.title="修改应用",e.dialog.submitTxt="修 改",d.appsDetail({AppName:t.AppName}).then(function(r){r.Status&&(t=r.Data,e.ruleForm.AppName=t.AppName,e.ruleForm.AppId=t.AppId,e.ruleForm.DockerVer=t.DockerVer,e.ruleForm.ShellScript=t.ShellScript,e.ruleForm.ClusterVer=t.ClusterVer,e.ruleForm.AppGit=t.AppGit,e.ruleForm.FrameworkGits=t.FrameworkGits,e.ruleForm.Dockerfile=t.Dockerfile,e.ruleForm.DockerfilePath=t.DockerfilePath,e.SelectItem=t.FrameworkGits,e.ruleForm.IsHealth=t.IsHealth,w(t.AppGit),b(t.FrameworkGits))})):(e.dialog.title="新增应用",e.dialog.submitTxt="新 增",e.ruleForm.AppName="",e.ruleForm.AppId="",e.ruleForm.DockerVer="",e.ruleForm.ShellScript="",e.ruleForm.ClusterVer="",e.ruleForm.AppGit=0,e.ruleForm.FrameworkGits=[],e.ruleForm.Dockerfile="",e.ruleForm.DockerfilePath=""),e.dialog.isShowDialog=!0},b=a=>{e.gitList=[];for(let t=0;t<a.length;t++)console.log(a[t]),d.gitInfo({gitId:a[t]}).then(function(r){r.Status?e.gitList.push(r.Data):e.gitList=[]})},w=a=>{d.gitInfo({gitId:a}).then(function(t){t.Status?e.ruleForm.AppGitName=t.Data.Name:e.ruleForm.AppGitName=""})},x=a=>{e.ruleForm.FrameworkGits=e.ruleForm.FrameworkGits.filter(t=>t!==parseInt(a.Id)),b(e.ruleForm.FrameworkGits)},D=()=>{e.dialog.isShowDialog=!1},T=()=>{D()},P=()=>{var a={AppName:e.ruleForm.AppName,AppId:e.ruleForm.AppId,ShellScript:e.ruleForm.ShellScript,AppGit:parseInt(e.ruleForm.AppGit),FrameworkGits:e.ruleForm.FrameworkGits};if(e.dialog.type==="add"){var t=JSON.stringify(a);d.appsAdd(t).then(function(r){r.Status?(g.success("添加成功"),D(),S("refresh")):g.error(r.StatusMessage)})}else e.dialog.type=="edit"&&d.appsEdit(a).then(function(r){r.Status?(g.success("修改成功"),D(),S("refresh")):g.error(r.StatusMessage)})},U=()=>{d.gitList({}).then(function(a){a.Status?(e.tableData.data=a.Data,e.tableData.total=a.Data.length):e.tableData.data=[]})},H=a=>a.Id,L=a=>{if(console.log(a),a.length!=0){e.SelectItem=[];for(let r=0;r<a.length;r++){var t=a[r];e.SelectItem.push(t.Id)}console.log(e.SelectItem)}},y=a=>{U(),e.gitType=a,e.gitDialogIsShow=!0,a==1&&e.tableData.data.forEach(function(t,r){var p=e.ruleForm.FrameworkGits.filter(u=>u==t.Id);p.length>0?F.value.setCurrentRow(t):F.value.setCurrentRow(t,!1)})},R=()=>{e.gitType==1?(e.ruleForm.FrameworkGits=e.SelectItem,b(e.ruleForm.FrameworkGits)):(e.ruleForm.AppGit=e.SelectItem[0],w(e.ruleForm.AppGit)),e.gitDialogIsShow=!1};return W(()=>{}),N({openDialog:z,delGit:x}),(a,t)=>{const r=s("el-input"),p=s("el-form-item"),u=s("el-col"),c=s("el-tag"),f=s("el-button"),n=s("el-table-column"),G=s("el-table"),B=s("el-row"),E=s("el-form"),A=s("el-dialog"),$=s("el-icon"),M=s("el-card"),O=j("loading");return _(),q("div",Z,[l(A,{title:e.dialog.title,modelValue:e.dialog.isShowDialog,"onUpdate:modelValue":t[6]||(t[6]=i=>e.dialog.isShowDialog=i),width:"900px"},{footer:o(()=>[V("span",ee,[l(f,{onClick:T,size:"default"},{default:o(()=>[m("取 消")]),_:1}),l(f,{type:"primary",onClick:P,size:"default"},{default:o(()=>[m(h(e.dialog.submitTxt),1)]),_:1})])]),default:o(()=>[l(E,{ref_key:"gitDialogFormRef",ref:C,model:e.ruleForm,size:"default","label-width":"120px"},{default:o(()=>[l(B,{gutter:35},{default:o(()=>[l(u,{xs:24,sm:24,md:24,lg:24,xl:24,class:"mb20"},{default:o(()=>[l(p,{label:"应用名称"},{default:o(()=>[l(r,{modelValue:e.ruleForm.AppName,"onUpdate:modelValue":t[0]||(t[0]=i=>e.ruleForm.AppName=i),placeholder:"请输入应用名称",clearable:"",style:{"max-width":"200px","margin-right":"5px"}},null,8,["modelValue"]),m(" 名称需要与应用的AppName完全一致，才能检查健康状态 ")]),_:1})]),_:1}),l(u,{xs:24,sm:12,md:12,lg:12,xl:12,class:"mb20"},{default:o(()=>[l(p,{label:"是否健康"},{default:o(()=>[e.ruleForm.IsHealth?(_(),k(c,{key:0,size:"mini",type:"success"},{default:o(()=>[m("健康")]),_:1})):(_(),k(c,{key:1,size:"mini",type:"warning"},{default:o(()=>[m("不健康")]),_:1}))]),_:1})]),_:1}),l(u,{xs:24,sm:12,md:12,lg:12,xl:12,class:"mb20"},{default:o(()=>[l(p,{label:"镜像版本"},{default:o(()=>[l(c,{size:"mini"},{default:o(()=>[m(h(e.ruleForm.DockerVer),1)]),_:1})]),_:1})]),_:1}),l(u,{xs:24,sm:24,md:24,lg:24,xl:24,class:"mb20"},{default:o(()=>[l(p,{label:"集群版本"},{default:o(()=>[l(c,{size:"mini"},{default:o(()=>[m(h(X(K.friendlyJSONstringify)(e.ruleForm.ClusterVer)),1)]),_:1})]),_:1})]),_:1}),l(u,{xs:24,sm:18,md:18,lg:18,xl:18,class:"mb20"},{default:o(()=>[l(p,{label:"Git主仓库",style:{float:"left"}},{default:o(()=>[l(c,{size:"mini"},{default:o(()=>[m(h(e.ruleForm.AppGitName),1)]),_:1})]),_:1}),l(f,{type:"primary",onClick:t[1]||(t[1]=i=>y(2)),size:"default",style:{"margin-left":"5px"}},{default:o(()=>[m("添加Git")]),_:1})]),_:1}),l(u,{xs:24,sm:24,md:24,lg:24,xl:24,class:"mb20"},{default:o(()=>[l(p,{label:"依赖仓库",style:{float:"left"}},{default:o(()=>[l(f,{type:"success",onClick:t[2]||(t[2]=i=>y(1)),size:"default",style:{"margin-left":"5px"}},{default:o(()=>[m("添加依赖的仓库")]),_:1}),l(G,{data:e.gitList,style:{width:"100%"}},{default:o(()=>[l(n,{prop:"Id",label:"编号",width:"60"}),l(n,{prop:"Name",label:"Git名称","show-overflow-tooltip":"",width:"120"}),l(n,{prop:"Hub",label:"托管地址","show-overflow-tooltip":""}),l(n,{label:"操作",width:"100"},{default:o(i=>[l(f,{size:"small",text:"",type:"primary",onClick:ae=>x(i.row)},{default:o(()=>[m("删除")]),_:2},1032,["onClick"])]),_:1})]),_:1},8,["data"])]),_:1})]),_:1}),l(u,{xs:24,sm:24,md:24,lg:24,xl:24,class:"mb20"},{default:o(()=>[l(p,{label:"Shell脚本"},{default:o(()=>[l(r,{modelValue:e.ruleForm.ShellScript,"onUpdate:modelValue":t[3]||(t[3]=i=>e.ruleForm.ShellScript=i),type:"textarea",placeholder:"请输入Shell脚本",clearable:""},null,8,["modelValue"])]),_:1})]),_:1}),l(u,{xs:24,sm:24,md:24,lg:24,xl:24,class:"mb20"},{default:o(()=>[l(p,{label:"Dockerfile内容"},{default:o(()=>[l(r,{modelValue:e.ruleForm.Dockerfile,"onUpdate:modelValue":t[4]||(t[4]=i=>e.ruleForm.Dockerfile=i),type:"textarea",placeholder:"请输入Dockerfile内容",clearable:""},null,8,["modelValue"])]),_:1})]),_:1}),l(u,{xs:24,sm:24,md:24,lg:24,xl:24,class:"mb20"},{default:o(()=>[l(p,{label:"Dockerfile"},{default:o(()=>[l(r,{modelValue:e.ruleForm.DockerfilePath,"onUpdate:modelValue":t[5]||(t[5]=i=>e.ruleForm.DockerfilePath=i),placeholder:"请输入Dockerfile路径",clearable:""},null,8,["modelValue"])]),_:1})]),_:1})]),_:1})]),_:1},8,["model"])]),_:1},8,["title","modelValue"]),l(A,{title:"Git列表",modelValue:e.gitDialogIsShow,"onUpdate:modelValue":t[8]||(t[8]=i=>e.gitDialogIsShow=i),width:"700px;",height:"300px;"},{default:o(()=>[l(M,{shadow:"hover",class:"layout-padding-auto"},{default:o(()=>[V("div",te,[l(f,{size:"default",type:"success",class:"ml10",onClick:t[7]||(t[7]=i=>R())},{default:o(()=>[l($),m(" 确认选择 ")]),_:1})]),Y((_(),k(G,{ref_key:"multipleTable",ref:F,data:e.tableData.data,style:{width:"100%"},"row-key":H,onSelectionChange:L},{default:o(()=>[l(n,{type:"selection","reserve-selection":!0,width:"55"}),l(n,{prop:"Id",label:"编号",width:"60"}),l(n,{prop:"Name",label:"Git名称","show-overflow-tooltip":""}),l(n,{prop:"Hub",label:"托管地址","show-overflow-tooltip":""}),l(n,{prop:"Branch",label:"Git分支","show-overflow-tooltip":""})]),_:1},8,["data"])),[[O,e.tableData.loading]])]),_:1})]),_:1},8,["modelValue"])])}}});export{Be as default};