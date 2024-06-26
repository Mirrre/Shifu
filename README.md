# 笔试作业
【笔试题目】

* 在Kubernetes中运行Shifu并编写一个应用

【具体任务】

1. 请参考以下指南，部署并运行Shifu：https://shifu.dev/docs/tutorials/demo-install/

2. 运行一个酶标仪的数字孪生：https://shifu.dev/docs/tutorials/demo-try/#3-interact-with-the-microplate-reader

3. 编写一个Go应用
- 定期轮询获取酶标仪的/get_measurement接口，并将返回值平均后打印出来，轮询时间可自定义 
- Go的应用需要容器化
- Go的应用需要运行在Shifu的k8s集群当中
- 最终通过kubectl logs命令可以查看打印的值

# 详细过程
## 任务一
下载Shifu
![](https://github.com/Mirrre/Shifu/blob/main/%E6%AD%A5%E9%AA%A4%E6%88%AA%E5%9B%BE/1%E3%80%81%E4%B8%8B%E8%BD%BDShifu.png)
## 任务二
与酶标仪交互、创建数字孪生

![](https://github.com/Mirrre/Shifu/blob/main/%E6%AD%A5%E9%AA%A4%E6%88%AA%E5%9B%BE/2%E3%80%81%E4%B8%8E%E9%85%B6%E6%A0%87%E4%BB%AA%E4%BA%A4%E4%BA%92%E3%80%81%E5%88%9B%E5%BB%BA%E6%95%B0%E5%AD%97%E5%AD%AA%E7%94%9F.png)

与数字孪生交互
![](https://github.com/Mirrre/Shifu/blob/main/%E6%AD%A5%E9%AA%A4%E6%88%AA%E5%9B%BE/3%E3%80%81%E4%B8%8E%E6%95%B0%E5%AD%97%E5%AD%AA%E7%94%9F%E4%BA%A4%E4%BA%92.png)
## 任务三
### 定期轮询获取酶标仪的/get_measurement接口，并将返回值平均后打印出来，轮询时间可自定义

**具体代码见`main.go`文件**

### Go的应用需要容器化

1. 编写好`main.go`文件
2. 制作包含应用程序代码的`Docker`镜像，编写`Dockerfile`文件
3. `build`镜像，在当前目录下执行下列代码：
```shell
docker build -t go-app-img1 . 
```

4. 推送镜像到`DockerHub`，到时候`k8s`在部署应用的时候可以根据指定的镜像名称从`DockerHub`上拉去镜像，代码如下（lips0715是我本人的`DockerHub`账号）：

```shell
docker build -t lips0715/kube-go-app1 .
docker push lips0715/kube-go-app1
```
至此Go应用的容器化完成
### Go的应用需要运行在Shifu的k8s集群当中
1. 在Shifu的k8s集群中创建一个新的部署`deployment`，部署是一个高层次的抽象，它管理一组无状态的`Pods`，确保指定数量的`Pods`始终运行。`--image`就是指定了部署中`Pods`所使用的容器镜像，这里就是我刚刚推送上去的镜像，执行代码：
```shell
sudo kubectl create deployment my-deployment --image=lips0715/kube-go-app1
```
2. 获取所有命名空间下的`Pods`信息，执行代码：
```shell
sudo kubectl get pods -A
```
![](https://github.com/Mirrre/Shifu/blob/main/%E6%AD%A5%E9%AA%A4%E6%88%AA%E5%9B%BE/4%E3%80%81%E8%8E%B7%E5%8F%96%E6%89%80%E6%9C%89%E5%91%BD%E5%90%8D%E7%A9%BA%E9%97%B4%E4%B8%8B%E7%9A%84Pods%E4%BF%A1%E6%81%AF.png)
`my-deployment-8c6f6867c-x9zj7`就是我刚刚运行的`Pod`
### 最终通过kubectl logs命令可以查看打印的值
```shell
sudo kubectl logs my-deployment-8c6f6867c-x9zj7
```
![](https://github.com/Mirrre/Shifu/blob/main/%E6%AD%A5%E9%AA%A4%E6%88%AA%E5%9B%BE/5%E3%80%81%E6%9C%80%E7%BB%88%E5%AE%8C%E6%88%90%E6%95%88%E6%9E%9C.png)


以上就是我本次笔试作业的所有内容，感谢您的阅读！