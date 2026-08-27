// callgraph prints a bounded set of deterministic edges from a Go CHA call graph.
//
// Run it from a Go module with golang.org/x/tools dependencies. Use -dir to
// select a different module only after this program has been compiled.
package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"sort"
	"strings"

	"golang.org/x/tools/go/callgraph"
	"golang.org/x/tools/go/callgraph/cha"
	"golang.org/x/tools/go/packages"
	"golang.org/x/tools/go/ssa"
	"golang.org/x/tools/go/ssa/ssautil"
)

const usage = `Usage:
  go run callgraph.go -package <pattern> -target <ssa-function> [options]

Options:
  -dir <path>                 Go module directory (default: current directory)
  -package <pattern>          Go package pattern to load (default: .)
  -target <ssa-function>      Exact SSA function name to inspect
  -direction callers|callees|both
                              Relationship direction to traverse (default: callers)
  -depth <n>                  Maximum number of call edges from the target (default: 4)
  -list                       Print available SSA function names and exit

Examples:
  go run callgraph.go -dir /src/project -package ./internal/store -list
  go run callgraph.go -dir /src/project -package ./internal/store \
    -target '(*example.com/project/internal/store.Store).Create' -direction callers -depth 4
`

func main() {
	var (
		dir       = flag.String("dir", ".", "Go module directory")
		pattern   = flag.String("package", ".", "Go package pattern to load")
		target    = flag.String("target", "", "exact SSA function name")
		direction = flag.String("direction", "callers", "callers, callees, or both")
		depth     = flag.Int("depth", 4, "maximum traversal depth")
		list      = flag.Bool("list", false, "list SSA function names and exit")
	)
	flag.Usage = func() { fmt.Fprint(os.Stderr, usage) }
	flag.Parse()

	if *depth < 1 {
		exitf("-depth must be at least 1")
	}
	if *direction != "callers" && *direction != "callees" && *direction != "both" {
		exitf("-direction must be callers, callees, or both")
	}
	if !*list && *target == "" {
		exitf("-target is required unless -list is set")
	}

	graph, err := loadCHA(*dir, *pattern)
	if err != nil {
		exitf("build call graph: %v", err)
	}

	if *list {
		for _, name := range functionNames(graph) {
			fmt.Println(name)
		}
		return
	}

	node, err := findNode(graph, *target)
	if err != nil {
		exitf("find target: %v", err)
	}

	fmt.Printf("TARGET %s\n", node.Func)
	if *direction == "callers" || *direction == "both" {
		fmt.Println("CALLERS")
		printEdges(node, *depth, callers)
	}
	if *direction == "callees" || *direction == "both" {
		fmt.Println("CALLEES")
		printEdges(node, *depth, callees)
	}
}

func loadCHA(dir, pattern string) (*callgraph.Graph, error) {
	cfg := &packages.Config{
		Mode: packages.NeedName | packages.NeedFiles | packages.NeedCompiledGoFiles |
			packages.NeedImports | packages.NeedDeps | packages.NeedSyntax |
			packages.NeedTypes | packages.NeedTypesInfo,
		Dir: dir,
	}
	pkgs, err := packages.Load(cfg, pattern)
	if err != nil {
		return nil, err
	}
	if len(pkgs) == 0 {
		return nil, errors.New("package pattern matched no packages")
	}
	if packages.PrintErrors(pkgs) > 0 {
		return nil, errors.New("package loading failed")
	}
	prog, _ := ssautil.AllPackages(pkgs, ssa.InstantiateGenerics)
	prog.Build()
	return cha.CallGraph(prog), nil
}

func functionNames(graph *callgraph.Graph) []string {
	names := make([]string, 0, len(graph.Nodes))
	for fn := range graph.Nodes {
		if fn != nil {
			names = append(names, fn.String())
		}
	}
	sort.Strings(names)
	return names
}

func findNode(graph *callgraph.Graph, target string) (*callgraph.Node, error) {
	var matches []*callgraph.Node
	for fn, node := range graph.Nodes {
		if fn != nil && fn.String() == target {
			matches = append(matches, node)
		}
	}
	if len(matches) == 1 {
		return matches[0], nil
	}
	if len(matches) > 1 {
		return nil, fmt.Errorf("target %q is ambiguous", target)
	}

	var suggestions []string
	for _, name := range functionNames(graph) {
		if strings.Contains(name, target) {
			suggestions = append(suggestions, name)
		}
	}
	if len(suggestions) == 0 {
		return nil, fmt.Errorf("target %q was not found; use -list to inspect available names", target)
	}
	return nil, fmt.Errorf("target %q was not found; possible matches:\n  %s", target, strings.Join(suggestions, "\n  "))
}

type traversal func(*callgraph.Node) []*callgraph.Edge

func callers(node *callgraph.Node) []*callgraph.Edge {
	return node.In
}

func callees(node *callgraph.Node) []*callgraph.Edge {
	return node.Out
}

func printEdges(target *callgraph.Node, depth int, next traversal) {
	edges := make(map[string]edge)
	collectEdges(target, depth, next, map[*callgraph.Node]bool{}, edges)

	ordered := make([]edge, 0, len(edges))
	for _, item := range edges {
		ordered = append(ordered, item)
	}
	sort.Slice(ordered, func(i, j int) bool {
		if ordered[i].caller == ordered[j].caller {
			return ordered[i].callee < ordered[j].callee
		}
		return ordered[i].caller < ordered[j].caller
	})
	for _, item := range ordered {
		fmt.Printf("%s -> %s\n", item.caller, item.callee)
	}
}

type edge struct {
	caller string
	callee string
}

func collectEdges(node *callgraph.Node, depth int, next traversal, path map[*callgraph.Node]bool, edges map[string]edge) {
	if depth == 0 || path[node] {
		return
	}
	path[node] = true
	defer delete(path, node)

	for _, item := range next(node) {
		if item == nil || item.Caller == nil || item.Callee == nil || item.Caller.Func == nil || item.Callee.Func == nil {
			continue
		}
		value := edge{caller: item.Caller.Func.String(), callee: item.Callee.Func.String()}
		edges[value.caller+"\x00"+value.callee] = value

		if node == item.Callee {
			collectEdges(item.Caller, depth-1, next, path, edges)
		} else {
			collectEdges(item.Callee, depth-1, next, path, edges)
		}
	}
}

func exitf(format string, args ...any) {
	fmt.Fprintf(os.Stderr, "callgraph: "+format+"\n", args...)
	os.Exit(2)
}
