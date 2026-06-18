---
title: Codex 配置
icon: robot
order: 3
author: FreeCode Team
date: 2026-01-01
category:
  - CC-Switch 使用
---

# Codex 配置（CC-Switch）

1. 打开你下载的 CC-Switch 软件，看到初始界面。
2. 在分组条中，将分组选择至 **Codex**。
3. 在供应商分组中，新增一个供应商，端点填写 `https://freecode.codes`（或选择已配置好的 FreeCode 供应商）。
4. 回顾 [创建 API 令牌](../register/4-token.md)，在 FreeCode 中创建 **Codex 分组**的令牌，点击复制按钮，复制 API Key 到剪切板。
5. 下拉配置项，找到 **API Key**，填入你刚才复制的 API Key，再点击右下角“添加”按钮。
6. 添加成功后，在主界面会看到我们配置的分组，在右侧点击“启用”按钮，显示“使用中”则配置完成。
7. 在终端运行 `codex`，看到对话界面并能正常回复即表示配置完成。
