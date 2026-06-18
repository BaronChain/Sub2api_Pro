import { hopeTheme } from "vuepress-theme-hope";
import navbar from "./navbar.js";
import sidebar from "./sidebar.js";

export default hopeTheme({
  hostname: "https://docs.freecode.codes",

  author: {
    name: "FreeCode Team",
    url: "https://freecode.codes",
  },

  logo: "/logo.png",

  repo: "",

  docsDir: "src",

  // 顶部导航栏
  navbar,

  // 侧边栏
  sidebar,

  // 页脚
  footer: "FreeCode | Copyright © 2026 FreeCode Team",
  displayFooter: true,
  copyright: "Copyright © 2026 FreeCode Team",

  // 文章元信息
  pageInfo: ["Author", "ReadingTime"],

  // 主题色
  themeColor: {
    blue: "#2196f3",
    purple: "#9c27b0",
    green: "#3eaf7c",
  },

  // 暗黑模式开关
  darkmode: "toggle",

  markdown: {
    align: true,
    tabs: true,
    codeTabs: true,
    figure: true,
    imgLazyload: true,
    imgMark: true,
    imgSize: true,
    mermaid: false,
  },

  plugins: {
    // 全文搜索
    slimsearch: {
      indexContent: true,
      suggestion: true,
    },

    components: {
      components: ["Badge", "VPCard"],
    },
  },
});
