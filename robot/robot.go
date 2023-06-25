package robot

import (
	//opentraceHandler "git.xq5.com/office/survey-backend/handler/opentrace"
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"github.com/amazingGan/pkg/log"
)

// SendRobotMsg 发送机器人通知
func SendRobotMsg(ctx context.Context, conf *Conf, msg string) {
	// 如果机器人配置为空，直接返回
	if conf == nil || len(conf.MentionedMobileList) == 0 {
		return
	}

	if !conf.IsOpen || msg == "" || conf.WebHoot == "" {
		// 如果关闭机器人，直接返回
		return
	}

	// traceID := opentraceHandler.GetTraceID(ctx)
	// log := opentraceHandler.GetTraceLogger(ctx)

	// 需要展示的内容，使用map进行装载
	contentMap := make(map[string]interface{})

	// 此处没有加锁，是因为contentMap是一个临时变量，不会被别人调用
	for k, v := range conf.Template {
		contentMap[k] = v
	}
	// 此处赋值traceID纯粹用作测试使用，以配置文件中配置的traceID作为输出
	// if c.Template["trace_id"] != nil {
	// 	traceID = c.Template["trace_id"].(string)
	// }

	contentMap["date"] = time.Now().Format(time.RFC3339)
	contentMap["service"] = conf.ServiceName
	contentMap["environment"] = conf.Environment
	contentMap["error"] = msg
	//contentMap["trace_id"] = traceID

	//traceURL := c.Template["trace_url"].(string)
	//contentMap["trace_url"] = traceURL + traceID

	bts, err := json.MarshalIndent(contentMap, "", "\t")
	if err != nil {
		log.Error("contentMap json转化报错", log.FieldErr(err))
	}

	reqBody := map[string]interface{}{
		"msgtype": "text",
		"text": map[string]interface{}{
			"content":               string(bts),
			"mentioned_mobile_list": conf.MentionedMobileList,
		},
	}
	b, _ := json.Marshal(reqBody)
	payload := strings.NewReader(string(b))
	client := &http.Client{}
	req, err := http.NewRequest(http.MethodPost, conf.WebHoot, payload)
	if err != nil {
		log.Error("robot webhoot new request err", "webhoot", conf.WebHoot, log.FieldErr(err))
		return
	}
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		log.Error("robot webhoot do request err", "webhoot", conf.WebHoot, log.FieldErr(err))
		return
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Error("robot webhoot res.Body read err", "webhoot", conf.WebHoot, log.FieldErr(err))
		return
	}
	if res.StatusCode != http.StatusOK {
		log.Error("robot webhoot 请求失败", "StatusCode", res.StatusCode, "Header", res.Header, "Body", string(body))
	}
}
