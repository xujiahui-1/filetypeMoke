package types

import "sync"

// 实际就是个并发安全的map，用来存储各种type
var Types sync.Map

// Add 方法帮我们把创建好的type加入到Types中
func Add(t Type) Type {
	Types.Store(t.Extension, t) //key:文件扩展名，value:Type实例
	return t                    //这里return完全是为了在NewType方法中返回给调用者的示例，我认为不需要这么写
}

// Get 通过extension查找一个Type
func Get(ext string) Type {
	//找到了，返回对应的Type对象
	if tmp, ok := Types.Load(ext); ok {
		kind := tmp.(Type)
		if kind.Extension != "" {
			return kind
		}
	}
	//找不到 返回Unknown
	return Unknown
}
