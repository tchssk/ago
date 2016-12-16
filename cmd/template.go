package cmd

import (
	"sort"
	"text/template"

	"github.com/goadesign/goa/design"
)

const (
	goHeaderT = `
package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)
`
	allT = `
{{template "goHeader" .}}
{{template "api" .}}
`

	// Components that have single value.
	basePathT            = `{{if .BasePath}}BasePath({{printf "%q" .BasePath}}){{end}}`                                  // This template expects APIDefinition or ResourceDefinition.
	canonicalActionNameT = `{{if .CanonicalActionName}}CanonicalActionName({{printf "%q" .CanonicalActionName}}){{end}}` // This template expects ResourceDefinition.
	contentTypeT         = `{{if .ContentType}}ContentType({{printf "%q" .ContentType}}){{end}}`                         // This template expects MediaTypeDefinition.
	credentialsT         = `{{if .Credentials}}Credentials(){{end}}`                                                     // This template expects CORSDefinition.
	descriptionT         = `{{if .Description}}Description({{printf "%q" .Description}}){{end}}`                         // This template expects APIDefinition or DocsDefinition or ResourceDefinition or ResponseDefinition.
	emailT               = `{{if .Email}}Email({{printf "%q" .Email}}){{end}}`                                           // This template expects ContactDefinition.
	functionT            = `{{if .Function}}Function({{printf "%q" .Function}}){{end}}`                                  // This template expects EncodingDefinition.
	hostT                = `{{if .Host}}Host({{printf "%q" .Host}}){{end}}`                                              // This template expects APIDefinition.
	maxAgeT              = `{{if .MaxAge}}MaxAge({{.MaxAge}}){{end}}`                                                    // This template expects CORSDefinition.
	maxLengthT           = `{{if .MaxLength}}MaxLength({{.MaxLength}}){{end}}`                                           // This template expects ValidationDefinition.
	minLengthT           = `{{if .MinLength}}MinLength({{.MinLength}}){{end}}`                                           // This template expects ValidationDefinition.
	nameT                = `{{if .Name}}Name({{printf "%q" .Name}}){{end}}`                                              // This template expects APIDefinition or ContactDefinition or LicenseDefinition.
	packageT             = `{{if .PackagePath}}Package({{printf "%q" .PackagePath}}){{end}}`                             // This template expects EncodingDefinition.
	statusT              = `{{if .Status}}Status({{.Status}}){{end}}`                                                    // This template expects ResponseDefinition.
	termsOfServiceT      = `{{if .TermsOfService}}TermsOfService({{printf "%q" .TermsOfService}}){{end}}`                // This template expects APIDefinition.
	titleT               = `{{if .Title}}Title({{printf "%q" .Title}}){{end}}`                                           // This template expects APIDefinition.
	typeNameT            = `{{if .TypeName}}TypeName({{printf "%q" .TypeName}}){{end}}`                                  // This template expects MediaTypeDefinition or UserTypeDefinition.
	urlT                 = `{{if .URL}}URL({{printf "%q" .URL}}){{end}}`                                                 // This template expects ContactDefinition or LicenseDefinition or DocsDefinition.
	versionT             = `{{if .Version}}Version({{printf "%q" .Version}}){{end}}`                                     // This template expects APIDefinition.

	// Components that have multiple values.
	exposeT = `{{if .Exposed}}Expose({{if (eq (len .Exposed) 1)}}{{range .Exposed}}{{printf "%q" .}}{{end}}){{else}}
{{range .Exposed}}{{printf "%q" .}},
{{end}}){{end}}{{end}}` // This template expects CORSDefinition.
	methodsT = `{{if .Methods}}Methods({{if (eq (len .Methods) 1)}}{{range .Methods}}{{printf "%q" .}}{{end}}){{else}}
{{range .Methods}}{{printf "%q" .}},
{{end}}){{end}}{{end}}` // This template expects CORSDefinition.
	schemeT = `{{if .Schemes}}Scheme({{if (eq (len .Schemes) 1)}}{{range .Schemes}}{{printf "%q" .}}{{end}}){{else}}
{{range .Schemes}}{{printf "%q" .}},
{{end}}){{end}}{{end}}` // This template expects APIDefinition or ResourceDefinition or ActionDefinition.

	// Containers.
	actionT = `{{if .Actions}}{{$actions := .Actions}}{{$keys := keys .Actions}}{{range $keys}}{{if (not (eq (index $keys 0) .))}}
{{end}}{{with index $actions .}}Action({{printf "%q" .Name}}, func() {
{{if .Description}}{{template "description" .}}
{{end}}{{if .Docs}}{{template "docs" .}}
{{end}}{{if .Schemes}}{{template "scheme" .}}
{{end}}{{if .Routes}}{{template "routing" .}}
{{end}}{{if .Payload}}{{template "payload" .}}
{{end}}}){{end}}{{end}}{{end}}`
	apiT = `{{if .}}var _ = API({{if .Name}}{{printf "%q" .Name}}{{else}}""{{end}}, func() {
{{if .Title}}{{template "title" .}}
{{end}}{{if .Description}}{{template "description" .}}
{{end}}{{if .Version}}{{template "version" .}}
{{end}}{{if .TermsOfService}}{{template "termsOfService" .}}
{{end}}{{if .Contact}}{{template "contact" .}}
{{end}}{{if .License}}{{template "license" .}}
{{end}}{{if .Docs}}{{template "docs" .}}
{{end}}{{if .Host}}{{template "host" .}}
{{end}}{{if .Schemes}}{{template "scheme" .}}
{{end}}{{if .BasePath}}{{template "basePath" .}}
{{end}}{{if .Origins}}{{template "origin" .}}
{{end}}{{if .Consumes}}{{template "consumes" .}}
{{end}}{{if .Produces}}{{template "produces" .}}
{{end}}}){{end}}`
	connectT  = `{{if and .Verb .Path}}{{if (eq .Verb "CONNECT")}}CONNECT({{printf "%q" .Path}}){{end}}{{end}}`
	consumesT = `{{if .Consumes}}{{range $index, $element := .Consumes}}{{with $element}}{{if (not (eq $index 0))}}
{{end}}Consumes({{if (or (gt (len .MIMETypes) 1) .Function .PackagePath)}}
{{range .MIMETypes}}{{printf "%q" .}},
{{end}}{{if (or .Function .PackagePath)}}func() {
{{if .Function}}{{template "function" .}}
{{end}}{{if .PackagePath}}{{template "package" .}}
{{end}}},
{{end}}{{else}}{{range .MIMETypes}}{{printf "%q" .}}{{end}}{{end}}){{end}}{{end}}{{end}}`
	contactT = `{{if .Contact}}{{with .Contact}}Contact(func() {
{{if .Name}}{{template "name" .}}
{{end}}{{if .Email}}{{template "email" .}}
{{end}}{{if .URL}}{{template "url" .}}
{{end}}}){{end}}{{end}}`
	deleteT = `{{if and .Verb .Path}}{{if (eq .Verb "DELETE")}}DELETE({{printf "%q" .Path}}){{end}}{{end}}`
	docsT   = `{{if .Docs}}{{with .Docs}}Docs(func() {
{{if .Description}}{{template "description" .}}
{{end}}{{if .URL}}{{template "url" .}}
{{end}}}){{end}}{{end}}`
	getT     = `{{if and .Verb .Path}}{{if (eq .Verb "GET")}}GET({{printf "%q" .Path}}){{end}}{{end}}`
	headT    = `{{if and .Verb .Path}}{{if (eq .Verb "HEAD")}}HEAD({{printf "%q" .Path}}){{end}}{{end}}`
	licenseT = `{{if .License}}{{with .License}}License(func() {
{{if .Name}}{{template "name" .}}
{{end}}{{if .URL}}{{template "url" .}}
{{end}}}){{end}}{{end}}`
	optionsT = `{{if and .Verb .Path}}{{if (eq .Verb "OPTIONS")}}OPTIONS({{printf "%q" .Path}}){{end}}{{end}}`
	originT  = `{{if .Origins}}{{$origins := .Origins}}{{$keys := keys .Origins}}{{range $keys}}{{if (not (eq (index $keys 0) .))}}
{{end}}{{with index $origins .}}Origin({{printf "%q" .Origin}}, func() {
{{if .Headers}}Headers(
{{range .Headers}}{{printf "%q" .}},
{{end}})
{{end}}{{if .Methods}}{{template "methods" .}}
{{end}}{{if .Exposed}}{{template "expose" .}}
{{end}}{{if .MaxAge}}{{template "maxAge" .}}
{{end}}{{if .Credentials}}{{template "credentials" .}}
{{end}}}){{end}}{{end}}{{end}}`
	patchT    = `{{if and .Verb .Path}}{{if (eq .Verb "PATCH")}}PATCH({{printf "%q" .Path}}){{end}}{{end}}`
	payloadT  = `{{if .Payload}}{{with .Payload}}Payload({{.TypeName}}){{end}}{{end}}`
	postT     = `{{if and .Verb .Path}}{{if (eq .Verb "POST")}}POST({{printf "%q" .Path}}){{end}}{{end}}`
	putT      = `{{if and .Verb .Path}}{{if (eq .Verb "PUT")}}PUT({{printf "%q" .Path}}){{end}}{{end}}`
	producesT = `{{if .Produces}}{{range $index, $element := .Produces}}{{with $element}}{{if (not (eq $index 0))}}
{{end}}Produces({{if (or (gt (len .MIMETypes) 1) .Function .PackagePath)}}
{{range .MIMETypes}}{{printf "%q" .}},
{{end}}{{if (or .Function .PackagePath)}}func() {
{{if .Function}}{{template "function" .}}
{{end}}{{if .PackagePath}}{{template "package" .}}
{{end}}},
{{end}}{{else}}{{range .MIMETypes}}{{printf "%q" .}}{{end}}{{end}}){{end}}{{end}}{{end}}`
	responseT = `{{if .Responses}}{{$responses := .Responses}}{{$keys := keys .Responses}}{{range $keys}}{{if (not (eq (index $keys 0) .))}}
{{end}}{{with index $responses .}}Response({{.Name}}){{end}}{{end}}{{end}}`
	routingT = `{{if .Routes}}Routing({{if (eq (len .Routes) 1)}}{{range .Routes}}{{template "connect" .}}{{template "delete" .}}{{template "get" .}}{{template "head" .}}{{template "options" .}}{{template "patch" .}}{{template "post" .}}{{template "put" .}}{{template "trace" .}}{{end}}){{else}}
{{range .Routes}}{{if (eq .Verb "CONNECT")}}{{template "connect" .}},
{{end}}{{if (eq .Verb "DELETE")}}{{template "delete" .}},
{{end}}{{if (eq .Verb "GET")}}{{template "get" .}},
{{end}}{{if (eq .Verb "HEAD")}}{{template "head" .}},
{{end}}{{if (eq .Verb "OPTIONS")}}{{template "options" .}},
{{end}}{{if (eq .Verb "PATCH")}}{{template "patch" .}},
{{end}}{{if (eq .Verb "POST")}}{{template "post" .}},
{{end}}{{if (eq .Verb "PUT")}}{{template "put" .}},
{{end}}{{if (eq .Verb "TRACE")}}{{template "trace" .}},
{{end}}{{end}}){{end}}{{end}}`
	traceT = `{{if and .Verb .Path}}{{if (eq .Verb "TRACE")}}TRACE({{printf "%q" .Path}}){{end}}{{end}}`
	typeT  = `{{if .Types}}{{$types := .Types}}{{$keys := keys .Types}}{{range $keys}}{{if (not (eq (index $keys 0) .))}}
{{end}}{{with index $types .}}var {{.TypeName}} = Type({{printf "%q" .TypeName}}, func() {
}){{end}}{{end}}{{end}}`
)

var (
	tmpl *template.Template
)

func init() {
	tmpl = template.New("")

	tmpl = tmpl.Funcs(template.FuncMap{
		"keys": func(x interface{}) interface{} {
			switch t := x.(type) {
			case map[string]*design.ActionDefinition:
				var keys []string
				for k := range t {
					keys = append(keys, k)
				}
				sort.Strings(keys)
				return keys
			case map[string]*design.CORSDefinition:
				var keys []string
				for k := range t {
					keys = append(keys, k)
				}
				sort.Strings(keys)
				return keys
			case map[string]*design.ResponseDefinition:
				var keys []string
				for k := range t {
					keys = append(keys, k)
				}
				sort.Strings(keys)
				return keys
			case map[string]*design.UserTypeDefinition:
				var keys []string
				for k := range t {
					keys = append(keys, k)
				}
				sort.Strings(keys)
				return keys
			default:
				return nil
			}
		},
	})

	tmpl = template.Must(tmpl.New("goHeader").Parse(goHeaderT))
	tmpl = template.Must(tmpl.New("all").Parse(allT))

	// Components that have single value.
	tmpl = template.Must(tmpl.New("basePath").Parse(basePathT))
	tmpl = template.Must(tmpl.New("canonicalActionName").Parse(canonicalActionNameT))
	tmpl = template.Must(tmpl.New("contentType").Parse(contentTypeT))
	tmpl = template.Must(tmpl.New("credentials").Parse(credentialsT))
	tmpl = template.Must(tmpl.New("description").Parse(descriptionT))
	tmpl = template.Must(tmpl.New("email").Parse(emailT))
	tmpl = template.Must(tmpl.New("function").Parse(functionT))
	tmpl = template.Must(tmpl.New("host").Parse(hostT))
	tmpl = template.Must(tmpl.New("maxAge").Parse(maxAgeT))
	tmpl = template.Must(tmpl.New("maxLength").Parse(maxLengthT))
	tmpl = template.Must(tmpl.New("minLength").Parse(minLengthT))
	tmpl = template.Must(tmpl.New("name").Parse(nameT))
	tmpl = template.Must(tmpl.New("package").Parse(packageT))
	tmpl = template.Must(tmpl.New("status").Parse(statusT))
	tmpl = template.Must(tmpl.New("termsOfService").Parse(termsOfServiceT))
	tmpl = template.Must(tmpl.New("title").Parse(titleT))
	tmpl = template.Must(tmpl.New("typeName").Parse(typeNameT))
	tmpl = template.Must(tmpl.New("url").Parse(urlT))
	tmpl = template.Must(tmpl.New("version").Parse(versionT))

	// Components that have multiple values.
	tmpl = template.Must(tmpl.New("expose").Parse(exposeT))
	tmpl = template.Must(tmpl.New("methods").Parse(methodsT))
	tmpl = template.Must(tmpl.New("scheme").Parse(schemeT))

	// Containers.
	tmpl = template.Must(tmpl.New("action").Parse(actionT))
	tmpl = template.Must(tmpl.New("api").Parse(apiT))
	tmpl = template.Must(tmpl.New("connect").Parse(connectT))
	tmpl = template.Must(tmpl.New("consumes").Parse(consumesT))
	tmpl = template.Must(tmpl.New("contact").Parse(contactT))
	tmpl = template.Must(tmpl.New("delete").Parse(deleteT))
	tmpl = template.Must(tmpl.New("docs").Parse(docsT))
	tmpl = template.Must(tmpl.New("get").Parse(getT))
	tmpl = template.Must(tmpl.New("head").Parse(headT))
	tmpl = template.Must(tmpl.New("license").Parse(licenseT))
	tmpl = template.Must(tmpl.New("options").Parse(optionsT))
	tmpl = template.Must(tmpl.New("origin").Parse(originT))
	tmpl = template.Must(tmpl.New("patch").Parse(patchT))
	tmpl = template.Must(tmpl.New("payload").Parse(payloadT))
	tmpl = template.Must(tmpl.New("post").Parse(postT))
	tmpl = template.Must(tmpl.New("put").Parse(putT))
	tmpl = template.Must(tmpl.New("produces").Parse(producesT))
	tmpl = template.Must(tmpl.New("response").Parse(responseT))
	tmpl = template.Must(tmpl.New("routing").Parse(routingT))
	tmpl = template.Must(tmpl.New("trace").Parse(traceT))
	tmpl = template.Must(tmpl.New("type").Parse(typeT))
}
