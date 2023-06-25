package robot

// Conf 机器人配置
type Conf struct {
	ServiceName         string // 服务名称
	Environment         string // 测试dev/生产环境prod
	IsOpen              bool   // 是否开启
	WebHoot             string
	MentionedMobileList []string
	Template            map[string]interface{} //模版内容
}
