---
title: GPT-Image-2
icon: image
order: 2
author: FreeCode Team
date: 2026-01-01
category:
  - 绘图模型教程
---

# GPT-Image-2 绘图教程

## 前置准备

`gpt-image-2` 模型属于 **Sora 分组**，使用前需要创建令牌分组为 **sora** 的令牌。参照 [创建 API 令牌](../register/4-token.md) 创建令牌，分组选择 `sora`。

## 调用方式

OpenAI 把图片相关能力分成三类，对 FreeCode 的 `gpt-image-2` 来说，出图请优先使用 **Images API**。

| API | 用途 | 是否支持出图 |
| --- | --- | --- |
| Responses API | 分析图片输入 | ❌ 不支持作为出图入口 |
| **Images API** | 文生图、图片编辑 | ✅ **推荐** |
| Chat Completions API | 分析图片输入并生成文本 | ❌ 不支持作为出图入口 |

## 方式一：Images API（推荐）

Images API 分为文生图和图片编辑两个接口：

- 文生图：`POST https://freecode.codes/v1/images/generations`
- 图片编辑 / 图生图：`POST https://freecode.codes/v1/images/edits`

### 文生图：/v1/images/generations

```bash
curl --location 'https://freecode.codes/v1/images/generations' \
--header 'Content-Type: application/json' \
--header 'Authorization: Bearer 你的Sora分组令牌' \
--data '{
    "model": "gpt-image-2",
    "prompt": "一只橘猫戴着橙色围巾抱着水獭，温暖插画风格",
    "size": "3840x2160",
    "quality": "high",
    "output_format": "png",
    "response_format": "url",
    "n": 1
}'
```

文生图主要参数：

| 参数 | 类型 | 支持情况 | 说明 |
| --- | --- | --- | --- |
| `model` | string | 支持 | 固定填写 `gpt-image-2` |
| `prompt` | string | 支持 | 描述主体、场景、风格、比例、文字内容 |
| `n` | integer | 仅支持 1 | 一次只返回 1 张图 |
| `size` | string | 支持 | 如 `1024x1024`、`1536x1024`、`3840x2160`、`auto` |
| `quality` | string | 支持 | `low` / `medium` / `high` / `auto` |
| `response_format` | string | 支持 | `url`（推荐）或 `b64_json` |
| `output_format` | string | 部分支持 | 推荐 `png` 或 `jpeg`，不建议 `webp` |
| `stream` | boolean | 不支持 | 请不要开启 |

### 图片编辑 / 图生图：/v1/images/edits

使用 `multipart/form-data` 上传图片，`image` 是二进制图片文件，`prompt` 写清楚如何修改。

```bash
curl --location 'https://freecode.codes/v1/images/edits' \
--header 'Authorization: Bearer 你的Sora分组令牌' \
--form 'model="gpt-image-2"' \
--form 'prompt="保留主体，在右上角加一枚红色小印章，写 DEMO"' \
--form 'image=@"/path/to/your-image.jpg"' \
--form 'size="1024x1024"' \
--form 'quality="high"' \
--form 'output_format="png"' \
--form 'response_format="url"'
```

局部修改时可额外传 `mask`（PNG，透明区域表示允许重点修改的位置）。

## 尺寸与质量说明

常用尺寸：`1024x1024`（正方形）、`1536x1024`（横向）、`1024x1536`（纵向）、`2048x2048`（2K）、`3840x2160`（4K 横向）、`auto`（默认）。

尺寸限制：最大边长 ≤ 3840 像素；宽高都必须是 16 的倍数；长短边比例不超过 3:1；总像素数不少于 655,360 且不超过 8,294,400。

## 返回结果

默认返回图片下载地址：

```json
{
  "created": 1776923999,
  "data": [
    {
      "url": "https://freecode.codes/file_download/xxxxxxxx",
      "revised_prompt": "..."
    }
  ]
}
```

`revised_prompt` 是模型实际使用前改写过的提示词，属于正常现象。如果请求传了 `"response_format": "b64_json"`，返回内容会变成 Base64 图片数据，需客户端自行解码保存。

## 在 Cherry Studio 中使用

1. 创建 **sora 分组**的令牌并复制 API Key。
2. 下载安装 Cherry Studio，进入 **模型服务**，点击 **添加** 新增提供商，类型选择 **New API**。
3. 将 sora 分组 API Key 填入 **API 密钥**，**API 地址** 填写 `https://freecode.codes`。
4. 点击 **获取模型列表**，添加 `gpt-image-2` 模型。
5. 点击该模型右侧编辑按钮，将 **端点类型** 设置为 **图像生成（OpenAI）** 后保存。
6. 回到首页，点击 `+`，选择 **绘画** 应用。
7. 左侧 **提供商** 选择刚添加的供应商，**模型** 选择 `gpt-image-2`。
8. **绘图** 模式可文生图；**编辑** 模式可上传参考图进行图生图或局部修改。

::: warning 长连接与代理设置
图片生成请求耗时较长（尤其编辑模式、高分辨率），本机代理或网关如果对长连接有 60 秒限制，可能在生成完成前断开，表现为 `Failed to fetch` 或请求超时。建议把 `freecode.codes` 加入代理工具的直连/白名单规则，让绘图请求直连，不再经过代理。
:::
