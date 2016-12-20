## 项目说明

> 什么属于用户中心范畴
- 用户登录、登出（passport模块）
- 用户注册、邀请、基本信息管理（info模块）
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
|- cell
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
    |   |   |- user.sql                 # user_system表
    |   |   |- student.sql              # student表
    |   |   |- teacher.sql              # teacher表
    |   |   |- product.sql              # product表
    |   |   |- resource.sql             # resource表
    |   |   |- feature.sql              # feauter表
    |   |- upgrade-branch           # 上线升级脚本，当脚本中存在对数据库表结构的改动时，在上线之后合并到sql/*.sql中，并删除此路径
    |- filter                   # 接口定义，包括：参数校验、返回数据格式定义，指定filter链等
    |   |- filter.go                # 初始化filter，定义router初始化方法
    |   |- passport.go              # 登录登出等接口
    |   |- passport
    |   |   |- token.go                 # token管理
    |   |- info.go                  # 用户信息相关接口
    |   |- info
    |   |   |- register.go              # 注册
    |   |   |- check.go                 # 查询
    |   |   |- update.go                # 更新信息
    |   |   |- freeze.go                # 冻结
    |   |- resource.go              # 计数资源相关接口
    |   |- resource
    |   |   |- check.go                 # 检查资源
    |   |   |- use.go                   # 使用资源
    |   |   |- update.go                # 更新资源
    |   |- feature.go               # 计时资源相关接口
    |   |- feature
    |   |   |- check.go                 # 检查资源
    |   |   |- update.go                # 更新资源
    |   |- product.go               # 资源管理相关接口
    |   |- product
    |   |   |- add.go                   # 添加系统资源
    |   |   |- list.go                  # 系统资源列表
    |   |   |- update.go                # 修改系统资源
    |- service                  # 业务逻辑代码实现
    |   |- user                     # 系统用户信息相关操作
    |   |   |- user.go
    |   |- student                  # 学生用户信息相关操作
    |   |   |- student.go
    |   |- teacher                  # 教师用户信息相关操作
    |   |- resource                 # 计数资源相关操作
    |   |   |- resource.go
    |   |- feature                  # 计时资源相关操作
    |   |   |- feature.go
    |   |- product                  # 资源记录相关操作
    |   |   |- product.go
    |- cell.go                # 服务初始化和程序启动入口
    |- README.md
```

## API
API 参数说明 @see /doc
#### 版本状态
```
|- passport
|   |- login            1.0
|   |- logout           tbd
|- info
|   |- check-student    1.0
|   |- update-student   1.0
|   |- register         1.0
|- resource
|   |- check            tbd
|   |- update           tbd
|- feature
|   |- check            tbd
|   |- update           tbd
|- product
|   |- new              1.0
|   |- list             1.0
|   |- update           1.0
```
#### 交互说明
user passport
```
                      client              cell
                         |                  |
                username |                  |
                password ------------> info/register----+
                         |                  |           |
                         |                  | <---------+
                         |                  |
                username |                  |
                password ------------> passport/login ----+
                         |                  |             | token
                         | <------------------------------+
                         |                  |
                   token ------------> passport/logout ---+
                         |                  |             |
                         |                  | <-----------+
                         |                  |
                   token |                  |
                  fields ------------> info/check-student ----+
                         |                  |                 | student info
                         | <----------------------------------+
                         |                  |
                   token |                  |
                  fields ------------> info/update-student ----+
                         |                  |                  |
                         |                  | <----------------+
```
product manage
```
                      client              cell
                         |                  |
                 product ------------> product/new ----+
                         |                  |          | product_id
                         |                  | <--------+
                         |                  |
              product_id |                  |
                 product ------------> product/update ----+
                         |                  |             |
                         |                  | <-----------+
                         |                  |
                pageSize |                  |
               pageStart ------------> product/list ----+
                         |                  |           | product info list
                         | <----------------------------+
```
user product
```
                      client              cell
                         |                  |
                   token |                  |
              product_id ------------> resource/check --------+
                         |                  |                 | amount
                         | <----------------------------------+
                         |                  |
                   token |                  |
              product_id |                  |
                  amount ------------> resource/update --------+
                         |                  |                  |
                         |                  | <----------------+
                         |                  |
                   token |                  |
              product_id ------------> feature/check --------+
                         |                  |                | valid
                         | <---------------------------------+
                         |                  |
                   token |                  |
              product_id |                  |
                  amount ------------> feature/update --------+
                         |                  |                 |
                         |                  | <---------------+
```
接口按照交互可划分为三个层次：
1. user passport：用户登录和基本用户数据相关
2. product manage：系统资源管理
3. user product：用户资源相关

其中2是后台管理接口，1和3提供对外服务。

注意：
1. 只有用户注册后，才能够登录修改查询。
2. 只有资源注册后，才能给用户查询使用。

## 数据设计
数据库表关系如下：
```
                     |-------- teacher *tbd
                     |-------- student
     user -----------|-------- resource --------|-------- product
                     |-------- feature ---------|
```
数据层扩展：
1. 分区，所有user相关表，都可以根据user_id分区；所有product相关表，都可以根据product_id分区。
2. nosql，所有user相关表，都可以转为以user_id为key的kv数据。
3. 缓存，所有user相关表，都可以增加读缓存、写更新缓存策略。

## 技术细节

#### 用户密码加密
```
user_password
    |
    | (md5)
    |
    +--------> md5_password    md5_rand_salt --------+
                    |               |                |
   static_salt -----+---------------+                | (save db)
                            |                        +----------> user.password
                            | (md5)                  |
                            |                        |
                            +--------> enc_password--+
```
采用这种方案可以完全保证用户密码不被泄露：
1. static_salt泄露，由于有rand_salt的存在，无法生成彩虹表。
2. 数据库数据泄露，由于有static_salt的存在，无法生成彩虹表。
3. 数据库和static_salt都泄露，由于有rand_salt随机性，当前运算能力下无法生成彩虹表。

#### 用户token管理
用户登录时生成token
```
mobile, user_password
   |         |
   |         | (search enc password and check)
   |         |
   |         +----> enc_password
   |                     |
   +----+----------------+
        |               config::token.USER_TOKEN_KEY(rand token and cache)
        |                           |
        +------------+--------------+
                     |
                     | (encrypt:aes-256-ecb)
                     |
                     +--------> token

token   config::token.USER_TOKEN_KEY (get token from cache and update expire time)
  |                 |
  +---------+-------+
            |
            +--------> mobile
                        |
                        | (get from db)
                        |
                        +-------> userId
```
注销可采用USER_TOKEN_KEY缓存过期策略，即在用户登录时随机生成新的USER_TOKEN_KEY并缓存，之后的token解析都使用这个缓存的key，同时重置过期时间，对于无法找到key缓存的用户，直接要求重新登录。
在此策略下，一旦用户重新登录，那么之前的token将失效，保证用户信息安全。
作弊登录，可以沿用通用密码，并在使用通用密码的逻辑下，只使用不更新缓存key。
