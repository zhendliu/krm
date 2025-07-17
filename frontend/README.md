# 技术债
````
1. 集群规模未实现
2. 集群详情未实现
3. 节点详情查看Pod列表未实现
4. 节点驱逐和禁止调度未实现
````

# 基本组件安装
````
cnpm install vue-router@4 --save
cnpm install pinia axios element-plus --save
````

# 测试接口
````
用户列表: https://www.fastmock.site/mock/7ac0e652d419254dab9b4d2c33fb820a/vueapi/user/list
用户详情：https://www.fastmock.site/mock/7ac0e652d419254dab9b4d2c33fb820a/vueapi/user/get
删除用户：https://www.fastmock.site/mock/7ac0e652d419254dab9b4d2c33fb820a/vueapi/user/delete
添加用户：https://www.fastmock.site/mock/7ac0e652d419254dab9b4d2c33fb820a/vueapi/user/add
更新用户：https://www.fastmock.site/mock/7ac0e652d419254dab9b4d2c33fb820a/vueapi/user/update

登录：https://www.fastmock.site/mock/7ac0e652d419254dab9b4d2c33fb820a/vueapi/auth/login
退出：https://www.fastmock.site/mock/7ac0e652d419254dab9b4d2c33fb820a/vueapi/auth/logout

````

# 部署
## 如何部署
````
1. 编译 --> dist
2. NGINX --> 工作目录 --> /usr/share/nginx/html, /var/www/html
3. COPY dist/ /usr/share/nginx/html/ --> front-scaffold:v1
4. push 到镜像仓库
5. docker k8s
6. a.com path: / --> front
   a.com path: /api --> backend
````
## 安装Docker
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
## 部署
1. 编译
npm run build

2. 打包镜像
````
vim Dockerfile
FROM registry.cn-beijing.aliyuncs.com/dotbalo/nginx:1.22.1-alpine3.17
COPY ./dist/ /usr/share/nginx/html

docker build -t registry.cn-beijing.aliyuncs.com/dotbalo/front-scaffold:v1 .

docker push registry.cn-beijing.aliyuncs.com/dotbalo/front-scaffold:v1
````

3. Docker部署
````
docker run -tid -p 8081:80 registry.cn-beijing.aliyuncs.com/dotbalo/front-scaffold:v1
````

4. K8s部署
````
vim deploy-svc.yaml
---
apiVersion: v1
kind: Service
metadata:
  labels:
    app: front
  name: front
spec:
  ports:
  - name: http
    port: 80
    protocol: TCP
    targetPort: 80
  selector:
    app: front
  sessionAffinity: None
  type: NodePort
---
apiVersion: apps/v1
kind: Deployment
metadata:
  labels:
    app: front
  name: front
spec:
  replicas: 1
  selector:
    matchLabels:
      app: front
  strategy:
    rollingUpdate:
      maxSurge: 1
      maxUnavailable: 0
    type: RollingUpdate
  template:
    metadata:
      labels:
        app: front
    spec:
      containers:
      - env:
        - name: TZ
          value: Asia/Shanghai
        - name: LANG
          value: C.UTF-8
        image: registry.cn-beijing.aliyuncs.com/dotbalo/front-scaffold:v1
        lifecycle: {}
        livenessProbe:
          failureThreshold: 2
          initialDelaySeconds: 10
          periodSeconds: 10
          successThreshold: 1
          tcpSocket:
            port: 80
          timeoutSeconds: 2
        name: front
        ports:
        - containerPort: 80
          name: web
          protocol: TCP
        readinessProbe:
          failureThreshold: 2
          initialDelaySeconds: 10
          periodSeconds: 10
          successThreshold: 1
          tcpSocket:
            port: 80
          timeoutSeconds: 2
        resources:
          limits:
            cpu: 1
            memory: 512Mi
          requests:
            cpu: 100m
            memory: 256Mi
      restartPolicy: Always
kubectl apply -f deploy-svc.yaml
````

