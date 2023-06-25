package robot

import (
	"context"
	"testing"
)

func TestSendRobotMsg(t *testing.T) {
	type args struct {
		ctx  context.Context
		conf *Conf
		msg  string
	}
	tests := []struct {
		name string
		args args
	}{

		{
			name: "测试发送机器人信息",
			args: args{
				ctx: context.Background(),
				conf: &Conf{
					ServiceName:         "test_server",
					Environment:         "test",
					IsOpen:              true,
					WebHoot:             "https://qyapi.weixin.qq.com/cgi-bin/webhook/send?key=b3abe54c-1ccf-4fa7-98fe-200e26ca5b00",
					MentionedMobileList: []string{"15527966819"},
					Template: map[string]interface{}{
						"trace_url": "接入trace服务的地址(目前暂未开启)",
						"trace_id":  "xxxxx",
					},
				},
				msg: "测试内容",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			SendRobotMsg(tt.args.ctx, tt.args.conf, tt.args.msg)
		})
	}
}
