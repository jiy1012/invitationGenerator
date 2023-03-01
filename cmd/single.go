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
	"fmt"
	"github.com/spf13/cobra"
	"path/filepath"
)

var imgTemplate, outDir, name, font string
var textAxisX, textAxisY, fontSize int
var colorRed, colorGreen, colorBlue, colorAlpha uint

// singleCmd represents the single command
var singleCmd = &cobra.Command{
	Use:   "single",
	Short: "单次执行",
	Long:  `命令行单次执行一次，无需生成配置文件`,
	Run: func(cmd *cobra.Command, args []string) {
		safeCreateFolder(outDir)
		im := loadTemplateImage(imgTemplate)
		drawGG(im, name, filepath.Ext(imgTemplate), outDir, font, fontSize, textAxisX, textAxisY, colorRed, colorGreen, colorBlue, colorAlpha)
		fmt.Println("completed。file at ", outDir)
	},
}

func init() {
	rootCmd.AddCommand(singleCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// singleCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// singleCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	singleCmd.Flags().StringVarP(&imgTemplate, "template", "t", "", "image template( required )")
	singleCmd.Flags().StringVarP(&outDir, "output", "o", "./", "output path (default is ./)")
	singleCmd.Flags().StringVarP(&name, "name", "n", "", "text of name (required)")

	singleCmd.Flags().StringVarP(&font, "font", "f", "fonts/STHeiti Medium.ttc", "font location (default is )")
	singleCmd.Flags().IntVarP(&fontSize, "size", "s", 0, "text of name (required)")

	singleCmd.Flags().IntVarP(&textAxisX, "axisX", "x", 0, "x-coordinate of text (required)")
	singleCmd.Flags().IntVarP(&textAxisY, "axisY", "y", 0, "y-coordinate of text (required)")

	singleCmd.Flags().UintVarP(&colorRed, "red", "r", 0, "color of text:r (required)")
	singleCmd.Flags().UintVarP(&colorGreen, "green", "g", 0, "color of text:g (required)")
	singleCmd.Flags().UintVarP(&colorBlue, "blue", "b", 0, "color of text:b (required)")
	singleCmd.Flags().UintVarP(&colorAlpha, "alpha", "a", 1, "color of text:a (required)")

}
