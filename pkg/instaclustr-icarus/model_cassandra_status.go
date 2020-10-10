/*
 * Instaclustr Icarus
 *
 * REST API for Instaclustr Icarus - a sidecar for Cassandra
 *
 * API version: 1.0.1
 * Contact: support@instaclustr.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package instaclustr-icarus

type CassandraStatus struct {
	// State of a Cassandra node as reported from StorageServiceMBean#getOperationMode
	NodeState string `json:"nodeState,omitempty"`
}
