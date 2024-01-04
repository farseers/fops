import{f as P}from"./index-59967c64.js";import{a as d}from"./element-plus-7da03efa.js";import{k as _,e as B,$ as S,ah as a,o as w,c as y,W as o,Q as r,a as H,U as F,V as A}from"./@vue-d9027515.js";import"./index-dd22bf27.js";import"./pinia-be7e992d.js";import"./vue-demi-71ba0ef2.js";import"./vue-router-a2fac16f.js";import"./vue-i18n-e0b9a81d.js";import"./@intlify-9e8a497c.js";import"./source-map-7d7e1c08.js";import"./vue-b46b26c7.js";import"./js-cookie-cf83ad76.js";import"./@element-plus-96a31696.js";import"./nprogress-8d2808ea.js";import"./axios-707ed124.js";import"./qs-bbfcbdf3.js";import"./side-channel-eafc5c70.js";import"./get-intrinsic-b9397c9a.js";import"./has-symbols-e8f3ca0e.js";import"./function-bind-22e7ee79.js";import"./has-26d28e02.js";import"./call-bind-9ceb8f5b.js";import"./object-inspect-7b08a46e.js";import"./mitt-f7ef348c.js";import"./vue-grid-layout-0ee01be3.js";import"./lodash-es-36eb724a.js";import"./@vueuse-2d9216d7.js";import"./@ctrl-f8748455.js";import"./dayjs-6aac7dfa.js";import"./async-validator-dee29e8b.js";import"./memoize-one-297ddbcb.js";import"./escape-html-1d60d822.js";import"./normalize-wheel-es-ed76fb12.js";import"./@floating-ui-463e90e0.js";import"./requestFSGet-14758c50.js";const G={class:"system-user-dialog-container"},k={class:"dialog-footer"},C=_({name:"fopsGitDialog"}),_e=_({...C,emits:["refresh"],setup(I,{expose:b,emit:p}){const f=P(),V=B(),e=S({ruleForm:{Id:0,Name:"",Hub:"",Branch:"main",UserName:"",UserPwd:"",Dir:"",PullAt:""},dialog:{isShowDialog:!1,type:"",title:"",submitTxt:""}}),U=(m,l)=>{e.dialog.type=m,m==="edit"?(e.ruleForm=l,e.dialog.title="修改Git",e.dialog.submitTxt="修 改",e.ruleForm.Id=l.Id,e.ruleForm.Name=l.Name,e.ruleForm.Hub=l.Hub,e.ruleForm.Branch=l.Branch,e.ruleForm.UserName=l.UserName,e.ruleForm.UserPwd=l.UserPwd,e.ruleForm.Dir=l.Dir,e.ruleForm.PullAt=l.PullAt):(e.dialog.title="新增Git",e.dialog.submitTxt="新 增",e.ruleForm.Id=0,e.ruleForm.Name="",e.ruleForm.Hub="",e.ruleForm.Branch="main",e.ruleForm.UserName="",e.ruleForm.UserPwd="",e.ruleForm.Dir="",e.ruleForm.PullAt=""),e.dialog.isShowDialog=!0},n=()=>{e.dialog.isShowDialog=!1},x=()=>{n()},c=()=>{var m={Id:e.ruleForm.Id,Name:e.ruleForm.Name,Hub:e.ruleForm.Hub,Branch:e.ruleForm.Branch,UserName:e.ruleForm.UserName,UserPwd:e.ruleForm.UserPwd,Dir:e.ruleForm.Dir};e.dialog.type==="add"?f.gitAdd(m).then(async function(l){l.Status?(d.success("添加成功"),n(),p("refresh")):d.error(l.StatusMessage)}):e.dialog.type=="edit"&&f.gitEdit(m).then(async function(l){l.Status?(d.success("修改成功"),n(),p("refresh")):d.error(l.StatusMessage)})};return b({openDialog:U,onSubmit:c}),(m,l)=>{const i=a("el-input"),s=a("el-form-item"),u=a("el-col"),D=a("el-row"),N=a("el-form"),g=a("el-button"),h=a("el-dialog");return w(),y("div",G,[o(h,{title:e.dialog.title,modelValue:e.dialog.isShowDialog,"onUpdate:modelValue":l[6]||(l[6]=t=>e.dialog.isShowDialog=t),width:"769px"},{footer:r(()=>[H("span",k,[o(g,{onClick:x,size:"default"},{default:r(()=>[F("取 消")]),_:1}),o(g,{type:"primary",onClick:c,size:"default"},{default:r(()=>[F(A(e.dialog.submitTxt),1)]),_:1})])]),default:r(()=>[o(N,{ref_key:"gitDialogFormRef",ref:V,model:e.ruleForm,size:"default","label-width":"90px"},{default:r(()=>[o(D,{gutter:35},{default:r(()=>[o(u,{xs:24,sm:24,md:24,lg:24,xl:24,class:"mb20"},{default:r(()=>[o(s,{label:"Git名称"},{default:r(()=>[o(i,{modelValue:e.ruleForm.Name,"onUpdate:modelValue":l[0]||(l[0]=t=>e.ruleForm.Name=t),placeholder:"请输入Git名称",clearable:""},null,8,["modelValue"])]),_:1})]),_:1}),o(u,{xs:24,sm:24,md:24,lg:24,xl:24,class:"mb20"},{default:r(()=>[o(s,{label:"托管地址"},{default:r(()=>[o(i,{modelValue:e.ruleForm.Hub,"onUpdate:modelValue":l[1]||(l[1]=t=>e.ruleForm.Hub=t),placeholder:"请输入托管地址",clearable:""},null,8,["modelValue"])]),_:1})]),_:1}),o(u,{xs:24,sm:24,md:24,lg:24,xl:24,class:"mb20"},{default:r(()=>[o(s,{label:"Git分支"},{default:r(()=>[o(i,{modelValue:e.ruleForm.Branch,"onUpdate:modelValue":l[2]||(l[2]=t=>e.ruleForm.Branch=t),placeholder:"请输入Git分支",clearable:""},null,8,["modelValue"])]),_:1})]),_:1}),o(u,{xs:24,sm:12,md:12,lg:12,xl:12,class:"mb20"},{default:r(()=>[o(s,{label:"账户名称"},{default:r(()=>[o(i,{modelValue:e.ruleForm.UserName,"onUpdate:modelValue":l[3]||(l[3]=t=>e.ruleForm.UserName=t),placeholder:"请输入账户名称",clearable:""},null,8,["modelValue"])]),_:1})]),_:1}),o(u,{xs:24,sm:12,md:12,lg:12,xl:12,class:"mb20"},{default:r(()=>[o(s,{label:"账户密码"},{default:r(()=>[o(i,{modelValue:e.ruleForm.UserPwd,"onUpdate:modelValue":l[4]||(l[4]=t=>e.ruleForm.UserPwd=t),placeholder:"请输入",clearable:""},null,8,["modelValue"])]),_:1})]),_:1}),o(u,{xs:24,sm:24,md:24,lg:24,xl:24,class:"mb20"},{default:r(()=>[o(s,{label:"存储目录"},{default:r(()=>[o(i,{modelValue:e.ruleForm.Dir,"onUpdate:modelValue":l[5]||(l[5]=t=>e.ruleForm.Dir=t),placeholder:"请输入存储目录",clearable:""},null,8,["modelValue"])]),_:1})]),_:1})]),_:1})]),_:1},8,["model"])]),_:1},8,["title","modelValue"])])}}});export{_e as default};