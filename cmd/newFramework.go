/**
 * @Author: BugZheng
 * @Description: 生成projectName项目的基础框架
 * @File:  projectName
 * @Version: 1.0.0
 * @Date: 2021/03/12 3:42 下午
 */
package cmd

import (
	"Buibui/asset"
	"fmt"
	"html/template"
	"io"
	"os"
	"strings"
)

func MakeprojectNameDir(params []string) {
	var f *os.File
	defer f.Close()
	projectName := "JoGo"
	if len(params) >= 3 {
		projectName = params[2]
	}
	os.MkdirAll(projectName+"/app/api", os.FileMode(0777))
	os.MkdirAll(projectName+"/app/model", os.FileMode(0777))
	os.MkdirAll(projectName+"/app/service", os.FileMode(0777))
	os.MkdirAll(projectName+"/boot", os.FileMode(0777))
	os.MkdirAll(projectName+"/config", os.FileMode(0777))
	os.MkdirAll(projectName+"/docs", os.FileMode(0777))
	os.MkdirAll(projectName+"/pkg/cmd", os.FileMode(0777))
	os.MkdirAll(projectName+"/pkg/middleware", os.FileMode(0777))
	os.MkdirAll(projectName+"/pkg/conf", os.FileMode(0777))
	os.MkdirAll(projectName+"/pkg/serializer", os.FileMode(0777))
	os.MkdirAll(projectName+"/route", os.FileMode(0777))
	os.MkdirAll(projectName+"/swaggers", os.FileMode(0777))
	/***************************** 生成Api的代码 ***********************************************/
	if checkFileIsExist(projectName + "/app/api/demo.go") { //如果文件存在
		f, _ = os.OpenFile(projectName+"/app/api/demo.go", os.O_APPEND, 0777) //打开文件
	} else {
		f, _ = os.Create(projectName + "/app/api/demo.go") //创建文件
	}
	apiTml, _ := asset.Asset("template/api.tml")
	io.WriteString(f, strings.Replace(string(apiTml), "ProjectName", projectName, -1)) //写入文件(字符串)
	/***************************** 生成service的代码 ***********************************************/
	if checkFileIsExist(projectName + "/app/service/demo.go") { //如果文件存在
		f, _ = os.OpenFile(projectName+"/app/service/demo.go", os.O_APPEND, 0777) //打开文件
	} else {
		f, _ = os.Create(projectName + "/app/service/demo.go") //创建文件
	}
	serviceTml, _ := asset.Asset("template/serviceDemo.tml")
	io.WriteString(f, strings.Replace(string(serviceTml), "ProjectName", projectName, -1)) //写入文件(字符串)
	/***************************** 生成Model的代码 ***********************************************/
	if checkFileIsExist(projectName + "/app/model/user.go") { //如果文件存在
		f, _ = os.OpenFile(projectName+"/app/model/user.go", os.O_APPEND, 0777) //打开文件
	} else {
		f, _ = os.Create(projectName + "/app/model/user.go") //创建文件
	}
	modelTml, _ := asset.Asset("template/userModel.tml")
	io.WriteString(f, strings.Replace(string(modelTml), "ProjectName", projectName, -1)) //写入文件(字符串)
	/***************************** 生成Boot的代码 ***********************************************/
	if checkFileIsExist(projectName + "/boot/user.go") { //如果文件存在
		f, _ = os.OpenFile(projectName+"/boot/boot.go", os.O_APPEND, 0777) //打开文件
	} else {
		f, _ = os.Create(projectName + "/boot/boot.go") //创建文件
	}
	bootTml, _ := asset.Asset("template/boot.tml")
	io.WriteString(f, strings.Replace(string(bootTml), "ProjectName", projectName, -1)) //写入文件(字符串)
	/***************************** 生成Conf的代码 ***********************************************/
	if checkFileIsExist(projectName + "/config/conf.go") { //如果文件存在
		f, _ = os.OpenFile(projectName+"/config/conf.go", os.O_APPEND, 0777) //打开文件
	} else {
		f, _ = os.Create(projectName + "/config/conf.go") //创建文件
	}
	configTml, _ := asset.Asset("template/conf.tml")
	io.WriteString(f, strings.Replace(string(configTml), "ProjectName", projectName, -1)) //写入文件(字符串)
	/***************************** 生成pkg/cmd的代码 ***********************************************/
	if checkFileIsExist(projectName + "/pkg/cmd/cmd.go") { //如果文件存在
		f, _ = os.OpenFile(projectName+"/pkg/cmd/cmd.go", os.O_APPEND, 0777) //打开文件
	} else {
		f, _ = os.Create(projectName + "/pkg/cmd/cmd.go") //创建文件
	}
	cmdTml, _ := asset.Asset("template/command.tml")
	io.WriteString(f, strings.Replace(string(cmdTml), "ProjectName", projectName, -1)) //写入文件(字符串)
	/***************************** 生成pkg/conf的代码 ***********************************************/
	if checkFileIsExist(projectName + "/pkg/conf/conf.go") { //如果文件存在
		f, _ = os.OpenFile(projectName+"/pkg/conf/conf.go", os.O_APPEND, 0777) //打开文件
	} else {
		f, _ = os.Create(projectName + "/pkg/conf/conf.go") //创建文件
	}
	confTml, _ := asset.Asset("template/conf.tml")
	io.WriteString(f, strings.Replace(string(confTml), "ProjectName", projectName, -1)) //写入文件(字符串)
	/***************************** 生成pkg/middleware的代码 ***********************************************/
	if checkFileIsExist(projectName + "/pkg/middleware/recovery.go") { //如果文件存在
		f, _ = os.OpenFile(projectName+"/pkg/middleware/recovery.go", os.O_APPEND, 0666) //打开文件
	} else {
		f, _ = os.Create(projectName + "/pkg/middleware/recovery.go") //创建文件
	}
	recoveryTml, _ := asset.Asset("template/recovery.tml")
	io.WriteString(f, strings.Replace(string(recoveryTml), "ProjectName", projectName, -1)) //写入文件(字符串)
	/***************************** 生成pkg/serializer的代码 ***********************************************/
	if checkFileIsExist(projectName + "/pkg/serializer/serializer.go") { //如果文件存在
		f, _ = os.OpenFile(projectName+"/pkg/serializer/serializer.go", os.O_APPEND, 0666) //打开文件
	} else {
		f, _ = os.Create(projectName + "/pkg/serializer/serializer.go") //创建文件
	}
	serializerTml, _ := asset.Asset("template/serializer.tml")
	io.WriteString(f, strings.Replace(string(serializerTml), "ProjectName", projectName, -1)) //写入文件(字符串)
	/***************************** 生成route的代码 ***********************************************/
	if checkFileIsExist(projectName + "/route/route.go") { //如果文件存在
		f, _ = os.OpenFile(projectName+"/route/route.go", os.O_APPEND, 0666) //打开文件
	} else {
		f, _ = os.Create(projectName + "/route/route.go") //创建文件
	}
	routeTml, _ := asset.Asset("template/router.tml")

	io.WriteString(f, strings.Replace(string(routeTml), "ProjectName", projectName, -1)) //写入文件(字符串)
	/***************************** 生成route的代码 ***********************************************/
	if checkFileIsExist(projectName + "/main.go") { //如果文件存在
		f, _ = os.OpenFile(projectName+"/main.go", os.O_APPEND, 0666) //打开文件
	} else {
		f, _ = os.Create(projectName + "/main.go") //创建文件
	}
	mainTml, _ := asset.Asset("template/main.tml")
	io.WriteString(f, strings.Replace(string(mainTml), "ProjectName", projectName, -1)) //写入文件(字符串)
}

/**
 * 判断文件是否存在  存在返回 true 不存在返回false
 */
func checkFileIsExist(filename string) bool {
	var exist = true
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		exist = false
	}
	return exist
}

var myTemplate *template.Template

type Person struct {
	Name string
	age  string
}

func userInfo() {
	initTemplate("template/main.tml")
	p := Person{Name: "Murphy", age: "28"}
	myTemplate.Execute(os.Stdout, p)
	file, err := os.OpenFile("template/main.tml", os.O_CREATE|os.O_WRONLY, 0755)
	if err != nil {
		fmt.Println("open failed err:", err)
		return
	}
	myTemplate.Execute(file, p)
}

func initTemplate(fileName string) (err error) {
	myTemplate, err = template.ParseFiles(fileName)
	if err != nil {
		fmt.Println("parse file err:", err)
		return
	}
	return
}
