package k3sscripts

import "embed"

type Context struct {
	manifests embed.FS
	tempFiles []string // Tracks temp files for cleanup
}

func NewContext(m embed.FS) *Context {
	return &Context{
		manifests: m,
		tempFiles: []string{},
	}
}
