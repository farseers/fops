import{_ as i}from"./index-a79ea617.js";import{a as E}from"./element-plus-7da03efa.js";import{k as h,aN as p,e as f,$ as x,ah as s,o as V,c as C,W as e,Q as o,u as c,a as P,U as R}from"./@vue-d9027515.js";import"./pinia-be7e992d.js";import"./vue-demi-71ba0ef2.js";import"./vue-router-a2fac16f.js";import"./vue-i18n-e0b9a81d.js";import"./@intlify-9e8a497c.js";import"./source-map-7d7e1c08.js";import"./vue-b46b26c7.js";import"./js-cookie-cf83ad76.js";import"./@element-plus-96a31696.js";import"./nprogress-8d2808ea.js";import"./axios-707ed124.js";import"./qs-bbfcbdf3.js";import"./side-channel-eafc5c70.js";import"./get-intrinsic-b9397c9a.js";import"./has-symbols-e8f3ca0e.js";import"./function-bind-22e7ee79.js";import"./has-26d28e02.js";import"./call-bind-9ceb8f5b.js";import"./object-inspect-7b08a46e.js";import"./mitt-f7ef348c.js";import"./vue-grid-layout-0ee01be3.js";import"./lodash-es-36eb724a.js";import"./@vueuse-2d9216d7.js";import"./@ctrl-f8748455.js";import"./dayjs-6aac7dfa.js";import"./async-validator-dee29e8b.js";import"./memoize-one-297ddbcb.js";import"./escape-html-1d60d822.js";import"./normalize-wheel-es-ed76fb12.js";import"./@floating-ui-463e90e0.js";const D={class:"form-rules-container layout-pd"},I={class:"flex-margin"},$=h({name:"pagesFormRules"}),_e=h({...$,setup(A){const v=p(()=>i(()=>import("./formRulesOne-ed38fb31.js"),["assets/js/formRulesOne-ed38fb31.js","assets/js/@vue-d9027515.js"])),F=p(()=>i(()=>import("./formRulesTwo-89c13dce.js"),["assets/js/formRulesTwo-89c13dce.js","assets/js/@vue-d9027515.js"])),g=p(()=>i(()=>import("./formRulesThree-228494ec.js"),["assets/js/formRulesThree-228494ec.js","assets/js/@vue-d9027515.js"])),a=f(),m=f(),l=f(),T=x({formRulesOneData:{name:"lyt",email:"lyt123@.com",autograph:"lyt123456",occupation:"1"}}),n=(u,_)=>new Promise(t=>{u.value.$refs[_].validate(r=>{r&&t(r)})}),w=()=>{a.value.$refs.formRulesOneRef.resetFields(),m.value.$refs.formRulesTwoRef.resetFields(),l.value.$refs.formRulesThreeRef.resetFields()},y=()=>{Promise.all([n(a,"formRulesOneRef"),n(m,"formRulesTwoRef"),n(l,"formRulesThreeRef")]).then(()=>{E.success("表单全部验证成功")})},O=()=>{w()};return(u,_)=>{const t=s("el-card"),r=s("SvgIcon"),d=s("el-button"),k=s("el-row");return V(),C("div",D,[e(t,{shadow:"hover",header:"表单组件1"},{default:o(()=>[e(c(v),{data:T.formRulesOneData,ref_key:"pagesFormRulesOneRef",ref:a},null,8,["data"])]),_:1}),e(t,{shadow:"hover",header:"表单组件2",class:"mt15"},{default:o(()=>[e(c(F),{ref_key:"pagesFormRulesTwoRef",ref:m},null,512)]),_:1}),e(t,{shadow:"hover",header:"表单组件3",class:"mt15"},{default:o(()=>[e(c(g),{ref_key:"pagesFormRulesThreeRef",ref:l},null,512)]),_:1}),e(k,{class:"flex mt15"},{default:o(()=>[P("div",I,[e(d,{size:"default",onClick:O},{default:o(()=>[e(r,{name:"ele-RefreshRight"}),R(" 重置表单 ")]),_:1}),e(d,{size:"default",type:"primary",onClick:y},{default:o(()=>[e(r,{name:"iconfont icon-shuxing"}),R(" 验证表单 ")]),_:1})])]),_:1})])}}});export{_e as default};
