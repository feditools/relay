package main

import (
	"github.com/feditools/relay/internal/log"
)

type empty struct{}

var logger = log.WithPackageField(empty{})
