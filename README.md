Cell
====

用户认证管理服务


使用方法
----

### 使用官方服务
https://cell.mapleque.com

###  使用Docker

### 下载发布包

### 自行编译安装

环境要求：
- go 1.11.2 darwin/amd64
- node 10.15.0

```sh
git clone https://github.com/mapleque/cell.git
cd cell
make install
cd bin
# edit .env for config
./server
```

### 开发环境调试

```sh
git clone https://github.com/mapleque/cell.git
cd cell
cp config/.env.example main/.env
# edit main/.env for config
make run
```
