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

type RestoreOperationResponse struct {
	// type of operation, one has to set it to 'restore' in case he wants this request to be considered as a backup one 
	Type_ string `json:"type"`
	// similar to field in backup request but used for telling from where files should be downloaded, not uploaded, in case globalRequest field is set to true, it does not matter what dc and node id is used, these components in storageLocation path will be automatically changed. 
	StorageLocation string `json:"storageLocation"`
	// similar to field in backup request but used for downloading files, not uploading them 
	ConcurrentConnections int32 `json:"concurrentConnections,omitempty"`
	// similar to field in backup request 
	CassandraDirectory string `json:"cassandraDirectory,omitempty"`
	// directory where one expects to find 'conf/cassandra.yaml' file in case we need to update it with initial tokens in case restoration strategy is IN_PLACE. 
	CassandraConfigDirectory string `json:"cassandraConfigDirectory,omitempty"`
	// a flag saying if we should restore system keyspaces as well, relevant only for IN_PLACE restoration 
	RestoreSystemKeyspace bool `json:"restoreSystemKeyspace,omitempty"`
	// name of snapshot to restore 
	SnapshotTag string `json:"snapshotTag"`
	// similar to field in backup request, when empty, all entities in given snapshot will be restored 
	Entities string `json:"entities,omitempty"`
	// flag telling if cassandra.yaml should be updated with initial_tokens, relevant only in case of IN_PLACE strategy 
	UpdateCassandraYaml bool `json:"updateCassandraYaml,omitempty"`
	// strategy telling how we should go about restoration, please refer to details in backup and sidecar documentation 
	RestorationStrategyType string `json:"restorationStrategyType,omitempty"`
	// phase telling what should we do, this field has to be set just once as DOWNLOAD if globalRequest if true and coordinator of that request will take care of all other phases automatically on its own 
	RestorationPhase string `json:"restorationPhase,omitempty"`
	// flag saying if we should not delete truncated SSTables after they are imported, as part of CLEANUP phase, defaults to false 
	NoDeleteTruncates bool `json:"noDeleteTruncates,omitempty"`
	// flag saying if we should not delete downloaded SSTables from remote location, as part of CLEANUP phase, defaults to false 
	NoDeleteDownloads bool `json:"noDeleteDownloads,omitempty"`
	// flag saying if we should not download data from remote location as we expect them to be there already, defaults to false, setting this to true has sense only in case noDeleteDownloads was set to true in previous restoration requests 
	NoDownloadData bool `json:"noDownloadData,omitempty"`
	// version of schema we want to restore from, upon backup, a schema version is automatically appended to snapshot name and its manifest is uploaded under that name (plus timestamp at the end). In case we have two snapshots having same name, we might distinguish between them by this schema version. If schema version is not specified, we expect that there will be one and only one backup taken with respective snapshot name. This schema version has to match the version of a Cassandra node we are doing restore for (hence, by proxy, when global request mode is used, all nodes have to be on exact same schema version). 
	SchemaVersion string `json:"schemaVersion,omitempty"`
	// flag saying if we indeed want a schema version of a running node match with schema version a snapshot is taken on. there might be cases when we want to restore a table for which its CQL schema has not changed but it has changed for other table / keyspace but a schema for that node has changed by doing that. 
	ExactSchemaVersion bool `json:"exactSchemaVersion,omitempty"`
	Import_ *AllOfRestoreOperationRequestImport_ `json:"import,omitempty"`
	// flag saying that this request is a global one, meaning that a Sidecar this request is sent to will act as a restoration coordinator sending all other requests to each node in a cluster, for each phase. 
	GlobalRequest bool `json:"globalRequest,omitempty"`
	// name of Kubernetes namespace to fetch Kubernetes secret for restores from, when not specified, it defaults to 'default' 
	K8sNamespace string `json:"k8sNamespace,omitempty"`
	// name of Kubernetes secret from which credentials used for the communication to cloud storage providers are read, if not specified, secret name to be read will be automatically derived in form 'cassandra-backup-restore-secret-cluster-{name-of-cluster}'. These secrets are used only in case protocol in storageLocation is gcp, azure or s3. 
	K8sSecretName string `json:"k8sSecretName,omitempty"`
	// number of hours to wait until restore is considered failed if not finished already 
	Timeout int32 `json:"timeout,omitempty"`
	// if set to true, host id of node to restore will be resolved from remote topology file located in a bucket by translating it from provided nodeId of storageLocation field 
	ResolveHostIdFromTopology bool `json:"resolveHostIdFromTopology,omitempty"`
	// Do not check the existence of a bucket. Some storage providers (e.g. S3) requires a special permissions to be able to list buckets or query their existence which might not be allowed. This flag will skip that check. Keep in mind that if that bucket does not exist, the whole backup operation will fail. 
	SkipBucketVerification bool `json:"skipBucketVerification,omitempty"`
	// unique identifier of an operation, a random id is assigned to each operation after a request is submitted, from caller's perspective, an id is sent back as a response to his request so he can further query state of that operation, referencing id, by operations/{id} endpoint 
	Id string `json:"id"`
	// state of an operation, operation might be in various states, PENDING - this operation is pending for being submitted. RUNNING - this operation is actively doing its job, COMPLETED - this operation has finished successfully, CANCELLED - this operation was interrupted while being run, FAILED - this operation has finished errorneously 
	State string `json:"state"`
	// float from 0.0 to 1.0, 1.0 telling that operation is completed, either successfully or with errors. 
	Progress float64 `json:"progress"`
	// timestamp telling when this operation was created on Sidecar's side 
	CreationTime string `json:"creationTime"`
	// timestamp telling when this operation was started by Sidecar, if an operation is created, it does not necessarily mean that it will be started right away, in most cases it is the case but if e.g. ExecutorService is full on its working thread, an execution of an operation is postponed and start time is updated only after that 
	StartTime string `json:"startTime,omitempty"`
	// timestamp telling when an operation has finished, irrelevant of its result, an operation can be failed and it would still have this field populated. 
	CompletionTime string `json:"completionTime,omitempty"`
	// This field contains serialized java.lang.Throwable in case this operation has failed 
	FailureCause *interface{} `json:"failureCause,omitempty"`
}
