package tls

import (
	"crypto/x509"
	"crypto/x509/pkix"
	"net"

	"github.com/anton-sidelnikov/otc-openshift-installer/pkg/asset"
	"github.com/anton-sidelnikov/otc-openshift-installer/pkg/asset/installconfig"
	openstacktypes "github.com/anton-sidelnikov/otc-openshift-installer/pkg/types/openstack"
)

// MCSCertKey is the asset that generates the MCS key/cert pair.
type MCSCertKey struct {
	SignedCertKey
}

var _ asset.Asset = (*MCSCertKey)(nil)

// Dependencies returns the dependency of the the cert/key pair, which includes
// the parent CA, and install config if it depends on the install config for
// DNS names, etc.
func (a *MCSCertKey) Dependencies() []asset.Asset {
	return []asset.Asset{
		&RootCA{},
		&installconfig.InstallConfig{},
	}
}

// Generate generates the cert/key pair based on its dependencies.
func (a *MCSCertKey) Generate(dependencies asset.Parents) error {
	ca := &RootCA{}
	installConfig := &installconfig.InstallConfig{}
	dependencies.Get(ca, installConfig)

	hostname := internalAPIAddress(installConfig.Config)

	cfg := &CertCfg{
		Subject:      pkix.Name{CommonName: "system:machine-config-server"},
		ExtKeyUsages: []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth},
		Validity:     ValidityTenYears,
	}

	var vips []string
	switch installConfig.Config.Platform.Name() {
	case openstacktypes.Name:
		vips = installConfig.Config.OpenStack.APIVIPs
	}

	cfg.IPAddresses = []net.IP{}
	cfg.DNSNames = []string{hostname}
	for _, vip := range vips {
		cfg.IPAddresses = append(cfg.IPAddresses, net.ParseIP(vip))
		cfg.DNSNames = append(cfg.DNSNames, vip)
	}

	return a.SignedCertKey.Generate(cfg, ca, "machine-config-server", DoNotAppendParent)
}

// Name returns the human-friendly name of the asset.
func (a *MCSCertKey) Name() string {
	return "Certificate (mcs)"
}
