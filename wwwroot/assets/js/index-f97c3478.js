import{_ as h}from"./index-1f38180f.js";import{k as d,aN as f,$ as v,ah as r,o as n,c as m,W as e,Q as a,u as p,F as b,a9 as g,P as x,U as y,V as k}from"./@vue-d9027515.js";import"./pinia-be7e992d.js";import"./vue-demi-71ba0ef2.js";import"./vue-router-a2fac16f.js";import"./vue-i18n-e0b9a81d.js";import"./@intlify-9e8a497c.js";import"./source-map-7d7e1c08.js";import"./vue-b46b26c7.js";import"./js-cookie-cf83ad76.js";import"./@element-plus-96a31696.js";import"./nprogress-8d2808ea.js";import"./axios-707ed124.js";import"./qs-bbfcbdf3.js";import"./side-channel-eafc5c70.js";import"./get-intrinsic-b9397c9a.js";import"./has-symbols-e8f3ca0e.js";import"./function-bind-22e7ee79.js";import"./has-26d28e02.js";import"./call-bind-9ceb8f5b.js";import"./object-inspect-7b08a46e.js";import"./element-plus-7da03efa.js";import"./lodash-es-36eb724a.js";import"./@vueuse-2d9216d7.js";import"./@ctrl-f8748455.js";import"./dayjs-6aac7dfa.js";import"./async-validator-dee29e8b.js";import"./memoize-one-297ddbcb.js";import"./escape-html-1d60d822.js";import"./normalize-wheel-es-ed76fb12.js";import"./@floating-ui-463e90e0.js";import"./mitt-f7ef348c.js";import"./vue-grid-layout-0ee01be3.js";const w={class:"notice-bar-container layout-pd"},B=d({name:"makeNoticeBar"}),ne=d({...B,setup(D){const i=f(()=>h(()=>import("./index-faf3cbde.js"),["assets/js/index-faf3cbde.js","assets/js/@vue-d9027515.js","assets/js/_plugin-vue_export-helper-c27b6911.js","assets/css/index-9f150520.css"])),l=v({noticeList:["🎉🎉🔥基于vue3.x 、Typescript、vite、Element plus等","适配手机、平板、pc的后台开源免费模板库（vue2.x请切换vue-prev-admin分支）","仓库地址：https://gitee.com/lyt-top/vue-next-admin","演示地址：https://lyt-top.gitee.io/vue-next-admin-preview/#/login"],tableData:[{a1:"mode",a2:"通知栏模式，用于右侧 icon 图标点击",a3:"string",a4:"closeable / link",a5:""},{a1:"text",a2:"通知文本内容，scrollable 为 false 时生效",a3:"string",a4:"",a5:""},{a1:"color",a2:"通知文本颜色",a3:"string",a4:"",a5:"#e6a23c"},{a1:"background",a2:"通知背景色",a3:"string",a4:"",a5:"#fdf6ec"},{a1:"size",a2:"字体大小，单位px",a3:"number / string",a4:"",a5:"14"},{a1:"height",a2:"通知栏高度，单位px",a3:"number / string",a4:"",a5:"40"},{a1:"delay",a2:"动画延迟时间 (s)",a3:"number / string",a4:"",a5:"1"},{a1:"speed",a2:"滚动速率 (px/s)",a3:"number / string",a4:"",a5:"100"},{a1:"scrollable",a2:"是否开启垂直滚动",a3:"boolean",a4:"true",a5:"false"},{a1:"leftIcon",a2:"自定义左侧图标",a3:"string",a4:"",a5:""},{a1:"rightIcon",a2:"自定义右侧图标",a3:"string",a4:"",a5:""}],tableData1:[{a1:"close",a2:"通知栏模式（mode）closeable 时回调事件",a3:"function",a4:""},{a1:"link",a2:"通知栏模式（mode）link 时回调事件",a3:"function",a4:""}]});return(E,N)=>{const o=r("el-card"),u=r("el-carousel-item"),_=r("el-carousel"),t=r("el-table-column"),s=r("el-table");return n(),m("div",w,[e(o,{shadow:"hover",header:"滚动通知栏：默认"},{default:a(()=>[e(p(i),{text:`🎉🎉🔥基于vue3.x 、Typescript、vite、Element plus等，适配手机、平板、pc
				的后台开源免费模板库（vue2.x请切换vue-prev-admin分支），仓库地址：https://gitee.com/lyt-top/vue-next-admin`})]),_:1}),e(o,{shadow:"hover",header:"滚动通知栏：设置样式",class:"mt15"},{default:a(()=>[e(p(i),{text:`🎉🎉🔥基于vue3.x 、Typescript、vite、Element plus等，适配手机、平板、pc
				的后台开源免费模板库（vue2.x请切换vue-prev-admin分支），仓库地址：https://gitee.com/lyt-top/vue-next-admin`,leftIcon:"iconfont icon-tongzhi2",rightIcon:"ele-ArrowRight",background:"#ecf5ff",color:"#409eff"})]),_:1}),e(o,{shadow:"hover",header:"滚动通知栏：搭配 NoticeBar 和 Carousel 走马灯 组件可以实现垂直滚动的效果",class:"mt15"},{default:a(()=>[e(p(i),{scrollable:!0},{default:a(()=>[e(_,{height:"40px",direction:"vertical",autoplay:!0,"indicator-position":"none",interval:3e3},{default:a(()=>[(n(!0),m(b,null,g(l.noticeList,c=>(n(),x(u,{key:c},{default:a(()=>[y(k(c),1)]),_:2},1024))),128))]),_:1})]),_:1})]),_:1}),e(o,{shadow:"hover",header:"滚动通知栏：参数",class:"mt15"},{default:a(()=>[e(s,{data:l.tableData,style:{width:"100%"}},{default:a(()=>[e(t,{prop:"a1",label:"参数"}),e(t,{prop:"a2",label:"说明"}),e(t,{prop:"a3",label:"类型"}),e(t,{prop:"a4",label:"可选值"}),e(t,{prop:"a5",label:"默认值"})]),_:1},8,["data"])]),_:1}),e(o,{shadow:"hover",header:"图标选择器(宽度自动)：事件",class:"mt15"},{default:a(()=>[e(s,{data:l.tableData1,style:{width:"100%"}},{default:a(()=>[e(t,{prop:"a1",label:"事件名称"}),e(t,{prop:"a2",label:"说明"}),e(t,{prop:"a3",label:"类型"}),e(t,{prop:"a4",label:"回调参数"})]),_:1},8,["data"])]),_:1})])}}});export{ne as default};