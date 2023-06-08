# filetypeMoke


## 各文件夹含义

- fixtures:各种文件的例子
- types:最基本的类型结构体
  - `defaults.go`:就是一个未知类型
  - `split.go`:切割MIME的方法，返回前半段和后半段，其中前半段就是MIME的TYPE，后半段就是SUBTYPE
  - `mime.go`:就是MIME结构体的声明和创建方法
  - `type.go`:就是type结构体的声明和创建方法
  - `types.go`:TYPES集合的声明，并且编辑了GET和ADD方法提供添加和获取TYPE操作
- matchers:处理器等文件
  - isobmff:处理image用的iso文件类型相关代码
  - `document.go`:文本文件doc等等的处理器代码
  - `matchers.go` 处理器的初始化等等
- `filetype.go`:封装各种方法
-  `kind.go`: 封装各种方法
-  `match.go`:封装各种方法
-------

1. 最底层结构体`Type`,他包含文件扩展名属性和MIME属性，`MIME`结构体又包括TYPE,SUBTYPE和VALUE，其中TYPE+SUBTYPE=VALUE
2. `TYPES`其实就是个并发安全的map，用来存储所有的TYPE类型，
3. 当我们创建新的TYPE类型的时候，在NewType方法中，就直接调用ADD方法将其放入了TYPES集合中了

------
>  各种文件的MIME请参见 https://developer.mozilla.org/zh-CN/docs/Web/HTTP/Basics_of_HTTP/MIME_types/Common_types