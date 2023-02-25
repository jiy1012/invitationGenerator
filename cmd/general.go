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
	"github.com/jiy1012/invitationGenerator/tpl"
	"io"
	"os"
	"path"

	"github.com/spf13/cobra"
)

var templateType, outputPath string

// generalCmd represents the general command
var generalCmd = &cobra.Command{
	Use:   "general",
	Short: "生成文件模板",
	Long:  `生成配置文件、处理文件等模板`,
	Run: func(cmd *cobra.Command, args []string) {
		//fmt.Println("templateType:", templateType, "outputPath:", outputPath)
		switch templateType {
		case "all":
			cobra.CheckErr(createConfig())
			cobra.CheckErr(createFont())
			cobra.CheckErr(createBG())
			fmt.Printf("all config、font、background-image template file created at %s\n", outputPath)
		case "config":
			cobra.CheckErr(createConfig())
			fmt.Printf("config file created at %s\n", outputPath)
		case "font":
			cobra.CheckErr(createFont())
			fmt.Printf("font created at %s\n", path.Join(outputPath, fontFolder))
		case "bg":
			cobra.CheckErr(createBG())
			fmt.Printf("background-image template created at %s\n", path.Join(outputPath, bgFolder))
		default:
			fmt.Println("nothing created")
		}
	},
}

func init() {
	rootCmd.AddCommand(generalCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	//generalCmd.PersistentFlags().String("template", "config", "conifg template")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	generalCmd.Flags().StringVarP(&templateType, "template", "t", "all", "config template(default is config)")
	generalCmd.Flags().StringVarP(&outputPath, "output", "o", "", "output path (default is $HOME/.ig/)")
}
func createBG() error {
	fillOutputFolder()
	safeCreateFolder(path.Join(outputPath, bgFolder))
	bgs, err := tpl.BGFileTemplate.ReadDir(bgFolder)
	if err != nil {
		return err
	}
	for _, bg := range bgs {
		fontFile, err := os.Create(path.Join(outputPath, bgFolder, bg.Name()))
		if err != nil {
			return err
		}
		defer fontFile.Close()
		fontContent, err := tpl.BGFileTemplate.Open(fmt.Sprintf("%s/%s", bgFolder, bg.Name()))
		if err != nil {
			return err
		}
		_, err = io.Copy(fontFile, fontContent)
		if err != nil {
			return err
		}
		return nil
	}
	return nil
}
func createFont() error {
	fillOutputFolder()
	safeCreateFolder(path.Join(outputPath, fontFolder))
	fonts, err := tpl.FontFileTemplate.ReadDir(fontFolder)
	if err != nil {
		return err
	}
	for _, font := range fonts {
		fontFile, err := os.Create(path.Join(outputPath, fontFolder, font.Name()))
		if err != nil {
			return err
		}
		defer fontFile.Close()
		fontContent, err := tpl.FontFileTemplate.Open(fmt.Sprintf("%s/%s", fontFolder, font.Name()))
		if err != nil {
			return err
		}
		_, err = io.Copy(fontFile, fontContent)
		if err != nil {
			return err
		}
		return nil
	}
	return nil
}

func createConfig() error {
	fillOutputFolder()
	safeCreateFolder(outputPath)
	configFile, err := os.Create(path.Join(outputPath, configFileName))
	if err != nil {
		return err
	}
	defer configFile.Close()
	configContent, err := tpl.ConfigFileTemplate.Open("config.yaml")
	if err != nil {
		return err
	}
	_, err = io.Copy(configFile, configContent)
	if err != nil {
		return err
	}
	return nil
}
func fillOutputFolder() {
	if outputPath == "" {
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)
		outputPath = path.Join(home, baseFolder)
	}
}
func safeCreateFolder(filePath string) {
	if !exists(filePath) {
		err := os.MkdirAll(filePath, os.ModePerm)
		cobra.CheckErr(err)
		fmt.Printf("%s created\n", filePath)
	}
}

func exists(path string) bool {
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false
	}
	return true
}
