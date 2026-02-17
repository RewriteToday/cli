package cmd

import (
	"go/parser"
	"go/token"
	"io/fs"
	"path/filepath"
	"runtime"
	"slices"
	"strings"
	"testing"
)

const (
	huhImportPath          = "github.com/charmbracelet/huh"
	legacyPromptImportPath = "github.com/rewritestudios/cli/internal/prompt"
	allowedHuhImportFile   = "internal/style/style.go"
)

func TestCmdPackageImportBoundaries(t *testing.T) {
	repoRoot := repoRoot(t)
	goFiles := collectGoFiles(t, filepath.Join(repoRoot, "cmd"))

	for _, file := range goFiles {
		imports := readImports(t, file)

		if slices.Contains(imports, huhImportPath) {
			t.Fatalf("cmd file must not import %q: %s", huhImportPath, relPath(repoRoot, file))
		}

		if slices.Contains(imports, legacyPromptImportPath) {
			t.Fatalf("cmd file must not import %q: %s", legacyPromptImportPath, relPath(repoRoot, file))
		}
	}
}

func TestHuhIntegrationIsCentralizedInStyleGo(t *testing.T) {
	repoRoot := repoRoot(t)
	goFiles := collectGoFiles(t, repoRoot)

	var huhImportFiles []string
	for _, file := range goFiles {
		imports := readImports(t, file)
		if slices.Contains(imports, huhImportPath) {
			huhImportFiles = append(huhImportFiles, relPath(repoRoot, file))
		}
	}

	if len(huhImportFiles) != 1 || huhImportFiles[0] != allowedHuhImportFile {
		t.Fatalf("expected only %q to import %q, got: %v", allowedHuhImportFile, huhImportPath, huhImportFiles)
	}
}

func repoRoot(t *testing.T) string {
	t.Helper()

	_, file, _, ok := runtime.Caller(0)
	if !ok {
		t.Fatal("failed to resolve caller path")
	}

	return filepath.Dir(filepath.Dir(file))
}

func collectGoFiles(t *testing.T, root string) []string {
	t.Helper()

	var files []string

	err := filepath.WalkDir(root, func(path string, d fs.DirEntry, walkErr error) error {
		if walkErr != nil {
			return walkErr
		}

		if d.IsDir() {
			name := d.Name()
			if name == ".git" || name == "vendor" {
				return filepath.SkipDir
			}
			return nil
		}

		if !strings.HasSuffix(path, ".go") {
			return nil
		}

		files = append(files, path)
		return nil
	})
	if err != nil {
		t.Fatalf("failed to walk %s: %v", root, err)
	}

	return files
}

func readImports(t *testing.T, path string) []string {
	t.Helper()

	parsed, err := parser.ParseFile(token.NewFileSet(), path, nil, parser.ImportsOnly)
	if err != nil {
		t.Fatalf("failed to parse imports from %s: %v", path, err)
	}

	imports := make([]string, len(parsed.Imports))
	for i, imp := range parsed.Imports {
		imports[i] = strings.Trim(imp.Path.Value, "\"")
	}

	return imports
}

func relPath(root, path string) string {
	rel, err := filepath.Rel(root, path)
	if err != nil {
		return path
	}
	return filepath.ToSlash(rel)
}
