// Plugins
import Components from 'unplugin-vue-components/vite'
import Vue from '@vitejs/plugin-vue'
import Vuetify, { transformAssetUrls } from 'vite-plugin-vuetify'
import ViteFonts from 'unplugin-fonts/vite'
import VueRouter from 'unplugin-vue-router/vite'

// Utilities
import { defineConfig } from 'vite'
import { fileURLToPath, URL } from 'node:url'

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [
    VueRouter(),
    Vue({
      template: { transformAssetUrls }
    }),
    // https://github.com/vuetifyjs/vuetify-loader/tree/master/packages/vite-plugin#readme
    Vuetify({
      autoImport: true,
      styles: {
        configFile: 'src/styles/settings.scss',
      },
    }),
    Components(),
    ViteFonts({
      google: {
        families: [{
          name: 'Roboto',
          styles: 'wght@100;300;400;500;700;900',
        }],
      },
    }),
  ],
  define: { 'process.env': {} },
  resolve: {
    alias: {
      '@': fileURLToPath(new URL('./src', import.meta.url))
    },
    extensions: [
      '.js',
      '.json',
      '.jsx',
      '.mjs',
      '.ts',
      '.tsx',
      '.vue',
    ],
  },
  server: {
    // detail: https://vitejs.cn/vite3-cn/config/server-options.html
    host: '0.0.0.0', // 指定服务器主机名，0.0.0.0将监听所有地址，包括局域网地址
    port: 3000,      // 指定服务器端口
    open: true,      // 运行项目时自动打开浏览器
    cors: true,      // 允许跨源资源共享
    strictPort: true,// 如果端口已被占用，则直接退出
    https: false,    // 是否启动https
    proxy: {         // 代理配置
      '/api': {
        target: `${process.env.VITE_BASE_PATH}/`, // 代理到 目标路径
        changeOrigin: false,
        rewrite: path => path.replace(new RegExp('^' + process.env.VITE_BASE_API), ''),
      }
    }
  },
})
