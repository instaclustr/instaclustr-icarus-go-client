/*
 * Instaclustr Icarus
 *
 * REST API for Instaclustr Icarus - a sidecar for Cassandra
 *
 * API version: 1.0.1
 * Contact: support@instaclustr.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package instaclustr_icarus

type AllOfRestoreOperationRequestImport_ struct {
	// has to be set to 'import' 
	Type_ string `json:"type"`
	Keyspace string `json:"keyspace,omitempty"`
	Table string `json:"table,omitempty"`
	SourceDir string `json:"sourceDir"`
	KeepLevel bool `json:"keepLevel,omitempty"`
	KeepRepaired bool `json:"keepRepaired,omitempty"`
	NoVerify bool `json:"noVerify,omitempty"`
	NoVerifyTokens bool `json:"noVerifyTokens,omitempty"`
	NoInvalidateCaches bool `json:"noInvalidateCaches,omitempty"`
	// defaults to false, if true, noVerifyTokens, noInvalidateCaches and noVerify will be set to true automatically 
	Quick bool `json:"quick,omitempty"`
	ExtendedVerify bool `json:"extendedVerify,omitempty"`
}