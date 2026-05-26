package kubeval

import (
	"github.com/xeipuuv/gojsonschema"
)

// ValidFormat is a type for quickly forcing
// new formats on the gojsonschema loader
type ValidFormat struct{}

// IsFormat always returns true and meets the
// gojsonschema.FormatChecker interface
func (f ValidFormat) IsFormat(input interface{}) bool {
	_ = "STUB: not implemented"

	// ValidationResult contains the details from
	// validating a given Kubernetes resource
	return false
}

type ValidationResult struct {
	FileName               string
	Kind                   string
	APIVersion             string
	ValidatedAgainstSchema bool
	Errors                 []gojsonschema.ResultError
	ResourceName           string
	ResourceNamespace      string
}

// VersionKind returns a string representation of this result's apiVersion and kind
func (v *ValidationResult) VersionKind() string { _ = "STUB: not implemented"; return "" }

// QualifiedName returns a string of the [namespace.]name of the k8s resource
func (v *ValidationResult) QualifiedName() string { _ = "STUB: not implemented"; return "" }

func determineSchemaURL(baseURL, kind, apiVersion string, config *Config) string {
	_ = "STUB: not implemented"
	// We have both the upstream Kubernetes schemas and the OpenShift schemas available
	// the tool can toggle between then using the config.OpenShift boolean flag and here we
	// use that to format the URL to match the required specification.
	return ""
}

// Most of the directories which store the schemas are prefixed with a v so as to
// match the tagging in the Kubernetes repository, apart from master.

// If we're using the openshift schemas, there's no further processing required

func determineSchemaBaseURL(config *Config) string {
	_ = "STUB: not implemented"
	// Order of precendence:
	// 1. If --openshift is passed, return the openshift schema location
	// 2. If a --schema-location is passed, use it
	// 3. If the KUBEVAL_SCHEMA_LOCATION is set, use it
	// 4. Otherwise, use the DefaultSchemaLocation
	return ""
}

// We only care that baseURL has a value after this call, so we can
// ignore LookupEnv's second return value

// validateResource validates a single Kubernetes resource against
// the relevant schema, detecting the type of resource automatically.
// Returns the result and raw YAML body as map.
func validateResource(data []byte, schemaCache map[string]*gojsonschema.Schema, config *Config) (ValidationResult, map[string]interface{}, error) {
	_ = "STUB: not implemented"
	return *new(ValidationResult), nil, nil
}

func validateAgainstSchema(body interface{}, resource *ValidationResult, schemaCache map[string]*gojsonschema.Schema, config *Config) ([]gojsonschema.ResultError, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Without forcing these types the schema fails to load
// Need to Work out proper handling for these types

// This error can only happen if the Object to validate is poorly formed. There's no hope of saving this one

// returned schema may be nil scehma is missing and missing schemas are allowed
func downloadSchema(resource *ValidationResult, schemaCache map[string]*gojsonschema.Schema, config *Config) (*gojsonschema.Schema, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// If the schema was previously cached, there's no work to be done

// We haven't cached this schema yet; look for one that works

// success! cache this and stop looking

// We couldn't find a schema for this URL, so take a note, then try the next URL

// We couldn't find a schema for this resource. Cache its lack of existence

func handleMissingSchema(err error, config *Config) ([]gojsonschema.ResultError, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// NewSchemaCache returns a new schema cache to be used with
// ValidateWithCache
func NewSchemaCache() map[string]*gojsonschema.Schema { _ = "STUB: not implemented"; return nil }

// Validate a Kubernetes YAML file, parsing out individual resources
// and validating them all according to the  relevant schemas
func Validate(input []byte, conf ...*Config) ([]ValidationResult, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ValidateWithCache validates a Kubernetes YAML file, parsing out individual resources
// and validating them all according to the relevant schemas
// Allows passing a kubeval.NewSchemaCache() to cache schemas in-memory
// between validations
func ValidateWithCache(input []byte, schemaCache map[string]*gojsonschema.Schema, conf ...*Config) ([]ValidationResult, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// split any list into its elements and add them to "bits"

// special case regexp for helm

// Save the fileName we were provided; if we detect a new fileName
// we'll use that, but we'll need to revert to the default afterward

// revert the filename back to the original

// set of [API version, kind, namespace, name]

// If resource has `metadata:name` attribute

func singleLineErrorFormat(es []error) string { _ = "STUB: not implemented"; return "" }
