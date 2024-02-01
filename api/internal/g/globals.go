package g

import (
	"github.com/alexedwards/scs/v2"
	"github.com/unrolled/render"
)

type Data map[string]any

var renderer = render.New()
var JSON = renderer.JSON
var Text = renderer.Text

var Session = scs.New()
