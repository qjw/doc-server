
将配置置于一个secret中，并mount到pod的环境变量中

``` bash
kubectl create secret generic -n tt doc-server \
	--from-literal=GITLAB_TOKEN=123456798123456798 \
	--from-literal=GITHUB_TOKEN=123456798123456798 \
	--from-literal=GIT_ORIGIN="https://oauth:token@gitlab.example.com/king/doc-server.git" \
	--from-literal=LOCAL_DIR=/tmp/data \
	--from-literal=REDIS_URL= \
	--from-literal=CORP_ID= \
	--from-literal=CORP_AGENT_SECRET= \
	--from-literal=CORP_AGENT_ID=
```

新建[部署],包含两个docker镜像的pod

``` yaml
apiVersion: extensions/v1beta1
kind: Deployment
metadata:
  name: doc-server
  namespace: tt
  labels:
    app: doc-server
spec:
  replicas: 1
  template:
    metadata:
      labels:
        app: doc-server
    spec:
      volumes:
      - name: config
        secret:
         secretName: doc-server
      - name: data
        emptyDir: {}
      containers:
      - name: doc-server
        image: qiujinwu/doc-server:0.1
        ports:
          - containerPort: 8888
            name: port-8888
        volumeMounts:
          - mountPath: /tmp/data
            name: data
        env:
          - name: GIT_ORIGIN
            valueFrom:
             secretKeyRef:
              name: doc-server 
              key: GIT_ORIGIN
          - name: LOCAL_DIR
            valueFrom:
             secretKeyRef:
              name: doc-server 
              key: LOCAL_DIR
          - name: GITLAB_TOKEN
            valueFrom:
             secretKeyRef:
              name: doc-server 
              key: GITLAB_TOKEN
          - name: GITHUB_TOKEN
            valueFrom:
             secretKeyRef:
              name: doc-server 
              key: GITHUB_TOKEN
          - name: REDIS_URL
            valueFrom:
             secretKeyRef:
              name: doc-server 
              key: REDIS_URL
          - name: CORP_ID
            valueFrom:
             secretKeyRef:
              name: doc-server 
              key: CORP_ID
          - name: CORP_AGENT_SECRET
            valueFrom:
             secretKeyRef:
              name: doc-server 
              key: CORP_AGENT_SECRET
          - name: CORP_AGENT_ID
            valueFrom:
             secretKeyRef:
              name: doc-server 
              key: CORP_AGENT_ID
      - name: redis
        image: redis:4
        ports:
          - containerPort: 6379
            name: port-6379

```

# 关于多副本的问题
在多副本的情况下，考虑到github/gitlab回调，通过nginx反向代理，或者kubernete的service，一般只有一台收到请求，这会导致不同的副本的数据不一致

解决这个问题的办法是，通过一个临时的目录[`volumes/emptyDir`]来共享多副本的数据