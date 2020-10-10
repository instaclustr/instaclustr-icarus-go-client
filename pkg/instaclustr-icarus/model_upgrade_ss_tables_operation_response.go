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

type UpgradeSsTablesOperationResponse struct {
	Type_ string `json:"type"`
	// keyspace to upgrade SSTables of 
	Keyspace string `json:"keyspace"`
	// an array of tables to upgrade SSTables of, empty or not provided array will default to upgrading of SSTables of all tables in respective keyspace 
	Tables []string `json:"tables,omitempty"`
	// the number of threads to use - 0 means use all available, it never uses more than concurrent_compactor threads 
	Jobs int32 `json:"jobs,omitempty"`
	// include all sstables, even those already on the current version, defaults to false
	IncludeAllSStables bool `json:"includeAllSStables,omitempty"`
	// unique identifier of an operation, a random id is assigned to each operation after a request is submitted, from caller's perspective, an id is sent back as a response to his request so he can further query state of that operation, referencing id, by operations/{id} endpoint 
	Id string `json:"id"`
	// state of an operation, operation might be in various states, PENDING - this operation is pending for being submitted. RUNNING - this operation is actively doing its job, COMPLETED - this operation has finished successfully, CANCELLED - this operation was interrupted while being run, FAILED - this operation has finished errorneously 
	State string `json:"state"`
	// float from 0.0 to 1.0, 1.0 telling that operation is completed, either successfully or with errors. 
	Progress string `json:"progress"`
	// timestamp telling when this operation was created on Sidecar's side 
	CreationTime time.Time `json:"creationTime"`
	// timestamp telling when this operation was started by Sidecar, if an operation is created, it does not necessarily mean that it will be started right away, in most cases it is the case but if e.g. ExecutorService is full on its working thread, an execution of an operation is postponed and start time is updated only after that 
	StartTime time.Time `json:"startTime,omitempty"`
	// timestamp telling when an operation has finished, irrelevant of its result, an operation can be failed and it would still have this field populated. 
	CompletionTime time.Time `json:"completionTime,omitempty"`
	// This field contains serialized java.lang.Throwable in case this operation has failed 
	FailureCause *interface{} `json:"failureCause,omitempty"`
}