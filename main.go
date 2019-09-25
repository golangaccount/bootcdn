package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/golangaccount/bootcdn/download"
	"github.com/urfave/cli"
)

func main() {
	if len(os.Args) > 1 {
		var app = cli.App{}
		app.Name = `从bootcdn下载对应框架的版本
		example:bootcdn -n jquery -v 1.4 - p jquery -t 0
		bootcdn jquey 1.4 jquey 0`

		//app.Usage = "\n  \n  "
		app.ArgsUsage = "-n {name} -v {version} -p {path} t {type 0:重新下载 1:已经下载就跳过}"
		app.HideVersion = true
		app.Flags = []cli.Flag{
			cli.StringFlag{Name: "n"},
			cli.StringFlag{Name: "v"},
			cli.StringFlag{Name: "p"},
			cli.StringFlag{Name: "t"},
		}
		app.Action = func(c *cli.Context) error {
			var name, version, path string
			var ty int
			name = c.String("n")
			version = c.String("v")
			path = c.String("p")
			ty = c.Int("t")
			if ty != 0 && ty != 1 {
				fmt.Fprintln(c.App.Writer, "type类型错误，请输入0或1 默认0")
			}
			if name != "" && version != "" && path != "" {
				err := download.Download(name, version, path, ty)
				if err != nil {
					fmt.Fprintln(c.App.Writer, err.Error())
				} else {
					fmt.Fprintln(c.App.Writer, "success")
				}
			} else if name == "" && version == "" && path == "" {
				var parms = c.Args()
				if len(parms) < 3 {
					fmt.Fprintln(c.App.Writer, "参数错误，请参照：bootcdn {name} {version} {path} [{type}]")
				} else if len(parms) >= 3 {
					name = parms[0]
					version = parms[1]
					path = parms[2]
					if len(parms) > 3 {
						ty, _ = strconv.Atoi(parms[3])
					}
				}
				err := download.Download(name, version, path, ty)
				if err != nil {
					fmt.Fprintln(c.App.Writer, err.Error())
				} else {
					fmt.Fprintln(c.App.Writer, "success")
				}
			} else {
				fmt.Fprintln(c.App.Writer, "参数错误，请参照：bootcdn -n name -v version -p path -t type ")
			}
			return nil
		}
		app.Run(os.Args)
	} else {
		var name, version, path string
		var ty int
		fmt.Print("请输入名称:")
		fmt.Scanln(&name)
		fmt.Print("请输入版本号:")
		fmt.Scanln(&version)
		fmt.Print("请输入储存路径:")
		fmt.Scanln(&path)
		fmt.Print("请输入保存类型:")
		fmt.Scanln(&ty)
		err := download.Download(name, version, path, ty)
		if err != nil {
			fmt.Println(err.Error())
		} else {
			fmt.Println("success")
		}
	}
}
