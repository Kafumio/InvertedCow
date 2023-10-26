import { createApp } from 'vue';
// 路由
import router from '@/router';
// 仓库
import pinia from '@/store';
import ElementPlus from 'element-plus';
import 'element-plus/dist/index.css';
// 配置element-plus国际化
// eslint-disable-next-line @typescript-eslint/ban-ts-comment
// @ts-ignore
import zhCn from 'element-plus/dist/locale/zh-cn.mjs';
// 引入路由鉴权文件
import './premisstion';
// 全局样式
import '@/styles/index.scss';
import App from './App.vue';
// 全局组件
import gloalComponent from '@/components';

const app = createApp(App);
app.use(ElementPlus, {
  //element-plus国际化
  locale: zhCn,
});
app.use(gloalComponent);
app.use(router);
app.use(pinia);
app.mount('#app');
