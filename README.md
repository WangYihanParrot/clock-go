# Clock-go

**Clock-go** 是一个使用 Go 开发的网页应用，主要用于显示当前时间。它特别适合在需要长时间查看时钟或进行专注任务时使用，因为它内置了防止设备自动息屏的功能。

## 特性

- 实时显示当前时间
- 防止设备息屏（通过集成 [NoSleep.js](https://github.com/richtr/NoSleep.js/tree/master) 实现）
- 轻量级、部署简单

防息屏是通过播放一个静音视频来实现的，这样可以保持屏幕持续亮着，不会自动变暗或进入休眠状态。

示例网站：  
[https://clock.parrotstudio.xyz/](https://clock.parrotstudio.xyz/)

## 使用教程

以 Linux 为例，使用步骤如下：

1. 下载最新版本的可执行文件（例如 `clock-go_v1.1_linux_amd64`）。
2. 下载网页所需的 HTML 文件，并确保它与可执行文件位于**同一目录**下。
3. 运行以下命令，启动服务：

```bash
./clock-go_v1.1_linux_amd64 -port 8080
```

上面这行的含义是：使用 `8080` 端口来监听访问。你可以替换为你想要监听的其他端口。

## 编译

```bash
./build.ps1
````

欢迎体验并提出建议！
