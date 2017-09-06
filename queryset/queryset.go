package queryset

import (
	"bytes"
	"fmt"
	"go/types"
	"io"
	"sort"
	"strings"
	"text/template"

	"golang.org/x/tools/go/loader"

	"github.com/jirfag/go-queryset/parser"
)

var qsTmpl = template.Must(
	template.New("generator").
		Funcs(template.FuncMap{
			"title": strings.Title,
		}).
		Parse(qsCode),
)

type querySetStructConfig struct {
	StructName string
	Name       string
	Methods    methodsSlice
}

type methodsSlice []method

func (s methodsSlice) Len() int { return len(s) }
func (s methodsSlice) Less(i, j int) bool {
	return strings.Compare(s[i].GetMethodName(), s[j].GetMethodName()) < 0
}
func (s methodsSlice) Swap(i, j int) { s[i], s[j] = s[j], s[i] }

type querySetStructConfigSlice []querySetStructConfig

func (s querySetStructConfigSlice) Len() int { return len(s) }
func (s querySetStructConfigSlice) Less(i, j int) bool {
	return strings.Compare(s[i].Name, s[j].Name) < 0
}
func (s querySetStructConfigSlice) Swap(i, j int) { s[i], s[j] = s[j], s[i] }

func getMethodsForField(pkgInfo *loader.PackageInfo, name string, typ fmt.Stringer, originalTypeName string) []method {
	typeName := typ.String()
	if originalTypeName != "" {
		// it's needed to preserver typedef's original name
		typeName = originalTypeName
	}

	basicTypeMethods := []method{
		newBinaryFilterMethod("eq", name, typeName),
		newBinaryFilterMethod("ne", name, typeName),
	}
	switch t := typ.(type) {
	case *types.Basic:
		if t.Info()&types.IsNumeric != 0 {
			return append(basicTypeMethods,
				newBinaryFilterMethod("lt", name, typeName),
				newBinaryFilterMethod("gt", name, typeName),
				newBinaryFilterMethod("lte", name, typeName),
				newBinaryFilterMethod("gte", name, typeName),
				newOrderByMethod(name))
		}
		// it's a string
		return basicTypeMethods
	case *types.Named:
		originalTypeName := t.Obj().Name()
		if t.Obj().Pkg() != pkgInfo.Pkg {
			originalTypeName = fmt.Sprintf("%s.%s", t.Obj().Pkg().Name(),
				originalTypeName)
		}
		return getMethodsForField(pkgInfo, name, t.Underlying(), originalTypeName)
	case *types.Pointer:
		if _, ok := t.Elem().Underlying().(*types.Struct); ok {
			if t.Underlying().String() == "*time.Time" {
				// XXX: maybe check struct tags?
				return nil
			}

			// Pointer usually used for assocations => preload them
			return []method{newPreloadMethod(name)}
		}
		// don't generate preloaders for usual pointers like DeletedAt *time.Time
		return nil
	default:
		// no filtering is needed
		return nil
	}
}

func getQuerySetFieldMethods(pkgInfo *loader.PackageInfo, fields []parser.StructField) []method {
	ret := []method{}
	for _, f := range fields {
		methods := getMethodsForField(pkgInfo, f.Name, f.Type, "")
		ret = append(ret, methods...)
	}

	return ret
}

// GenerateQuerySetsForStructs is an internal method to retrieve querysets
// generated code from parsed structs
func GenerateQuerySetsForStructs(pkgInfo *loader.PackageInfo, structs parser.ParsedStructs) (io.Reader, error) {
	querySetStructConfigs := querySetStructConfigSlice{}

	for structTypeName, ps := range structs {
		doc := ps.Doc
		if doc == nil {
			continue
		}

		ok := false
		for _, c := range doc.List {
			parts := strings.Split(strings.TrimSpace(c.Text), ":")
			ok = len(parts) == 2 &&
				strings.TrimSpace(strings.TrimPrefix(parts[0], "//")) == "gen" &&
				strings.TrimSpace(parts[1]) == "qs"
			if ok {
				break
			}
		}
		if !ok {
			continue
		}

		methods := []method{newLimitMethod(), newAllMethod(structTypeName),
			newOneMethod(structTypeName)}
		fieldMethods := getQuerySetFieldMethods(pkgInfo, ps.Fields)
		methods = append(methods, fieldMethods...)

		qsConfig := querySetStructConfig{
			StructName: structTypeName,
			Name:       structTypeName + "QuerySet",
			Methods:    methods,
		}
		sort.Sort(qsConfig.Methods)
		querySetStructConfigs = append(querySetStructConfigs, qsConfig)
	}

	if len(querySetStructConfigs) == 0 {
		return nil, nil
	}

	sort.Sort(querySetStructConfigs)

	var b bytes.Buffer
	err := qsTmpl.Execute(&b, struct {
		Configs querySetStructConfigSlice
	}{
		Configs: querySetStructConfigs,
	})

	if err != nil {
		return nil, fmt.Errorf("can't generate structs query sets: %s", err)
	}

	return &b, nil
}

const qsCode = `
// ===== BEGIN of all query sets

{{ range .Configs }}
  // ===== BEGIN of query set {{ .Name }}

	// {{ .Name }} is an queryset type for {{ .StructName }}
  type {{ .Name }} struct {
	  db *gorm.DB
  }

  // New{{ .Name }} constructs new {{ .Name }}
  func New{{ .Name }}(db *gorm.DB) {{ .Name }} {
	  return {{ .Name }}{
		  db: db,
	  }
  }

  {{ $qSTypeName := .Name }}

	{{ range .Methods }}
		{{ .GetDoc .GetMethodName }}
		func (qs {{ $qSTypeName }}) {{ .GetMethodName }}({{ .GetArgsDeclaration }})
		{{- .GetReturnValuesDeclaration $qSTypeName }} {
      {{ .GetBody }}
		}
	{{ end }}

  // ===== END of query set {{ .Name }}
{{ end }}

// ===== END of all query sets
`
