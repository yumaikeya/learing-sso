package fonts

import (
	"embed"
)

//go:embed embed/*.ttf
var FontFiles embed.FS
