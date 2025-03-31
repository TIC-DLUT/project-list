# 项目主页生成器

这是一个基于JSON配置文件自动生成单页面主页的工具。通过简单的JSON配置，可以快速创建一个包含公告、会议安排和项目列表的响应式网页。

## 功能特点

- 支持公告展示
- 会议日程管理
- 项目分类展示
- 响应式布局设计
- 简单易用的JSON配置

## 使用说明

### 1. 配置文件

项目使用`config.json`文件进行配置，包含以下主要部分：

#### organization（组织信息）
包含组织的基本信息配置：
- `name`：组织名称，将显示在页面顶部
- `description`：组织简介，用于描述组织的主要情况和目标
- `link`：组织主页或相关链接，必须是完整的URL地址（如：https://github.com/example）
- `logo`：组织logo图片的URL地址，建议使用清晰的PNG或SVG格式图片

#### announcements（公告列表）
用于展示组织的公告信息，为数组类型，每个公告包含：
- `title`：公告标题，建议简明扼要
- `content`：公告详细内容，支持多行文本
- `date`：发布日期，使用YYYY-MM-DD格式（如：2024-01-15）

#### meetings（会议安排）
用于管理组织的会议日程，为数组类型，每个会议包含：
- `date`：会议日期，使用YYYY-MM-DD格式
- `time`：会议时间，使用24小时制（如：14:00-16:00）
- `location`：会议地点，可以是实体地点或线上会议链接
- `topic`：会议主题，简要描述会议内容和目的

#### projects（项目列表）
用于展示组织的项目信息，按类别组织：
- 键名为项目分类名称（如："研发项目"、"比赛项目"等）
- 值为该分类下的项目数组，每个项目包含：
  - `name`：项目名称
  - `description`：项目描述，简要说明项目目标和特点
  - `link`：项目链接，通常指向项目的代码仓库或文档

配置示例：

```json
{
  "organization": {
    "name": "创新技术研究团队",
    "description": "致力于探索和实践前沿技术的学生研究团队",
    "link": "https://github.com/example/tech-team",
    "logo": "https://example.com/logo.png"
  },
  "announcements": [
    {
      "title": "2024年第一季度工作计划",
      "content": "我们将在第一季度重点关注项目进度跟踪和技术创新",
      "date": "2024-01-15"
    }
  ],
  "meetings": [
    {
      "date": "2024-01-20",
      "time": "14:00-16:00",
      "location": "图书馆会议室A",
      "topic": "项目进度汇报与技术分享"
    }
  ],
  "projects": {
    "研发项目": [
      {
        "name": "技术分享平台",
        "description": "用于团队内部技术知识分享和学习的平台",
        "link": "https://github.com/example/tech-sharing"
      }
    ]
  }
}
```

### 2. 配置示例

```json
{
  "organization": {
    "name": "创新技术研究团队",
    "description": "致力于探索和实践前沿技术的学生研究团队",
    "link": "https://github.com/example/tech-team",
    "logo": "https://www.baidu.com/img/flexible/logo/pc/peak-result.png"
  },
  "announcements": [
    {
      "title": "2024年第一季度工作计划",
      "content": "我们将在第一季度重点关注项目进度跟踪和技术创新",
      "date": "2024-01-15"
    },
    {
      "title": "新成员招募",
      "content": "欢迎对技术感兴趣的同学加入我们的团队",
      "date": "2024-01-10"
    }
  ],
  "meetings": [
    {
      "date": "2024-01-20",
      "time": "14:00-16:00",
      "location": "图书馆会议室A",
      "topic": "项目进度汇报与技术分享"
    },
    {
      "date": "2024-02-05",
      "time": "15:00-17:00",
      "location": "线上会议",
      "topic": "代码评审与最佳实践讨论"
    }
  ],
  "projects": {
    "例会项目": [
      {
        "name": "技术分享平台",
        "description": "用于团队内部技术知识分享和学习的平台",
        "link": "https://github.com/example/tech-sharing"
      },
      {
        "name": "项目管理系统",
        "description": "团队项目进度追踪和任务分配系统",
        "link": "https://github.com/example/project-management"
      }
    ],
    "比赛项目": [
      {
        "name": "智能算法竞赛",
        "description": "基于机器学习的图像识别系统",
        "link": "https://github.com/example/ml-competition"
      },
      {
        "name": "创新应用大赛",
        "description": "基于区块链的供应链追踪系统",
        "link": "https://github.com/example/blockchain-tracking"
      }
    ]
  }
}
```

### 3. 运行项目

1. 确保已安装所需的依赖
2. 将配置文件`config.json`放置在项目根目录
3. 运行项目
4. 在浏览器中访问生成的页面

## 项目结构

```
.
├── config.json     # 配置文件
└── main.go         # 主程序
```

## 自动部署说明

将本项目fork到你的仓库，修改`config.json`，push即可

每次push都会自动更新github pages

## 效果展示

生成的页面包含以下主要部分：

- 顶部公告栏：展示最新公告信息
- 会议安排区：显示即将举行的会议
- 项目列表区：按分类展示所有项目

所有内容都采用响应式设计，可以在不同设备上良好显示。