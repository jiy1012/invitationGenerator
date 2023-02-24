package tpl

func ConfigTemplate() []byte {
	return []byte(`#basic config
# input
# names.txt contains one name in a line
Names: "~/.ig/names.txt"
ImageTemplate: '~/.ig/Templates/sl.jpg'
# output
OutDir: './Output/'

# display 
Font: "~/.ig/Fonts/SourceHanSansSC-Bold.otf"
FontSize: 140
TextAxisX: 1000
TextAxisY: 2000
TextColor:
  R: 255
  G: 222
  B: 184
  A: 0
`)
}
