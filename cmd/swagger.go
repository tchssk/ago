package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"go/format"
	"io/ioutil"
	"log"
	"os"

	"github.com/goadesign/goa/design"
	"github.com/goadesign/goa/goagen/gen_swagger"
	"github.com/spf13/cobra"
)

// swaggerCmd represents the swagger command
var swaggerCmd = &cobra.Command{
	Use:   "swagger",
	Short: "Generate design from swagger definitions",
	Long:  `Generate design from swagger definitions`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			log.Fatal("invalid file path")
			return
		}
		data, err := ioutil.ReadFile(args[0])
		if err != nil {
			log.Fatal(err)
		}
		var swagger genswagger.Swagger
		if err := json.Unmarshal(data, &swagger); err != nil {
			log.Fatal(err)
		}
		buf := new(bytes.Buffer)
		if err := tmpl.ExecuteTemplate(buf, "all", swaggerToAPI(swagger)); err != nil {
			log.Fatal(err)
		}
		formated, err := format.Source(buf.Bytes())
		if err != nil {
			fmt.Println(string(buf.Bytes()))
			log.Fatal(err)
		}
		if _, err := bytes.NewBuffer(formated).WriteTo(os.Stdout); err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	RootCmd.AddCommand(swaggerCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// swaggerCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// swaggerCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}

func swaggerToAPI(swagger genswagger.Swagger) *design.APIDefinition {
	api := design.APIDefinition{
		Host:     swagger.Host,
		Schemes:  swagger.Schemes,
		BasePath: swagger.BasePath,
		//		Params *AttributeDefinition
		Consumes: []*design.EncodingDefinition{
			&design.EncodingDefinition{
				MIMETypes: swagger.Consumes,
				Encoder:   false,
			},
		},
		Produces: []*design.EncodingDefinition{
			&design.EncodingDefinition{
				MIMETypes: swagger.Produces,
				Encoder:   true,
			},
		},
		//		Origins map[string]*CORSDefinition
		//		Resources map[string]*ResourceDefinition
		//		Types map[string]*UserTypeDefinition
		//		MediaTypes map[string]*MediaTypeDefinition
		//		Traits map[string]*dslengine.TraitDefinition
		//		Responses map[string]*ResponseDefinition
		//		ResponseTemplates map[string]*ResponseTemplateDefinition
		//		DefaultResponses map[string]*ResponseDefinition
		//		DefaultResponseTemplates map[string]*ResponseTemplateDefinition
		//		DSLFunc func()
		//		Metadata dslengine.MetadataDefinition
		//		SecuritySchemes []*SecuritySchemeDefinition
		//		Security *SecurityDefinition
		//		NoExamples bool
	}
	if swagger.Info != nil {
		api.Name = swagger.Info.Title
		api.Title = swagger.Info.Title
		api.Description = swagger.Info.Description
		api.Version = swagger.Info.Version
		api.TermsOfService = swagger.Info.TermsOfService
		api.Contact = swagger.Info.Contact
		api.License = swagger.Info.License
	}
	if swagger.ExternalDocs != nil {
		api.Docs = &design.DocsDefinition{
			Description: swagger.ExternalDocs.Description,
			URL:         swagger.ExternalDocs.URL,
		}
	}
	return &api
}
