package static

import "embed"

//go:embed *
var StaticFS embed.FS
