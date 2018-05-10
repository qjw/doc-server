![logo](images/1.png)

现在越来越流行[swagger](https://swagger.io/)，配合[swagger ui](https://github.com/swagger-api/swagger-ui)，对于前后端沟通接口，以及后续查阅，新人学习都有很大的帮助，不过也存在一些问题

1. spec.json管理不方便，存储不集中，访问也分散
2. 对于spec.json自动生成的场景，查阅还需要启动后台服务，例如<https://github.com/qjw/flasgger>

doc-server解决

1. 统一入口，查阅所有文档
2. 通过企业微信做授权访问，避免核心API泄密
3. 不依赖持久化DB，所有的spec.json放在私有的gitlab/github仓库，并支持查看各个历史版本

体验环境 <http://docserver.t.self.kim>

# spec
doc-server 不依赖mysql等持久化的数据库存储配置和文档，而是从一个git仓库中获取所有信息，对于这个git仓库，有一些要求

1. 根目录必须存在spec.json
2. spec.json中的路径，基于根目录

## spec.json
``` json
{
    "title":"BAT接口文档中心",
    "description": "后端接口文档",
    "company": "BAT宇宙无限公司",
    "groups": [
        {
            "name": "BB产品线",
            "description": "AA做的牛BB产品TT",
            "projects": [
                {
                    "name": "wxqq",
                    "description": "你懂的",
                    "spec": "bb/wxqq.json"
                },
                {
                    "name": "wxbd",
                    "description": "你懂的2",
                    "spec": "bb/wxbd.json"
                }
            ]
        },
        {
            "name": "TT产品线",
            "description": "BB做的牛TT产品AA",
            "projects": [
                {
                    "name": "apay",
                    "description": "你懂的3",
                    "spec": "aa/pay.json"
                }
            ]
        }
    ]
}
```

1. 整体分成两级， **产品线** 和 **项目**
2. 项目需要指定它的spec，参见swagger的spec.json

git仓库的目录结构如下
``` bash
├── bb
│   └── wxqq.json
│   └── wxbd.json
├── aa
│   └── pay.json
└── spec.json
```

# 编译前端
``` bash
cd frontend
npm i
# 在frontend/build目录生成前端目标文件
npm run-script build

# 直接运行
npm start
```
后端默认启用了跨域，所以前后端并不需要在同一个域下，注意`.roadhogrc`下代理配置

*gateway用于github/gitlab webhook回调，以及可能的（企业）微信回调*

``` json
{
  "proxy": {
     "/api/v1": {
       "target": "http://127.0.0.1:8888",
       "changeOrigin": true,
       "pathRewrite": { "^/api/v1" : "/api/v1" }
     },
     "/gateway": {
       "target": "http://127.0.0.1:8888",
       "changeOrigin": true,
       "pathRewrite": { "^/gateway" : "/gateway" }
     },
     "/swagger": {
       "target": "http://127.0.0.1:8888",
       "changeOrigin": true,
       "pathRewrite": { "^/swagger" : "/swagger" }
     }
  }
}
```

# 编译Go
``` bash
go build .

# 静态编译
make
```

# 运行
## 环境变量
1. LOCAL_DIR ： 从git仓库clone下来存放的路径
2. GIT_ORIGIN ： git远程地址,可以为空，若为空，就直接从`$LOCAL_DIR`读取（*务必确保`$LOCAL_DIR`是一个合法的git仓库地址*）
3. SWAGGER_UI : swagger ui前端的本地路径
4. FRONTEND ： 前端资源的本地路径，参见`编译前端`
5. PORT : 绑定的端口
6. REDIS_URL : redis主机地址（或者ip）默认"redis://localhost:6379/1"，若指定密码，使用格式`redis://:password@localhost:6379/1`

## 私有仓库
对于私有仓库需要授权访问，gitlab可以新建一个仅有存储doc spec的git仓库读权限的账户，然后在该账户下新建`personal_access_tokens`

访问格式如下<https://oauth:personal_access_tokens@gitlab.example.com/king/doc-server.git>

## 直接编译
设置好环境变量之后
``` bash
./doc-server

# 或者
FRONTEND=frontend/dist/ SWAGGER_UI=swaggerui ./doc-server

# 直接从本地读取，没必要每次都clone
GIT_ORIGIN= ./doc-server
```

## Docker运行
``` bash
# 静态编译
make
# 创建docker镜像
docker build -t doc_server:0.1 .
# 删除重复的
docker images | grep "^<none>" | awk '{print $3}' | xargs -i docker rmi {}
# 直接测试运行
docker run --rm -it -e REDIS_URL=redis://redis_host:6379/1 -p 8888:8888 doc_server:0.1

# push
docker login
docker tag doc_server:0.1 qiujinwu/doc-server:0.1
docker push qiujinwu/docl-server:0.1
```

docker镜像见<https://hub.docker.com/r/qiujinwu/doc-server/>

关于kubernet配置，参见[这里](Kubernete.md)


# 企业号登陆
## 环境变量
1. CORP_ID: 企业号corpid
2. CORP_AGENT_SECRET: 企业号应用secret
3. CORP_AGENT_ID: 企业号应用ID

## Gitlab/Github hook
1. GITLAB_TOKEN: 回调的Token，用于验证来源。回调：<http://YOUR_SITE/gateway/gitlab/hook>
1. GITHUB_TOKEN: 回调的Token，用于验证来源，Content type选择application/json。回调：<http://YOUR_SITE/gateway/github/hook>


手动pull最新代码
``` bash
kill -s SIGUSR2 `pidof doc-server`
```
