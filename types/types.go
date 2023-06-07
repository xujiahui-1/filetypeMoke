package types

import "sync"

// 实际就是个并发安全的map，用来存储各种type
var Types sync.Map

// Add 方法帮我们把创建好的type加入到Types中
func Add(t Type) Type {
	Types.Store(t.Extension, t) //key:type类型，value:Type实例
	return t                    //这里return完全是为了在NewType方法中返回给调用者的示例，我认为不需要这么写
}
