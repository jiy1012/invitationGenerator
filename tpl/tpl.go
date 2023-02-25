package tpl

import "embed"

//go:embed "config.yaml"
var ConfigFileTemplate embed.FS

//go:embed "fonts/*"
var FontFileTemplate embed.FS

//go:embed "templates/background.png"
var BGFileTemplate embed.FS
