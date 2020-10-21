package XMPushSDK

//type Message struct {
//	Payload					string 					//消息内容
//	RestrictedPackageName	string					//支持多包名
//	PassThrough				int32					//是否透传给app(1 透传 0 通知栏信息)
//	NotifyType				int                     //通知类型 可组合 (-1 Default_all,1 Default_sound,2 Default_vibrate(震动),4 Default_lights)
//	NotifyId  				int              		//0-4同一个notifyId在通知栏只会保留一条
//	Description				string                  //在通知栏的描述，长度小于128
//	Title					string                  //在通知栏的标题，长度小于16
//	TimeToLive				int            			//可选项，当用户离线是，消息保留时间，默认两周，单位ms
//	TimeToSend				int            			//可选项，定时发送消息，用自1970年1月1日以来00:00:00.0 UTC时间表示（以毫秒为单位的时间）。
//	Extra					map[string]string   	//可选项，额外定义一些key value（字符不能超过1024，key不能超过10个）
//  Fields					map[string]interface{}  //含有本条消息所有属性的数组
//	JsonInfos				map[string]interface{}
//}

type Message struct {
	Fields map[string]interface{} //含有本条消息所有属性的数组
}

func InitMessage(title string, desc string, restrictedPackageName string, payloadStr string, passThrough string) *Message {
	fields := map[string]interface{}{
		"registration_id":         "",
		"title":                   title,
		"description":             desc,
		"restricted_package_name": restrictedPackageName,
		"payload":                 payloadStr,
		"notify_type":             "-1",
		"pass_through":            passThrough,
	}
	return &Message{
		Fields: fields,
	}
}

func (m *Message) GetFields() map[string]interface{} {
	return m.Fields
}
