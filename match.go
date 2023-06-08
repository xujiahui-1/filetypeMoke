package filetypeMoke

import (
	"filetyoe/matchers"
	"filetyoe/types"
	"io"
	"os"
)

// Matchers  给 matchers.Matchers创建一个别名
var Matchers = matchers.Matchers

// MatcherKeys  &matchers.MatcherKeys也取个别名
var MatcherKeys = &matchers.MatcherKeys

// NewMatcher matchers.NewMatcher也取个别名
var NewMatcher = matchers.NewMatcher

// Match 匹配
func Match(buf []byte) (types.Type, error) {
	length := len(buf) //获取长度
	//为空肯定不行
	if length == 0 {
		return types.Unknown, ErrEmptyBuffer
	}
	//遍历MatcherKeys
	for _, kind := range *MatcherKeys {
		//拉取所有处理器，挨个进行匹配
		checker := Matchers[kind]
		match := checker(buf)
		//如果找到有一个匹配，说明找到了
		if match != types.Unknown && match.Extension != "" {
			return match, nil
		}
	}
	return types.Unknown, nil
}

// Get is an alias to Match()
func Get(buf []byte) (types.Type, error) {
	return Match(buf)
}

// MatchFile 推断一个文件类型
func MatchFile(filepath string) (types.Type, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return types.Unknown, err
	}
	defer file.Close()

	return MatchReader(file)
}

// 读取文件
func MatchReader(reader io.Reader) (types.Type, error) {
	buffer := make([]byte, 8192) // 8K makes msooxml tests happy and allows for expanded custom file checks

	_, err := reader.Read(buffer)
	if err != nil && err != io.EOF {
		return types.Unknown, err
	}

	return Match(buffer)
}

// 自定义matcher用
func AddMatcher(fileType types.Type, matcher matchers.Matcher) matchers.TypeMatcher {
	return matchers.NewMatcher(fileType, matcher)
}

func Matches(buf []byte) bool {
	kind, _ := Match(buf)
	return kind != types.Unknown
}

// MatchMap performs a file matching against a map of match functions
func MatchMap(buf []byte, matchers matchers.Map) types.Type {
	for kind, matcher := range matchers {
		if matcher(buf) {
			return kind
		}
	}
	return types.Unknown
}

// MatchesMap is an alias to Matches() but using matching against a map of match functions
func MatchesMap(buf []byte, matchers matchers.Map) bool {
	return MatchMap(buf, matchers) != types.Unknown
}
