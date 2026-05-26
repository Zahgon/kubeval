package kubeval

import (
	"github.com/spf13/cobra"
)

// DefaultSchemaLocation is the default location to search for schemas
const DefaultSchemaLocation = "https://kubernetesjsonschema.dev"

// OpenShiftSchemaLocation is the alternative location for OpenShift specific schemas
const OpenShiftSchemaLocation = "https://raw.githubusercontent.com/garethr/openshift-json-schema/master"

// A Config object contains various configuration data for kubeval
type Config struct {
	// DefaultNamespace is the namespace to assume in resources
	// if no namespace is set in `metadata:namespace` (as used with
	// `kubectl apply --namespace ...` or `helm install --namespace ...`,
	// for example)
	DefaultNamespace string

	// KubernetesVersion represents the version of Kubernetes
	// for which we should load the schema
	KubernetesVersion string

	// SchemaLocation is the base URL from which to search for schemas.
	// It can be either a remote location or a local directory
	SchemaLocation string

	// AdditionalSchemaLocations is a list of alternative base URLs from
	// which to search for schemas, given that the desired schema was not
	// found at SchemaLocation
	AdditionalSchemaLocations []string

	// OpenShift represents whether to test against
	// upstream Kubernetes or the OpenShift schemas
	OpenShift bool

	// Strict tells kubeval whether to prohibit properties not in
	// the schema. The API allows them, but kubectl does not
	Strict bool

	// IgnoreMissingSchemas tells kubeval whether to skip validation
	// for resource definitions without an available schema
	IgnoreMissingSchemas bool

	// ExitOnError tells kubeval whether to halt processing upon the
	// first error encountered or to continue, aggregating all errors
	ExitOnError bool

	// KindsToSkip is a list of kubernetes resources types with which to skip
	// schema validation
	KindsToSkip []string

	// KindsToReject is a list of case-sensitive prohibited kubernetes resources types
	KindsToReject []string

	// FileName is the name to be displayed when testing manifests read from stdin
	FileName string

	// OutputFormat is the name of the output formatter which will be used when
	// reporting results to the user.
	OutputFormat string

	// Quiet indicates whether non-results output should be emitted to the applications
	// log.
	Quiet bool

	// InsecureSkipTLSVerify controls whether to skip TLS certificate validation
	// when retrieving schema content over HTTPS
	InsecureSkipTLSVerify bool
}

// NewDefaultConfig creates a Config with default values
func NewDefaultConfig() *Config { _ = "STUB: not implemented"; return nil }

// AddKubevalFlags adds the default flags for kubeval to cmd
func AddKubevalFlags(cmd *cobra.Command, config *Config) *cobra.Command {
	_ = "STUB: not implemented"
	return nil
}
