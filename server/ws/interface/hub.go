package abstract

type AbstractHub interface {
	GenPath(topicName string) string
	Decoding(path string) string
	Encoding(path string) string
}

var ghub AbstractHub

func GetHubWithAdapter() AbstractHub {
	return ghub
}

func InitGHub(hub AbstractHub) {
	ghub = hub
}