// Copyright 2012 The z2y Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package z2y_test

import (
	. "github.com/AllenDang/z2y"
	"testing"
)

var data = []struct {
	origin, result string
}{
	{"", ""},
	{"啊薄车单饿飞干很见空冷吗嫩哦盘群润三图玩先烟总", "abcdefghjklmnopqrstwxyz"},
	{"中华人民共和国", "zhrmghg"},
	{"汉语拼音都是浮云", "hypydsfy"},
	{"世界你好！hello world!", "sjnh！hello world!"},
	{"！，。；……（）你好￥：、——", "！，。；……（）nh￥：、——"},
	{"all english", "all english"},
	{"1234567890", "1234567890"},
	{"!@#$%^&*()~,.<>/?';:[]{}", "!@#$%^&*()~,.<>/?';:[]{}"},
	{"转换原理是把UTF8字符先转换成gbk然后查表", "zhylsbUTF8zfxzhcgbkrhcb"},
}

func TestConvertToPY(t *testing.T) {
	for _, d := range data {
		r := ConvertToPY(d.origin)
		if r != d.result {
			t.Errorf("Expecting %s but %s", d.result, r)
		}
	}
}
