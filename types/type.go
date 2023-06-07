package types

// 最底层结构体
type Type struct {
	MIME      MIME   //文件MIME
	Extension string //文件扩展名
}

// NewType 创建 Type
func NewType(ext, mime string) Type {
	t := Type{MIME: NewMIME(mime), Extension: ext}
	return Add(t)
}
