// Copyright 2012 The z2y Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package z2y

import (
	"code.google.com/p/mahonia"
)

// ConvertToPY converts (and only converts) chinese character in utf8 encoding into initial Pinyin string.
// E.g.
// "啊薄车单饿飞干很见空冷吗嫩哦盘群润三图玩先烟总" -> "abcdefghjklmnopqrstwxyz"
// "世界你好！hello world!" -> "sjnh"
// "把UTF8字符转换成gbk" -> "bzfzhc"
// "all english" -> ""
// "1234890" -> ""
// "!@#$%^&*()~,.<>/?';:[]{}" -> ""
func ConvertToPY(s string) string {
	//Convert utf8 to gbk
	e := mahonia.NewEncoder("gbk")
	str := e.ConvertString(s)
	println(str)

	mapfunc := func(r int32) rune {
		switch {
		case r >= -20319 && r <= -20284:
			return 'a'
		case r >= -20283 && r <= -19776:
			return 'b'
		case r >= -19775 && r <= -19219:
			return 'c'
		case r >= -19218 && r <= -18711:
			return 'd'
		case r >= -18710 && r <= -18527:
			return 'e'
		case r >= -18526 && r <= -18240:
			return 'f'
		case r >= -18239 && r <= -17923:
			return 'g'
		case r >= -17922 && r <= -17418:
			return 'h'
		case r >= -17417 && r <= -16475:
			return 'j'
		case r >= -16474 && r <= -16213:
			return 'k'
		case r >= -16212 && r <= -15641:
			return 'l'
		case r >= -15640 && r <= -15166:
			return 'm'
		case r >= -15165 && r <= -14923:
			return 'n'
		case r >= -14922 && r <= -14915:
			return 'o'
		case r >= -14914 && r <= -14631:
			return 'p'
		case r >= -14630 && r <= -14150:
			return 'q'
		case r >= -14149 && r <= -14091:
			return 'r'
		case r >= -14090 && r <= -13319:
			return 's'
		case r >= -13318 && r <= -12839:
			return 't'
		case r >= -12838 && r <= -12557:
			return 'w'
		case r >= -12556 && r <= -11848:
			return 'x'
		case r >= -11847 && r <= -11056:
			return 'y'
		case r >= -11055 && r <= -10247:
			return 'z'
		}

		return 0
	}

	l := len(str)
	buf := make([]rune, 0)

	for i := 0; i < l && (i+1) < l; {
		if str[i] <= 'z' {
			i += 1
			continue
		}
		a := int32(str[i])*256 + int32(str[i+1]) - 65536
		if r := mapfunc(a); r != 0 { // r will be 0 is input is not a Chinese character.
			buf = append(buf, r)
		}
		i += 2
	}

	return string(buf)
}
