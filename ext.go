package binext

import (
	"regexp"
	"strings"
)

const BinaryFileExtensions = `
3dm 3ds 3g2 3gp 7z a aac adp ai aif
aiff alz ape apk appimage ar arj asf au
avi bak baml bh bin bk bmp btif bz2
bzip2 cab caf cgm class cmx cpio cr2
cur dat dcm deb dex djvu dll dmg dng
doc docm docx dot dotm dra DS_Store dsk
dts dtshd dvb dwg dxf ecelp4800 ecelp7470
ecelp9600 egg eol eot epub exe f4v fbs
fh fla flac flatpak fli flv fpx fst fvt
g3 gh gif graffle gz gzip h261 h263
h264 icns ico ief img ipa iso jar jpeg
jpg jpgv jpm jxr key ktx lha lib lvp
lz lzh lzma lzo m3u m4a m4v mar mdi
mht mid midi mj2 mka mkv mmr mng mobi
mov movie mp3 mp4 mp4a mpeg mpg mpga
mxu nef npx numbers nupkg o odp ods odt
oga ogg ogv otf ott pages pbm pcx pdb
pdf pea pgm pic png pnm pot potm potx
ppa ppam ppm pps ppsm ppsx ppt pptm
pptx psd pya pyc pyo pyv qt rar ras
raw resources rgb rip rlc rmf rmvb rpm
rtf rz s3m s7z scpt sgi shar snap sil
sketch slk smv snk so stl suo sub swf
tar tbz tbz2 tga tgz thmx tif tiff tlz
ttc ttf txz udf uvh uvi uvm uvp uvs
uvu viv vob war wav wax wbmp wdp weba
webm webp whl wim wm wma wmv wmx woff
woff2 wrm wvx xbm xif xla xlam xls xlsb
xlsm xlsx xlt xltm xltx xm xmind xpi
xpm xwd xz z zip zipx`

func init() {
	extensions := strings.Fields(BinaryFileExtensions)
	// prefix with .
	for i := range extensions {
		extensions[i] = `\.` + extensions[i]
	}
	expr := "(" + strings.Join(extensions, "|") + ")$"
	binaryExt = regexp.MustCompile(expr)
}

var binaryExt *regexp.Regexp

// IsBinary returns true if the given value is binary according
// to defined BinaryFileExtensions.
//
// The value can be of type string, []byte or implement interface{
// Name() string
// }
func IsBinary(v interface{}) bool {
	switch v := v.(type) {
	case string:
		return binaryExt.MatchString(v)
	case []byte:
		return binaryExt.Match(v)
	case interface{ Name() string }:
		return binaryExt.MatchString(v.Name())
	}
	return false
}
