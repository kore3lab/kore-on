package model

type KoreonToml struct {
	Koreon struct {
		ClusterName      string `toml:"cluster-name,omitempty"`
		ClusterID        string `toml:"cluster-id,omitempty"`
		InstallDir       string `toml:"install-dir,omitempty"`
		CertValidityDays int    `toml:"cert-validity-days,omitempty"`

		//#Airgap
		ClosedNetwork              bool   `toml:"closed-network,omitempty"`
		LocalRepository            string `toml:"local-repository,omitempty"`
		LocalRepositoryArchiveFile string `toml:"local-repository-archive-file"`
		DebugMode                  bool   `toml:"debug-mode,omitempty"`
	} `toml:"koreon,omitempty"`

	Kubernetes struct {
		Version          string   `toml:"version,omitempty"`
		ContainerRuntime string   `toml:"container-runtime"`
		KubeProxyMode    string   `toml:"kube-proxy-mode"`
		VxlanMode        bool     `toml:"vxlan-mode"`
		ServiceCidr      string   `toml:"service-cidr,omitempty"`
		PodCidr          string   `toml:"pod-cidr,omitempty"`
		NodePortRange    string   `toml:"node-port-range,omitempty"`
		AuditLogEnable   bool     `toml:"audit-log-enable"`
		ApiSans          []string `toml:"api-sans,omitempty"`

		Etcd struct {
			IP            []string `toml:"ip,omitempty"`
			PrivateIP     []string `toml:"private-ip,omitempty"`
			EncryptSecret bool     `toml:"encrypt-secret,omitempty"`
		} `toml:"etcd,omitempty"`
	} `toml:"kubernetes,omitempty"`

	NodePool struct {
		DataDir string `toml:"data-dir,omitempty"`

		Security struct {
			SSHUserID      string `toml:"ssh-user-id,omitempty"`
			SSHPort        int    `toml:"ssh-port,omitempty"`
			PrivateKeyPath string `toml:"private-key-path,omitempty"`
		} `toml:"security,omitempty"`

		Master struct {
			Name           string   `toml:"name,omitempty"`
			IP             []string `toml:"ip,omitempty"`
			PrivateIP      []string `toml:"private-ip,omitempty"`
			LbIP           string   `toml:"lb-ip,omitempty"`
			Isolated       bool     `toml:"isolated"`
			HaproxyInstall bool     `toml:"haproxy-install"`
		} `toml:"master,omitempty"`

		Node StrNode `toml:"node,omitempty"`
	} `toml:"node-pool,omitempty"`

	SharedStorage struct {
		Install    bool   `toml:"install"`
		StorageIP  string `toml:"storage-ip,omitempty"`
		PrivateIP  string `toml:"private-ip,omitempty"`
		VolumeDir  string `toml:"volume-dir,omitempty"`
		VolumeSize int    `toml:"volume-size,omitempty"`
		//StorageType       string `toml:"storage-type,omitempty"`

	} `toml:"shared-storage,omitempty"`

	PrivateRegistry struct {
		Install             bool   `toml:"install"`
		RegistryIP          string `toml:"registry-ip,omitempty"`
		RegistryDomain      string `toml:"registry-domain,omitempty"`
		PrivateIP           string `toml:"private-ip,omitempty"`
		DataDir             string `toml:"data-dir,omitempty"`
		PublicCert          bool   `toml:"public-cert"`
		RegistryArchiveFile string `toml:"registry-archive-file"`
		CertFile            struct {
			SslCertificate    string `toml:"ssl-certificate,omitempty"`
			SslCertificateKey string `toml:"ssl-certificate-key,omitempty"`
		} `toml:"cert-file,omitempty"`
	} `toml:"private-registry,omitempty"`
}

type StrNode struct {
	IP        []string `toml:"ip,omitempty"`
	PrivateIP []string `toml:"private-ip,omitempty"`
}
