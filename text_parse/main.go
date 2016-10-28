package main

import (
	"encoding/json"
	"fmt"
	"regexp"
	"net/http"
	"io/ioutil"
	"strings"
)

type DayLive struct {
	TotalCount int
	LogDay     string
}

type Rs struct {
	DayLives []DayLive
}

func main() {
	testRegx()
}

func testRegx() {
	resp, err := http.Get("http://www.so.com")
	checkErr(err)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	checkErr(err)
	src := string(body)
	re, _ := regexp.Compile("\\<[\\S\\s]+?\\>")
	//将html标签转换成小写
	src = re.ReplaceAllStringFunc(src, strings.ToLower)
	sre, _ := regexp.Compile("\\<style[\\S\\s]+?\\</style\\>")
	//过滤掉所有style
	src = sre.ReplaceAllString(src, "")
	jsRe, _ := regexp.Compile("\\<script[\\S\\s]+?\\</script\\>")
	src = jsRe.ReplaceAllString(src, "")
	hre, _ := regexp.Compile("\\<[\\S\\s]+?\\>")
	src =hre.ReplaceAllString(src, "\n")
	nre,_ := regexp.Compile("\\s{2,}")
	src = nre.ReplaceAllString(src, "\n")
	fmt.Println(strings.TrimSpace(src))
}

func IsIP(ip string) bool {
	m, _ := regexp.MatchString("^[0-9]{1,3}.[0-9]{1,3}.[0-9]{1,3}.[0-9]{1,3}$", ip);
	return m

}

func jsonDemo() {
	var rs Rs
	str := `{"DayLives":[{"TotalCount":189393,"LogDay":"2016-10-21"},{"TotalCount":222218,"LogDay":"2016-10-22"},{"TotalCount":225080,"LogDay":"2016-10-23"},{"TotalCount":213959,"LogDay":"2016-10-24"},{"TotalCount":220441,"LogDay":"2016-10-25"},{"TotalCount":235729,"LogDay":"2016-10-26"},{"TotalCount":246066,"LogDay":"2016-10-27"},{"TotalCount":178436,"LogDay":"2016-10-28"}]}`
	err := json.Unmarshal([]byte(str), &rs)
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Printf("%+v", len(rs.DayLives))
}
func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}