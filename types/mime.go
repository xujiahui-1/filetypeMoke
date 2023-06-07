package types

// 文件的MIME结构体
type MIME struct {
	Type    string
	Subtype string
	Value   string
}

// NewMIME MIME的创建方法,传入string,通过对string进行切割，确定他的Type和Subtype
func NewMIME(mime string) MIME {
	kind, subtype := splitMime(mime)
	return MIME{Type: kind, Subtype: subtype, Value: mime}
}
