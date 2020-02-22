## 1 如何安装 `oann` 命令行工具

### 现阶段安装方式

1. 安装go，go module管理，要求go版本至少go 1.11+

2. 手动安装protoc以及protoc-gen-go

3. 安装该oann命令行程序

    由于 `oann` 除了可执行程序之外，还依赖不同语言的模板文件、oann本身配置文件中的信息，`go install` 只安装oann命令行工具是无法正常工作的。

    现阶段的安装方式：

    ```bash
    git clone https://github.com/oann-go-cmdline
    cd oann-go-cmdline
    make && make install