package checker

import (
	"go/ast"
	"go/token"
	"go/types"
	"regexp"
	"strings"
	"unicode"

	"golang.org/x/tools/go/analysis"
)

var Analyzer = &analysis.Analyzer{
	Name: "loglint",
	Doc:  "checks slog and zap log messages",
	Run:  run,
}

var (
	sensitiveRegex = regexp.MustCompile(`(?i)(password|token|secret|apikey|api_key|pwd|key=)`)
)

func run(pass *analysis.Pass) (interface{}, error) {

	for _, file := range pass.Files {
		ast.Inspect(file, func(n ast.Node) bool {

			call, ok := n.(*ast.CallExpr)
			if !ok {
				return true
			}

			sel, ok := call.Fun.(*ast.SelectorExpr)
			if !ok {
				return true
			}

			// Получаем тип получателя (logger.Info -> logger)
			receiverType := pass.TypesInfo.TypeOf(sel.X)

			isLoggerMethod := false

			if receiverType != nil && isSupportedLogger(receiverType) {
				isLoggerMethod = true
			}

			if isSlogPackageCall(pass, sel) {
				isLoggerMethod = true
			}

			if !isLoggerMethod {
				return true
			}

			if len(call.Args) == 0 {
				return true
			}

			if len(call.Args) == 0 {
				return true
			}

			// передаем expr напрямую, extractString уже умеет обрабатывать BinaryExpr
			validateMessage(pass, call.Args[0], call.Args[0])
			return true
		})
	}

	return nil, nil
}

func isSupportedLogger(t types.Type) bool {

	ptr, ok := t.(*types.Pointer)
	if ok {
		t = ptr.Elem()
	}

	named, ok := t.(*types.Named)
	if !ok {
		return false
	}

	obj := named.Obj()
	if obj == nil || obj.Pkg() == nil {
		return false
	}

	pkgPath := obj.Pkg().Path()
	typeName := obj.Name()

	// slog.Logger
	if pkgPath == "log/slog" && typeName == "Logger" {
		return true
	}

	// zap.Logger
	if pkgPath == "go.uber.org/zap" && typeName == "Logger" {
		return true
	}

	return false
}

func validateMessage(pass *analysis.Pass, node ast.Node, expr ast.Expr) {
	msg := extractString(expr)
	if msg == "" {
		return
	}

	// sensitive data — проверяем первым
	if sensitiveRegex.MatchString(msg) && (strings.Contains(msg, ":") || strings.Contains(msg, "=")) {
		pass.Reportf(node.Pos(), "log message may contain sensitive data")
		return // больше не проверяем другие правила
	}

	// ASCII check
	for _, r := range msg {
		if r > unicode.MaxASCII {
			pass.Reportf(node.Pos(), "log message must contain only english ascii characters with lowercase letters and allowed symbols")
			return
		}
	}

	// allowed characters + lowercase
	for _, r := range msg {
		if !(unicode.IsLower(r) || unicode.IsDigit(r) || strings.ContainsRune(" .,:-", r)) {
			pass.Reportf(node.Pos(), "log message must contain only lowercase english letters and allowed characters")
			return
		}
	}
}

func isSlogPackageCall(pass *analysis.Pass, sel *ast.SelectorExpr) bool {
	ident, ok := sel.X.(*ast.Ident)
	if !ok {
		return false
	}

	obj := pass.TypesInfo.Uses[ident]
	if obj == nil {
		return false
	}

	pkgName, ok := obj.(*types.PkgName)
	if !ok {
		return false
	}

	return pkgName.Imported().Path() == "log/slog"
}

func extractString(expr ast.Expr) string {
	switch v := expr.(type) {
	case *ast.BasicLit:
		if v.Kind == token.STRING {
			return strings.Trim(v.Value, `"`)
		}
	case *ast.BinaryExpr:
		return extractString(v.X) + extractString(v.Y)
	}
	return ""
}
