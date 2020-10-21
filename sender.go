package XMPushSDK

import "strings"

func SendToOneRegId(appSecret string, message *Message, regId string) (*Result, error) {
	fields := message.GetFields()
	fields["registration_id"] = regId
	return PostResult(MessageRegIdURL, fields, 1, appSecret)
}

func SendRegIds(appSecret string, message *Message, regIds []string) (*Result, error) {
	fields := message.GetFields()
	fields["registration_id"] = strings.Join(regIds, ",")
	return PostResult(MessageRegIdURL, fields, 1, appSecret)
}
