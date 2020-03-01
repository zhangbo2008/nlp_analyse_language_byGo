package whatlanggo

import "unicode"

var (// 设置unicode编码
	_HiraganaKatakana = &unicode.RangeTable{
		R16: append(unicode.Hiragana.R16, unicode.Katakana.R16...),
		R32: append(unicode.Hiragana.R32, unicode.Katakana.R32...),
	}
)
