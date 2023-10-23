import { defineConfig } from 'vite';
import vue from '@vitejs/plugin-vue';
import path from 'path';

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [vue()],
  resolve: {
    alias: {
      '@': path.resolve('./src'),
    },
  },
  // scss全局变量配置
  css: {
    preprocessorOptions: {
      scss: {
        javascriptEnables: true,
        additionalData: '@import "./src/styles/variable.scss";',
      },
    },
  },
});
