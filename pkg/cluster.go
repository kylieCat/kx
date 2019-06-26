package pkg

import "errors"

type ClusterConf struct {
	CertificateAuthorityData string `yaml:"certificate-authority-data,omitempty"`
	CertificateAuthority     string `yaml:"certificate-authority,omitempty"`
	Server                   string `yaml:"server,omitempty"`
	InsecureSkipVerify bool `yaml:"insecure-skip-tls-verify,omitempty"`
}

type Cluster struct {
	ClusterConf ClusterConf `yaml:"cluster"`
	Name        string      `yaml:"name"`
}

type Clusters []Cluster

func (cl Clusters) Len() int           { return len(cl) }
func (cl Clusters) Swap(i, j int)      { cl[i], cl[j] = cl[j], cl[i] }
func (cl Clusters) Less(i, j int) bool { return cl[i].Name < cl[j].Name }

func (cl Clusters) Get(name string) (Cluster, error) {
	for _, cluster := range cl {
		if cluster.Name == name {
			return cluster, nil
		}
	}
	return Cluster{}, errors.New("cluster not found: " + name)
}

type ClusterFilter func(ctx Cluster) bool

func (cl Clusters) Filter(filter ClusterFilter) Clusters {
	contexts := make(Clusters, 0, len(cl))
	for _, ctx := range cl {
		if filter(ctx) {
			contexts = append(contexts, ctx)
		}
	}
	return contexts
}