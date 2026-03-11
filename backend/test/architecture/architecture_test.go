package architecture_test

import (
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

const (
	suffixUseCaseFile    = "-use-case.go"
	suffixControllerFile = "-controller.go"
	suffixRepositoryFile = "-repository.go"
	suffixFactoryFile    = "-factory.go"
	suffixUseCase        = "UseCase"
	suffixController     = "Controller"
	suffixRepository     = "Repository"
	suffixFactory        = "Factory"
	suffixInput          = "Input"
	suffixOutput         = "Output"
)

// TestArchitecture runs all custom architecture rules as file naming, struct naming, use case I/O, etc.
// Excludes test folder and *_test.go files from checking.
// Run from backend with "go test ./test/architecture/... -run TestArchitecture"
func TestArchitecture(t *testing.T) {
	root := findModuleRoot(t)
	dirs := []string{filepath.Join(root, "internal"), filepath.Join(root, "cmd")}

	for _, dir := range dirs {
		if _, err := os.Stat(dir); os.IsNotExist(err) {
			continue
		}

		err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}

			if info.IsDir() {
				if info.Name() == "test" {
					return filepath.SkipDir
				}

				return nil
			}

			if !strings.HasSuffix(info.Name(), ".go") || strings.HasSuffix(info.Name(), "_test.go") {
				return nil
			}

			rel, _ := filepath.Rel(root, path)
			rel = filepath.ToSlash(rel)

			checkFile(t, root, path, rel, info.Name())

			return nil
		})

		if err != nil {
			t.Fatalf("walk %s: %v", dir, err)
		}
	}
}

func findModuleRoot(t *testing.T) string {
	t.Helper()
	dir, err := os.Getwd()

	if err != nil {
		t.Fatalf("getwd: %v", err)
	}

	for {
		if _, err := os.Stat(filepath.Join(dir, "go.mod")); err == nil {
			return dir
		}

		parent := filepath.Dir(dir)
		if parent == dir {
			t.Fatal("go.mod not found")
		}

		dir = parent
	}
}

func checkFile(t *testing.T, root, path, relPath, name string) {
	t.Helper()

	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, path, nil, parser.ParseComments)

	if err != nil {
		t.Errorf("parse %s: %v", relPath, err)
		return
	}

	switch {
	case strings.Contains(relPath, "use-cases"):
		if !strings.HasSuffix(name, suffixUseCaseFile) {
			t.Errorf("%s: use-cases file must end with %s", relPath, suffixUseCaseFile)
		}

		checkUseCaseFile(t, relPath, f)

	case strings.Contains(relPath, "controllers") && strings.Contains(relPath, "infrastructure"):
		if !strings.HasSuffix(name, suffixControllerFile) {
			t.Errorf("%s: controllers file must end with %s", relPath, suffixControllerFile)
		}

		checkControllerFile(t, relPath, f)

	case strings.Contains(relPath, "repositories") && strings.Contains(relPath, "infrastructure"):
		if !strings.HasSuffix(name, suffixRepositoryFile) {
			t.Errorf("%s: infrastructure repositories file must end with %s", relPath, suffixRepositoryFile)
		}

		checkRepositoryFile(t, relPath, f)

	case strings.Contains(relPath, "factories"):
		if !strings.HasSuffix(name, suffixFactoryFile) {
			t.Errorf("%s: factories file must end with %s", relPath, suffixFactoryFile)
		}

		checkFactoryFile(t, relPath, f)

	case strings.Contains(relPath, "application") && strings.Contains(relPath, "repositories"):
		if !strings.HasSuffix(name, suffixRepositoryFile) {
			t.Errorf("%s: application repositories file must end with %s", relPath, suffixRepositoryFile)
		}
	}
}

func checkUseCaseFile(t *testing.T, relPath string, f *ast.File) {
	t.Helper()
	var hasUseCaseStruct bool
	var hasInput, hasOutput bool

	for _, decl := range f.Decls {
		switch d := decl.(type) {
		case *ast.GenDecl:
			if d.Tok != token.TYPE {
				continue
			}

			for _, spec := range d.Specs {
				ts, ok := spec.(*ast.TypeSpec)
				if !ok {
					continue
				}

				name := ts.Name.Name
				if strings.HasSuffix(name, suffixUseCase) && isStruct(ts.Type) {
					hasUseCaseStruct = true
				}

				if strings.HasSuffix(name, suffixInput) {
					hasInput = true
				}

				if strings.HasSuffix(name, suffixOutput) {
					hasOutput = true
				}
			}
		}
	}

	if !hasUseCaseStruct {
		t.Errorf("%s: use-case file must define a struct ending with %s", relPath, suffixUseCase)
	}

	if !hasInput {
		t.Errorf("%s: use-case file must define an input type (struct name ending with %s)", relPath, suffixInput)
	}

	if !hasOutput {
		t.Errorf("%s: use-case file must define an output type (struct name ending with %s)", relPath, suffixOutput)
	}
}

func checkControllerFile(t *testing.T, relPath string, f *ast.File) {
	t.Helper()
	var hasController bool

	for _, decl := range f.Decls {
		switch d := decl.(type) {
		case *ast.GenDecl:
			if d.Tok != token.TYPE {
				continue
			}

			for _, spec := range d.Specs {
				ts, ok := spec.(*ast.TypeSpec)
				if !ok {
					continue
				}

				if strings.HasSuffix(ts.Name.Name, suffixController) && isStruct(ts.Type) {
					hasController = true
					break
				}
			}
		}
	}

	if !hasController {
		t.Errorf("%s: controller file must define a struct ending with %s", relPath, suffixController)
	}
}

func checkRepositoryFile(t *testing.T, relPath string, f *ast.File) {
	t.Helper()
	var hasRepository bool

	for _, decl := range f.Decls {
		switch d := decl.(type) {
		case *ast.GenDecl:
			if d.Tok != token.TYPE {
				continue
			}

			for _, spec := range d.Specs {
				ts, ok := spec.(*ast.TypeSpec)
				if !ok {
					continue
				}

				if strings.HasSuffix(ts.Name.Name, suffixRepository) && isStruct(ts.Type) {
					hasRepository = true
					break
				}
			}
		}
	}

	if !hasRepository {
		t.Errorf("%s: repository file must define a struct ending with %s", relPath, suffixRepository)
	}
}

func checkFactoryFile(t *testing.T, relPath string, f *ast.File) {
	t.Helper()
	var hasFactoryFunc bool

	for _, decl := range f.Decls {
		fn, ok := decl.(*ast.FuncDecl)
		if !ok || !fn.Name.IsExported() {
			continue
		}

		if strings.HasSuffix(fn.Name.Name, suffixFactory) {
			hasFactoryFunc = true
			break
		}
	}

	if !hasFactoryFunc {
		t.Errorf("%s: factory file must define an exported function ending with %s", relPath, suffixFactory)
	}
}

func isStruct(e ast.Expr) bool {
	_, ok := e.(*ast.StructType)
	return ok
}
