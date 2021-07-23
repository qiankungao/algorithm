package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

//func init() {
//	// 加载配置文件
//	viper.SetConfigFile(".env")
//	_ = viper.ReadInConfig()
//}

const (
	//WEBHOOK_URL = "https://qyapi.weixin.qq.com/cgi-bin/webhook/send?key=5f0f5c00-bc07-4e99-8ea5-48cede416bc7"
	WEBHOOK_URL = "https://qyapi.weixin.qq.com/cgi-bin/webhook/send?key=75c5a745-b93e-43f9-8917-2077656d01b8"
	LISTEN      = ":80"
)

type Notify struct {
	Msgtype string        `json:"msgtype"`
	Content NotifyContent `json:"text"`
}
type NotifyContent struct {
	Content             string   `json:"content"`
	MentionedList       []string `json:"mentioned_list"`
	MentionedMobileList []string `json:"mentioned_mobile_list"`
}


//MMobo0wvybbDhKLg9ew_wxONmrfpctRDWUJHYlZK8g_afLMjcMq3hI3MCqfJ_KqObfXcES82fTINESe5ip0zYpagYgwWLeDf0wwLFJ201D80KLaLSad4tp_P9lYXiSKAJiJIOJQuJt5x-doPi8FR2mDA_55pjfidhlEQemAvNW9Lx_RjZrbYFaQLmWn9HcA_Mm3IXYkiOgM61m9Wt_IPTA
func main() {

	notify := &Notify{
		Msgtype: "text",
		Content: NotifyContent{
			Content:             "你好，helix2-Bug",
			MentionedList:       []string{"qiankun", "@all"},
			MentionedMobileList: []string{"15216884556", "@all"},
		},
	}
	payload, err := json.Marshal(notify)
	fmt.Println("player:", string(payload))
	response, err := http.Post(WEBHOOK_URL, "application/json", bytes.NewReader(payload))
	if err != nil {
		fmt.Println(err)
		return
	}

	all, err := ioutil.ReadAll(response.Body)
	fmt.Println(string(all), err)

	//})

	//_ = http.ListenAndServe(viper.GetString("LISTEN"), nil)
}
