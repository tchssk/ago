package cmd

import (
	"bytes"
	"testing"

	"github.com/goadesign/goa/design"
	"github.com/goadesign/goa/dslengine"
)

func TestGoHeaderTmpl(t *testing.T) {
	cases := map[string]struct {
		expected string
	}{
		"pattern 1": {
			expected: `
package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)
`,
		},
	}
	for k, tc := range cases {
		buf := new(bytes.Buffer)
		if err := tmpl.ExecuteTemplate(buf, "goHeader", nil); err != nil {
			t.Fatalf("Execute returned %s", err)
		}
		actual := buf.String()
		if actual != tc.expected {
			t.Errorf("%s: got %v, expected %v", k, actual, tc.expected)
		}
	}
}

func TestAllTmpl(t *testing.T) {
	// TODO Write.
}

// Components have single values.
func TestBasePathTmpl(t *testing.T) {
	cases := map[string]struct {
		definition interface{}
		expected   string
	}{
		"with api definition": {
			definition: design.APIDefinition{
				BasePath: "/v1",
			},
			expected: `BasePath("/v1")`,
		},
		"with resource definition": {
			definition: design.ResourceDefinition{
				BasePath: "/resource",
			},
			expected: `BasePath("/resource")`,
		},
		"without definition": {
			definition: design.APIDefinition{},
			expected:   ``,
		},
	}
	for k, tc := range cases {
		buf := new(bytes.Buffer)
		if err := tmpl.ExecuteTemplate(buf, "basePath", tc.definition); err != nil {
			t.Fatalf("Execute returned %s", err)
		}
		actual := buf.String()
		if actual != tc.expected {
			t.Errorf("%s: got %v, expected %v", k, actual, tc.expected)
		}
	}
}

func TestCanonicalActionNameTmpl(t *testing.T) {
	cases := map[string]struct {
		definition interface{}
		expected   string
	}{
		"with definition": {
			definition: design.ResourceDefinition{
				CanonicalActionName: "get",
			},
			expected: `CanonicalActionName("get")`,
		},
		"without definition": {
			definition: design.ResourceDefinition{},
			expected:   ``,
		},
	}
	for k, tc := range cases {
		buf := new(bytes.Buffer)
		if err := tmpl.ExecuteTemplate(buf, "canonicalActionName", tc.definition); err != nil {
			t.Fatalf("Execute returned %s", err)
		}
		actual := buf.String()
		if actual != tc.expected {
			t.Errorf("%s: got %v, expected %v", k, actual, tc.expected)
		}
	}
}

func TestContentTypeTmpl(t *testing.T) {
	cases := map[string]struct {
		definition interface{}
		expected   string
	}{
		"with definition": {
			definition: design.MediaTypeDefinition{
				ContentType: "application/json",
			},
			expected: `ContentType("application/json")`,
		},
		"without definition": {
			definition: design.MediaTypeDefinition{},
			expected:   ``,
		},
	}
	for k, tc := range cases {
		buf := new(bytes.Buffer)
		if err := tmpl.ExecuteTemplate(buf, "contentType", tc.definition); err != nil {
			t.Fatalf("Execute returned %s", err)
		}
		actual := buf.String()
		if actual != tc.expected {
			t.Errorf("%s: got %v, expected %v", k, actual, tc.expected)
		}
	}
}

func TestCredentialsTmpl(t *testing.T) {
	cases := map[string]struct {
		definition interface{}
		expected   string
	}{
		"with definition": {
			definition: design.CORSDefinition{
				Credentials: true,
			},
			expected: `Credentials()`,
		},
		"without definition": {
			definition: design.CORSDefinition{},
			expected:   ``,
		},
	}
	for k, tc := range cases {
		buf := new(bytes.Buffer)
		if err := tmpl.ExecuteTemplate(buf, "credentials", tc.definition); err != nil {
			t.Fatalf("Execute returned %s", err)
		}
		actual := buf.String()
		if actual != tc.expected {
			t.Errorf("%s: got %v, expected %v", k, actual, tc.expected)
		}
	}
}

func TestDescriptionTmpl(t *testing.T) {
	cases := map[string]struct {
		definition interface{}
		expected   string
	}{
		"with api definition": {
			definition: design.APIDefinition{
				Description: "Description of API",
			},
			expected: `Description("Description of API")`,
		},
		"with docs definition": {
			definition: design.DocsDefinition{
				Description: "Description of docs",
			},
			expected: `Description("Description of docs")`,
		},
		"with resource definition": {
			definition: design.ResourceDefinition{
				Description: "Description of resource",
			},
			expected: `Description("Description of resource")`,
		},
		"with response definition": {
			definition: design.ResponseDefinition{
				Description: "Description of response",
			},
			expected: `Description("Description of response")`,
		},
		"with multi lines in definition": {
			definition: design.APIDefinition{
				Description: `Description
of
API`,
			},
			expected: `Description("Description\nof\nAPI")`,
		},
		"without definition": {
			definition: design.APIDefinition{},
			expected:   ``,
		},
	}
	for k, tc := range cases {
		buf := new(bytes.Buffer)
		if err := tmpl.ExecuteTemplate(buf, "description", tc.definition); err != nil {
			t.Fatalf("Execute returned %s", err)
		}
		actual := buf.String()
		if actual != tc.expected {
			t.Errorf("%s: got %v, expected %v", k, actual, tc.expected)
		}
	}
}

func TestEmailTmpl(t *testing.T) {
	cases := map[string]struct {
		definition interface{}
		expected   string
	}{
		"with definition": {
			definition: design.ContactDefinition{
				Email: "foo@bar.com",
			},
			expected: `Email("foo@bar.com")`,
		},
		"without definition": {
			definition: design.ContactDefinition{},
			expected:   ``,
		},
	}
	for k, tc := range cases {
		buf := new(bytes.Buffer)
		if err := tmpl.ExecuteTemplate(buf, "email", tc.definition); err != nil {
			t.Fatalf("Execute returned %s", err)
		}
		actual := buf.String()
		if actual != tc.expected {
			t.Errorf("%s: got %v, expected %v", k, actual, tc.expected)
		}
	}
}

func TestFunctionTmpl(t *testing.T) {
	cases := map[string]struct {
		definition interface{}
		expected   string
	}{
		"with definition": {
			definition: design.EncodingDefinition{
				Function: "newFunc",
			},
			expected: `Function("newFunc")`,
		},
		"without definition": {
			definition: design.EncodingDefinition{},
			expected:   ``,
		},
	}
	for k, tc := range cases {
		buf := new(bytes.Buffer)
		if err := tmpl.ExecuteTemplate(buf, "function", tc.definition); err != nil {
			t.Fatalf("Execute returned %s", err)
		}
		actual := buf.String()
		if actual != tc.expected {
			t.Errorf("%s: got %v, expected %v", k, actual, tc.expected)
		}
	}
}

func TestHostTmpl(t *testing.T) {
	cases := map[string]struct {
		definition interface{}
		expected   string
	}{
		"with definition": {
			definition: design.APIDefinition{
				Host: "http://localhost/",
			},
			expected: `Host("http://localhost/")`,
		},
		"without definition": {
			definition: design.APIDefinition{},
			expected:   ``,
		},
	}
	for k, tc := range cases {
		buf := new(bytes.Buffer)
		if err := tmpl.ExecuteTemplate(buf, "host", tc.definition); err != nil {
			t.Fatalf("Execute returned %s", err)
		}
		actual := buf.String()
		if actual != tc.expected {
			t.Errorf("%s: got %v, expected %v", k, actual, tc.expected)
		}
	}
}

func TestMaxAgeTmpl(t *testing.T) {
	cases := map[string]struct {
		definition interface{}
		expected   string
	}{
		"with definition": {
			definition: design.CORSDefinition{
				MaxAge: 600,
			},
			expected: `MaxAge(600)`,
		},
		"without definition": {
			definition: design.CORSDefinition{},
			expected:   ``,
		},
	}
	for k, tc := range cases {
		buf := new(bytes.Buffer)
		if err := tmpl.ExecuteTemplate(buf, "maxAge", tc.definition); err != nil {
			t.Fatalf("Execute returned %s", err)
		}
		actual := buf.String()
		if actual != tc.expected {
			t.Errorf("%s: got %v, expected %v", k, actual, tc.expected)
		}
	}
}

func TestMaxLengthTmpl(t *testing.T) {
	cases := map[string]struct {
		definition interface{}
		expected   string
	}{
		"with definition": {
			definition: dslengine.ValidationDefinition{
				MaxLength: &[]int{5}[0],
			},
			expected: `MaxLength(5)`,
		},
		"without definition": {
			definition: dslengine.ValidationDefinition{},
			expected:   ``,
		},
	}
	for k, tc := range cases {
		buf := new(bytes.Buffer)
		if err := tmpl.ExecuteTemplate(buf, "maxLength", tc.definition); err != nil {
			t.Fatalf("Execute returned %s", err)
		}
		actual := buf.String()
		if actual != tc.expected {
			t.Errorf("%s: got %v, expected %v", k, actual, tc.expected)
		}
	}
}

func TestMinLengthTmpl(t *testing.T) {
	cases := map[string]struct {
		definition interface{}
		expected   string
	}{
		"with definition": {
			definition: dslengine.ValidationDefinition{
				MinLength: &[]int{3}[0],
			},
			expected: `MinLength(3)`,
		},
		"without definition": {
			definition: dslengine.ValidationDefinition{},
			expected:   ``,
		},
	}
	for k, tc := range cases {
		buf := new(bytes.Buffer)
		if err := tmpl.ExecuteTemplate(buf, "minLength", tc.definition); err != nil {
			t.Fatalf("Execute returned %s", err)
		}
		actual := buf.String()
		if actual != tc.expected {
			t.Errorf("%s: got %v, expected %v", k, actual, tc.expected)
		}
	}
}

func TestNameTmpl(t *testing.T) {
	cases := map[string]struct {
		definition interface{}
		expected   string
	}{
		"with api definition": {
			definition: design.APIDefinition{
				Name: "Name of API",
			},
			expected: `Name("Name of API")`,
		},
		"with contact definition": {
			definition: design.ContactDefinition{
				Name: "Name of contact",
			},
			expected: `Name("Name of contact")`,
		},
		"with license definition": {
			definition: design.LicenseDefinition{
				Name: "Name of license",
			},
			expected: `Name("Name of license")`,
		},
		"without definition": {
			definition: design.ContactDefinition{},
			expected:   ``,
		},
	}
	for k, tc := range cases {
		buf := new(bytes.Buffer)
		if err := tmpl.ExecuteTemplate(buf, "name", tc.definition); err != nil {
			t.Fatalf("Execute returned %s", err)
		}
		actual := buf.String()
		if actual != tc.expected {
			t.Errorf("%s: got %v, expected %v", k, actual, tc.expected)
		}
	}
}

func TestPackageTmpl(t *testing.T) {
	cases := map[string]struct {
		definition interface{}
		expected   string
	}{
		"with definition": {
			definition: design.EncodingDefinition{
				PackagePath: "encoding/json",
			},
			expected: `Package("encoding/json")`,
		},
		"without definition": {
			definition: design.EncodingDefinition{},
			expected:   ``,
		},
	}
	for k, tc := range cases {
		buf := new(bytes.Buffer)
		if err := tmpl.ExecuteTemplate(buf, "package", tc.definition); err != nil {
			t.Fatalf("Execute returned %s", err)
		}
		actual := buf.String()
		if actual != tc.expected {
			t.Errorf("%s: got %v, expected %v", k, actual, tc.expected)
		}
	}
}

func TestStatusTmpl(t *testing.T) {
	cases := map[string]struct {
		definition interface{}
		expected   string
	}{
		"with definition": {
			definition: design.ResponseDefinition{
				Status: 200,
			},
			expected: `Status(200)`,
		},
		"without definition": {
			definition: design.ResponseDefinition{},
			expected:   ``,
		},
	}
	for k, tc := range cases {
		buf := new(bytes.Buffer)
		if err := tmpl.ExecuteTemplate(buf, "status", tc.definition); err != nil {
			t.Fatalf("Execute returned %s", err)
		}
		actual := buf.String()
		if actual != tc.expected {
			t.Errorf("%s: got %v, expected %v", k, actual, tc.expected)
		}
	}
}

func TestTermsOfServiceTmpl(t *testing.T) {
	cases := map[string]struct {
		definition interface{}
		expected   string
	}{
		"with definition": {
			definition: design.APIDefinition{
				TermsOfService: "terms of service",
			},
			expected: `TermsOfService("terms of service")`,
		},
		"without definition": {
			definition: design.APIDefinition{},
			expected:   ``,
		},
	}
	for k, tc := range cases {
		buf := new(bytes.Buffer)
		if err := tmpl.ExecuteTemplate(buf, "termsOfService", tc.definition); err != nil {
			t.Fatalf("Execute returned %s", err)
		}
		actual := buf.String()
		if actual != tc.expected {
			t.Errorf("%s: got %v, expected %v", k, actual, tc.expected)
		}
	}
}

func TestTitleTmpl(t *testing.T) {
	cases := map[string]struct {
		definition interface{}
		expected   string
	}{
		"with definition": {
			definition: design.APIDefinition{
				Title: "Title of API",
			},
			expected: `Title("Title of API")`,
		},
		"without definition": {
			definition: design.APIDefinition{},
			expected:   ``,
		},
	}
	for k, tc := range cases {
		buf := new(bytes.Buffer)
		if err := tmpl.ExecuteTemplate(buf, "title", tc.definition); err != nil {
			t.Fatalf("Execute returned %s", err)
		}
		actual := buf.String()
		if actual != tc.expected {
			t.Errorf("%s: got %v, expected %v", k, actual, tc.expected)
		}
	}
}

func TestTypeNameTmpl(t *testing.T) {
	cases := map[string]struct {
		definition interface{}
		expected   string
	}{
		"with media type definition": {
			definition: design.MediaTypeDefinition{
				UserTypeDefinition: &design.UserTypeDefinition{
					TypeName: "TypeName of media type",
				},
			},
			expected: `TypeName("TypeName of media type")`,
		},
		"with user type definition": {
			definition: design.UserTypeDefinition{
				TypeName: "TypeName of user type",
			},
			expected: `TypeName("TypeName of user type")`,
		},
		"without definition": {
			definition: design.UserTypeDefinition{},
			expected:   ``,
		},
	}
	for k, tc := range cases {
		buf := new(bytes.Buffer)
		if err := tmpl.ExecuteTemplate(buf, "typeName", tc.definition); err != nil {
			t.Fatalf("Execute returned %s", err)
		}
		actual := buf.String()
		if actual != tc.expected {
			t.Errorf("%s: got %v, expected %v", k, actual, tc.expected)
		}
	}
}

func TestURLTmpl(t *testing.T) {
	cases := map[string]struct {
		definition interface{}
		expected   string
	}{
		"with contact definition": {
			definition: design.ContactDefinition{
				URL: "http://localhost/contact",
			},
			expected: `URL("http://localhost/contact")`,
		},
		"with license definition": {
			definition: design.LicenseDefinition{
				URL: "http://localhost/license",
			},
			expected: `URL("http://localhost/license")`,
		},
		"with docs definition": {
			definition: design.DocsDefinition{
				URL: "http://localhost/docs",
			},
			expected: `URL("http://localhost/docs")`,
		},
		"without definition": {
			definition: design.ContactDefinition{},
			expected:   ``,
		},
	}
	for k, tc := range cases {
		buf := new(bytes.Buffer)
		if err := tmpl.ExecuteTemplate(buf, "url", tc.definition); err != nil {
			t.Fatalf("Execute returned %s", err)
		}
		actual := buf.String()
		if actual != tc.expected {
			t.Errorf("%s: got %v, expected %v", k, actual, tc.expected)
		}
	}
}

func TestVersionTmpl(t *testing.T) {
	cases := map[string]struct {
		definition interface{}
		expected   string
	}{
		"with definition": {
			definition: design.APIDefinition{
				Version: "1.0.0",
			},
			expected: `Version("1.0.0")`,
		},
		"without definition": {
			definition: design.APIDefinition{},
			expected:   ``,
		},
	}
	for k, tc := range cases {
		buf := new(bytes.Buffer)
		if err := tmpl.ExecuteTemplate(buf, "version", tc.definition); err != nil {
			t.Fatalf("Execute returned %s", err)
		}
		actual := buf.String()
		if actual != tc.expected {
			t.Errorf("%s: got %v, expected %v", k, actual, tc.expected)
		}
	}
}

// Components have multiple values.
func TestExposeTmpl(t *testing.T) {
	cases := map[string]struct {
		definition interface{}
		expected   string
	}{
		"with single definition": {
			definition: design.CORSDefinition{
				Exposed: []string{
					"Authorization",
				},
			},
			expected: `Expose("Authorization")`,
		},
		"with multi definition": {
			definition: design.CORSDefinition{
				Exposed: []string{
					"Authorization",
					"Link",
				},
			},
			expected: `Expose(
"Authorization",
"Link",
)`,
		},
		"without definition": {
			definition: design.CORSDefinition{},
			expected:   ``,
		},
	}
	for k, tc := range cases {
		buf := new(bytes.Buffer)
		if err := tmpl.ExecuteTemplate(buf, "expose", tc.definition); err != nil {
			t.Fatalf("Execute returned %s", err)
		}
		actual := buf.String()
		if actual != tc.expected {
			t.Errorf("%s: got %v, expected %v", k, actual, tc.expected)
		}
	}
}

func TestMethodsTmpl(t *testing.T) {
	cases := map[string]struct {
		definition interface{}
		expected   string
	}{
		"with single definition": {
			definition: design.CORSDefinition{
				Methods: []string{
					"GET",
				},
			},
			expected: `Methods("GET")`,
		},
		"with multi definition": {
			definition: design.CORSDefinition{
				Methods: []string{
					"GET",
					"POST",
				},
			},
			expected: `Methods(
"GET",
"POST",
)`,
		},
		"without definition": {
			definition: design.CORSDefinition{},
			expected:   ``,
		},
	}
	for k, tc := range cases {
		buf := new(bytes.Buffer)
		if err := tmpl.ExecuteTemplate(buf, "methods", tc.definition); err != nil {
			t.Fatalf("Execute returned %s", err)
		}
		actual := buf.String()
		if actual != tc.expected {
			t.Errorf("%s: got %v, expected %v", k, actual, tc.expected)
		}
	}
}

func TestSchemeTmpl(t *testing.T) {
	cases := map[string]struct {
		definition interface{}
		expected   string
	}{
		"with single definition": {
			definition: design.APIDefinition{
				Schemes: []string{
					"http",
				},
			},
			expected: `Scheme("http")`,
		},
		"with multi definition": {
			definition: design.APIDefinition{
				Schemes: []string{
					"http",
					"https",
				},
			},
			expected: `Scheme(
"http",
"https",
)`,
		},
		"without definition": {
			definition: design.APIDefinition{},
			expected:   ``,
		},
	}
	for k, tc := range cases {
		buf := new(bytes.Buffer)
		if err := tmpl.ExecuteTemplate(buf, "scheme", tc.definition); err != nil {
			t.Fatalf("Execute returned %s", err)
		}
		actual := buf.String()
		if actual != tc.expected {
			t.Errorf("%s: got %v, expected %v", k, actual, tc.expected)
		}
	}
}

// Containers.
func TestActionTmpl(t *testing.T) {
	cases := map[string]struct {
		definition interface{}
		expected   string
	}{
		"with single definition": {
			definition: design.ResourceDefinition{
				Actions: map[string]*design.ActionDefinition{
					"": &design.ActionDefinition{
						Description: "Description of action",
						Docs: &design.DocsDefinition{
							Description: "Description of docs",
							URL:         "http://localhost/docs",
						},
						Schemes: []string{
							"http",
						},
						Routes: []*design.RouteDefinition{
							&design.RouteDefinition{
								Verb: "GET",
								Path: "/",
							},
						},
						Payload: &design.UserTypeDefinition{
							TypeName: "FooPayload",
						},
					},
				},
			},
			expected: `Action("", func() {
Description("Description of action")
Docs(func() {
Description("Description of docs")
URL("http://localhost/docs")
})
Scheme("http")
Routing(GET("/"))
Payload(FooPayload)
})`,
		},
		"without definition": {
			definition: design.ResourceDefinition{},
			expected:   ``,
		},
	}
	for k, tc := range cases {
		buf := new(bytes.Buffer)
		if err := tmpl.ExecuteTemplate(buf, "action", tc.definition); err != nil {
			t.Fatalf("Execute returned %s", err)
		}
		actual := buf.String()
		if actual != tc.expected {
			t.Errorf("%s: \ngot:\n%v\nexpected:\n%v", k, actual, tc.expected)
		}
	}
}

func TestAPITmpl(t *testing.T) {
	cases := map[string]struct {
		definition interface{}
		expected   string
	}{
		"with definition": {
			definition: design.APIDefinition{
				Name:           "Name of API",
				Title:          "Title of API",
				Description:    "Description of API",
				Version:        "1.0.0",
				TermsOfService: "terms of service",
				Contact: &design.ContactDefinition{
					Name:  "Name of contact",
					Email: "foo@bar.com",
					URL:   "http://localhost/contact",
				},
				License: &design.LicenseDefinition{
					Name: "Name of license",
					URL:  "http://localhost/license",
				},
				Docs: &design.DocsDefinition{
					Description: "Description of docs",
					URL:         "http://localhost/docs",
				},
				Host: "http://localhost/",
				Schemes: []string{
					"http",
					"https",
				},
				BasePath: "/v1",
				Origins: map[string]*design.CORSDefinition{
					"*": &design.CORSDefinition{
						Origin: "*",
						Headers: []string{
							"Authorization",
							"Link",
						},
						Methods: []string{
							"GET",
						},
						Exposed: []string{
							"Authorization",
						},
						MaxAge:      600,
						Credentials: true,
						Regexp:      false,
					},
				},
				Consumes: []*design.EncodingDefinition{
					&design.EncodingDefinition{
						MIMETypes: []string{"application/json"},
						Encoder:   false,
					},
				},
				Produces: []*design.EncodingDefinition{
					&design.EncodingDefinition{
						MIMETypes: []string{"application/json"},
						Encoder:   true,
					},
				},
			},
			expected: `var _ = API("Name of API", func() {
Title("Title of API")
Description("Description of API")
Version("1.0.0")
TermsOfService("terms of service")
Contact(func() {
Name("Name of contact")
Email("foo@bar.com")
URL("http://localhost/contact")
})
License(func() {
Name("Name of license")
URL("http://localhost/license")
})
Docs(func() {
Description("Description of docs")
URL("http://localhost/docs")
})
Host("http://localhost/")
Scheme(
"http",
"https",
)
BasePath("/v1")
Origin("*", func() {
Headers(
"Authorization",
"Link",
)
Methods("GET")
Expose("Authorization")
MaxAge(600)
Credentials()
})
Consumes("application/json")
Produces("application/json")
})`,
		},
		"without definition": {
			definition: design.APIDefinition{},
			expected: `var _ = API("", func() {
})`,
		},
	}
	for k, tc := range cases {
		buf := new(bytes.Buffer)
		if err := tmpl.ExecuteTemplate(buf, "api", tc.definition); err != nil {
			t.Fatalf("Execute returned %s", err)
		}
		actual := buf.String()
		if actual != tc.expected {
			t.Errorf("%s: \ngot:\n%v\nexpected:\n%v", k, actual, tc.expected)
		}
	}
}

func TestCONNECTTmpl(t *testing.T) {
	cases := map[string]struct {
		definition interface{}
		expected   string
	}{
		"with connect definition": {
			definition: design.RouteDefinition{
				Verb: "CONNECT",
				Path: "/",
			},
			expected: `CONNECT("/")`,
		},
		"with get definition": {
			definition: design.RouteDefinition{
				Verb: "GET",
				Path: "/",
			},
			expected: ``,
		},
		"without definition": {
			definition: design.RouteDefinition{},
			expected:   ``,
		},
	}
	for k, tc := range cases {
		buf := new(bytes.Buffer)
		if err := tmpl.ExecuteTemplate(buf, "connect", tc.definition); err != nil {
			t.Fatalf("Execute returned %s", err)
		}
		actual := buf.String()
		if actual != tc.expected {
			t.Errorf("%s: \ngot:\n%v\nexpected:\n%v", k, actual, tc.expected)
		}
	}
}

func TestConsumesTmpl(t *testing.T) {
	cases := map[string]struct {
		definition interface{}
		expected   string
	}{
		"with complex definition": {
			definition: design.APIDefinition{
				Consumes: []*design.EncodingDefinition{
					&design.EncodingDefinition{
						MIMETypes: []string{"application/json"},
						Encoder:   false,
					},
					&design.EncodingDefinition{
						MIMETypes: []string{
							"application/json",
							"application/xml",
						},
						Encoder: false,
					},
					&design.EncodingDefinition{
						MIMETypes: []string{"application/json"},
						Function:  "newFunc",
						Encoder:   false,
					},
					&design.EncodingDefinition{
						MIMETypes:   []string{"application/json"},
						PackagePath: "encoding/json",
						Encoder:     false,
					},
				},
			},
			expected: `Consumes("application/json")
Consumes(
"application/json",
"application/xml",
)
Consumes(
"application/json",
func() {
Function("newFunc")
},
)
Consumes(
"application/json",
func() {
Package("encoding/json")
},
)`,
		},
		"with multi consumes definition": {
			definition: design.APIDefinition{
				Consumes: []*design.EncodingDefinition{
					&design.EncodingDefinition{
						MIMETypes: []string{"application/json"},
						Encoder:   false,
					},
					&design.EncodingDefinition{
						MIMETypes: []string{
							"application/json",
							"application/xml",
						},
						Encoder: false,
					},
				},
			},
			expected: `Consumes("application/json")
Consumes(
"application/json",
"application/xml",
)`,
		},
		"with single consumes definition": {
			definition: design.APIDefinition{
				Consumes: []*design.EncodingDefinition{
					&design.EncodingDefinition{
						MIMETypes: []string{"application/json"},
						Encoder:   false,
					},
				},
			},
			expected: `Consumes("application/json")`,
		},
		"with function definition": {
			definition: design.APIDefinition{
				Consumes: []*design.EncodingDefinition{
					&design.EncodingDefinition{
						MIMETypes: []string{"application/json"},
						Function:  "newFunc",
						Encoder:   false,
					},
				},
			},
			expected: `Consumes(
"application/json",
func() {
Function("newFunc")
},
)`,
		},
		"with package path definition": {
			definition: design.APIDefinition{
				Consumes: []*design.EncodingDefinition{
					&design.EncodingDefinition{
						MIMETypes:   []string{"application/json"},
						PackagePath: "encoding/json",
						Encoder:     false,
					},
				},
			},
			expected: `Consumes(
"application/json",
func() {
Package("encoding/json")
},
)`,
		},
		"without definition": {
			definition: design.APIDefinition{},
			expected:   ``,
		},
	}
	for k, tc := range cases {
		buf := new(bytes.Buffer)
		if err := tmpl.ExecuteTemplate(buf, "consumes", tc.definition); err != nil {
			t.Fatalf("Execute returned %s", err)
		}
		actual := buf.String()
		if actual != tc.expected {
			t.Errorf("%s: \ngot:\n%v\nexpected:\n%v", k, actual, tc.expected)
		}
	}
}

func TestContactTmpl(t *testing.T) {
	cases := map[string]struct {
		definition interface{}
		expected   string
	}{
		"with name/email/url definition": {
			definition: design.APIDefinition{
				Contact: &design.ContactDefinition{
					Name:  "Name of contact",
					Email: "foo@bar.com",
					URL:   "http://localhost/contact",
				},
			},
			expected: `Contact(func() {
Name("Name of contact")
Email("foo@bar.com")
URL("http://localhost/contact")
})`,
		},
		"with name definition": {
			definition: design.APIDefinition{
				Contact: &design.ContactDefinition{
					Name: "Name of contact",
				},
			},
			expected: `Contact(func() {
Name("Name of contact")
})`,
		},
		"with email definition": {
			definition: design.APIDefinition{
				Contact: &design.ContactDefinition{
					Email: "foo@bar.com",
				},
			},
			expected: `Contact(func() {
Email("foo@bar.com")
})`,
		},
		"with url definition": {
			definition: design.APIDefinition{
				Contact: &design.ContactDefinition{
					URL: "http://localhost/contact",
				},
			},
			expected: `Contact(func() {
URL("http://localhost/contact")
})`,
		},
		"without empty contact definition": {
			definition: design.APIDefinition{
				Contact: &design.ContactDefinition{},
			},
			expected: `Contact(func() {
})`,
		},
		"without definition": {
			definition: design.APIDefinition{},
			expected:   ``,
		},
	}
	for k, tc := range cases {
		buf := new(bytes.Buffer)
		if err := tmpl.ExecuteTemplate(buf, "contact", tc.definition); err != nil {
			t.Fatalf("Execute returned %s", err)
		}
		actual := buf.String()
		if actual != tc.expected {
			t.Errorf("%s: \ngot:\n%v\nexpected:\n%v", k, actual, tc.expected)
		}
	}
}

func TestDELETETmpl(t *testing.T) {
	cases := map[string]struct {
		definition interface{}
		expected   string
	}{
		"with delete definition": {
			definition: design.RouteDefinition{
				Verb: "DELETE",
				Path: "/",
			},
			expected: `DELETE("/")`,
		},
		"with get definition": {
			definition: design.RouteDefinition{
				Verb: "GET",
				Path: "/",
			},
			expected: ``,
		},
		"without definition": {
			definition: design.RouteDefinition{},
			expected:   ``,
		},
	}
	for k, tc := range cases {
		buf := new(bytes.Buffer)
		if err := tmpl.ExecuteTemplate(buf, "delete", tc.definition); err != nil {
			t.Fatalf("Execute returned %s", err)
		}
		actual := buf.String()
		if actual != tc.expected {
			t.Errorf("%s: \ngot:\n%v\nexpected:\n%v", k, actual, tc.expected)
		}
	}
}

func TestDocsTmpl(t *testing.T) {
	cases := map[string]struct {
		definition interface{}
		expected   string
	}{
		"with api definition": {
			definition: design.APIDefinition{
				Docs: &design.DocsDefinition{
					Description: "Description of API",
					URL:         "http://localhost/api",
				},
			},
			expected: `Docs(func() {
Description("Description of API")
URL("http://localhost/api")
})`,
		},
		"with action definition": {
			definition: design.ActionDefinition{
				Docs: &design.DocsDefinition{
					Description: "Description of action",
					URL:         "http://localhost/action",
				},
			},
			expected: `Docs(func() {
Description("Description of action")
URL("http://localhost/action")
})`,
		},
		"with name definition": {
			definition: design.APIDefinition{
				Docs: &design.DocsDefinition{
					Description: "Description of action",
				},
			},
			expected: `Docs(func() {
Description("Description of action")
})`,
		},
		"with url definition": {
			definition: design.APIDefinition{
				Docs: &design.DocsDefinition{
					URL: "http://localhost/api",
				},
			},
			expected: `Docs(func() {
URL("http://localhost/api")
})`,
		},
		"without empty docs definition": {
			definition: design.APIDefinition{
				Docs: &design.DocsDefinition{},
			},
			expected: `Docs(func() {
})`,
		},
		"without definition": {
			definition: design.APIDefinition{},
			expected:   ``,
		},
	}
	for k, tc := range cases {
		buf := new(bytes.Buffer)
		if err := tmpl.ExecuteTemplate(buf, "docs", tc.definition); err != nil {
			t.Fatalf("Execute returned %s", err)
		}
		actual := buf.String()
		if actual != tc.expected {
			t.Errorf("%s: \ngot:\n%v\nexpected:\n%v", k, actual, tc.expected)
		}
	}
}

func TestGETTmpl(t *testing.T) {
	cases := map[string]struct {
		definition interface{}
		expected   string
	}{
		"with get definition": {
			definition: design.RouteDefinition{
				Verb: "GET",
				Path: "/",
			},
			expected: `GET("/")`,
		},
		"with post definition": {
			definition: design.RouteDefinition{
				Verb: "POST",
				Path: "/",
			},
			expected: ``,
		},
		"without definition": {
			definition: design.RouteDefinition{},
			expected:   ``,
		},
	}
	for k, tc := range cases {
		buf := new(bytes.Buffer)
		if err := tmpl.ExecuteTemplate(buf, "get", tc.definition); err != nil {
			t.Fatalf("Execute returned %s", err)
		}
		actual := buf.String()
		if actual != tc.expected {
			t.Errorf("%s: \ngot:\n%v\nexpected:\n%v", k, actual, tc.expected)
		}
	}
}

func TestHEADTmpl(t *testing.T) {
	cases := map[string]struct {
		definition interface{}
		expected   string
	}{
		"with head definition": {
			definition: design.RouteDefinition{
				Verb: "HEAD",
				Path: "/",
			},
			expected: `HEAD("/")`,
		},
		"with get definition": {
			definition: design.RouteDefinition{
				Verb: "GET",
				Path: "/",
			},
			expected: ``,
		},
		"without definition": {
			definition: design.RouteDefinition{},
			expected:   ``,
		},
	}
	for k, tc := range cases {
		buf := new(bytes.Buffer)
		if err := tmpl.ExecuteTemplate(buf, "head", tc.definition); err != nil {
			t.Fatalf("Execute returned %s", err)
		}
		actual := buf.String()
		if actual != tc.expected {
			t.Errorf("%s: \ngot:\n%v\nexpected:\n%v", k, actual, tc.expected)
		}
	}
}

func TestLicenseTmpl(t *testing.T) {
	cases := map[string]struct {
		definition interface{}
		expected   string
	}{
		"with name/url definition": {
			definition: design.APIDefinition{
				License: &design.LicenseDefinition{
					Name: "Name of license",
					URL:  "http://localhost/license",
				},
			},
			expected: `License(func() {
Name("Name of license")
URL("http://localhost/license")
})`,
		},
		"with name definition": {
			definition: design.APIDefinition{
				License: &design.LicenseDefinition{
					Name: "Name of license",
				},
			},
			expected: `License(func() {
Name("Name of license")
})`,
		},
		"with url definition": {
			definition: design.APIDefinition{
				License: &design.LicenseDefinition{
					URL: "http://localhost/license",
				},
			},
			expected: `License(func() {
URL("http://localhost/license")
})`,
		},
		"without empty license definition": {
			definition: design.APIDefinition{
				License: &design.LicenseDefinition{},
			},
			expected: `License(func() {
})`,
		},
		"without definition": {
			definition: design.APIDefinition{},
			expected:   ``,
		},
	}
	for k, tc := range cases {
		buf := new(bytes.Buffer)
		if err := tmpl.ExecuteTemplate(buf, "license", tc.definition); err != nil {
			t.Fatalf("Execute returned %s", err)
		}
		actual := buf.String()
		if actual != tc.expected {
			t.Errorf("%s: \ngot:\n%v\nexpected:\n%v", k, actual, tc.expected)
		}
	}
}

func TestOPTIONSTmpl(t *testing.T) {
	cases := map[string]struct {
		definition interface{}
		expected   string
	}{
		"with options definition": {
			definition: design.RouteDefinition{
				Verb: "OPTIONS",
				Path: "/",
			},
			expected: `OPTIONS("/")`,
		},
		"with get definition": {
			definition: design.RouteDefinition{
				Verb: "GET",
				Path: "/",
			},
			expected: ``,
		},
		"without definition": {
			definition: design.RouteDefinition{},
			expected:   ``,
		},
	}
	for k, tc := range cases {
		buf := new(bytes.Buffer)
		if err := tmpl.ExecuteTemplate(buf, "options", tc.definition); err != nil {
			t.Fatalf("Execute returned %s", err)
		}
		actual := buf.String()
		if actual != tc.expected {
			t.Errorf("%s: \ngot:\n%v\nexpected:\n%v", k, actual, tc.expected)
		}
	}
}

func TestOriginTmpl(t *testing.T) {
	cases := map[string]struct {
		definition interface{}
		expected   string
	}{
		"with multi definition": {
			definition: design.APIDefinition{
				Origins: map[string]*design.CORSDefinition{
					"1": &design.CORSDefinition{},
					"2": &design.CORSDefinition{},
				},
			},
			expected: `Origin("", func() {
})
Origin("", func() {
})`,
		},
		"with single definition": {
			definition: design.APIDefinition{
				Origins: map[string]*design.CORSDefinition{
					"*": &design.CORSDefinition{
						Origin: "*",
						Headers: []string{
							"Authorization",
							"Link",
						},
						Methods: []string{
							"GET",
						},
						Exposed: []string{
							"Authorization",
						},
						MaxAge:      600,
						Credentials: true,
						Regexp:      false,
					},
				},
			},
			expected: `Origin("*", func() {
Headers(
"Authorization",
"Link",
)
Methods("GET")
Expose("Authorization")
MaxAge(600)
Credentials()
})`,
		},
		"without definition": {
			definition: design.APIDefinition{},
			expected:   ``,
		},
	}
	for k, tc := range cases {
		buf := new(bytes.Buffer)
		if err := tmpl.ExecuteTemplate(buf, "origin", tc.definition); err != nil {
			t.Fatalf("Execute returned %s", err)
		}
		actual := buf.String()
		if actual != tc.expected {
			t.Errorf("%s: \ngot:\n%v\nexpected:\n%v", k, actual, tc.expected)
		}
	}
}

func TestPATCHTmpl(t *testing.T) {
	cases := map[string]struct {
		definition interface{}
		expected   string
	}{
		"with patch definition": {
			definition: design.RouteDefinition{
				Verb: "PATCH",
				Path: "/",
			},
			expected: `PATCH("/")`,
		},
		"with get definition": {
			definition: design.RouteDefinition{
				Verb: "GET",
				Path: "/",
			},
			expected: ``,
		},
		"without definition": {
			definition: design.RouteDefinition{},
			expected:   ``,
		},
	}
	for k, tc := range cases {
		buf := new(bytes.Buffer)
		if err := tmpl.ExecuteTemplate(buf, "patch", tc.definition); err != nil {
			t.Fatalf("Execute returned %s", err)
		}
		actual := buf.String()
		if actual != tc.expected {
			t.Errorf("%s: \ngot:\n%v\nexpected:\n%v", k, actual, tc.expected)
		}
	}
}

func TestPayloadTmpl(t *testing.T) {
	cases := map[string]struct {
		definition interface{}
		expected   string
	}{
		"with definition": {
			definition: design.ActionDefinition{
				Payload: &design.UserTypeDefinition{
					TypeName: "FooPayload",
				},
			},
			expected: `Payload(FooPayload)`,
		},
		"without definition": {
			definition: design.ActionDefinition{},
			expected:   ``,
		},
	}
	for k, tc := range cases {
		buf := new(bytes.Buffer)
		if err := tmpl.ExecuteTemplate(buf, "payload", tc.definition); err != nil {
			t.Fatalf("Execute returned %s", err)
		}
		actual := buf.String()
		if actual != tc.expected {
			t.Errorf("%s: \ngot:\n%v\nexpected:\n%v", k, actual, tc.expected)
		}
	}
}

func TestPOSTTmpl(t *testing.T) {
	cases := map[string]struct {
		definition interface{}
		expected   string
	}{
		"with post definition": {
			definition: design.RouteDefinition{
				Verb: "POST",
				Path: "/",
			},
			expected: `POST("/")`,
		},
		"with get definition": {
			definition: design.RouteDefinition{
				Verb: "GET",
				Path: "/",
			},
			expected: ``,
		},
		"without definition": {
			definition: design.RouteDefinition{},
			expected:   ``,
		},
	}
	for k, tc := range cases {
		buf := new(bytes.Buffer)
		if err := tmpl.ExecuteTemplate(buf, "post", tc.definition); err != nil {
			t.Fatalf("Execute returned %s", err)
		}
		actual := buf.String()
		if actual != tc.expected {
			t.Errorf("%s: \ngot:\n%v\nexpected:\n%v", k, actual, tc.expected)
		}
	}
}

func TestPUTTmpl(t *testing.T) {
	cases := map[string]struct {
		definition interface{}
		expected   string
	}{
		"with put definition": {
			definition: design.RouteDefinition{
				Verb: "PUT",
				Path: "/",
			},
			expected: `PUT("/")`,
		},
		"with get definition": {
			definition: design.RouteDefinition{
				Verb: "GET",
				Path: "/",
			},
			expected: ``,
		},
		"without definition": {
			definition: design.RouteDefinition{},
			expected:   ``,
		},
	}
	for k, tc := range cases {
		buf := new(bytes.Buffer)
		if err := tmpl.ExecuteTemplate(buf, "put", tc.definition); err != nil {
			t.Fatalf("Execute returned %s", err)
		}
		actual := buf.String()
		if actual != tc.expected {
			t.Errorf("%s: \ngot:\n%v\nexpected:\n%v", k, actual, tc.expected)
		}
	}
}

func TestProducesTmpl(t *testing.T) {
	cases := map[string]struct {
		definition interface{}
		expected   string
	}{
		"with complex definition": {
			definition: design.APIDefinition{
				Produces: []*design.EncodingDefinition{
					&design.EncodingDefinition{
						MIMETypes: []string{"application/json"},
						Encoder:   true,
					},
					&design.EncodingDefinition{
						MIMETypes: []string{
							"application/json",
							"application/xml",
						},
						Encoder: true,
					},
					&design.EncodingDefinition{
						MIMETypes: []string{"application/json"},
						Function:  "newFunc",
						Encoder:   true,
					},
					&design.EncodingDefinition{
						MIMETypes:   []string{"application/json"},
						PackagePath: "encoding/json",
						Encoder:     true,
					},
				},
			},
			expected: `Produces("application/json")
Produces(
"application/json",
"application/xml",
)
Produces(
"application/json",
func() {
Function("newFunc")
},
)
Produces(
"application/json",
func() {
Package("encoding/json")
},
)`,
		},
		"with multi produces definition": {
			definition: design.APIDefinition{
				Produces: []*design.EncodingDefinition{
					&design.EncodingDefinition{
						MIMETypes: []string{"application/json"},
						Encoder:   true,
					},
					&design.EncodingDefinition{
						MIMETypes: []string{
							"application/json",
							"application/xml",
						},
						Encoder: true,
					},
				},
			},
			expected: `Produces("application/json")
Produces(
"application/json",
"application/xml",
)`,
		},
		"with single produces definition": {
			definition: design.APIDefinition{
				Produces: []*design.EncodingDefinition{
					&design.EncodingDefinition{
						MIMETypes: []string{"application/json"},
						Encoder:   true,
					},
				},
			},
			expected: `Produces("application/json")`,
		},
		"with function definition": {
			definition: design.APIDefinition{
				Produces: []*design.EncodingDefinition{
					&design.EncodingDefinition{
						MIMETypes: []string{"application/json"},
						Function:  "newFunc",
						Encoder:   true,
					},
				},
			},
			expected: `Produces(
"application/json",
func() {
Function("newFunc")
},
)`,
		},
		"with package path definition": {
			definition: design.APIDefinition{
				Produces: []*design.EncodingDefinition{
					&design.EncodingDefinition{
						MIMETypes:   []string{"application/json"},
						PackagePath: "encoding/json",
						Encoder:     true,
					},
				},
			},
			expected: `Produces(
"application/json",
func() {
Package("encoding/json")
},
)`,
		},
		"without definition": {
			definition: design.APIDefinition{},
			expected:   ``,
		},
	}
	for k, tc := range cases {
		buf := new(bytes.Buffer)
		if err := tmpl.ExecuteTemplate(buf, "produces", tc.definition); err != nil {
			t.Fatalf("Execute returned %s", err)
		}
		actual := buf.String()
		if actual != tc.expected {
			t.Errorf("%s: \ngot:\n%v\nexpected:\n%v", k, actual, tc.expected)
		}
	}
}

func TestResponseTmpl(t *testing.T) {
	cases := map[string]struct {
		definition interface{}
		expected   string
	}{
		"with multi definition": {
			definition: design.ActionDefinition{
				Responses: map[string]*design.ResponseDefinition{
					"FooMedia": &design.ResponseDefinition{
						Name: "FooMedia",
					},
					"BarMedia": &design.ResponseDefinition{
						Name: "BarMedia",
					},
				},
			},
			expected: `Response(BarMedia)
Response(FooMedia)`,
		},
		"with single definition": {
			definition: design.ActionDefinition{
				Responses: map[string]*design.ResponseDefinition{
					"FooMediat": &design.ResponseDefinition{
						Name: "FooMedia",
					},
				},
			},
			expected: `Response(FooMedia)`,
		},
		"without definition": {
			definition: design.ActionDefinition{},
			expected:   ``,
		},
	}
	for k, tc := range cases {
		buf := new(bytes.Buffer)
		if err := tmpl.ExecuteTemplate(buf, "response", tc.definition); err != nil {
			t.Fatalf("Execute returned %s", err)
		}
		actual := buf.String()
		if actual != tc.expected {
			t.Errorf("%s: \ngot:\n%v\nexpected:\n%v", k, actual, tc.expected)
		}
	}
}

func TestRoutingTmpl(t *testing.T) {
	cases := map[string]struct {
		definition interface{}
		expected   string
	}{
		"with multi definition": {
			definition: design.ActionDefinition{
				Routes: []*design.RouteDefinition{
					&design.RouteDefinition{
						Verb: "CONNECT",
						Path: "/",
					},
					&design.RouteDefinition{
						Verb: "DELETE",
						Path: "/",
					},
					&design.RouteDefinition{
						Verb: "GET",
						Path: "/",
					},
					&design.RouteDefinition{
						Verb: "HEAD",
						Path: "/",
					},
					&design.RouteDefinition{
						Verb: "OPTIONS",
						Path: "/",
					},
					&design.RouteDefinition{
						Verb: "PATCH",
						Path: "/",
					},
					&design.RouteDefinition{
						Verb: "POST",
						Path: "/",
					},
					&design.RouteDefinition{
						Verb: "PUT",
						Path: "/",
					},
					&design.RouteDefinition{
						Verb: "TRACE",
						Path: "/",
					},
				},
			},
			expected: `Routing(
CONNECT("/"),
DELETE("/"),
GET("/"),
HEAD("/"),
OPTIONS("/"),
PATCH("/"),
POST("/"),
PUT("/"),
TRACE("/"),
)`,
		},
		"with single definition": {
			definition: design.ActionDefinition{
				Routes: []*design.RouteDefinition{
					&design.RouteDefinition{
						Verb: "GET",
						Path: "/",
					},
				},
			},
			expected: `Routing(GET("/"))`,
		},
		"without definition": {
			definition: design.ActionDefinition{},
			expected:   ``,
		},
	}
	for k, tc := range cases {
		buf := new(bytes.Buffer)
		if err := tmpl.ExecuteTemplate(buf, "routing", tc.definition); err != nil {
			t.Fatalf("Execute returned %s", err)
		}
		actual := buf.String()
		if actual != tc.expected {
			t.Errorf("%s: got %v, expected %v", k, actual, tc.expected)
		}
	}
}

func TestTRACETmpl(t *testing.T) {
	cases := map[string]struct {
		definition interface{}
		expected   string
	}{
		"with trace definition": {
			definition: design.RouteDefinition{
				Verb: "TRACE",
				Path: "/",
			},
			expected: `TRACE("/")`,
		},
		"with get definition": {
			definition: design.RouteDefinition{
				Verb: "GET",
				Path: "/",
			},
			expected: ``,
		},
		"without definition": {
			definition: design.RouteDefinition{},
			expected:   ``,
		},
	}
	for k, tc := range cases {
		buf := new(bytes.Buffer)
		if err := tmpl.ExecuteTemplate(buf, "trace", tc.definition); err != nil {
			t.Fatalf("Execute returned %s", err)
		}
		actual := buf.String()
		if actual != tc.expected {
			t.Errorf("%s: \ngot:\n%v\nexpected:\n%v", k, actual, tc.expected)
		}
	}
}

func TestTypeTmpl(t *testing.T) {
	cases := map[string]struct {
		definition interface{}
		expected   string
	}{
		"with multi definition": {
			definition: design.APIDefinition{
				Types: map[string]*design.UserTypeDefinition{
					"FooPayload": &design.UserTypeDefinition{
						TypeName: "FooPayload",
					},
					"BarPayload": &design.UserTypeDefinition{
						TypeName: "BarPayload",
					},
				},
			},
			expected: `var BarPayload = Type("BarPayload", func() {
})
var FooPayload = Type("FooPayload", func() {
})`,
		},
		"with single definition": {
			definition: design.APIDefinition{
				Types: map[string]*design.UserTypeDefinition{
					"FooPayload": &design.UserTypeDefinition{
						TypeName: "FooPayload",
					},
				},
			},
			expected: `var FooPayload = Type("FooPayload", func() {
})`,
		},
		"without definition": {
			definition: design.APIDefinition{},
			expected:   ``,
		},
	}
	for k, tc := range cases {
		buf := new(bytes.Buffer)
		if err := tmpl.ExecuteTemplate(buf, "type", tc.definition); err != nil {
			t.Fatalf("Execute returned %s", err)
		}
		actual := buf.String()
		if actual != tc.expected {
			t.Errorf("%s: \ngot:\n%v\nexpected:\n%v", k, actual, tc.expected)
		}
	}
}
