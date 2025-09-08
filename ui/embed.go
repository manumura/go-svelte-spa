package ui

import (
	"embed"
	"io/fs"
)

var IndexFilePath = "ui/dist/index.html"

//go:embed all:dist
var DistDir embed.FS

// DistDirFS contains the embedded dist directory files (without the "dist" prefix)
var DistDirFS, _ = fs.Sub(DistDir, "dist")
