<div align="center">
    <h1>🚀 Gin-Casbin-Admin</h1>
    <p>基于 Go 语言开发的企业级后台管理系统，采用领域驱动设计（DDD）架构</p>
    <p>
        <a href="https://golang.org/">
            <img src="https://img.shields.io/badge/Go-1.20%2B-blue" alt="Go version">
        </a>
        <a href="https://github.com/gin-gonic/gin">
            <img src="https://img.shields.io/badge/Gin-1.11.0-brightgreen" alt="Gin version">
        </a>
        <a href="https://gorm.io">
            <img src="https://img.shields.io/badge/GORM-1.31.1-red" alt="GORM version">
        </a>
        <a href="https://github.com/casbin/casbin">
            <img src="https://img.shields.io/badge/Casbin-2.134.0-orange" alt="Casbin version">
        </a>
        <a href="https://github.com/google/wire">
            <img src="https://img.shields.io/badge/Wire-0.7.0-yellowgreen" alt="Wire version">
        </a>
        <a href="https://github.com/wxlbd/gin-casbin-admin/blob/main/LICENSE">
            <img src="https://img.shields.io/github/license/wxlbd/gin-casbin-admin" alt="License">
        </a>
    </p>
</div>

## ✨ 核心特性

### 🏗️ 架构设计
- **DDD 领域驱动设计**: 采用经典四层架构，业务逻辑与技术实现分离
- **依赖注入**: 基于 Google Wire 的编译时依赖注入
- **分层架构**: 清晰的职责边界，高内聚低耦合
- **领域模型**: 纯净的业务实体，不依赖技术框架

### 🔐 权限管理
- **RBAC 权限控制**: 基于 Casbin 的细粒度权限管理
- **JWT 身份认证**: 支持 Token 自动续期和刷新
- **菜单权限**: 动态菜单加载和权限验证
- **多维度授权**: 支持角色、用户、资源的多维度权限控制

### 🚀 技术特性
- **多数据库支持**: PostgreSQL (主推)、MySQL、SQLite
- **高性能缓存**: Redis 会话存储和数据缓存
- **RESTful API**: 规范的接口设计和响应格式
- **统一错误处理**: 标准化的错误响应和日志记录
- **数据库迁移**: 自动化数据库版本管理

### 📊 业务功能
- **用户管理**: 完整的用户生命周期管理
- **角色管理**: 角色权限分配和管理
- **菜单管理**: 动态菜单树形结构管理
- **字典管理**: 系统配置和数据字典
- **日志审计**: 操作日志和登录日志追踪

## 🏗️ 架构设计

### DDD 分层架构

本项目采用领域驱动设计（DDD）的四层架构模式：

```
┌─────────────────────────────────────────────────────────────┐
│                    接口层 (Interfaces)                      │
│              处理外部请求，API 控制器等                     │
├─────────────────────────────────────────────────────────────┤
│                   应用层 (Application)                      │
│              协调领域对象，实现用例逻辑                     │
├─────────────────────────────────────────────────────────────┤
│                     领域层 (Domain)                        │
│              业务实体、值对象、领域服务                     │
├─────────────────────────────────────────────────────────────┤
│                 基础设施层 (Infrastructure)                 │
│            技术实现，数据库、缓存、外部服务                 │
└─────────────────────────────────────────────────────────────┘
```

### 核心领域模型

#### 用户域 (User Domain)
- 用户实体管理
- 用户认证和授权
- 用户角色关联

#### 角色域 (Role Domain)
- 角色权限管理
- 角色菜单关联
- 角色层级结构

#### 菜单域 (Menu Domain)
- 菜单树形结构
- 菜单权限控制
- 动态菜单加载

#### 字典域 (Dict Domain)
- 数据字典类型
- 字典数据管理
- 系统配置项

## 🚀 快速开始

### 环境要求

- **Go**: 1.20+
- **数据库**: PostgreSQL 10+ / MySQL 5.7+ / SQLite 3
- **缓存**: Redis 6.0+
- **操作系统**: Linux / macOS / Windows

### 安装部署

```bash
# 1. 克隆项目
git clone https://github.com/wxlbd/gin-casbin-admin.git
cd gin-casbin-admin

# 2. 安装依赖
go mod download

# 3. 配置文件
cp configs/config.yaml.example configs/config.yaml
# 编辑配置文件，设置数据库和Redis连接信息

# 4. 编译项目
go build -o app cmd/server/main.go

# 5. 数据库迁移
./app migrate up

# 6. 启动服务
./app start
```

### Docker 部署

```bash
# 使用 Docker Compose 一键部署
docker-compose up -d
```

### 开发模式

```bash
# 热重载开发模式
air

# 或者使用内置的开发模式
./app start --dev
```

## 📦 项目结构

基于 DDD 架构的项目结构：

```
.
├── cmd/                              # 应用程序入口
│   └── server/                       # HTTP 服务器
│       ├── main.go                   # 主程序入口
│       ├── wire/                     # 依赖注入配置
│       │   ├── wire.go              # Wire 配置
│       │   └── wire_gen.go          # 自动生成
│       └── migrations/              # 数据库迁移
│           ├── 000001_init_schema.up.sql
│           ├── 000002_init_data.up.sql
│           └── mysql_backup/        # MySQL 版本备份
├── configs/                          # 配置文件
│   ├── config.yaml                  # 主配置文件
│   └── casbin/                      # Casbin 规则配置
├── internal/                         # 核心业务逻辑（DDD 架构）
│   ├── application/                 # 应用层
│   │   ├── captcha/                 # 验证码应用服务
│   │   ├── dict/                    # 字典应用服务
│   │   ├── menu/                    # 菜单应用服务
│   │   ├── role/                    # 角色应用服务
│   │   └── user/                    # 用户应用服务
│   ├── domain/                      # 领域层
│   │   ├── captcha/                 # 验证码领域
│   │   ├── dict/                    # 字典领域
│   │   ├── menu/                    # 菜单领域
│   │   ├── role/                    # 角色领域
│   │   └── user/                    # 用户领域
│   ├── infrastructure/              # 基础设施层
│   │   ├── captcha/                 # 验证码基础设施
│   │   ├── persistence/             # 数据持久化
│   │   │   ├── models/              # 数据模型
│   │   │   ├── dict_repository.go   # 字典仓储实现
│   │   │   ├── menu_repository.go   # 菜单仓储实现
│   │   │   ├── role_repository.go   # 角色仓储实现
│   │   │   └── user_repository.go   # 用户仓储实现
│   ├── interfaces/                  # 接口层
│   │   └── api/v1/                  # API 接口
│   │       ├── captcha_handler.go   # 验证码接口
│   │       ├── dict_handler.go      # 字典接口
│   │       ├── menu_handler.go      # 菜单接口
│   │       ├── role_handler.go      # 角色接口
│   │       └── user_handler.go      # 用户接口
│   ├── middleware/                  # 中间件
│   │   ├── auth.go                  # JWT 认证
│   │   ├── casbin.go                # 权限控制
│   │   ├── cors.go                  # 跨域处理
│   │   └── logger.go                # 日志记录
│   ├── server/                      # 服务器配置
│   │   └── router.go                # 路由配置
│   └── types/                       # 共享类型
├── pkg/                              # 公共工具包
│   ├── config/                      # 配置管理
│   ├── errors/                      # 错误处理
│   ├── ginx/                        # Gin 扩展
│   ├── jwtx/                        # JWT 工具
│   ├── log/                         # 日志工具
│   └── utils/                       # 通用工具
├── docs/                             # 文档
│   └── swagger/                     # API 文档
├── scripts/                          # 脚本文件
├── test/                             # 测试文件
├── Makefile                          # 构建脚本
├── docker-compose.yml               # Docker 编排
└── go.mod                           # Go 模块定义
```

## 🔨 技术栈详解

### 核心框架
- **Web 框架**: [Gin 1.11.0](https://github.com/gin-gonic/gin) - 高性能、轻量级 HTTP Web 框架
- **ORM 框架**: [GORM 1.31.1](https://gorm.io/) - 强大的对象关系映射框架
- **依赖注入**: [Google Wire 0.7.0](https://github.com/google/wire) - 编译时依赖注入工具
- **权限管理**: [Casbin 2.134.0](https://casbin.org/) - 灵活的访问控制框架

### 数据存储
- **数据库**: PostgreSQL (主推) / MySQL / SQLite 多数据库支持
- **缓存**: [Redis](https://redis.io/) - 高性能键值存储
- **迁移**: [golang-migrate 4.19.1](https://github.com/golang-migrate/migrate) - 数据库版本管理

### 认证授权
- **JWT**: [golang-jwt 5.2.2](https://github.com/golang-jwt/jwt) - JSON Web Token 实现
- **会话**: Redis 会话存储，支持分布式部署
- **验证码**: [base64Captcha 1.3.8](https://github.com/mojocn/base64Captcha) - 图形验证码

### 工具组件
- **配置**: [Viper 1.21.0](https://github.com/spf13/viper) - 灵活的配置管理
- **日志**: [Zap 1.27.1](https://github.com/uber-go/zap) + Lumberjack - 结构化日志
- **验证**: [Validator 10.28.0](https://github.com/go-playground/validator) - 请求参数验证
- **CLI**: [urfave/cli 2.27.5](https://github.com/urfave/cli) - 命令行接口

## 📚 API 文档

项目集成了 Swagger API 文档，启动服务后访问：

```
http://localhost:8080/swagger/index.html
```

### 主要接口分类

```
POST   /api/v1/auth/login          # 用户登录
POST   /api/v1/auth/logout         # 用户登出
POST   /api/v1/auth/refresh        # Token 刷新
GET    /api/v1/captcha             # 获取验证码

GET    /api/v1/system/user/me      # 获取当前用户信息
GET    /api/v1/system/user/list    # 获取用户列表
POST   /api/v1/system/user         # 创建用户
PUT    /api/v1/system/user         # 更新用户
DELETE /api/v1/system/user         # 删除用户

GET    /api/v1/system/role/list    # 获取角色列表
POST   /api/v1/system/role         # 创建角色
PUT    /api/v1/system/role         # 更新角色
DELETE /api/v1/system/role         # 删除角色

GET    /api/v1/system/menu/tree    # 获取菜单树
GET    /api/v1/system/menu/list    # 获取菜单列表
POST   /api/v1/system/menu         # 创建菜单
PUT    /api/v1/system/menu         # 更新菜单
DELETE /api/v1/system/menu         # 删除菜单

GET    /api/v1/system/dict/type    # 获取字典类型
GET    /api/v1/system/dict/data    # 获取字典数据
```

## 🔧 配置说明

### 配置文件结构

```yaml
server:
  port: 8080                    # 服务端口
  mode: release                 # 运行模式

jwt:
  secret: your-secret-key       # JWT 密钥
  access_expire: 3600          # Access Token 过期时间（秒）
  refresh_expire: 86400        # Refresh Token 过期时间（秒）

redis:
  addr: localhost:6379          # Redis 地址
  password: ""                  # Redis 密码
  db: 0                         # Redis 数据库

database:
  driver: postgres              # 数据库类型（postgres/mysql/sqlite）
  host: localhost               # 数据库主机
  port: 5432                    # 数据库端口
  database: gin_admin           # 数据库名称
  username: postgres            # 数据库用户
  password: postgres            # 数据库密码
  log_level: info               # SQL 日志级别

log:
  level: info                   # 日志级别
  format: json                  # 日志格式
  output: stdout                # 日志输出
```

## 🧪 开发指南

### 环境搭建

```bash
# 安装开发依赖
go install github.com/cosmtrek/air@latest
go install github.com/swaggo/swag/cmd/swag@latest

# 初始化数据库
createdb gin_admin_development

# 启动开发服务器
air
```

### 代码规范

- 遵循 [Uber Go Style Guide](https://github.com/uber-go/guide/blob/master/style.md)
- 使用 `gofmt` 和 `golint` 进行代码格式化
- 遵循 DDD 分层架构原则
- 保持领域层的纯净性

### 添加新功能

1. **定义领域模型** 在 `internal/domain/` 目录下
2. **实现仓储接口** 在 `internal/infrastructure/persistence/` 目录下
3. **创建应用服务** 在 `internal/application/` 目录下
4. **实现接口层** 在 `internal/interfaces/api/v1/` 目录下
5. **配置依赖注入** 更新 `cmd/server/wire/wire.go`

## 🔍 监控和运维

### 健康检查

```
GET /health          # 服务健康检查
GET /metrics         # Prometheus 指标（待实现）
```

### 日志监控

- 结构化 JSON 日志输出
- 支持日志级别动态调整
- 集成日志轮转和归档

### 性能监控

- 支持接口响应时间监控
- 数据库查询性能分析
- Redis 缓存命中率统计

## 🐛 常见问题

### Q: 数据库连接失败？
A: 检查数据库配置是否正确，确保数据库服务已启动

### Q: JWT Token 无效？
A: 检查 JWT 密钥配置，确保前后端使用的密钥一致

### Q: 权限验证失败？
A: 检查 Casbin 规则配置，确保角色权限已正确分配

### Q: Redis 连接失败？
A: 检查 Redis 服务状态和连接配置

## 🤝 贡献指南

欢迎提交 Issue 和 Pull Request！

### 开发流程

1. Fork 项目到您的仓库
2. 创建新的功能分支 (`git checkout -b feature/amazing-feature`)
3. 提交您的变更 (`git commit -m 'Add some amazing feature'`)
4. 推送到分支 (`git push origin feature/amazing-feature`)
5. 创建 Pull Request

### 代码贡献

- 遵循项目的代码规范
- 为新功能添加测试用例
- 更新相关文档
- 确保所有测试通过

## 📄 许可证

[MIT License](LICENSE) - 详见 [LICENSE](LICENSE) 文件

## 📧 联系方式

- **作者**: wxl
- **邮箱**: gopher095@gmail.com
- **项目地址**: [https://github.com/wxlbd/gin-casbin-admin](https://github.com/wxlbd/gin-casbin-admin)

## 🙏 致谢

感谢以下开源项目的贡献：

- [Gin](https://github.com/gin-gonic/gin) - HTTP Web 框架
- [GORM](https://gorm.io/) - ORM 框架
- [Casbin](https://casbin.org/) - 权限管理框架
- [Viper](https://github.com/spf13/viper) - 配置管理
- [Zap](https://github.com/uber-go/zap) - 日志框架

---

<p align="center">如果这个项目对您有帮助，请给个 ⭐️ 支持一下！</p>

<p align="center">
  <a href="#top">🔝 返回顶部</a>
</p>