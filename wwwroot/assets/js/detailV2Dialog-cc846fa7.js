import{f as N}from"./index-3be0d155.js";import{k as h,e as w,$ as V,ah as i,o as l,c as s,W as d,Q as p,P as f,U as R,V as a,a as n,F as u,a9 as y,L as m,u as D}from"./@vue-d9027515.js";import{s as _}from"./@intlify-9e8a497c.js";import"./index-1fb9eb7d.js";import"./pinia-be7e992d.js";import"./vue-demi-71ba0ef2.js";import"./vue-router-a2fac16f.js";import"./vue-i18n-e0b9a81d.js";import"./vue-b46b26c7.js";import"./js-cookie-cf83ad76.js";import"./@element-plus-96a31696.js";import"./nprogress-8d2808ea.js";import"./axios-707ed124.js";import"./qs-bbfcbdf3.js";import"./side-channel-eafc5c70.js";import"./get-intrinsic-b9397c9a.js";import"./has-symbols-e8f3ca0e.js";import"./function-bind-22e7ee79.js";import"./has-26d28e02.js";import"./call-bind-9ceb8f5b.js";import"./object-inspect-7b08a46e.js";import"./element-plus-7da03efa.js";import"./lodash-es-36eb724a.js";import"./@vueuse-2d9216d7.js";import"./@ctrl-f8748455.js";import"./dayjs-6aac7dfa.js";import"./async-validator-dee29e8b.js";import"./memoize-one-297ddbcb.js";import"./escape-html-1d60d822.js";import"./normalize-wheel-es-ed76fb12.js";import"./@floating-ui-463e90e0.js";import"./mitt-f7ef348c.js";import"./vue-grid-layout-0ee01be3.js";import"./source-map-7d7e1c08.js";import"./requestFSGet-d236e20b.js";const E={class:"system-user-dialog-container"},U={class:"custom-list mt20"},v=["title"],C={key:0},F={key:1},B=["title"],J={key:0},L={key:1},O=n("span",{style:{width:"100%"}},"暂无数据",-1),z=h({name:"fopsLinkDetailV2Dialog"}),bt=h({...z,emits:["refresh"],setup(P,{expose:T,emit:Q}){const A=N(),I=w(),t=V({ruleForm:{},tableData:[{Rgba:"",AppId:0,AppIp:"",AppName:"",StartTs:0,UseTs:0,Caption:"",Desc:"",UseDesc:"",Exception:""}],TraceId:0,totalTs:0,Rgba:"",AppId:0,AppIp:"",AppName:"",dialog:{isShowDialog:!1,type:"",title:"",submitTxt:""}});return T({openDialog:r=>{t.dialog.title="链路追踪详情(TraceId："+r.TraceId+")",t.TraceId=r.TraceId,A.linkTraceInfo(r.TraceId).then(function(o){console.log(_.friendlyJSONstringify(o)),o.Status&&(t.tableData=o.Data,o.Data.length>0&&(t.totalTs=o.Data[0].UseTs,t.Rgba=o.Data[0].Rgba,t.AppId=o.Data[0].AppId,t.AppIp=o.Data[0].AppIp,t.AppName=o.Data[0].AppName))}),t.dialog.isShowDialog=!0}}),(r,o)=>{const x=i("el-tag"),g=i("el-col"),S=i("el-row"),b=i("el-form"),k=i("el-dialog");return l(),s("div",E,[d(k,{title:t.dialog.title,modelValue:t.dialog.isShowDialog,"onUpdate:modelValue":o[0]||(o[0]=e=>t.dialog.isShowDialog=e),width:"70%"},{default:p(()=>[d(b,{ref_key:"gitDialogFormRef",ref:I,size:"default","label-width":"100px"},{default:p(()=>[d(S,{gutter:35},{default:p(()=>[t.tableData.length>0?(l(),f(g,{key:0,xs:24,sm:24,md:24,lg:24,xl:24,class:"mb20"},{default:p(()=>[d(x,null,{default:p(()=>[R("Rgba:"+a(t.Rgba)+"   AppId:"+a(t.AppId)+"   AppName:"+a(t.AppName)+"   AppIp:"+a(t.AppIp),1)]),_:1}),n("ul",U,[(l(!0),s(u,null,y(t.tableData,(e,c)=>(l(),s("li",{style:{float:"left"},key:c},[n("span",{style:m({width:e.StartTs/t.totalTs+"%"})},null,4),n("span",{style:m({"padding-left":e.StartTs/t.totalTs+"%"}),title:e.Desc},"("+a(e.StartTs)+")微妙，("+a(e.UseTs)+")微妙，"+a(e.Caption),13,v),e.Exception!=null?(l(),s("span",C,"异常："+a(D(_.friendlyJSONstringify)(e.Exception)),1)):(l(),s("span",F))]))),128)),(l(!0),s(u,null,y(t.tableData,(e,c)=>(l(),s("li",{style:{clear:"both"},key:c},[n("span",{style:m({"padding-left":e.StartTs/t.totalTs+"%"}),title:e.Desc},"("+a(e.StartTs)+")微妙，("+a(e.UseTs)+")微妙，"+a(e.Caption),13,B),e.Exception!=null?(l(),s("span",J,"异常："+a(D(_.friendlyJSONstringify)(e.Exception)),1)):(l(),s("span",L))]))),128))])]),_:1})):(l(),f(g,{key:1,style:{"text-align":"center"}},{default:p(()=>[O]),_:1}))]),_:1})]),_:1},512)]),_:1},8,["title","modelValue"])])}}});export{bt as default};
