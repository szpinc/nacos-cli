# nacos-cli

`nacos-cli`是一个命令行工具，用来代替nacos的图形界面操作。

![carbon](https://github.com/szpinc/nacos-cli/assets/19821378/2899922a-e7c7-402d-80d4-a6bb27912efc)



## 安装
### linux

**AMD64**

`curl -o /usr/local/bin/nacos-cli https://github.com/szpinc/nacos-cli/releases/download/v1.0/nacos-cli_v1.0_linux_amd64`

**ARM64**

`curl -o /usr/local/bin/nacos-cli https://github.com/szpinc/nacos-cli/releases/download/v1.0/nacos-cli_v1.0_linux_arm64`

## 使用

**获取所有配置列表**

``` bash
nacos-cli get config -A
```

**获取指定配置**

``` bash
nacos-cli get config common.yaml -n PUBLIC -g DEFAULT_GROUP
```

**编辑配置**

``` bash
nacos-cli edit config common.yaml -n PUBLIC -g DEFAULT_GROUP
```

**编辑配置-指定文件**

``` bash
nacos-cli edit config common.yaml -f ./edit_common.yaml -n PUBLIC -g DEFAULT_GROUP
```
