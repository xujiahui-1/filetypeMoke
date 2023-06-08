package matchers

import "filetyoe/types"

// 把newType方法复制过来
var newType = types.NewType

// Matcher 这是matcher的结构体、文件类型匹配起
type Matcher func([]byte) bool //传入一个字节数组，返回布尔

// Map  定义一个map接口，各种类型的文件创建他并存入
type Map map[types.Type]Matcher //键是Type实例，值是处理器方法

// TypeMatcher 处理器
type TypeMatcher func([]byte) types.Type

// 存储已注册的文件类型匹配器
var Matchers = make(map[types.Type]TypeMatcher)
var MatcherKeys []types.Type //只存Matchers键的数组

// NewMatcher 创建并存储一个新的 type matcher function
func NewMatcher(kind types.Type, fn Matcher) TypeMatcher { //传入Type和Matcher，返回TypeMatcher
	//返回的这个TypeMatcher里面就包含了Matcher！！！！
	matcher := func(buf []byte) types.Type {
		if fn(buf) { //这里的意思是如果把Type传入到了Matcher中，返回的是true，说明匹配，就返回type
			return kind
		}
		return types.Unknown //不匹配，返回Unknow
	}
	//存到Matchers map中
	Matchers[kind] = matcher
	//在存到只存Matchers键的数组里
	MatcherKeys = append([]types.Type{kind}, MatcherKeys...)
	return matcher
}

// 在我们将所有小map存到Matchers的时候，
func register(matchers ...Map) {
	MatcherKeys = MatcherKeys[:0]
	for _, m := range matchers {
		for kind, matcher := range m {
			NewMatcher(kind, matcher)
		}
	}
}

// 初始化的时候，把所有的类型的map，全部通过register方法，遍历创建新的NewMatcher，存到Matchers map中
func init() {
	// Arguments order is intentional
	// Archive files will be checked last due to prepend above in func NewMatcher
	register(Document)
}
