package template

const (
	// TableQueryStruct table query struct
	TableQueryGenericStruct = `
	{{.QueryStructComment}}
	type {{.QueryStructName}}Fields struct {
		` + fieldsGeneric + `
	}

	` + defineMethodStructGeneric + createMethodGeneric + tableMethodGeneric + asMethondGeneric + updateFieldMethodGeneric + getFieldMethod + fillFieldMapMethodGeneric + cloneMethodGeneric + replaceMethodGeneric + relationship
)

const (
	createMethodGeneric = `
	func new{{.ModelStructName}}(db *gorm.DB, opts ...gen.DOOption) {{.QueryStructName}} {
		_{{.QueryStructName}} := {{.QueryStructName}}{}
	
		_{{.QueryStructName}}.DO.UseDB(db,opts...)
		_{{.QueryStructName}}.DO.UseModel(&{{.StructInfo.Package}}.{{.StructInfo.Type}}{})
	
		tableName := _{{.QueryStructName}}.DO.TableName()
		_{{$.QueryStructName}}.Fields.ALL = field.NewAsterisk(tableName)
		{{range .Fields -}}
		{{if not .IsRelation -}}
			{{- if .ColumnName -}}_{{$.QueryStructName}}.Fields.{{.Name}} = field.New{{.GenType}}(tableName, "{{.ColumnName}}"){{- end -}}
		{{- else -}}
			_{{$.QueryStructName}}.Fields.{{.Relation.Name}} = {{$.QueryStructName}}{{.Relation.RelationshipName}}{{.Relation.Name}}{
				db: db.Session(&gorm.Session{}),

				{{.Relation.StructFieldInit}}
			}
		{{end}}
		{{end}}

		_{{$.QueryStructName}}.fillFieldMap()
		return _{{.QueryStructName}}
	}
	`
	fieldsGeneric = `
	ALL field.Asterisk
	{{range .Fields -}}
		{{if not .IsRelation -}}
			{{if .MultilineComment -}}
			/*
{{.ColumnComment}}
    		*/
			{{end -}}
			{{- if .ColumnName -}}{{.Name}} field.{{.GenType}}{{if not .MultilineComment}}{{if .ColumnComment}}// {{.ColumnComment}}{{end}}{{end}}{{- end -}}
		{{- else -}}
			{{.Relation.Name}} {{$.QueryStructName}}{{.Relation.RelationshipName}}{{.Relation.Name}}
		{{end}}
	{{end}}
`
	tableMethodGeneric = `
func ({{.S}} {{.QueryStructName}}) Table(newTableName string) *{{.QueryStructName}} { 
	{{.S}}.DO.UseTable(newTableName)
	return {{.S}}.updateTableName(newTableName)
}
`

	asMethondGeneric = `	
func ({{.S}} {{.QueryStructName}}) As(alias string) *{{.QueryStructName}} { 
	{{.S}}.DO = *({{.S}}.DO.As(alias).(*gen.DO))
	return {{.S}}.updateTableName(alias)
}
`
	updateFieldMethodGeneric = `
func ({{.S}} *{{.QueryStructName}}) updateTableName(table string) *{{.QueryStructName}} { 
	{{.S}}.Fields.ALL = field.NewAsterisk(table)
	{{range .Fields -}}
	{{if not .IsRelation -}}
		{{- if .ColumnName -}}{{$.S}}.Fields.{{.Name}} = field.New{{.GenType}}(table, "{{.ColumnName}}"){{- end -}}
	{{end}}
	{{end}}

	{{.S}}.fillFieldMap()
	return {{.S}}
}
`

	cloneMethodGeneric = `
func ({{.S}} {{.QueryStructName}}) clone(db *gorm.DB) {{.QueryStructName}} {
	{{.S}}.DO.ReplaceConnPool(db.Statement.ConnPool){{range .Fields }}{{if .IsRelation}}
  {{$.S}}.{{.Relation.Name}}.db = db.Session(&gorm.Session{Initialized: true})
  {{$.S}}.{{.Relation.Name}}.db.Statement.ConnPool = db.Statement.ConnPool{{end}}{{end}}
	return {{.S}}
}
`
	replaceMethodGeneric = `
func ({{.S}} {{.QueryStructName}}) replaceDB(db *gorm.DB) {{.QueryStructName}} {
	{{.S}}.DO.ReplaceDB(db){{range .Fields}}{{if .IsRelation}}
  {{$.S}}.{{.Relation.Name}}.db = db.Session(&gorm.Session{}){{end}}{{end}}
	return {{.S}}
}
`
	getFieldMethodGeneric = `
func ({{.S}} *{{.QueryStructName}}) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := {{.S}}.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe,ok := _f.(field.OrderExpr)
	return _oe,ok
}
`
	relationshipGeneric = `{{range .Fields}}{{if .IsRelation}}` +
		`{{- $relation := .Relation }}{{- $relationship := $relation.RelationshipName}}` +
		relationStruct + relationTx +
		`{{end}}{{end}}`

	defineMethodStructGeneric = `
	type {{.QueryStructName}}Base = {{.GenericTypeName}}[{{.QueryStructName}}Fields, {{.StructInfo.Package}}.{{.StructInfo.Type}}]

	type {{.QueryStructName}} struct {
		{{.QueryStructName}}Base
		fieldMap map[string]field.Expr
	}`

	fillFieldMapMethodGeneric = `
func ({{.S}} *{{.QueryStructName}}) fillFieldMap() {
	{{.S}}.fieldMap =  make(map[string]field.Expr, {{len .Fields}})
	{{range .Fields -}}
	{{if not .IsRelation -}}
		{{- if .ColumnName -}}{{$.S}}.fieldMap["{{.ColumnName}}"] = {{$.S}}.Fields.{{.Name}}{{- end -}}
	{{end}}
	{{end -}}
}
`
)

const (
	relationStructGeneric = `
type {{$.QueryStructName}}{{$relationship}}{{$relation.Name}} struct{
	db *gorm.DB
	
	field.RelationField
	
	{{$relation.StructField}}
}

func (a {{$.QueryStructName}}{{$relationship}}{{$relation.Name}}) Where(conds ...field.Expr) *{{$.QueryStructName}}{{$relationship}}{{$relation.Name}} {
	if len(conds) == 0 {
		return &a
	}

	exprs := make([]clause.Expression, 0, len(conds))
	for _, cond := range conds {
		exprs = append(exprs, cond.BeCond().(clause.Expression))
	}
	a.db = a.db.Clauses(clause.Where{Exprs: exprs})
	return &a
}

func (a {{$.QueryStructName}}{{$relationship}}{{$relation.Name}}) WithContext(ctx context.Context) *{{$.QueryStructName}}{{$relationship}}{{$relation.Name}} {
	a.db = a.db.WithContext(ctx)
	return &a
}

func (a {{$.QueryStructName}}{{$relationship}}{{$relation.Name}}) Session(session *gorm.Session) *{{$.QueryStructName}}{{$relationship}}{{$relation.Name}} {
	a.db = a.db.Session(session)
	return &a
}

func (a {{$.QueryStructName}}{{$relationship}}{{$relation.Name}}) Model(m *{{$.StructInfo.Package}}.{{$.StructInfo.Type}}) *{{$.QueryStructName}}{{$relationship}}{{$relation.Name}}Tx {
	return &{{$.QueryStructName}}{{$relationship}}{{$relation.Name}}Tx{a.db.Model(m).Association(a.Name())}
}

func (a {{$.QueryStructName}}{{$relationship}}{{$relation.Name}}) Unscoped() *{{$.QueryStructName}}{{$relationship}}{{$relation.Name}} {
	a.db = a.db.Unscoped()
	return &a
}

`
	relationTxGeneric = `
type {{$.QueryStructName}}{{$relationship}}{{$relation.Name}}Tx struct{ tx *gorm.Association }

func (a {{$.QueryStructName}}{{$relationship}}{{$relation.Name}}Tx) Find() (result {{if eq $relationship "HasMany" "ManyToMany"}}[]{{end}}*{{$relation.Type}}, err error) {
	return result, a.tx.Find(&result)
}

func (a {{$.QueryStructName}}{{$relationship}}{{$relation.Name}}Tx) Append(values ...*{{$relation.Type}}) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Append(targetValues...)
}

func (a {{$.QueryStructName}}{{$relationship}}{{$relation.Name}}Tx) Replace(values ...*{{$relation.Type}}) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Replace(targetValues...)
}

func (a {{$.QueryStructName}}{{$relationship}}{{$relation.Name}}Tx) Delete(values ...*{{$relation.Type}}) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Delete(targetValues...)
}

func (a {{$.QueryStructName}}{{$relationship}}{{$relation.Name}}Tx) Clear() error {
	return a.tx.Clear()
}

func (a {{$.QueryStructName}}{{$relationship}}{{$relation.Name}}Tx) Count() int64 {
	return a.tx.Count()
}

func (a {{$.QueryStructName}}{{$relationship}}{{$relation.Name}}Tx) Unscoped() *{{$.QueryStructName}}{{$relationship}}{{$relation.Name}}Tx {
	a.tx = a.tx.Unscoped()
	return &a
}
`
)
