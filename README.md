# Tellus

开发环境部署，请参考：[项目部署](https://git.oschina.net/SusuanServer/Tellus/wikis/项目部署)

提交代码前务必熟读并背诵：[编码规范](https://git.oschina.net/SusuanServer/Tellus/wikis/编码规范)

## 项目说明


```
|- Tellus
    |- common                   # 公共方法封装和工具方法
    |- config
    |   |- config.ini.example   # 配置文件，包含数据库链接、redis链接等
    |- constant
    |   |- define.go            # 代码中可能用到的任何独立常量
    |   |- enum.go              # 代码中的枚举常量
    |   |- cache.go             # 全部cache key定义
    |   |- status.go            # 返回数据中的status定义
    |- deploy
    |   |- sql                  # 当前数据库初始化sql
    |   |- upgrade-branch       # 上线升级脚本，当脚本中存在对数据库表结构的改动时，在上线之后合并到sql/*.sql中，并删除此路径
    |- filter
    |   |- filter.go            # 初始化文件，定义router初始化方法
    |   |- <对应业务>.go        # 接口定义，包括：参数校验、返回数据格式定义，指定filter链等
    |   |- <对应业务>           # filter的代码实现
    |- service
    |   |- <对应数据>           # 业务逻辑代码实现
    |- tellus.go                # 服务初始化和程序启动入口
    |- README.md
```
