package main

import (
	"encoding/json"
	"html/template"
	"io/ioutil"
	"log"
	"os"
)

type Announcement struct {
	Title   string `json:"title"`
	Content string `json:"content"`
	Date    string `json:"date"`
}

type Meeting struct {
	Date     string `json:"date"`
	Time     string `json:"time"`
	Location string `json:"location"`
	Topic    string `json:"topic"`
}

type Project struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Link        string `json:"link"`
}

type Organization struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Link        string `json:"link"`
	Logo        string `json:"logo"`
}

type Config struct {
	Copyright     string               `json:"copyright"`
	Organization  Organization         `json:"organization"`
	Announcements []Announcement       `json:"announcements"`
	Meetings      []Meeting            `json:"meetings"`
	Projects      map[string][]Project `json:"projects"`
}

const htmlTemplate = `
<!DOCTYPE html>
<html lang="zh">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>项目主页</title>
    <link href="https://cdn.jsdelivr.net/npm/bootstrap@5.3.0/dist/css/bootstrap.min.css" rel="stylesheet">
    <style>
        :root {
            --primary-color: #4a90e2;
            --secondary-color: #f8f9fa;
            --text-color: #2c3e50;
            --border-radius: 12px;
            --transition-speed: 0.3s;
        }
        body {
            font-family: 'Segoe UI', system-ui, -apple-system, sans-serif;
            color: var(--text-color);
            line-height: 1.6;
            background-color: #f5f7fa;
        }
        .navbar {
            background: linear-gradient(135deg, var(--primary-color), #357abd);
            backdrop-filter: blur(10px);
            -webkit-backdrop-filter: blur(10px);
            box-shadow: 0 4px 30px rgba(0, 0, 0, 0.1);
            border-bottom: 1px solid rgba(255, 255, 255, 0.1);
            padding: 1.2rem 0;
        }
        .navbar-brand {
            font-weight: 600;
            letter-spacing: 0.5px;
        }
        .navbar-brand img {
            transition: transform var(--transition-speed) ease;
        }
        .navbar-brand:hover img {
            transform: scale(1.1);
        }
        .container {
            padding: 2rem 1rem;
        }
        section {
            margin-bottom: 5rem;
        }
        .card {
            border: none;
            border-radius: var(--border-radius);
            background: rgba(255, 255, 255, 0.9);
            backdrop-filter: blur(10px);
            -webkit-backdrop-filter: blur(10px);
            box-shadow: 0 8px 32px rgba(0, 0, 0, 0.1);
            transition: all var(--transition-speed) ease;
        }
        .announcement-card {
            transition: transform var(--transition-speed) ease, box-shadow var(--transition-speed) ease;
            margin-bottom: 2rem;
        }
        .announcement-card:hover {
            transform: translateY(-8px);
            box-shadow: 0 12px 24px rgba(0, 0, 0, 0.15);
        }
        .meeting-row {
            transition: background-color var(--transition-speed) ease;
            border-radius: var(--border-radius);
        }
        .meeting-row:hover {
            background-color: rgba(74, 144, 226, 0.08);
        }
        .project-card {
            height: 100%;
            transition: transform var(--transition-speed) ease, box-shadow var(--transition-speed) ease;
        }
        .project-card:hover {
            transform: translateY(-8px) scale(1.02);
            box-shadow: 0 12px 24px rgba(0, 0, 0, 0.15);
        }
        .btn-primary {
            background: linear-gradient(135deg, var(--primary-color), #357abd);
            border: none;
            border-radius: var(--border-radius);
            padding: 0.6rem 1.8rem;
            font-weight: 500;
            letter-spacing: 0.5px;
            transition: all var(--transition-speed) ease;
        }
        .btn-primary:hover {
            background: linear-gradient(135deg, #357abd, var(--primary-color));
            transform: translateY(-2px);
            box-shadow: 0 8px 15px rgba(74, 144, 226, 0.3);
        }
        h2 {
            color: var(--text-color);
            font-weight: 700;
            margin-bottom: 2.5rem;
            position: relative;
            padding-bottom: 1rem;
            letter-spacing: 0.5px;
        }
        h2:after {
            content: '';
            position: absolute;
            bottom: 0;
            left: 0;
            width: 80px;
            height: 4px;
            background: linear-gradient(90deg, var(--primary-color), #357abd);
            border-radius: 4px;
        }
        .table {
            border-radius: var(--border-radius);
            overflow: hidden;
            box-shadow: 0 0 30px rgba(0, 0, 0, 0.08);
            margin-bottom: 2rem;
        }
        .table thead {
            background: linear-gradient(135deg, var(--primary-color), #357abd);
            color: white;
        }
        .table th {
            font-weight: 600;
            padding: 1.2rem 1rem;
            letter-spacing: 0.5px;
        }
        .table td {
            padding: 1rem;
            vertical-align: middle;
        }
        footer {
            background: linear-gradient(135deg, #f8f9fa, #e9ecef);
            padding: 3rem 0;
            text-align: center;
            margin-top: 5rem;
            box-shadow: 0 -4px 30px rgba(0, 0, 0, 0.05);
            border-top: 1px solid rgba(0, 0, 0, 0.05);
        }
    </style>
</head>
<body>
    <nav class="navbar navbar-expand-lg navbar-dark bg-primary">
        <div class="container">
            <a class="navbar-brand" href="{{.Organization.Link}}" target="_blank">
                {{if .Organization.Logo}}
                <img src="{{.Organization.Logo}}" alt="{{.Organization.Name}}" height="30" class="d-inline-block align-text-top me-2">
                {{end}}
                {{.Organization.Name}}
            </a>
        </div>
    </nav>

    <div class="container mt-4">
        <!-- 组织介绍部分 -->
        <section class="mb-5">
            <div class="card">
                <div class="card-body">
                    <h2 class="card-title mb-3">关于我们</h2>
                    <p class="card-text">{{.Organization.Description}}</p>
                </div>
            </div>
        </section>
        <!-- 公告部分 -->
        <section class="mb-5">
            <h2 class="mb-4">近期公告</h2>
            <div class="row">
                {{range .Announcements}}
                <div class="col-md-6 mb-4">
                    <div class="card announcement-card">
                        <div class="card-body">
                            <h5 class="card-title">{{.Title}}</h5>
                            <p class="card-text">{{.Content}}</p>
                            <small class="text-muted">{{.Date}}</small>
                        </div>
                    </div>
                </div>
                {{end}}
            </div>
        </section>

        <!-- 例会安排部分 -->
        <section class="mb-5">
            <h2 class="mb-4">本学期例会安排</h2>
            <div class="table-responsive">
                <table class="table table-hover">
                    <thead class="table-light">
                        <tr>
                            <th>日期</th>
                            <th>时间</th>
                            <th>地点</th>
                            <th>内容</th>
                        </tr>
                    </thead>
                    <tbody>
                        {{range .Meetings}}
                        <tr class="meeting-row">
                            <td>{{.Date}}</td>
                            <td>{{.Time}}</td>
                            <td>{{.Location}}</td>
                            <td>{{.Topic}}</td>
                        </tr>
                        {{end}}
                    </tbody>
                </table>
            </div>
        </section>

        <!-- 项目列表部分 -->
        <section>
            <h2 class="mb-4">项目列表</h2>
            {{range $category, $projects := .Projects}}
            <div class="mb-5">
                <h3 class="mb-3">{{$category}}</h3>
                <div class="row">
                    {{range $projects}}
                    <div class="col-md-6 mb-4">
                        <div class="card project-card">
                            <div class="card-body">
                                <h5 class="card-title">{{.Name}}</h5>
                                <p class="card-text">{{.Description}}</p>
                                <a href="{{.Link}}" class="btn btn-primary" target="_blank">查看项目</a>
                            </div>
                        </div>
                    </div>
                    {{end}}
                </div>
            </div>
            {{end}}
        </section>
    </div>

    <footer class="mt-5">
        <div class="container">
            <p class="mb-0">{{.Copyright}}</p>
        </div>
    </footer>

    <script src="https://cdn.jsdelivr.net/npm/bootstrap@5.3.0/dist/js/bootstrap.bundle.min.js"></script>
</body>
</html>
`

func main() {
	// 读取配置文件
	configFile, err := ioutil.ReadFile("config.json")
	if err != nil {
		log.Fatal("Error reading config file:", err)
	}

	// 解析配置文件
	var config Config
	if err := json.Unmarshal(configFile, &config); err != nil {
		log.Fatal("Error parsing config file:", err)
	}

	// 创建模板
	tmpl, err := template.New("page").Parse(htmlTemplate)
	if err != nil {
		log.Fatal("Error parsing template:", err)
	}

	// 创建输出文件
	outputFile, err := os.Create("build/index.html")
	if err != nil {
		log.Fatal("Error creating output file:", err)
	}
	defer outputFile.Close()

	// 生成HTML
	if err := tmpl.Execute(outputFile, config); err != nil {
		log.Fatal("Error generating HTML:", err)
	}
}
