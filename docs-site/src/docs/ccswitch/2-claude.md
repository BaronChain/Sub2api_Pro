---
title: Claude Code 配置
icon: robot
order: 2
author: FreeCode Team
date: 2026-01-01
category:
  - CC-Switch 使用
---

# Claude Code 配置（CC-Switch）

1. 打开你下载的 CC-Switch 软件，看到初始界面。
2. 在分组条中，将分组选择至 **Claude**。
3. 在供应商分组中，新增一个供应商，端点填写 `https://freecode.codes`（或选择已配置好的 FreeCode 供应商）。
4. 回顾 [创建 API 令牌](../register/4-token.md)，在 FreeCode 中创建 **CC 分组**的令牌，点击复制按钮，复制 API Key 到剪切板。
5. 下拉配置项，找到 **API Key**，填入你刚才复制的 API Key，再点击右下角“添加”按钮。
6. 添加成功后，在主界面会看到我们配置的分组，在右侧点击“启用”按钮，显示“使用中”则配置完成。
7. 在终端运行 `claude`，看到对话界面并能正常回复即表示配置完成。

::: warning 使用提醒
如果你使用的是 **CC 分组**，请注意该分组不支持第三方接入，因此无法在 CC-Switch 中完成完整的调用测试。这类配置是否生效，请直接以 Claude Code 内的实际对话结果为准，并在 Claude Code 中完成最终测试。
:::
