package data

import _ "embed"

// since go 1.16

//go:embed data.json
var Feeds []byte
