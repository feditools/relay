package static

import (
	"github.com/feditools/relay/internal/config"
	ihttp "github.com/feditools/relay/internal/http"
	"github.com/feditools/relay/internal/path"
	"github.com/feditools/relay/web"
	"io/fs"
	nethttp "net/http"
)

const DirStatic = "static"

func New() (*Module, error) {
	staticFS, err := fs.Sub(web.Files, DirStatic)
	if err != nil {
		return nil, err
	}

	return &Module{
		fs: staticFS,
	}, nil
}

type Module struct {
	fs fs.FS
}

// Name return the module name
func (m *Module) Name() string {
	return config.ServerRoleActivityPub
}

// SetServer adds a reference to the server to the module.
func (*Module) SetServer(_ *ihttp.Server) {}

// Route attaches routes to the web server.
func (m *Module) Route(s *ihttp.Server) error {
	s.PathPrefix(path.Static).Handler(nethttp.StripPrefix(path.Static, nethttp.FileServer(nethttp.FS(m.fs))))

	return nil
}
