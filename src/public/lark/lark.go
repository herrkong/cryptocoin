package lark

import (
	"cryptocoin/src/public/request"
	"fmt"
	"time"
)


type Lark struct{}

// const (
// 	ALARM_SEND_TO_DARWIN  = "e070e5e4-7a22-4302-b1d9-60d5a4dd8503"
// )
//https://open.larksuite.com/open-apis/bot/v2/hook/e070e5e4-7a22-4302-b1d9-60d5a4dd8503

//fmt.Println(time.Now().Format("2006-01-02 15:04:05"))

func SendLark(alarmId string , title string , content string )(err error){
	url := fmt.Sprintf("https://open.larksuite.com/open-apis/bot/v2/hook/%s", alarmId)
	body := map[string]interface{}{
		"msg_type": `post`,
		"content": map[string]interface{}{
			"post": map[string]interface{}{
				"zh_cn": map[string]interface{}{
					"title": fmt.Sprintf("【%s】%s", title,time.Now().Format("2006-01-02 15:04:05")),
					"content": [][]map[string]interface{}{
						{
							{
								"tag":  `text`,
								"text": content,
							},
						},
					},
				}, 
			},
		},
	}
	if response, err := request.PostJson(url, body); err != nil {
		fmt.Print("send alarm text error", url, body, err)
		return err
	} else {
		fmt.Print("send alarm text", url, body, response)
		return nil
	}
}
