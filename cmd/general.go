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
	"html/template"
	"os"
	"path"

	"github.com/spf13/cobra"
)

var templateType, outputPath string

const configFileName = "config.yaml"
const baseFolder = ".ig"

// generalCmd represents the general command
var generalCmd = &cobra.Command{
	Use:   "general",
	Short: "生成文件模板",
	Long:  `生成配置文件、处理文件等模板`,
	Run: func(cmd *cobra.Command, args []string) {
		//fmt.Println("templateType:", templateType, "outputPath:", outputPath)
		switch templateType {
		case "config":
			cobra.CheckErr(createConfig())
			fmt.Printf("%s created at %s\n", configFileName, outputPath)
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
	generalCmd.Flags().StringVarP(&templateType, "template", "t", "config", "config template")
	generalCmd.Flags().StringVarP(&outputPath, "output", "o", "", "output path")
}

func createConfig() error {
	if outputPath == "" {
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)
		outputPath = path.Join(home, baseFolder)
	}
	err := os.MkdirAll(outputPath, os.ModePerm)
	fmt.Sprintln(outputPath, err)
	if err != nil {
		return err
	}
	configFile, err := os.Create(path.Join(outputPath, configFileName))
	if err != nil {
		return err
	}
	defer configFile.Close()
	configTemplate := template.Must(template.New("config").Parse(string(tpl.ConfigTemplate())))
	err = configTemplate.Execute(configFile, nil)
	if err != nil {
		return err
	}
	return nil
}
