# nacos-cli

`nacos-cli`是一个命令行工具，用来代替nacos的图形界面操作。

## 安装
### linux

`curl -o /usr/local/bin/nacos-cli https://github.com/szpinc/nacos-cli/releases/download/v1.0/nacos-cli_linux_amd64`

## 使用

**获取所有配置列表**

`nacos-cli get config -A`

**获取指定配置**

`nacos-cli get config common.yaml -n PUBLIC -g DEFAULT_GROUP`

**编辑配置**

`nacos-cli edit config common.yaml -n PUBLIC -g DEFAULT_GROUP`

**编辑配置-指定文件**

`nacos-cli edit config common.yaml -f ./edit_common.yaml -n PUBLIC -g DEFAULT_GROUP`
