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
import (
	"time"
)

type SidecarVersion struct {
	// textual representation of Sidecar version
	Version string `json:"version,omitempty"`
	// timestamp this Sidecar was built
	BuildTime time.Time `json:"buildTime,omitempty"`
	// git commit hash this Sidecar was built from
	GitCommit string `json:"gitCommit,omitempty"`
}