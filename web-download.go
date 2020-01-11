package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os/exec"
)

func main() {

	//第一个参数是接口名，第二个参数 http handle func
	http.HandleFunc("/", requestHandle)
	//服务器要监听的主机地址和端口号
	http.ListenAndServe(":3000", nil)

}

func requestHandle(w http.ResponseWriter, r *http.Request) {
	var err error
	switch r.Method {
	case "POST":
		err = handlePost(w, r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func handlePost(w http.ResponseWriter, r *http.Request) (err error) {
	var body map[string]interface{}
	content, _ := ioutil.ReadAll(r.Body)
	fmt.Println(string(content))

	json.Unmarshal(content, &body)

	url := body["url"].(string)
	// folderName := body["name"]
	fmt.Println("URL:", url)

	// cmd := exec.Command("you-get", "-o /opt/", url.(string), "--debug")

	// s := fmt.Sprintf("you-get %s", url)
	// fmt.Println(s)
	cmd := exec.Command("you-get", "-o", "/download", url)

	var output bytes.Buffer
	var stderr bytes.Buffer

	cmd.Stdout = &output
	cmd.Stderr = &stderr
	fmt.Println(cmd)
	e := cmd.Run()

	if e != nil {
		fmt.Println(fmt.Sprint(err) + ": " + stderr.String())
	} else {
		fmt.Println(output.String())
	}

	return
}
