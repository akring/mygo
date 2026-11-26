//go:build darwin

package darwin

import (
	"bytes"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"testing"
)

// privateName is what a private selector of AppKit or WebKit looks like.
var privateName = regexp.MustCompile(`^_[a-z][A-Za-z0-9]*(:[A-Za-z0-9]*)*:?$`)

// TestSPIOnlyInSPIFiles keeps every private selector in spi.go, so that a
// build for the Mac App Store, which takes spi_appstore.go instead, names
// none: a string literal that looks like one anywhere else fails.
func TestSPIOnlyInSPIFiles(t *testing.T) {
	files, err := filepath.Glob("*.go")
	if err != nil {
		t.Fatal(err)
	}
	fset := token.NewFileSet()
	for _, name := range files {
		if strings.HasSuffix(name, "_test.go") || name == "spi.go" {
			continue
		}
		f, err := parser.ParseFile(fset, name, nil, 0)
		if err != nil {
			t.Fatal(err)
		}
		ast.Inspect(f, func(n ast.Node) bool {
			lit, ok := n.(*ast.BasicLit)
			if !ok || lit.Kind != token.STRING {
				return true
			}
			if s, err := strconv.Unquote(lit.Value); err == nil && privateName.MatchString(s) {
				t.Errorf("%s: private selector %q outside spi.go", fset.Position(lit.Pos()), s)
			}
			return true
		})
	}
}

// spiNames returns the private selectors spi.go names.
func spiNames(t *testing.T) []string {
	t.Helper()
	src, err := os.ReadFile("spi.go")
	if err != nil {
		t.Fatal(err)
	}
	var names []string
	for _, m := range regexp.MustCompile(`=\s*"(_[^"]+)"`).FindAllSubmatch(src, -1) {
		names = append(names, string(m[1]))
	}
	if len(names) < 30 {
		t.Fatalf("found %d private selectors in spi.go", len(names))
	}
	return names
}

// TestAppStoreBuildHasNoSPI builds examples with the mygo_appstore tag and
// fails when a binary holds the name of a private selector of spi.go,
// which App Review's scan of a binary would find; built without the tag,
// they must hold some, else the scan proves nothing.
func TestAppStoreBuildHasNoSPI(t *testing.T) {
	if testing.Short() {
		t.Skip("builds examples")
	}
	names := spiNames(t)
	dir := t.TempDir()
	for _, example := range []string{"hello", "counter-native"} {
		for _, appStore := range []bool{false, true} {
			out := filepath.Join(dir, example)
			args := []string{"build", "-trimpath", "-o", out}
			if appStore {
				out += "-appstore"
				args = []string{"build", "-trimpath", "-tags", "mygo_appstore", "-o", out}
			}
			cmd := exec.Command("go", append(args, "./examples/"+example)...)
			cmd.Dir = filepath.Join("..", "..")
			if b, err := cmd.CombinedOutput(); err != nil {
				t.Fatalf("go %s: %v\n%s", strings.Join(cmd.Args[1:], " "), err, b)
			}
			bin, err := os.ReadFile(out)
			if err != nil {
				t.Fatal(err)
			}
			var found []string
			for _, name := range names {
				if bytes.Contains(bin, []byte(name)) {
					found = append(found, name)
				}
			}
			switch {
			case appStore && len(found) > 0:
				t.Errorf("%s built for the Mac App Store holds %s", example, strings.Join(found, ", "))
			case !appStore && len(found) == 0:
				t.Errorf("%s built without mygo_appstore holds no private selector of spi.go", example)
			}
		}
	}
}
