# 📦 goe - 一行命令快速设置 Go 编译环境变量

`goe` 是一个轻量级命令行工具，用于便捷地设置 Go 的交叉编译环境变量：`GOOS` 和 `GOARCH`。

相比原生的 `go env -w` 命令，`goe` 提供了更快速直观的方式：

---

## ✨ 特性

- ✅ 简洁语法：`goe linux/amd64`
- ✅ 支持只设置系统或架构：`goe windows`、`goe arm64`
- ✅ 自动验证参数合法性
- ✅ 跨平台支持：macOS / Linux / Windows

---

## 📦 安装

确保你已经安装了 Go（版本 ≥ 1.16）：

```bash
go install github.com/你的用户名/goe@latest
```
## 🚀 使用示例
- ✅ 同时设置 GOOS 和 GOARCH
```bash
goe linux/amd64
goe windows/arm64
```
- ✅ 只设置 GOOS
```bash
goe darwin
goe linux
```

- ✅ 只设置 GOARCH
```bash
goe amd64
goe arm64
```
设置后会自动执行 go env 输出结果确认。

## ⚙️ 支持平台列表
**GOOS:** linux, windows, darwin, freebsd, android, ios, js, wasip1, plan9

**GOARCH:** amd64, arm64, 386, arm, ppc64, ppc64le, mips, mips64, wasm

## 📄 LICENSE
MIT License