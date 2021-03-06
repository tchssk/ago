## Swagger

`swaggerToAPI()` in [cmd/swagger.go](cmd/swagger.go) converts `swagger.Swagger` to `design.APIDefinition`.

- [x] Name
- [x] Title
- [x] Description
- [x] Version
- [x] Host
- [x] Schemes
- [x] BasePath
- [ ] Params
- [x] Consumes
- [x] Produces
- [x] Origins
- [x] TermsOfService
- [x] Contact
- [x] License
- [x] Docs
- [ ] Resources
- [ ] Types
- [ ] MediaTypes
- [ ] Traits
- [ ] Responses
- [ ] ResponseTemplates
- [ ] DefaultResponses
- [ ] DefaultResponseTemplates
- [ ] DSLFunc
- [ ] Metadata
- [ ] SecuritySchemes
- [ ] Security
- [ ] NoExamples

## Templates for DSL

All goa DSL functions are generated by templates in [cmd/template.go](cmd/template.go). Templates are grouped into the following three types.

- Components that have single value
- Components that have multiple values
- Containers

### DSL functions

- [ ] func API(name string, dsl func()) *design.APIDefinition
- [ ] func APIKeySecurity(name string, dsl ...func()) *design.SecuritySchemeDefinition
- [ ] func AccessCodeFlow(authorizationURL, tokenURL string)
- [ ] func Action(name string, dsl func())
- [ ] func ApplicationFlow(tokenURL string)
- [ ] func ArrayOf(v interface{}, dsl ...func()) *design.Array
- [ ] func Attribute(name string, args ...interface{})
- [ ] func Attributes(apidsl func())
- [x] func BasePath(val string)
- [ ] func BasicAuthSecurity(name string, dsl ...func()) *design.SecuritySchemeDefinition
- [x] func CONNECT(path string, dsl ...func()) *design.RouteDefinition
- [x] func CanonicalActionName(a string)
- [ ] func CollectionOf(v interface{}, apidsl ...func()) *design.MediaTypeDefinition
- [x] func Consumes(args ...interface{})
- [x] func Contact(dsl func())
- [x] func ContentType(typ string)
- [x] func Credentials()
- [x] func DELETE(path string, dsl ...func()) *design.RouteDefinition
- [ ] func Default(def interface{})
- [ ] func DefaultMedia(val interface{}, viewName ...string)
- [x] func Description(d string)
- [x] func Docs(dsl func())
- [x] func Email(email string)
- [ ] func Enum(val ...interface{})
- [ ] func Example(exp interface{})
- [x] func Expose(vals ...string)
- [ ] func Files(path, filename string, dsls ...func())
- [ ] func Format(f string)
- [x] func Function(fn string)
- [x] func GET(path string, dsl ...func()) *design.RouteDefinition
- [x] func HEAD(path string, dsl ...func()) *design.RouteDefinition
- [ ] func HashOf(k, v design.DataType) *design.Hash
- [ ] func Header(name string, args ...interface{})
- [ ] func Headers(params ...interface{})
- [x] func Host(host string)
- [ ] func ImplicitFlow(authorizationURL string)
- [ ] func JWTSecurity(name string, dsl ...func()) *design.SecuritySchemeDefinition
- [x] func License(dsl func())
- [ ] func Link(name string, view ...string)
- [ ] func Links(apidsl func())
- [x] func MaxAge(val uint)
- [x] func MaxLength(val int)
- [ ] func Maximum(val interface{})
- [ ] func Media(val interface{}, viewName ...string)
- [ ] func MediaType(identifier string, apidsl func()) *design.MediaTypeDefinition
- [ ] func Member(name string, args ...interface{})
- [ ] func Metadata(name string, value ...string)
- [x] func Methods(vals ...string)
- [x] func MinLength(val int)
- [ ] func Minimum(val interface{})
- [x] func Name(name string)
- [ ] func NoExample()
- [ ] func NoSecurity()
- [ ] func OAuth2Security(name string, dsl ...func()) *design.SecuritySchemeDefinition
- [x] func OPTIONS(path string, dsl ...func()) *design.RouteDefinition
- [ ] func OptionalPayload(p interface{}, dsls ...func())
- [x] func Origin(origin string, dsl func())
- [x] func PATCH(path string, dsl ...func()) *design.RouteDefinition
- [x] func POST(path string, dsl ...func()) *design.RouteDefinition
- [x] func PUT(path string, dsl ...func()) *design.RouteDefinition
- [x] func Package(path string)
- [ ] func Param(name string, args ...interface{})
- [ ] func Params(dsl func())
- [ ] func Parent(p string)
- [ ] func PasswordFlow(tokenURL string)
- [ ] func Pattern(p string)
- [x] func Payload(p interface{}, dsls ...func())
- [x] func Produces(args ...interface{})
- [ ] func Query(parameterName string)
- [ ] func Reference(t design.DataType)
- [ ] func Required(names ...string)
- [ ] func Resource(name string, dsl func()) *design.ResourceDefinition
- [ ] func Response(name string, paramsAndDSL ...interface{})
- [ ] func ResponseTemplate(name string, p interface{})
- [x] func Routing(routes ...*design.RouteDefinition)
- [x] func Scheme(vals ...string)
- [ ] func Scope(name string, desc ...string)
- [ ] func Security(scheme interface{}, dsl ...func())
- [x] func Status(status int)
- [x] func TRACE(path string, dsl ...func()) *design.RouteDefinition
- [x] func TermsOfService(terms string)
- [x] func Title(val string)
- [ ] func TokenURL(tokenURL string)
- [ ] func Trait(name string, val ...func())
- [ ] func Type(name string, dsl func()) *design.UserTypeDefinition
- [x] func TypeName(name string)
- [x] func URL(url string)
- [ ] func UseTrait(names ...string)
- [x] func Version(ver string)
- [ ] func View(name string, apidsl ...func())
