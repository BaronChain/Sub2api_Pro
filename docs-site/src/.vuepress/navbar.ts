import { navbar } from "vuepress-theme-hope";

export default navbar([
  {
    text: "快速开始",
    link: "/docs/register/",
    icon: "rocket",
  },
  {
    text: "用户中心",
    link: "/docs/user/",
    icon: "user",
  },
  {
    text: "API 接入",
    link: "/docs/api/",
    icon: "code",
  },
  {
    text: "部署运维",
    link: "/docs/deploy/",
    icon: "server",
  },
  {
    text: "FreeCode 官网",
    link: "https://freecode.codes",
    icon: "house",
  },
  {
    text: "服务监控",
    link: "https://freecode.codes/status",
    icon: "chart-line",
  },
]);
