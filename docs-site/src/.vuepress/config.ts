import { defineUserConfig } from "vuepress";
import { viteBundler } from "@vuepress/bundler-vite";
import theme from "./theme.js";

export default defineUserConfig({
  base: "/",

  bundler: viteBundler(),
  lang: "zh-CN",
  title: "FreeCode 使用文档",
  description: "FreeCode 中转服务文档，提供最稳定、最便捷的 AI 模型中转服务。",

  head: [
    ["link", { rel: "icon", href: "/logo.png" }],
    ["link", { rel: "apple-touch-icon", href: "/logo.png" }],
  ],

  theme,

  // 禁用 PWA / 调试输出，保持构建简单
  shouldPrefetch: false,
});
