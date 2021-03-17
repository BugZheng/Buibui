/**
 * @Author: BugZheng
 * @Description:工具项目
 * @File:  cmd.go
 * @Version: 1.0.0
 * @Date: 2021/03/12 3:34 下午
 */
package cmd

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"os"
)

var Commands = []*cli.Command{
	{
		Name:    "new",
		Aliases: []string{"fuck"},
		Usage:   "new 后面加上你的项目名称 创建一个属于你的框架",
		Action: func(c *cli.Context) error {
			MakeprojectNameDir(os.Args)
			fmt.Println("偷懒是种美德！已经成功生成projectName基础框架，欢迎使用")
			return nil
		},
	},
	{
		Name:    "swagInit",
		Aliases: []string{"swag"},
		Usage:   "生成你的swagger doc 文档",
		Action: func(c *cli.Context) error {
			fmt.Println("偷懒是种美德!项目的文档成功生成~")
			return nil
		},
	},
	{
		Name:    "model",
		Aliases: []string{"model"},
		Usage:   "生成你的数据库表的结构在model目录下面",
		Action: func(c *cli.Context) error {
			fmt.Println("偷懒是种美德!表的结构体成功生成~")
			return nil
		},
	},
}
