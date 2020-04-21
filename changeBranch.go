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

func logError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	var start = time.Now()
	pathUrl := "D:/WorkeSpace2.0/"
	//获取文件或目录相关信息
	parentList, err := ioutil.ReadDir(pathUrl)
	logError(err)
	// 查找所有的父级目录
	for _, parent := range parentList {
		if parent.IsDir() && !strings.Contains(parent.Name(), ".") {
			projectList, err := ioutil.ReadDir(pathUrl + parent.Name())
			logError(err)
			// 查找父级目录下的项目
			for _, project := range projectList {
				if project.IsDir() {
					// 执行切换分支命令
					command := exec.Command("git", "checkout", "-b", "2.x", "origin/2.x")
					command.Dir = pathUrl + parent.Name() + "/" + project.Name()
					//运行脚本
					_, err := command.Output()
					if err != nil {
						fmt.Println(project.Name() + " 执行失败，请检查")
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
