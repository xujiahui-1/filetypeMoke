package filetypeMoke

import (
	"errors"
	"filetyoe/matchers"
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
		// 调用IsType方法，比较缓冲和type一不一样
		return IsType(buf, kind)
	}
	return false
}

// IsType 判断缓冲区中的和Type一不一样
func IsType(buf []byte, kind types.Type) bool {
	//根据Type去查Matchers中，找到TypeMatcher
	//TypeMatcher就是一个方法，返回Type
	matcher := matchers.Matchers[kind]
	//找不到，就说明判断不了你这个Type
	if matcher == nil {
		return false
	}
	//找到了，那么就要去看类型是否匹配
	return matcher(buf) != types.Unknown
}

// IsExtension semantic alias to Is()
func IsExtension(buf []byte, ext string) bool {
	return Is(buf, ext)
}

// IsMIME checks if a given buffer matches with the given MIME type
func IsMIME(buf []byte, mime string) bool {
	result := false
	types.Types.Range(func(k, v interface{}) bool {
		kind := v.(types.Type)
		if kind.MIME.Value == mime {
			matcher := matchers.Matchers[kind]
			result = matcher(buf) != types.Unknown
			return false
		}
		return true
	})

	return result
}

// IsSupported checks if a given file extension is supported
func IsSupported(ext string) bool {
	result := false
	types.Types.Range(func(k, v interface{}) bool {
		key := k.(string)
		if key == ext {
			result = true
			return false
		}
		return true
	})

	return result
}

// IsMIMESupported checks if a given MIME type is supported
func IsMIMESupported(mime string) bool {
	result := false
	types.Types.Range(func(k, v interface{}) bool {
		kind := v.(types.Type)
		if kind.MIME.Value == mime {
			result = true
			return false
		}
		return true
	})

	return result
}

// GetType retrieves a Type by file extension
func GetType(ext string) types.Type {
	return types.Get(ext)
}
