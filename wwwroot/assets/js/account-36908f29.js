import{u as R,b as U}from"./vue-router-a2fac16f.js";import{u as $}from"./vue-i18n-e0b9a81d.js";import{a as q}from"./js-cookie-cf83ad76.js";import{s as B}from"./pinia-be7e992d.js";import{u as D,S as y,k as E,l as j,N as z}from"./index-a79ea617.js";import{a as A}from"./formatTime-29ac8c52.js";import{a as v}from"./element-plus-7da03efa.js";import{k as V,$ as M,j as O,ah as a,ar as J,o as _,P as f,Q as e,W as o,a as N,K,R as P,U as L,V as Q}from"./@vue-d9027515.js";import{_ as W}from"./_plugin-vue_export-helper-c27b6911.js";import"./@intlify-9e8a497c.js";import"./source-map-7d7e1c08.js";import"./vue-b46b26c7.js";import"./vue-demi-71ba0ef2.js";import"./@element-plus-96a31696.js";import"./nprogress-8d2808ea.js";import"./axios-707ed124.js";import"./qs-bbfcbdf3.js";import"./side-channel-eafc5c70.js";import"./get-intrinsic-b9397c9a.js";import"./has-symbols-e8f3ca0e.js";import"./function-bind-22e7ee79.js";import"./has-26d28e02.js";import"./call-bind-9ceb8f5b.js";import"./object-inspect-7b08a46e.js";import"./mitt-f7ef348c.js";import"./vue-grid-layout-0ee01be3.js";import"./lodash-es-36eb724a.js";import"./@vueuse-2d9216d7.js";import"./@ctrl-f8748455.js";import"./dayjs-6aac7dfa.js";import"./async-validator-dee29e8b.js";import"./memoize-one-297ddbcb.js";import"./escape-html-1d60d822.js";import"./normalize-wheel-es-ed76fb12.js";import"./@floating-ui-463e90e0.js";const G=V({name:"loginAccount"}),H=V({...G,setup(X){const{t:S}=$(),k=D(),{themeConfig:C}=B(k),c=R(),g=U(),t=M({isShowPassword:!1,ruleForm:{userName:"admin",password:"123456",code:"1234"},loading:{signIn:!1}}),I=O(()=>A(new Date)),F=async()=>{if(t.loading.signIn=!0,y.set("token",Math.random().toString(36).substr(0)),q.set("userName",t.ruleForm.userName),C.value.isRequestRoutes){const n=await j();w(n)}else{const n=await E();w(n)}},w=n=>{var s,u,r,l;if(n)v.warning("抱歉，您没有登录权限"),y.clear();else{let i=I.value;(s=c.query)!=null&&s.redirect?g.push({path:(u=c.query)==null?void 0:u.redirect,query:Object.keys((r=c.query)==null?void 0:r.params).length>0?JSON.parse((l=c.query)==null?void 0:l.params):""}):g.push("/");const p=S("message.signInText");v.success(`${i}，${p}`),z.start()}t.loading.signIn=!1};return(n,s)=>{const u=a("ele-User"),r=a("el-icon"),l=a("el-input"),i=a("el-form-item"),p=a("ele-Unlock"),T=a("ele-Position"),d=a("el-col"),h=a("el-button"),b=a("el-form"),x=J("waves");return _(),f(b,{size:"large",class:"login-content-form"},{default:e(()=>[o(i,{class:"login-animation1"},{default:e(()=>[o(l,{text:"",placeholder:n.$t("message.account.accountPlaceholder1"),modelValue:t.ruleForm.userName,"onUpdate:modelValue":s[0]||(s[0]=m=>t.ruleForm.userName=m),clearable:"",autocomplete:"off"},{prefix:e(()=>[o(r,{class:"el-input__icon"},{default:e(()=>[o(u)]),_:1})]),_:1},8,["placeholder","modelValue"])]),_:1}),o(i,{class:"login-animation2"},{default:e(()=>[o(l,{type:t.isShowPassword?"text":"password",placeholder:n.$t("message.account.accountPlaceholder2"),modelValue:t.ruleForm.password,"onUpdate:modelValue":s[2]||(s[2]=m=>t.ruleForm.password=m),autocomplete:"off"},{prefix:e(()=>[o(r,{class:"el-input__icon"},{default:e(()=>[o(p)]),_:1})]),suffix:e(()=>[N("i",{class:K(["iconfont el-input__icon login-content-password",t.isShowPassword?"icon-yincangmima":"icon-xianshimima"]),onClick:s[1]||(s[1]=m=>t.isShowPassword=!t.isShowPassword)},null,2)]),_:1},8,["type","placeholder","modelValue"])]),_:1}),o(i,{class:"login-animation3"},{default:e(()=>[o(d,{span:15},{default:e(()=>[o(l,{text:"",maxlength:"4",placeholder:n.$t("message.account.accountPlaceholder3"),modelValue:t.ruleForm.code,"onUpdate:modelValue":s[3]||(s[3]=m=>t.ruleForm.code=m),clearable:"",autocomplete:"off"},{prefix:e(()=>[o(r,{class:"el-input__icon"},{default:e(()=>[o(T)]),_:1})]),_:1},8,["placeholder","modelValue"])]),_:1}),o(d,{span:1}),o(d,{span:8},{default:e(()=>[P((_(),f(h,{class:"login-content-code"},{default:e(()=>[L("1234")]),_:1})),[[x]])]),_:1})]),_:1}),o(i,{class:"login-animation4"},{default:e(()=>[P((_(),f(h,{type:"primary",class:"login-content-submit",round:"",onClick:F,loading:t.loading.signIn},{default:e(()=>[N("span",null,Q(n.$t("message.account.accountBtnText")),1)]),_:1},8,["loading"])),[[x]])]),_:1})]),_:1})}}});const Ue=W(H,[["__scopeId","data-v-58ea2e44"]]);export{Ue as default};
