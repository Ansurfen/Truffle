package etcd

import "strconv"

func RegisterService(name string, pod int) string {
	return "/truffle/" + name + "/pod" + strconv.Itoa(pod)
}
