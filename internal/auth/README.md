# 权限管理说明

### 基于RBAC管理权限

使用casbin框架实现的RBAC权限管理，角色内置：`guest`, `dev`, `ops`, `root`，其中，`root`包含所有权限，`ops`包含`dev`的所有权限，`dev`包含`guest`的所有权限

### 权限配置

权限由路由权限（位于`routers.go`)和资源点(位于`resource.go`)定义，在项目启动时识别是否存在数据表`casbin_rules`，如果不存在，则会自动创建表，并将路由权限和资源点同步写入数据库中，以后在进行鉴权的时候只读取数据库中的值，因此建议做法是自行管理`casbin_rules`的数据而不是修改源代码