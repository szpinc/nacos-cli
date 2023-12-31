# nacos-cli

`nacos-cli`是一个命令行工具，用来代替nacos的图形界面操作。

![carbon](https://github.com/szpinc/nacos-cli/assets/19821378/2899922a-e7c7-402d-80d4-a6bb27912efc)



## 安装
### linux

**AMD64**

`curl -L -o /usr/local/bin/nacos-cli https://github.com/szpinc/nacos-cli/releases/download/v1.2/nacos-cli_linux_amd64`

**ARM64**

`curl -L -o /usr/local/bin/nacos-cli https://github.com/szpinc/nacos-cli/releases/download/v1.2/nacos-cli_linux_arm64`

## 使用

**环境变量**

export NACOS_ADDR="http://127.0.0.1:8848/nacos"

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

**从文件更新配置**

``` bash
nacos-cli apply -f common.yaml -n public -g DEFAULT_GROUP --id common.yaml
```
