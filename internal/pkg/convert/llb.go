package convert

import (
	"github.com/moby/buildkit/client/llb"

	"github.com/talos-systems/bldr/internal/pkg/environment"
	"github.com/talos-systems/bldr/internal/pkg/solver"
)

// BuildLLB translates package graph into LLB DAG
func BuildLLB(graph *solver.PackageGraph, options *environment.Options) (llb.State, error) {
	return NewGraphLLB(graph, options).Build()
}
