# Tellus

开发环境部署，请参考：[项目部署](https://git.oschina.net/SusuanServer/Tellus/wikis/项目部署)

提交代码前务必熟读并背诵：[编码规范](https://git.oschina.net/SusuanServer/Tellus/wikis/编码规范)

## 项目说明

> 什么属于用户中心范畴
- 用户登录、登出（passport模块）
- 用户注册、邀请管理（register模块）
- 用户基本信息管理（info模块）
- 用户资源体系（source模块）
- 用户特权体系（feature模块）

> 什么不属于用户中心范畴
- 用户支付相关信息（支付中心）
- 用户作业相关信息（业务线）
- 用户游戏相关信息（业务线）
- 用户学习相关信息（业务线）
- 其他业务相关信息（业务线）

## 文件组织
```
|- Tellus
    |- common                   # 公共方法封装和工具方法
    |   |- auth.go                  # 公共filter，用于接口请求权限校验
    |   |- crypto.go                # 加密工具封装
    |- config                   # 配置，只允许写example
    |   |- config.ini.example       # 配置文件，包含数据库链接、redis链接等
    |- constant                 # 常量定义
    |   |- define.go                # 代码中可能用到的任何独立常量
    |   |- enum.go                  # 代码中的枚举常量
    |   |- cache.go                 # 全部cache key定义
    |   |- status.go                # 返回数据中的status定义
    |- deploy                   # 部署相关
    |   |- sql                      # 当前数据库初始化sql
    |   |   |- user_system.sql          # user_system表
    |   |   |- user_info.sql            # user_info表
    |   |   |- user_resource.sql        # user_resourcel表
    |   |   |- user_feature.sql         # user_feauter表
    |   |- upgrade-branch           # 上线升级脚本，当脚本中存在对数据库表结构的改动时，在上线之后合并到sql/*.sql中，并删除此路径
    |- filter                   # 接口定义，包括：参数校验、返回数据格式定义，指定filter链等
    |   |- filter.go                # 初始化filter，定义router初始化方法
    |   |- passport.go              # 登录登出等接口
    |   |- passport
    |   |   |- token.go                 # token转换
    |   |- info.go                  # 用户信息相关接口
    |   |- info
    |   |   |- register.go              # 注册
    |   |   |- update.go                # 更新信息
    |   |   |- freeze.go                # 冻结
    |   |- resource.go              # 计数资源相关接口
    |   |- resource
    |   |   |- check.go                 # 检查资源
    |   |   |- use.go                   # 使用资源
    |   |   |- update.go                # 更新资源
    |   |   |- manage.go                # 管理系统资源
    |   |- feature.go               # 计时资源相关接口
    |   |- feature
    |   |   |- check.go                 # 检查资源
    |   |   |- update.go                # 更新资源
    |   |   |- manage.go                # 管理系统资源
    |- service                  # 业务逻辑代码实现
    |   |- system                   # 用户系统信息相关操作
    |   |- info                     # 用户个人信息相关操作
    |   |- resource                 # 计数资源相关操作
    |   |- feature                  # 计时资源相关操作
    |- tellus.go                # 服务初始化和程序启动入口
    |- README.md
```

## 需求
当前速算学生端业务中，与用户中心相关的内容整理如下：
- 学生登录：学生登录，换取token，在之后的请求中携带token，系统解析token获取学生id
- 学生注册：手机号注册，获取验证码，写入数据
- 获取学生信息：通过token换id，通过id，获取信息
- 学生资源（所有计数的东西）：体力，积分，钥匙，金币，各种道具
- 学生特权（所有计时的东西）：学习资源包（永久），VIP

## API状态
@see /doc
```
|- passport
|   |- login tbd
|   |- logout tbd
|- info
|   |- check tbd
|   |- update tbd
|   |- register tbd
|- resource
|   |- check tbd
|   |- update tbd
|   |- manage tbd
|- feature
|   |- check tbd
|   |- update tbd
|   |- manage tbd
```

## 数据设计
所有user相关的表视情况根据user_id分表。
```
                     |-------- teacher *tbd
                     |-------- student
     user_system ----|-------- user_resource -------- resource
                     |-------- user_feature ----------- feature
```

- user
type 类型：学生、老师、家长...

- user_system
id
username 用户名
password 密码
inviter_id 邀请人id
channel 注册渠道
status 状态：正常、冻结、注销
device_code 
device_code 设备码
version 客户端版本
last_login_time 最后登录时间

- user_info
user_id [FK]
name 姓名
mobile 手机号
sex 性别
head_photo 头像

- resource
id
name 名称
type 类型
status 状态
desc 描述
time 添加时间

- user_resource
user_id [FK]
source_id [FK]
amount 数量
time 最后修改时间

- feature
id
name 名称
type 类型
status 状态
desc 描述
time 添加时间

- user_feature
user_id [FK]
privilege_id [FK]
activated_time 激活时间
expired_time 过期时间
time 最后修改时间
