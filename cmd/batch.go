/*
Copyright © 2023 jiy1012 <jiy1012@163.com>

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package cmd

import (
	"errors"
	"fmt"
	"github.com/cheggaaa/pb/v3"
	"github.com/fogleman/gg"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"image"
	"os"
	"path"
	"path/filepath"
	"strings"
)

// batchCmd represents the batch command
var batchCmd = &cobra.Command{
	Use:   "batch",
	Short: "批量操作",
	Long:  `根据配置文件、模板文件批量处理生成图片`,
	Run: func(cmd *cobra.Command, args []string) {

		names, maxLength := getNames()
		count := len(names)

		imagePath := getTemplateImagePath()
		im := loadTemplateImage(imagePath)
		outputPath = getOutput()
		/*
			= and > for indicating progress
			- and . for indicating incomplete progress
			█ and ▒ for indicating progress with a visual bar
			* and . for indicating progress with stars or dots
		*/

		tmpl := `{{string . "now_process_name" | blue}} {{ bar . " " (green "█") (black "█") (black "█")  " " }} {{speed . | rndcolor }} {{percent .}}`

		// start bar based on our template
		bar := pb.ProgressBarTemplate(tmpl).Start64(int64(count))
		bar.Set("now_process_name", strings.Repeat(" ", maxLength))
		for i := 0; i < count; i++ {
			bar.Increment()
			bar.Set("now_process_name", fmt.Sprintf("%s%s", names[i], strings.Repeat(" ", maxLength-len(names[i]))))
			drawGG(im, names[i], filepath.Ext(imagePath), outputPath)
			//time.Sleep(time.Second * 1)
		}
		bar.Set("now_process_name", fmt.Sprintf("%s%s", "complete!", strings.Repeat(" ", maxLength-9)))
		// finish bar
		bar.Finish()

		fmt.Println("batch completed。files at ", outputPath)
	},
}

func init() {
	rootCmd.AddCommand(batchCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// batchCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// batchCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func getNames() ([]string, int) {
	fileName := viper.GetString("Names")
	fileNameAbs, err := filepath.Abs(fileName)
	cobra.CheckErr(err)
	fileContent, err := os.ReadFile(fileNameAbs)
	cobra.CheckErr(err)
	if len(string(fileContent)) <= 0 {
		cobra.CheckErr(errors.New("names file is empty"))
	}
	var names []string
	maxLength := 9 // which complete! is
	for _, n := range strings.Split(string(fileContent), "\n") {
		if len(n) > maxLength {
			maxLength = len(n)
		}
		names = append(names, n)
	}
	return names, maxLength
}

func getTemplateImagePath() string {
	fileName := viper.GetString("ImageTemplate")
	fileNameAbs, err := filepath.Abs(fileName)
	cobra.CheckErr(err)
	return fileNameAbs
}

func getFont() (fontName string, fontSize int, x, y int, r, g, b, a uint) {
	fontName = viper.GetString("Font")
	fontSize = viper.GetInt("FontSize")
	x = viper.GetInt("TextAxisX")
	y = viper.GetInt("TextAxisY")
	r, g, b, a = viper.GetUint("TextColor.R"), viper.GetUint("TextColor.G"), viper.GetUint("TextColor.B"), viper.GetUint("TextColor.A")
	return
}
func getOutput() string {
	outDir := viper.GetString("OutDir")
	safeCreateFolder(outDir)
	return outDir
}
func loadTemplateImage(imgPath string) image.Image {
	im, err := gg.LoadImage(imgPath)
	cobra.CheckErr(err)
	return im
}
func drawGG(img image.Image, word string, ext string, output string) {
	dc := gg.NewContextForImage(img)
	fontName, fontSize, x, y, r, g, b, a := getFont()
	err := dc.LoadFontFace(fontName, float64(fontSize))
	cobra.CheckErr(err)
	////1920x1200
	//for i := 0; i < 100; i++ {
	//	dc.SetRGBA(0, 0, 0, float64(a))
	//	dc.SetLineWidth(float64(1))
	//	dc.DrawLine(0, float64(50*i), 1920, float64(50*i))
	//	dc.Stroke()
	//}
	//for i := 0; i < 100; i++ {
	//	dc.SetRGBA(255, 255, 255, float64(a))
	//	dc.SetLineWidth(float64(1))
	//	dc.DrawLine(float64(50*i), 0, float64(50*i), 1920)
	//	dc.Stroke()
	//}
	dc.SetRGBA(float64(r), float64(g), float64(b), float64(a))
	dc.DrawStringAnchored(word, float64(x), float64(y), 0.5, 0.5)
	err = dc.SavePNG(path.Join(output, fmt.Sprintf("%s%s", word, ext)))
	cobra.CheckErr(err)
}
