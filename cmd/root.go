package cmd

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

var (
	osFlag   string
	archFlag string
)

var rootCmd = &cobra.Command{
	Use:   "goenv-setter",
	Short: "设置当前终端 go env 中的 GOOS 和 GOARCH 环境变量",
	Run: func(cmd *cobra.Command, args []string) {
		if osFlag == "" && archFlag == "" {
			fmt.Println("请至少传入 --os 或 --arch 参数")
			return
		}

		// 读取当前的环境变量（原样保留）
		env := os.Environ()

		if osFlag != "" {
			fmt.Printf("设置 GOOS=%s\n", osFlag)
			env = append(env, "GOOS="+osFlag)
		}
		if archFlag != "" {
			fmt.Printf("设置 GOARCH=%s\n", archFlag)
			env = append(env, "GOARCH="+archFlag)
		}

		// 验证设置是否成功
		checkCmd := exec.Command("go", "env")
		checkCmd.Env = env
		checkCmd.Stdout = os.Stdout
		checkCmd.Stderr = os.Stderr

		if err := checkCmd.Run(); err != nil {
			fmt.Printf("设置 go env 失败: %v\n", err)
			os.Exit(1)
		}
	},
}

func Execute() {
	rootCmd.PersistentFlags().StringVarP(&osFlag, "os", "o", "", "目标 GOOS，例如 linux, windows, darwin")
	rootCmd.PersistentFlags().StringVarP(&archFlag, "arch", "a", "", "目标 GOARCH，例如 amd64, arm64")
	if err := rootCmd.Execute(); err != nil {
		fmt.Println("执行失败:", err)
		os.Exit(1)
	}
}
