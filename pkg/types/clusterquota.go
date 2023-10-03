package types

// ClusterQuota contains the size, in cloud quota, of
// the cluster that was created by installer.
type ClusterQuota struct {
	Stub *StubQuota `json:"gcp,omitempty"`
}

type StubQuota struct {
}
