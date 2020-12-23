package ip_proxy

var constIPList = []string{
	"http://47.99.195.197:8080",
	"http://127.0.0.1:1087",
}
var constIPIndex = 0

func GetNextConstIP() string {
	if constIPIndex == len(constIPList) {
		constIPIndex = 0
	}
	r := constIPList[index]
	constIPIndex++
	return r
}
