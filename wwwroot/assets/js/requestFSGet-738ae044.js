import{a}from"./axios-707ed124.js";import{S as s}from"./index-a508280f.js";import{l as c}from"./qs-bbfcbdf3.js";import{E as l,a as t}from"./element-plus-7da03efa.js";const i=a.create({baseURL:"/",timeout:5e4,headers:{"Content-Type":"application/x-www-form-urlencoded","Access-Control-Allow-Origin":"*"},paramsSerializer:{serialize(e){return c.stringify(e,{allowDots:!0})}}});i.interceptors.request.use(e=>(s.get("token")&&(e.headers.Authorization=`${s.get("token")}`),e),e=>Promise.reject(e));i.interceptors.response.use(e=>{const r=e.data;return r.code&&r.code!==0?((r.code===401||r.code===4001)&&(s.clear(),window.location.href="/",l.alert("你已被登出，请重新登录","提示",{}).then(()=>{}).catch(()=>{})),Promise.reject(i.interceptors.response)):r},e=>(e.message.indexOf("timeout")!=-1?t.error("网络超时"):e.message=="Network Error"?t.error("网络连接错误"):e.response.data?t.error(e.response.statusText):t.error("接口路径找不到"),Promise.reject(e)));const o=a.create({baseURL:"https://fschedule.fsgit.cc/",timeout:5e4,headers:{"Content-Type":"application/json","Access-Control-Allow-Origin":"*"},paramsSerializer:{serialize(e){return c.stringify(e,{allowDots:!0})}}});o.interceptors.request.use(e=>(s.get("token")&&(e.headers.Authorization=`${s.get("token")}`),e),e=>Promise.reject(e));o.interceptors.response.use(e=>{const r=e.data;return r.code&&r.code!==0?((r.code===401||r.code===4001)&&(s.clear(),window.location.href="/",l.alert("你已被登出，请重新登录","提示",{}).then(()=>{}).catch(()=>{})),Promise.reject(o.interceptors.response)):r},e=>(e.message.indexOf("timeout")!=-1?t.error("网络超时"):e.message=="Network Error"?t.error("网络连接错误"):e.response.data?t.error(e.response.statusText):t.error("接口路径找不到"),Promise.reject(e)));const m=o,n=a.create({baseURL:"https://fschedule.fsgit.cc/",timeout:5e4,headers:{"Content-Type":"application/x-www-form-urlencoded","Access-Control-Allow-Origin":"*"},paramsSerializer:{serialize(e){return c.stringify(e,{allowDots:!0})}}});n.interceptors.request.use(e=>(s.get("token")&&(e.headers.Authorization=`${s.get("token")}`),e),e=>Promise.reject(e));n.interceptors.response.use(e=>{const r=e.data;return r.code&&r.code!==0?((r.code===401||r.code===4001)&&(s.clear(),window.location.href="/",l.alert("你已被登出，请重新登录","提示",{}).then(()=>{}).catch(()=>{})),Promise.reject(n.interceptors.response)):r},e=>(e.message.indexOf("timeout")!=-1?t.error("网络超时"):e.message=="Network Error"?t.error("网络连接错误"):e.response.data?t.error(e.response.statusText):t.error("接口路径找不到"),Promise.reject(e)));export{i as a,m as r,n as s};
