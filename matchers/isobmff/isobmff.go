package isobmff

import "encoding/binary"

// ISO 基本媒体文件格式 https://en.wikipedia.org/wiki/ISO_base_media_file_format
/*
在 ISOBMFF 中，通常使用大端字节序（Big-Endian）来表示多字节的整数值。
大端字节序是一种字节排列方式，其中最高有效字节位于最低内存地址处。
通过使用 binary.BigEndian.Uint32 函数，将字节数组的前 4 个字节解析为大端字节序的无符号 32 位整数。
这是因为 ISOBMFF 规范指定了长度字段应该使用大端字节序进行编码。
使用大端字节序解析长度字段的好处是确保与规范一致的解析和处理。
无论程序运行在哪种字节序的机器上，使用统一的大端字节序解析方法可以保证解析结果的一致性。
*/
// IsISOBMFF 确认文件 检查给定的缓冲区是否表示ISO基本媒体文件格式数据
func IsISOBMFF(buf []byte) bool { //传入字节数组，返回是不是
	if len(buf) < 16 || string(buf[4:8]) != "ftyp" {
		return false
	}
	//将字节数组的前 4 个字节解析为大端字节序的无符号 32 位整数
	if ftypLength := binary.BigEndian.Uint32(buf[0:4]); len(buf) < int(ftypLength) {
		return false
	}

	return true
}

// GetFtyp returns the major brand, minor version and compatible brands of the ISO-BMFF data
func GetFtyp(buf []byte) (string, string, []string) {
	if len(buf) < 17 {
		return "", "", []string{""}
	}

	ftypLength := binary.BigEndian.Uint32(buf[0:4])

	majorBrand := string(buf[8:12])
	minorVersion := string(buf[12:16])

	compatibleBrands := []string{}
	for i := 16; i < int(ftypLength); i += 4 {
		if len(buf) >= (i + 4) {
			compatibleBrands = append(compatibleBrands, string(buf[i:i+4]))
		}
	}

	return majorBrand, minorVersion, compatibleBrands
}
