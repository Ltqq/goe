package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

var validGOOS = map[string]bool{
	"linux": true, "windows": true, "darwin": true, "freebsd": true,
	"android": true, "ios": true, "plan9": true, "js": true, "wasip1": true,
}

var validGOARCH = map[string]bool{
	"amd64": true, "arm64": true, "386": true, "arm": true,
	"wasm": true, "ppc64": true, "ppc64le": true, "mips": true, "mips64": true,
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("用法: goe <GOOS>/<GOARCH> | <GOOS> | <GOARCH>")
		fmt.Println("示例: goe linux/amd64 或 goe windows 或 goe arm64")
		os.Exit(1)
	}

	input := os.Args[1]
	var goos, goarch string

	if strings.Contains(input, "/") {
		parts := strings.Split(input, "/")
		if len(parts) != 2 {
			fmt.Println("格式错误，应为 GOOS/GOARCH，例如 linux/amd64")
			os.Exit(1)
		}
		goos, goarch = parts[0], parts[1]
	} else {
		if validGOOS[input] {
			goos = input
		} else if validGOARCH[input] {
			goarch = input
		} else {
			fmt.Printf("无法识别的参数: %s\n", input)
			os.Exit(1)
		}
	}

	if goos != "" {
		fmt.Printf("设置 GOOS=%s\n", goos)
		if err := runCmd("go", "env", "-w", "GOOS="+goos); err != nil {
			fmt.Println("设置 GOOS 失败:", err)
			os.Exit(1)
		}
	}

	if goarch != "" {
		fmt.Printf("设置 GOARCH=%s\n", goarch)
		if err := runCmd("go", "env", "-w", "GOARCH="+goarch); err != nil {
			fmt.Println("设置 GOARCH 失败:", err)
			os.Exit(1)
		}
	}

	fmt.Println("当前 go env 设置如下：")
	_ = runCmd("go", "env")
}

func runCmd(name string, args ...string) error {
	cmd := exec.Command(name, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
