package provision

// AWS provisioner for creating and destroying infrastructure on AWS.
type AWS struct {
	Terraform
	AccessKeyID     string
	SecretAccessKey string
}

// AWSTerraformData provider for creating and destroying infrastructure on AWS
type AWSTerraformData struct {
	Region            string `json:"region,omitempty"`
	PrivateSSHKeyPath string `json:"private_ssh_key_path"`
	PublicSSHKeyPath  string `json:"public_ssh_key_path"`
	ClusterName       string `json:"cluster_name"`
	MasterCount       int    `json:"master_count"`
	EtcdCount         int    `json:"etcd_count"`
	WorkerCount       int    `json:"worker_count"`
	IngressCount      int    `json:"ingress_count"`
	StorageCount      int    `json:"storage_count"`
}