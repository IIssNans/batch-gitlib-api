package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"strconv"
	"syscall"
	"time"
)

func httpGet(url string) *http.Response {
	client := &http.Client{}
	resp, err := client.Get(url)
	if err != nil {
		fmt.Println(err)
	}

	return resp
}

func printError(err error) {
	if err != nil {
		fmt.Print(err)
	}
}

func PathExists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return false
}

func main() {
	token := "K-c_19vAXjNxnMUFK7RP"
	ip := "http://XXX"
	dir := "/Users/gopher/zlx/workspace2.0/"

	url := ip + "/api/v4/namespaces?private_token=" + token
	var start = time.Now()
	resp := httpGet(url)
	body, _ := ioutil.ReadAll(resp.Body)
	var d1 []interface{}
	if err := json.Unmarshal([]byte(string(body)), &d1); err == nil {
		for _, val := range d1 {
			d2 := val.(map[string]interface{})
			var kind = d2["kind"].(string)
			var name = d2["name"].(string)
			var id = d2["id"].(float64)
			if name != "document" && kind == "group" {
				groupDir := dir + name
				// 如果文件夹不存在
				if !PathExists(groupDir) {
					err := os.Mkdir(groupDir, os.ModePerm)
					printError(err)
				}
				purl := ip + "/api/v4/groups/" + strconv.FormatInt(int64(id), 10) + "/projects?private_token=" + token
				presp := httpGet(purl)
				pbody, _ := ioutil.ReadAll(presp.Body)
				var pd []interface{}
				if err := json.Unmarshal([]byte(string(pbody)), &pd); err == nil {
					for _, val := range pd {
						d2 := val.(map[string]interface{})
						var repo = d2["ssh_url_to_repo"].(string)

						command := exec.Command("git", "clone", repo) //初始化Cmd
						command.Dir = groupDir
						_, err := command.Output() //运行脚本
						printError(err)
						if command.ProcessState.Sys().(syscall.WaitStatus).ExitStatus() == 0 {

						}
					}
				}
			}

		}
		var end = time.Now()
		fmt.Println("耗时：", end.Sub(start))
		fmt.Print("********************** 全部下载完成 *************************************")
	}
}
