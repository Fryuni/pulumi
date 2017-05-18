// Copyright 2017 Pulumi, Inc. All rights reserved.

package core

import (
	"github.com/pulumi/lumi/pkg/diag"
)

// Phase represents a compiler phase.
type Phase interface {
	// Diag fetches the diagnostics sink used by this compiler pass.
	Diag() diag.Sink
}
