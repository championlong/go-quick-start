package message

type MarkdownMessage struct {
	Title string `json:"title"` //首屏会话透出的展示内容
	Text  string `json:"text"`  //markdown格式的消息
}
