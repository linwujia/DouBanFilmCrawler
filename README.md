# DouBanFilmCrawler
豆瓣电影Top250 爬虫抓取, 从豆瓣网中爬取电影Top250排行, 每页25条数据, 页数从第一页开始,最大页数支持10页， 抓取后每页数据保存到目录下以`page [页数].json`为名的json文件中

支持的操作系统：
- Windows
- Linux
- 其他 [golang](https://go.dev/) 编程语言支持的系统

提示：
- 豆瓣电影Top250 爬虫抓取， 只是本人学习go编程的例子，如需使用到项目中，自己酌情使用，本人不对结果负责。

## 下载

你可以在此处下载 DouBanFilmCrawler 的可执行文件：

https://github.com/linwujia/DouBanFilmCrawler/releases

下载适用于你系统的`DouBanFilmCrawler-xxx-xxx`可执行文件，然后给`DouBanFilmCrawler-xxx-xxx`执行权限（Linux/macOS需要）并重命名为`DouBanFilmCrawler`。

给执行权限和重命名示例：
```bash
chmod +x DouBanFilmCrawler-linux-x64
mv DouBanFilmCrawler-linux-x64 DouBanFilmCrawler
```

我该下载哪个可执行文件？
* 32位Windows系统：`DouBanFilmCrawler-windows-x86.exe`
* 64位Windows系统：`DouBanFilmCrawler-windows-x64.exe`
* 32位Linux系统：`DouBanFilmCrawler-linux-x86`
* 64位Linux系统：`DouBanFilmCrawler-linux-x64`
* 运行32位系统的树莓派：`DouBanFilmCrawler-linux-arm`
* 运行64位系统的树莓派：`DouBanFilmCrawler-linux-arm64`
* 英特尔CPU的Mac：`DouBanFilmCrawler-macos-x64`
* M1芯片的Mac：`DouBanFilmCrawler-macos-arm64`

## 运行

在包含`DouBanFilmCrawler`的文件夹中运行以下命令。

如果你使用图形界面，可以在文件管理器空白处右击，选择“在此处打开终端”。

```bash
# 创建日志文件夹
mkdir log

# 启动BTCAgent
# -s 抓取的起始页  -e 抓取的结束页（最大为10） -s <= -e
DouBanFilmCrawler -s 1 -e 10 -l log -alsologtostderr
```

按 Ctrl + C 可停止`DouBanFilmCrawler`。

## 编译安装

### 源码编译

适用于开发者。

如果能[下载](https://github.com/linwujia/DouBanFilmCrawler/releases) 到适合自己系统的可执行文件，就不需要编译安装。

1. 从 https://go.dev/ 安装 golang

2. 从 https://git-scm.com/ 安装 git

3. 运行以下命令:
   ```bash
   git clone https://github.com/linwujia/DouBanFilmCrawler.git
   cd DouBanFilmCrawler
   go build
   ```

4. 然后就能得到可执行文件`DouBanFilmCrawler`（Windows中为`DouBanFilmCrawler.exe`）。

#### 在 Windows 中编译

建议在 [WSL](https://aka.ms/wsl) 中通过 [./build-all.sh](./build-all.sh) 编译。

### docker编译
