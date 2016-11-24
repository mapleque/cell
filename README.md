# Tellus

[编码规范](https://git.oschina.net/SusuanServer/Tellus/wikis/编码规范)

[环境部署](https://git.oschina.net/SusuanServer/Tellus/wikis/项目部署)

[接口文档](https://git.oschina.net/SusuanServer/Tellus/wikis/接口文档)

## 项目说明


```
|- Tellus
    |- config
        |- config.ini.example # 总配置，包含数据库链接、redis链接等
    |- constant
        |- define.go     # 代码中可能用到的任何独立常量
        |- enum.go       # 代码中的枚举常量 
        |- redis.go      # 全部redis key定义
        |- status.go     # 返回数据中的status定义
    |- deploy
        |- sql               # 当前数据库初始化sql
        |- upgrade-branch    # 上线升级脚本，当脚本中存在对数据库表结构的改动时，在上线之后合并到sql/*.sql中，并删除此路径
    |- filters
        |- router.go    # 接口定义，包括：参数校验、返回数据格式定义，指定filter链等
        |- <对应业务>
    |- service
        |- <对应数据>    # 业务逻辑代码
    |- tellus.go   # 服务初始化和程序启动入口
    |- README.md
```