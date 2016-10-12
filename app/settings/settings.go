package settings

import (
	"fmt"
	"github.com/lcaballero/kublai/conf"
	"github.com/lcaballero/mime"
	"strings"
)

type Settings interface {
	Endpoint() string
	RootPage() string
	AssetRoot() string
	LoginPage() string
	MimeType(ext string) (string, bool)
}

type settings struct {
	conf *conf.Config
	keys *conf.Keys
	exts mime.ExtensionToType
}

func NewSettings(c *conf.Config, k *conf.Keys) (Settings, error) {
	exts, err := mime.LoadMimeTypes(c.MimeTypes)
	if err != nil {
		return nil, err
	}
	s := &settings{
		conf: c,
		keys: k,
		exts: exts,
	}
	return s, nil
}

func (c *settings) AssetRoot() string {
	return c.conf.AssetRoot
}
func (c *settings) LoginPage() string {
	return c.conf.LoginPage
}
func (c *settings) MimeType(ext string) (string, bool) {
	mime, ok := c.exts[ext]
	return mime, ok
}
func (c *settings) RootPage() string {
	if strings.TrimSpace(c.conf.DefaultPage) != "" {
		return c.conf.DefaultPage
	} else {
		return "index.html"
	}
}

func (c *settings) Endpoint() string {
	// TODO: check the values of IP and PORT for something reasonable.
	return fmt.Sprintf("%s:%d", c.conf.Ip, c.conf.Port)
}
