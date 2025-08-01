package meat

import "laiwh/hello/logprint"

func MakeDish(name string) {
	logprint.Info("开始制作肉菜" + name)
}
