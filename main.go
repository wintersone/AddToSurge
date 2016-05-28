package main

import (
	"fmt"
	"os/exec"
	"log"
	"net/url"

	"io/ioutil"
	"strings"
	"os"
)

var (
	osscript = "tell application \"Safari\" to return URL of front document"
	osscriptClose = "do shell script \"killall Surge\""
	osscriptOpen = "run application \"Surge\""
	filePath = "/Users/huanghuan/.surge.conf"
	ruleSuffix = ",Proxy"
	rulePrefix = "\nDOMAIN-SUFFIX,"
	flag = "[Rule]"
)

func main() {



	
	cmdOutput, err := exec.Command("osascript", "-e", osscript).Output()

	if err != nil {
		log.Fatal(err)
	}


	
	urlEntity,_ := url.Parse(string(cmdOutput))

	host := strings.TrimPrefix(urlEntity.Host, "www.")

	if HasRule(host) {
		fmt.Printf("Host %s Already Exist",host)
	} else {
		AddToConf(host)

		fmt.Printf("Host %s Add To Your %s", host, filePath)

		exec.Command("killall","Surge").Run()
		exec.Command("osascript", "-e", osscriptOpen).Run()
		exec.Command("osascript", "reload.scpt").Run()
	}
}

func HasRule(host string) bool {

	input, _ := ioutil.ReadFile(filePath)
	
	lines := strings.Split(string(input), "\n")

	for _, line := range lines {
		if strings.Contains(line, host) {
			return true
		}
		
	}

	return false
}

func AddToConf(host string) {

	rule := rulePrefix + host + ruleSuffix
	
	fmt.Println(rule)

	input, _ := ioutil.ReadFile(filePath)
	
	lines := strings.Split(string(input), "\n")

	for i, line := range lines {
		if strings.Contains(line, "[Rule]") {
			lines[i] = line + rule

		}
		
	}

	output := strings.Join(lines, "\n")

	fileInfo,_ := os.Stat(filePath)

	ioutil.WriteFile(filePath, []byte(output),fileInfo.Mode())
}


