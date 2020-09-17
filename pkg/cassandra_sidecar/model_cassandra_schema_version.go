/*
 * Instaclustr Cassandra Sidecar
 *
 * REST API for Cassandra Sidecar from Instaclustr
 *
 * API version: 2.0.0
 * Contact: support@instaclustr.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package cassandra_sidecar

type CassandraSchemaVersion struct {
	// UUID representing a version of Cassandra node as reported by its StorageServiceMBean#getSchemaVersion
	SchemaVersion string `json:"schemaVersion,omitempty"`
}