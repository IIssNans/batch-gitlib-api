package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os/exec"
	"strings"
	"syscall"
	"time"
)

/**
只能查找2级目录

*/
func main() {
	var start = time.Now()
	pathUrl := "/Users/gopher/zlx/workspace2.0/"
	//获取文件或目录相关信息
	parentList, err := ioutil.ReadDir(pathUrl)
	if err != nil {
		log.Fatal(err)
	}
	for _, parent := range parentList {
		if parent.IsDir() && !strings.Contains(parent.Name(), ".") {
			fmt.Println(parent.Name())
			projectList, err := ioutil.ReadDir(pathUrl + parent.Name())
			if err != nil {
				log.Fatal(err)
			}
			for _, project := range projectList {
				if project.IsDir() {
					//fmt.Println(project.Name())
					command := exec.Command("git", "checkout", "-b", "2.x", "origin/2.x") //初始化Cmd
					command.Dir = pathUrl + parent.Name() + "/" + project.Name()
					_, err := command.Output() //运行脚本
					//	printError(err)
					if err != nil {
						fmt.Println(project.Name() + "执行失败，请检查")
					}
					if command.ProcessState.Sys().(syscall.WaitStatus).ExitStatus() == 0 {

					}
					time.Sleep(1000)
				}
			}
		}
	}
	var end = time.Now()
	fmt.Println("耗时：", end.Sub(start))
	fmt.Println("全部切换完成")
}
