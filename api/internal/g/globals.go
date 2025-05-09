package g

import (
	"github.com/alexedwards/scs/v2"
	"github.com/unrolled/render"
)

type Data map[string]any

var renderer = render.New()
var JSON = renderer.JSON
var Text = renderer.Text

var Session *scs.SessionManager

type User struct {
	Username      string   `mapstructure:"username"`
	PasswordClear string   `mapstructure:"password-clear"`
	FileAreas     []string `mapstructure:"file-areas"`
}

type FileArea struct {
	Id   string `mapstructure:"id"`
	Name string `mapstructure:"name"`
	Path string `mapstructure:"path"`
}

var Users = []User{}
var FileAreas = []FileArea{}
