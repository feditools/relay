package web

import "embed"

// Files contains static files required by the application
//go:embed static/css/default.min.css
//go:embed static/css/error.min.css
//go:embed static/css/login.min.css
//go:embed template/*
var Files embed.FS
