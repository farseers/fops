import{_ as T,l as V,u as w,o as M,e as N}from"./index-dd22bf27.js";import{k as S,aN as j,$ as E,j as F,H as O,ah as l,o as s,c as p,W as r,Q as a,F as h,a9 as v,u as k,P as C,a as g,V as u,ab as U,U as y,a4 as Q}from"./@vue-d9027515.js";import{u as W,o as q}from"./vue-router-a2fac16f.js";import{s as L}from"./pinia-be7e992d.js";import{_ as G}from"./_plugin-vue_export-helper-c27b6911.js";import"./vue-i18n-e0b9a81d.js";import"./@intlify-9e8a497c.js";import"./source-map-7d7e1c08.js";import"./vue-b46b26c7.js";import"./js-cookie-cf83ad76.js";import"./@element-plus-96a31696.js";import"./nprogress-8d2808ea.js";import"./axios-707ed124.js";import"./qs-bbfcbdf3.js";import"./side-channel-eafc5c70.js";import"./get-intrinsic-b9397c9a.js";import"./has-symbols-e8f3ca0e.js";import"./function-bind-22e7ee79.js";import"./has-26d28e02.js";import"./call-bind-9ceb8f5b.js";import"./object-inspect-7b08a46e.js";import"./element-plus-7da03efa.js";import"./lodash-es-36eb724a.js";import"./@vueuse-2d9216d7.js";import"./@ctrl-f8748455.js";import"./dayjs-6aac7dfa.js";import"./async-validator-dee29e8b.js";import"./memoize-one-297ddbcb.js";import"./escape-html-1d60d822.js";import"./normalize-wheel-es-ed76fb12.js";import"./@floating-ui-463e90e0.js";import"./mitt-f7ef348c.js";import"./vue-grid-layout-0ee01be3.js";import"./vue-demi-71ba0ef2.js";const J={class:"el-menu-horizontal-warp"},K=S({name:"navMenuHorizontal"}),X=S({...K,props:{menuList:{type:Array,default:()=>[]}},setup(A){const b=A,x=j(()=>T(()=>import("./subItem-7605eee0.js"),["assets/js/subItem-7605eee0.js","assets/js/@vue-d9027515.js","assets/js/index-dd22bf27.js","assets/js/pinia-be7e992d.js","assets/js/vue-demi-71ba0ef2.js","assets/js/vue-router-a2fac16f.js","assets/js/vue-i18n-e0b9a81d.js","assets/js/@intlify-9e8a497c.js","assets/js/source-map-7d7e1c08.js","assets/js/vue-b46b26c7.js","assets/js/js-cookie-cf83ad76.js","assets/js/@element-plus-96a31696.js","assets/js/nprogress-8d2808ea.js","assets/css/nprogress-8b89e2e0.css","assets/js/axios-707ed124.js","assets/js/qs-bbfcbdf3.js","assets/js/side-channel-eafc5c70.js","assets/js/get-intrinsic-b9397c9a.js","assets/js/has-symbols-e8f3ca0e.js","assets/js/function-bind-22e7ee79.js","assets/js/has-26d28e02.js","assets/js/call-bind-9ceb8f5b.js","assets/js/object-inspect-7b08a46e.js","assets/js/element-plus-7da03efa.js","assets/js/lodash-es-36eb724a.js","assets/js/@vueuse-2d9216d7.js","assets/js/@ctrl-f8748455.js","assets/js/dayjs-6aac7dfa.js","assets/js/async-validator-dee29e8b.js","assets/js/memoize-one-297ddbcb.js","assets/js/escape-html-1d60d822.js","assets/js/normalize-wheel-es-ed76fb12.js","assets/js/@floating-ui-463e90e0.js","assets/js/mitt-f7ef348c.js","assets/js/vue-grid-layout-0ee01be3.js","assets/css/index-c2520709.css"])),R=V(),$=w(),{routesList:B}=L(R),{themeConfig:d}=L($),I=W(),c=E({defaultActive:""}),z=F(()=>b.menuList),_=o=>o.filter(t=>{var e;return!((e=t.meta)!=null&&e.isHide)}).map(t=>(t=Object.assign({},t),t.children&&(t.children=_(t.children)),t)),D=o=>{const t=o.split("/");let e={children:[]};return _(B.value).map((i,m)=>{i.path===`/${t[1]}`&&(i.k=m,e.item={...i},e.children=[{...i}],i.children&&(e.children=i.children))}),e},f=o=>{const{path:t,meta:e}=o;if(d.value.layout==="classic")c.defaultActive=`/${t==null?void 0:t.split("/")[1]}`;else{const i=e!=null&&e.isDynamic?e.isDynamicPath.split("/"):t.split("/");i.length>=4&&(e!=null&&e.isHide)?c.defaultActive=i.splice(0,3).join("/"):c.defaultActive=t}},H=o=>{M.handleOpenLink(o)};return O(()=>{f(I)}),q(o=>{f(o);let{layout:t,isClassicSplitMenu:e}=d.value;t==="classic"&&e&&N.emit("setSendClassicChildren",D(o.path))}),(o,t)=>{const e=l("SvgIcon"),i=l("el-sub-menu"),m=l("el-menu-item"),P=l("el-menu");return s(),p("div",J,[r(P,{router:"","default-active":c.defaultActive,"background-color":"transparent",mode:"horizontal"},{default:a(()=>[(s(!0),p(h,null,v(k(z),n=>(s(),p(h,null,[n.children&&n.children.length>0?(s(),C(i,{index:n.path,key:n.path},{title:a(()=>[r(e,{name:n.meta.icon},null,8,["name"]),g("span",null,u(o.$t(n.meta.title)),1)]),default:a(()=>[r(k(x),{chil:n.children},null,8,["chil"])]),_:2},1032,["index"])):(s(),C(m,{index:n.path,key:n.path},U({_:2},[!n.meta.isLink||n.meta.isLink&&n.meta.isIframe?{name:"title",fn:a(()=>[r(e,{name:n.meta.icon},null,8,["name"]),y(" "+u(o.$t(n.meta.title)),1)]),key:"0"}:{name:"title",fn:a(()=>[g("a",{class:"w100",onClick:Q(Y=>H(n),["prevent"])},[r(e,{name:n.meta.icon},null,8,["name"]),y(" "+u(o.$t(n.meta.title)),1)],8,["onClick"])]),key:"1"}]),1032,["index"]))],64))),256))]),_:1},8,["default-active"])])}}});const Pe=G(X,[["__scopeId","data-v-bfa3ac3f"]]);export{Pe as default};