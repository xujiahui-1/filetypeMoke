package filetypeMoke

import (
	"errors"
	"filetyoe/types"
)

// filetype 对

// 把这三个东西从types包复制过来，方便地下写代码
var Types = types.Types

var NewType = types.NewType //这里为什么复制一遍不懂

var Unknown = types.Unknown

// ErrEmptyBuffer 创建一个缓冲为空的错误
var ErrEmptyBuffer = errors.New("Empty buffer")

// ErrUnknownBuffer 找不到缓冲的错误
var ErrUnknownBuffer = errors.New("Unknown buffer type")

// AddType 创建一个type并加到Types中
func AddType(ext, mime string) types.Type {
	return types.NewType(ext, mime)
}

// 检查给定缓冲区是否与给定文件类型扩展名匹配
func Is(buf []byte, ext string) bool {
	kind := types.Get(ext)     //找map中有没有这个扩展名
	if kind != types.Unknown { //有，那么看看和buf中的一不一样
		//TODO 调用IsType方法，比较缓冲和type一不一样
	}
	return false
}

// IsType 判断缓冲区中的和kind一不一样 TODO
//func IsType(buf []byte, kind types.Type) bool {
//
//}
