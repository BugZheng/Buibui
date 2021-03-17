/**
 * @Author: BugZheng
 * @Description:工具项目
 * @File:  main.go
 * @Version: 1.0.0
 * @Date: 2021/03/17 01:34 上午
 */
package main

import (
	"Buibui/cmd"
	"fmt"
	"github.com/urfave/cli/v2"
	"os"
)

func main() {
	app := cli.NewApp()
	app.Name = "E-tools"
	app.Usage = "一切为了简化开发(偷懒,摸鱼)！ 建设性懒惰是程序员的基本美德之一"
	app.Commands = cmd.Commands
	app.Action = func(c *cli.Context) error {
		fmt.Println("Buibui~工具项目已经启动V1.0")
		return nil
	}
	if err := app.Run(os.Args); err != nil {
		fmt.Println("app running error", err)
	}

}
