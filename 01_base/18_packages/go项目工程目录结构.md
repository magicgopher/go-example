# Go 项目工程目录结构

## 典型 Go 工程目录结构

```text
myproject/
├── cmd/                # 存放主应用程序入口
│   ├── myapp/         # 一个具体的可执行程序
│   │   ├── main.go   # 程序入口文件
│   └── anotherapp/    # 另一个可执行程序（可选）
│       ├── main.go   # 另一个程序入口文件
├── internal/          # 内部包，限制外部访问
│   ├── db/           # 数据库相关逻辑
│   │   ├── mysql.go
│   │   └── postgres.go
│   └── auth/         # 认证相关逻辑
│       ├── auth.go
│       └── token.go
├── pkg/               # 可复用的库代码（可选，供外部使用）
│   ├── logger/       # 日志工具包
│   │   └── logger.go
│   └── utils/        # 通用工具包
│       └── utils.go
├── api/               # API 定义（例如 REST 或 gRPC）
│   ├── swagger/      # Swagger 文档（可选）
│   └── handlers.go   # HTTP 处理函数
├── config/            # 配置文件或配置加载逻辑
│   ├── config.go     # 配置解析逻辑
│   └── config.yaml   # 配置文件示例
├── scripts/           # 脚本文件（构建、部署等）
│   ├── build.sh
│   └── deploy.sh
├── test/              # 测试相关文件（可选）
│   ├── integration/  # 集成测试
│   └── unit/         # 单元测试
├── go.mod            # Go Modules 依赖管理文件
├── go.sum            # 依赖校验文件
├── README.md         # 项目说明文档
└── .gitignore        # Git 忽略文件
```

各目录和文件的作用
1. `cmd/`
   - 用于存放项目的可执行程序入口（main 函数所在文件）。
   - 每个子目录对应一个独立的可执行程序，例如 myapp。
   - 这种结构允许多个应用程序共享同一代码库。
   - 示例：cmd/myapp/main.go 是程序的入口，调用其他包的逻辑。

2. `internal/`
   - 存放项目内部的私有代码，其他外部项目无法导入这些包。
   - 常用于封装数据库、认证、服务逻辑等模块。
   - 优点是强制模块隔离，避免外部误用。

3. `pkg/`（可选）
   - 存放可以被外部项目复用的公共库代码。
   - 如果你的项目是一个库，或者希望提供一些通用工具，可以放在这里。
   - 注意：并非所有项目都需要这个目录，小型项目可以省略。

4. `api/`（可选）
   - 存放 API 相关代码，例如 HTTP 处理函数、gRPC 定义或 Swagger 文档。
   - 适用于 Web 服务或微服务项目。

5. `config/`
   - 存放配置文件的定义和加载逻辑。
   - 可以包含 YAML、JSON 等格式的配置文件，以及解析代码。

6. `scripts/`（可选）
   - 存放构建、测试、部署相关的脚本文件。
   - 例如编译二进制文件或 Docker 部署的脚本。

7. `test/`（可选）
   - 存放测试相关的文件，可以按类型分为单元测试和集成测试。
   - Go 的测试文件通常以 _test.go 结尾，也可以直接放在对应包内。

8. `go.mod` 和 `go.sum`
   - go.mod 定义了项目的模块路径和依赖。
   - go.sum 是依赖的校验文件，确保依赖版本一致性。

9. `README.md`
   - 项目说明文档，描述项目功能、安装步骤、使用方法等。

10. `.gitignore`
    - Git 忽略文件，用于排除不需要版本控制的文件（例如 vendor/ 或编译产物）。

## 小型项目的简化结构

对于小型项目，结构可以更简单：

```text
simpleproject/
├── main.go          # 主入口文件
├── config.go        # 配置加载逻辑
├── db.go            # 数据库相关代码
├── go.mod          # 模块定义
├── go.sum          # 依赖校验
└── README.md       # 项目说明
```

这种结构适用于简单的命令行工具或单文件应用。

## Web 服务项目

假设我们要开发一个 RESTful 服务，目录结构可能是这样的：

```text
restapi/
├── cmd/
│   └── server/
│       └── main.go      # 服务启动入口
├── internal/
│   ├── db/             # 数据库操作
│   │   └── db.go
│   ├── models/         # 数据模型
│   │   └── user.go
│   └── service/        # 业务逻辑
│       └── user.go
├── api/
│   ├── handlers.go     # HTTP 处理函数
│   └── swagger.yaml    # API 文档
├── config/
│   ├── config.go       # 配置加载
│   └── config.yaml     # 默认配置
├── go.mod
├── go.sum
└── README.md
```

## DDD（Domain-Driven Design）

在 Go 语言中实现领域驱动设计（DDD, Domain-Driven Design）时，目录结构会围绕 DDD 的核心概念（如领域模型、聚合、实体、值对象、应用服务、基础设施等）进行组织。DDD 强调将业务逻辑放在核心位置，同时隔离基础设施和外部依赖。以下是一个基于 DDD 的 Go 项目目录结构的推荐设计，适用于中大型项目（例如微服务或复杂业务系统）。

```text
mydddproject/
├── cmd/                    # 主应用程序入口
│   └── server/            # 可执行程序目录
│       └── main.go        # 程序入口
├── domain/                # 领域层（核心业务逻辑）
│   ├── model/            # 领域模型（实体、值对象、聚合根）
│   │   ├── user.go      # 用户实体
│   │   └── order.go     # 订单聚合根
│   ├── repository/       # 领域仓储接口（抽象）
│   │   ├── user_repository.go
│   │   └── order_repository.go
│   └── service/          # 领域服务（复杂业务逻辑）
│       └── order_service.go
├── application/           # 应用层（协调领域逻辑）
│   ├── dto/              # 数据传输对象（输入输出）
│   │   ├── user_dto.go
│   │   └── order_dto.go
│   └── service/          # 应用服务
│       ├── user_service.go
│       └── order_service.go
├── infrastructure/        # 基础设施层（外部依赖实现）
│   ├── persistence/      # 持久化实现
│   │   ├── mysql/       # MySQL 具体实现
│   │   │   ├── user_repository.go
│   │   │   └── order_repository.go
│   │   └── redis/       # Redis 实现（如缓存）
│   │       └── cache.go
│   ├── logger/          # 日志实现
│   │   └── logger.go
│   └── httpclient/      # 外部 HTTP 调用
│       └── client.go
├── interfaces/            # 接口层（用户交互入口）
│   ├── http/             # HTTP 接口
│   │   ├── handler/     # 处理函数
│   │   │   ├── user_handler.go
│   │   │   └── order_handler.go
│   │   └── router.go    # 路由定义
│   └── cli/             # 命令行接口（可选）
│       └── cli.go
├── config/                # 配置相关
│   ├── config.go         # 配置加载逻辑
│   └── config.yaml       # 默认配置文件
├── scripts/               # 构建、部署脚本
│   ├── build.sh
│   └── deploy.sh
├── test/                  # 测试目录
│   ├── domain/          # 领域层测试
│   ├── application/     # 应用层测试
│   └── integration/     # 集成测试
├── go.mod                # Go Modules 文件
├── go.sum                # 依赖校验文件
└── README.md             # 项目说明
```

各层职责与目录说明

1. `cmd/`
   - **作用**：程序入口，负责依赖注入和启动应用程序。
   - 内容：通常只有一个 main.go，初始化配置、数据库、HTTP 服务等。
   - 示例：
   - ```go
     package main
     
     import (
         "mydddproject/config"
         "mydddproject/interfaces/http"
     )
     
     func main() {
         cfg := config.LoadConfig()
         router := http.SetupRouter(cfg)
         router.Run(":8080")
     }
     ```
   
2. `domain/`

   - 作用：DDD 的核心，包含业务逻辑和领域模型，完全独立于外部依赖。
   - 子目录：
     - **`model/`**：定义实体（Entity）、值对象（Value Object）和聚合根（Aggregate Root）。
       
       - 示例：user.go
       
       - ```go
         package model
         
         type User struct {
             ID    int
             Name  string
             Email string
         }
         ```
     
     - **`repository/`**：定义仓储接口，声明领域对象的存储和查询方法。
     
       - 示例：user_repository.go
     
       - ```go
         package repository
         
         type UserRepository interface {
             FindByID(id int) (*model.User, error)
             Save(user *model.User) error
         }
         ```
     
     - **`service/`**：领域服务，处理跨聚合的复杂业务逻辑。
     
       - 示例：order_service.go
     
       - ```go
         package service
         
         func ProcessOrder(order *model.Order) error {
             // 领域逻辑
             return nil
         }
         ```

3. `application/`

   - **作用**：协调领域层和外部请求，定义应用服务和 DTO。

   - **子目录**：

     - **`dto/`**：数据传输对象，用于输入和输出。

       - 示例：user_dto.go

       - ```go
         package dto
         
         type UserDTO struct {
             ID    int    `json:"id"`
             Name  string `json:"name"`
             Email string `json:"email"`
         }
         ```

     - **`service/`**：应用服务，调用领域服务并与仓储交互。

       - 示例：user_service.go

       - ```go
         package service
         
         type UserService struct {
             repo domain.UserRepository
         }
         
         func (s *UserService) GetUser(id int) (*dto.UserDTO, error) {
             user, err := s.repo.FindByID(id)
             if err != nil {
                 return nil, err
             }
             return &dto.UserDTO{ID: user.ID, Name: user.Name, Email: user.Email}, nil
         }
         ```

4. `infrastructure/`

   - **作用**：实现外部依赖（如数据库、日志、第三方服务）。

   - **子目录**：

     - **persistence/**：实现 domain/repository 中的接口。
       - 示例：mysql/user_repository.go
     - **logger/**：日志工具的具体实现。
     - **httpclient/**：调用外部 API 的实现。

5. `interfaces/`

   - **作用**：处理用户交互（如 HTTP、gRPC、CLI）。
   
   - **子目录**：
   
     - **`http/`**：HTTP 接口实现。
   
       - 示例：handler/user_handler.go
   
       - ```go
         package handler
         
         func GetUser(svc *service.UserService) http.HandlerFunc {
             return func(w http.ResponseWriter, r *http.Request) {
                 id := 1 // 假设从 URL 获取
                 user, err := svc.GetUser(id)
                 if err != nil {
                     http.Error(w, err.Error(), http.StatusInternalServerError)
                     return
                 }
                 json.NewEncoder(w).Encode(user)
             }
         }
         ```
   
6. `config/`**
   
    - **作用**：管理配置加载。
    - **内容**：配置文件和解析逻辑。
   
7. **`test/`**
   
    - **作用**：按层组织测试代码，确保领域逻辑和集成行为正确。

---

DDD 在 Go 中的设计特点

1. 依赖倒挂（Dependency Inversion）
   - `domain/` 定义接口（如 `UserRepository`），`infrastructure/` 提供实现。
2. 层次隔离
   - `domain/` 不依赖其他层，保持纯净。
   - `application/` 协调领域逻辑和外部输入。
   - `infrastructure/` 和 `interfaces/` 依赖其他层。
3. 模块化
   - 每个模块（如用户、订单）可以有独立的 `model`、`repository` 和 `service`。

---

示例：电商系统的简化结构

假设我们要实现一个简单的电商系统：

```text
ecommerce/
├── cmd/
│   └── server/
│       └── main.go
├── domain/
│   ├── model/
│   │   ├── product.go
│   │   └── order.go
│   ├── repository/
│   │   ├── product_repository.go
│   │   └── order_repository.go
│   └── service/
│       └── order_service.go
├── application/
│   ├── dto/
│   │   ├── product_dto.go
│   │   └── order_dto.go
│   └── service/
│       ├── product_service.go
│       └── order_service.go
├── infrastructure/
│   ├── persistence/
│   │   └── mysql/
│   │       ├── product_repository.go
│   │       └── order_repository.go
├── interfaces/
│   ├── http/
│   │   ├── handler/
│   │   │   ├── product_handler.go
│   │   │   └── order_handler.go
│   │   └── router.go
├── config/
│   ├── config.go
│   └── config.yaml
├── go.mod
└── README.md
```

优点与适用场景

- 优点：
  - 业务逻辑清晰，领域模型独立于技术细节。
  - 易于测试和维护，适合复杂业务系统。
  - 支持微服务架构，每个服务可以用类似结构。
- 适用场景：
  - 中大型项目，尤其是业务逻辑复杂的系统（如电商、金融、企业应用）。
  - 需要长期迭代和多人协作的项目。

如果你的项目较小，DDD 的完整结构可能显得繁琐，可以适当精简（例如合并 application 和 interfaces）。有什么具体需求或疑问吗？我可以进一步调整示例！