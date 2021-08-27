package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"regexp"
	"time"
)

var (
	dkhpEndpoint     = "https://dkhp.hcmue.edu.vn/Login"
	reTitle          = regexp.MustCompile(`<title>[\n\t]*(.*)[\n\t]*<\/title>`)
	reError          = regexp.MustCompile(`<span style="color:red">[\n\t\r]*(.*)[\n\t\r]*<\/span>`)
	unexpectedErrMsg = "some thing went wrong\ndetails:\n%s"
)

func dkhp(username, password string) error {
	// Sending dkhp request
	data := url.Values{}
	data.Set("username", username)
	data.Set("password", password)

	resp, err := http.PostForm(dkhpEndpoint, data)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	raw, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	res := string(raw)

	switch resp.StatusCode {
	case http.StatusMovedPermanently:
		log.Println("Đăng nhập thành công")
		return nil
	case http.StatusOK:
		matches := reTitle.FindStringSubmatch(res)
		if len(matches) != 2 {
			log.Fatalf(unexpectedErrMsg, res)
		}
		title := matches[1]

		matches = reError.FindStringSubmatch(res)
		var msg string
		if len(matches) == 2 {
			msg = matches[1]
		}
		return fmt.Errorf("%s: %s", title, msg)
	default:
		log.Fatalf(unexpectedErrMsg, res)
		return nil
	}
}

func main() {
	cre := getCredentials()
	for {
		err := dkhp(cre.Mssv, cre.Password)
		if err == nil {
			if err := notify("Đăng ký học phần thôi !!!", cre.BotID, cre.ChatID); err != nil {
				log.Println(err)
			}
			break
		}
		log.Println(err)
		time.Sleep(time.Duration(2) * time.Minute)
	}
}
