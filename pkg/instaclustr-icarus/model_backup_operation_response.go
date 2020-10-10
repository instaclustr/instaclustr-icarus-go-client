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

type BackupOperationResponse struct {
	// type of operation, one has to set it to 'backup' in case he wants this request to be considered as a backup one 
	Type_ string `json:"type"`
	// location where SSTables will be uploaded. A value of the storageLocation property has to have exact format which is 'protocol://bucket/clusterName/dcName/nodeName'. protocol is either 'gcp', 's3', 'azure', 'oracle' or 'file:/'. For global requests, clusterName, dcName and nodeName do not need to be specified, so location would look like 's3://my-bucket'. 
	StorageLocation string `json:"storageLocation"`
	// directory of Cassandra, by default it is /var/lib/cassandra, in this path, one expects there is 'data' directory 
	CassandraDirectory string `json:"cassandraDirectory,omitempty"`
	// Based on this field, there will be throughtput per second computed based on what size data we want to upload we have. The formula is \"size / duration\". The lower the duration is, the higher throughput per second we will need and vice versa. This will influence e.g. responsiveness of a node to its business requests so one can control how much bandwidth is used for backup purposes in case a cluster is fully operational. The format of this field is \"amount unit\". 'unit' is just a (case-insensitive) java.util.concurrent.TimeUnit enum value. If not used, there will not be any restrictions as how fast an upload can be. 
	Duration string `json:"duration,omitempty"`
	Bandwidth *DataRate `json:"bandwidth,omitempty"`
	// number of threads used for upload, there might be at most so many uploading threads at any given time, when not set, it defaults to 10 
	ConcurrentConnections int32 `json:"concurrentConnections,omitempty"`
	// name of snapshot to make so this snapshot will be uploaded to storage location. If not specified, the name of snapshot will be automatically generated and it will have name 'autosnap-milliseconds-since-epoch'. There is a schema version automatically appended to this string, as well as a timestamp what that snapshot was taken. Both information are necessary in order to be able to distinguish between two backups with same name done against same schema version - but they are done at different time. The end user does not have to specify schema nor timestamp, it will be resovled automatically. 
	SnapshotTag string `json:"snapshotTag,omitempty"`
	// name of datacenter to backup, nodes in the other datacenter(s) will not be involved 
	Dc string `json:"dc,omitempty"`
	// database entities to backup, it might be either only keyspaces or only tables (from different keyspaces if needed), e.g. 'k1,k2' if one wants to backup whole keyspaces and 'ks1.t1,ks2,t2' if one wants to backup tables. These formats can not be used together so 'k1,k2.t2' is invalid. If this field is empty, all keyspaces are backed up.
	Entities string `json:"entities,omitempty"`
	// name of Kubernetes namespace to fetch Kubernetes secret for backups from, when not specified, it defaults to 'default' 
	K8sNamespace string `json:"k8sNamespace,omitempty"`
	// name of Kubernetes secret from which credentials used for the communication to cloud storage providers are read, if not specified, secret name to be read will be automatically derived in form 'cassandra-backup-restore-secret-cluster-{name-of-cluster}'. These secrets are used only in case protocol in storageLocation is gcp, azure or s3. 
	K8sSecretName string `json:"k8sSecretName,omitempty"`
	// flag saying if this request is meant to be global or not, once a global backup request is submitted to Sidecar, it will coordinate backup for all other nodes in a cluster (including itself) so from a point of view of a caller, one might just backup whole cluster by one request and repeatedly query its status based on returned operation id. 
	GlobalRequest bool `json:"globalRequest,omitempty"`
	// number of hours to wait until backup is considered failed if not finished already 
	Timeout int32 `json:"timeout,omitempty"`
	// Relevant during upload to S3-like bucket only. Specifies whether the metadata is copied from the source object or replaced with metadata provided in the request. Defaults to COPY. Consult com.amazonaws.services.s3.model.MetadatDirective for more information. 
	MetadataDirective string `json:"metadataDirective,omitempty"`
	// Relevant during upload to S3-like bucket only. If true, communication is done via HTTP instead of HTTPS. Defaults to false. 
	Insecure bool `json:"insecure,omitempty"`
	// Automatically creates a bucket if it does not exist. If a bucket does not exist, backup operation will fail. Defaults to false. 
	CreateMissingBucket bool `json:"createMissingBucket,omitempty"`
	// Do not check the existence of a bucket. Some storage providers (e.g. S3) requires a special permissions to be able to list buckets or query their existence which might not be allowed. This flag will skip that check. Keep in mind that if that bucket does not exist, the whole backup operation will fail. 
	SkipBucketVerification bool `json:"skipBucketVerification,omitempty"`
	// If set, a cluster topology file will be uploaded alongside a backup, defaults to false. This flag is implicitly set to true if a request is global - coordinator node will upload this file every time but no other nodes will. 
	UploadClusterTopology bool `json:"uploadClusterTopology,omitempty"`
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
