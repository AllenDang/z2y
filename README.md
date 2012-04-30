##About

z2y converts (and only converts) chinese character in utf8 encoding into initial Pinyin string.
z2y 会把utf8格式的汉字字符串转换成拼音首字母字符串

	E.g.
	"啊薄车单饿飞干很见空冷吗嫩哦盘群润三图玩先烟总" -> "abcdefghjklmnopqrstwxyz"
	"中华人民共和国" -> "zhrmghg"
	"汉语拼音都是浮云" -> "hypydsfy"
	"世界你好！hello world!" -> "sjnh！hello world!"
	"！，。；……（）你好￥：、——" -> "！，。；……（）nh￥：、——"
	"all english" -> "all english"
	"1234567890" -> "1234567890"
	"!@#$%^&*()~,.<>/?';:[]{}" -> "!@#$%^&*()~,.<>/?';:[]{}"
	"转换原理是把UTF8字符先转换成gbk然后查表" -> "zhylsbUTF8zfxzhcgbkrhcb"

##Setup

###1. Make sure you have a working Go installation, you can get it from:
   http://golang.org/
   
###2. Create a "gopath" directory if you do not have one yet and set the
   GOPATH variable accordingly. For example:
   mkdir -p go-externals/src
   export GOPATH=${PWD}/go-externals

###3. go get github.com/AllenDang/z2y

###4. go install github.com/AllenDang/z2y

##Usage
	py := z2y.ConvertToPy("你好世界")
	println(py) //output is "nhsj"

##Contribute

Contributions in form of design, code, documentation, bug reporting or other
ways you see fit are very welcome.

Thank You!