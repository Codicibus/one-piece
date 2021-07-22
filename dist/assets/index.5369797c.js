import{a2 as t,d as e,p as a,b as s,r,o as d,c,a as o,f as n,al as l,e as u}from"./vendor.9783560a.js";let p=t({id:"dashboard",state:()=>({loading:!1,cpu_info:[],cpu_percent:null,mem_stat:{},net_stat:[]}),getters:{cpuPercent(){let t=this.cpu_percent;return t?Number(t.toFixed(0)):0},RAMPercent(){let t=this.mem_stat.usedPercent;return t?Number(t.toFixed(0)):0}},actions:{cpuInfo(t){return this.cpu_info.forEach((e=>e[t]))},RAMStat(t){return((t,e)=>{if(0===t)return"0 B";let a=t/1024/1024,s=t/1024/1024/1024;switch(e){case"MB":return a;case"GB":return s;default:return a}})(this.mem_stat[t]).toFixed(0)+" MB"}}})();const i=e({name:"Dashboard",setup:()=>({dashboardStore:p})}),h=u();a("data-v-7885d14e");const f=o("div",{class:"title"},"CPU使用情况",-1),b={class:"detail"},_=o("div",{class:"title"},"内存使用情况",-1),m={class:"detail"};s();const v=h(((t,e,a,s,u,p)=>{const i=r("a-progress"),v=r("a-divider"),S=r("a-col"),M=r("a-row"),y=r("a-card");return d(),c(y,{title:"状态"},{default:h((()=>[o(M,{gutter:16,class:"serverState"},{default:h((()=>[o(S,{span:6,class:"state"},{default:h((()=>[f,o(i,{type:"dashboard",percent:t.dashboardStore.cpuPercent},null,8,["percent"]),o("div",b,[n(l(t.dashboardStore.cpuInfo("cores")+"个核心")+" ",1),o(v,{type:"vertical"}),n(" "+l(t.dashboardStore.cpuInfo("mhz")),1)])])),_:1}),o(S,{span:6,class:"state"},{default:h((()=>[_,o(i,{type:"dashboard",percent:t.dashboardStore.RAMPercent},null,8,["percent"]),o("div",m,[n(l(t.dashboardStore.RAMStat("used"))+" ",1),o(v,{type:"vertical"}),n(" "+l(t.dashboardStore.RAMStat("total")),1)])])),_:1}),o(S,{span:6,class:"state"},{default:h((()=>[o(i,{type:"dashboard",percent:75})])),_:1}),o(S,{span:6,class:"state"},{default:h((()=>[o(i,{type:"dashboard",percent:75})])),_:1})])),_:1})])),_:1})}));i.render=v,i.__scopeId="data-v-7885d14e";export default i;
