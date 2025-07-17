## 用来描述项目的一些信息
````
这是一个脚手架项目，可以根据这个项目去生成一个基础的框架
````
## Go语言服务部署
### 如何部署
1. 拿到我们的代码
2. 编译成一个二进制包
3. ./二进制包  --> 镜像
4. 镜像推送到镜像仓库
5. docker run
6. yaml文件
### 1.1 使用Docker部署
#### Docker安装
````
关闭防火墙：
systemctl disable --now firewalld 
systemctl disable --now dnsmasq
setenforce 0
sed -i 's#SELINUX=enforcing#SELINUX=disabled#g' /etc/sysconfig/selinux
sed -i 's#SELINUX=enforcing#SELINUX=disabled#g' /etc/selinux/config
关闭swap分区，fstab注释swap：
swapoff -a && sysctl -w vm.swappiness=0
sed -ri '/^[^#]*swap/s@^@#@' /etc/fstab

curl -o /etc/yum.repos.d/CentOS-Base.repo https://mirrors.aliyun.com/repo/Centos-7.repo
yum install -y lrzsz yum-utils device-mapper-persistent-data lvm2
yum-config-manager --add-repo https://mirrors.aliyun.com/docker-ce/linux/centos/docker-ce.repo
sed -i -e '/mirrors.cloud.aliyuncs.com/d' -e '/mirrors.aliyuncs.com/d' /etc/yum.repos.d/CentOS-Base.repo
配置内核参数：
cat <<EOF > /etc/sysctl.d/docker.conf
net.ipv4.ip_forward = 1
net.bridge.bridge-nf-call-iptables = 1
net.bridge.bridge-nf-call-ip6tables = 1
EOF
sysctl --system
安装Docker：
yum install docker-ce-20.10.* docker-ce-cli-20.10.* containerd.io -y
启动Docker：
systemctl enable --now docker
````

#### 代码编译
````
cmd
SET CGO_ENABLED=0
SET GOOS=linux
SET GOARCH=amd64
go build -o demo main.go
SET CGO_ENABLED=1
SET GOOS=windows
SET GOARCH=amd64
上传二进制包至服务器
chmod +x demo
````

#### 制作镜像
````
# cat Dockerfile 
FROM registry.cn-beijing.aliyuncs.com/dotbalo/alpine:3.9-tomcat
COPY demo ./
ENTRYPOINT ./demo
# docker build -t registry.cn-beijing.aliyuncs.com/dotbalo/demo:v1 .
启动测试
docker run -ti --rm -p 8080:8080 registry.cn-beijing.aliyuncs.com/dotbalo/demo:v1
````

#### 部署启动
````
docker run -tid -p 8080:8080 -e LOG_LEVEL=info -e USERNAME=34C7357D6D1C641D4CA1E369E7244F61 -e PASSWORD=A1884CD37D07E3CC64AE2299CCD4F597 -e GIN_MODE=release registry.cn-beijing.aliyuncs.com/dotbalo/demo:v1
查看日志
docker logs -f xxx
使用postman测试
````

#### 推送镜像
````
docker push registry.cn-beijing.aliyuncs.com/dotbalo/demo:v1
````

### 1.2 使用K8s部署
````
vim deploy-svc.yaml
---
apiVersion: v1
kind: Service
metadata:
  labels:
    app: demo
  name: demo
spec:
  ports:
  - name: http
    port: 8080
    protocol: TCP
    targetPort: 8080
  selector:
    app: demo
  sessionAffinity: None
  type: NodePort
---
apiVersion: apps/v1
kind: Deployment
metadata:
  labels:
    app: demo
  name: demo
spec:
  replicas: 1
  selector:
    matchLabels:
      app: demo
  strategy:
    rollingUpdate:
      maxSurge: 1
      maxUnavailable: 0
    type: RollingUpdate
  template:
    metadata:
      labels:
        app: demo
    spec:
      containers:
      - env:
        - name: TZ
          value: Asia/Shanghai
        - name: LANG
          value: C.UTF-8
        - name: GIN_MODE
          value: release
        - name: LOG_LEVEL
          value: info
        - name: USERNAME
          value: 34C7357D6D1C641D4CA1E369E7244F61
        - name: PASSWORD
          value: A1884CD37D07E3CC64AE2299CCD4F597
        image: registry.cn-beijing.aliyuncs.com/dotbalo/demo:v2
        lifecycle: {}
        livenessProbe:
          failureThreshold: 2
          initialDelaySeconds: 10
          periodSeconds: 10
          successThreshold: 1
          tcpSocket:
            port: 8080
          timeoutSeconds: 2
        name: demo
        ports:
        - containerPort: 8080
          name: web
          protocol: TCP
        readinessProbe:
          failureThreshold: 2
          initialDelaySeconds: 10
          periodSeconds: 10
          successThreshold: 1
          tcpSocket:
            port: 8080
          timeoutSeconds: 2
        resources:
          limits:
            cpu: 1
            memory: 1024Mi
          requests:
            cpu: 200m
            memory: 256Mi
      restartPolicy: Always
kubectl apply -f deploy-svc.yaml
````