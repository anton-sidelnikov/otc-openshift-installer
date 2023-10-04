module github.com/anton-sidelnikov/otc-openshift-installer

go 1.20

require (
	github.com/AlecAivazis/survey/v2 v2.3.5
	github.com/apparentlymart/go-cidr v1.1.0
	github.com/awalterschulze/gographviz v0.0.0-20190522210029-fa59802746ab
	github.com/cavaliercoder/go-cpio v0.0.0-20180626203310-925f9528c45e
	github.com/clarketm/json v1.14.1
	github.com/containers/image v3.0.2+incompatible
	github.com/coreos/ignition/v2 v2.14.0
	github.com/coreos/stream-metadata-go v0.1.8
	github.com/daixiang0/gci v0.9.0
	github.com/diskfs/go-diskfs v1.4.0
	github.com/go-openapi/errors v0.20.3
	github.com/go-openapi/strfmt v0.21.5
	github.com/go-openapi/swag v0.22.3
	github.com/gofrs/flock v0.8.1
	github.com/golang/mock v1.7.0-rc.1
	github.com/google/uuid v1.3.1
	github.com/gophercloud/gophercloud v1.6.0
	github.com/gophercloud/utils v0.0.0-20230523080330-de873b9cf00d
	github.com/h2non/filetype v1.0.12
	github.com/hashicorp/go-multierror v1.1.1
	github.com/hashicorp/terraform-exec v0.17.3
	github.com/metal3-io/baremetal-operator/apis v0.2.0
	github.com/openshift/api v0.0.0-20230831095235-8891815aa476
	github.com/openshift/assisted-image-service v0.0.0-20230829160050-0b98ec74397b
	github.com/openshift/assisted-service/api v0.0.0
	github.com/openshift/assisted-service/client v0.0.0
	github.com/openshift/assisted-service/models v0.0.0
	github.com/openshift/client-go v0.0.0-20221019143426-16aed247da5c
	github.com/openshift/hive/apis v0.0.0-20220222213051-def9088fdb5a
	github.com/openshift/library-go v0.0.0-20220920133651-093893cf326b
	github.com/openshift/machine-config-operator v0.0.0
	github.com/pborman/uuid v1.2.0
	github.com/pelletier/go-toml v1.9.5
	github.com/pkg/diff v0.0.0-20210226163009-20ebb0f2a09e
	github.com/pkg/errors v0.9.1
	github.com/pkg/sftp v1.10.1
	github.com/prometheus/client_golang v1.16.0
	github.com/prometheus/common v0.44.0
	github.com/rogpeppe/go-internal v1.11.0
	github.com/shurcooL/vfsgen v0.0.0-20181202132449-6a9ea43bcacd
	github.com/sirupsen/logrus v1.9.3
	github.com/spf13/cobra v1.6.1
	github.com/stretchr/testify v1.8.2
	github.com/thedevsaddam/retry v0.0.0-20200324223450-9769a859cc6d
	github.com/ulikunitz/xz v0.5.11
	github.com/vincent-petithory/dataurl v1.0.0
	github.com/werf/lockgate v0.1.1
	golang.org/x/crypto v0.7.0
	golang.org/x/sync v0.3.0
	golang.org/x/sys v0.10.0
	golang.org/x/term v0.6.0
	gopkg.in/ini.v1 v1.66.6
	gopkg.in/yaml.v2 v2.4.0
	k8s.io/api v0.27.2
	k8s.io/apiextensions-apiserver v0.26.1
	k8s.io/apimachinery v0.27.2
	k8s.io/client-go v12.0.0+incompatible
	k8s.io/klog v1.0.0
	k8s.io/klog/v2 v2.90.1
	k8s.io/utils v0.0.0-20230209194617-a36077c30491
	sigs.k8s.io/controller-tools v0.10.0
	sigs.k8s.io/yaml v1.3.0
)

require (
	github.com/BurntSushi/toml v0.3.1 // indirect
	github.com/asaskevich/govalidator v0.0.0-20230301143203-a9d515a09cc2 // indirect
	github.com/beorn7/perks v1.0.1 // indirect
	github.com/cespare/xxhash/v2 v2.2.0 // indirect
	github.com/coreos/go-semver v0.3.0 // indirect
	github.com/coreos/go-systemd/v22 v22.3.2 // indirect
	github.com/coreos/vcontext v0.0.0-20211021162308-f1dbbca7bef4 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/elliotwutingfeng/asciiset v0.0.0-20230602022725-51bbb787efab // indirect
	github.com/emicklei/go-restful/v3 v3.10.1 // indirect
	github.com/fatih/color v1.13.0 // indirect
	github.com/go-logr/logr v1.2.4 // indirect
	github.com/go-openapi/analysis v0.21.2 // indirect
	github.com/go-openapi/jsonpointer v0.19.6 // indirect
	github.com/go-openapi/jsonreference v0.20.2 // indirect
	github.com/go-openapi/loads v0.21.1 // indirect
	github.com/go-openapi/runtime v0.23.0 // indirect
	github.com/go-openapi/spec v0.20.7 // indirect
	github.com/go-openapi/validate v0.22.0 // indirect
	github.com/gobuffalo/flect v0.2.5 // indirect
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/golang/protobuf v1.5.3 // indirect
	github.com/google/gnostic v0.6.9 // indirect
	github.com/google/go-cmp v0.5.9 // indirect
	github.com/google/gofuzz v1.2.0 // indirect
	github.com/hashicorp/errwrap v1.0.0 // indirect
	github.com/hashicorp/go-uuid v1.0.3 // indirect
	github.com/hashicorp/go-version v1.6.0 // indirect
	github.com/hashicorp/terraform-json v0.14.0 // indirect
	github.com/hexops/gotextdiff v1.0.3 // indirect
	github.com/imdario/mergo v0.3.13 // indirect
	github.com/inconshreveable/mousetrap v1.0.1 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.4 // indirect
	github.com/josharian/intern v1.0.0 // indirect
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/kballard/go-shellquote v0.0.0-20180428030007-95032a82bc51 // indirect
	github.com/kr/fs v0.1.0 // indirect
	github.com/mailru/easyjson v0.7.7 // indirect
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.16 // indirect
	github.com/matttproud/golang_protobuf_extensions v1.0.4 // indirect
	github.com/metal3-io/baremetal-operator/pkg/hardwareutils v0.2.0 // indirect
	github.com/mgutz/ansi v0.0.0-20170206155736-9520e82c474b // indirect
	github.com/mitchellh/go-homedir v1.1.0 // indirect
	github.com/mitchellh/mapstructure v1.5.0 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/munnerz/goautoneg v0.0.0-20191010083416-a7dc8b61c822 // indirect
	github.com/oklog/ulid v1.3.1 // indirect
	github.com/onsi/ginkgo/v2 v2.11.0 // indirect
	github.com/opencontainers/go-digest v1.0.0 // indirect
	github.com/opencontainers/image-spec v1.0.3-0.20211202193544-a5463b7f9c84 // indirect
	github.com/openshift/custom-resource-status v1.1.2 // indirect
	github.com/opentracing/opentracing-go v1.2.0 // indirect
	github.com/pierrec/lz4/v4 v4.1.17 // indirect
	github.com/pkg/xattr v0.4.9 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/prometheus/client_model v0.4.0 // indirect
	github.com/prometheus/procfs v0.11.0 // indirect
	github.com/shurcooL/httpfs v0.0.0-20190707220628-8d4bc4ba7749 // indirect
	github.com/spaolacci/murmur3 v1.1.0 // indirect
	github.com/spf13/pflag v1.0.6-0.20210604193023-d5e0c0615ace // indirect
	github.com/zclconf/go-cty v1.11.0 // indirect
	go.mongodb.org/mongo-driver v1.11.3 // indirect
	go.uber.org/atomic v1.7.0 // indirect
	go.uber.org/multierr v1.6.0 // indirect
	go.uber.org/zap v1.24.0 // indirect
	golang.org/x/mod v0.10.0 // indirect
	golang.org/x/net v0.12.0 // indirect
	golang.org/x/oauth2 v0.8.0 // indirect
	golang.org/x/text v0.11.0 // indirect
	golang.org/x/time v0.3.0 // indirect
	golang.org/x/tools v0.9.3 // indirect
	google.golang.org/appengine v1.6.7 // indirect
	google.golang.org/protobuf v1.31.0 // indirect
	gopkg.in/djherbis/times.v1 v1.3.0 // indirect
	gopkg.in/inf.v0 v0.9.1 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
	gorm.io/gorm v1.24.5 // indirect
	k8s.io/kube-openapi v0.0.0-20230501164219-8b0f38b5fd1f // indirect
	sigs.k8s.io/controller-runtime v0.14.5 // indirect
	sigs.k8s.io/json v0.0.0-20221116044647-bc3834ca7abd // indirect
	sigs.k8s.io/structured-merge-diff/v4 v4.2.3 // indirect
)

// OpenShift Forks
replace (
	github.com/metal3-io/baremetal-operator => github.com/openshift/baremetal-operator v0.0.0-20230531194024-8dde0991ffdd
	github.com/metal3-io/baremetal-operator/apis => github.com/openshift/baremetal-operator/apis v0.0.0-20230531194024-8dde0991ffdd
	github.com/metal3-io/baremetal-operator/pkg/hardwareutils => github.com/openshift/baremetal-operator/pkg/hardwareutils v0.0.0-20230531194024-8dde0991ffdd
	k8s.io/cloud-provider-vsphere => github.com/openshift/cloud-provider-vsphere v1.19.1-0.20211222185833-7829863d0558
	sigs.k8s.io/cluster-api => sigs.k8s.io/cluster-api v0.4.5
	sigs.k8s.io/cluster-api-provider-aws => github.com/openshift/cluster-api-provider-aws v0.2.1-0.20200929152424-eab2e087f366 // Indirect dependency through MAO from cluster API providers
	sigs.k8s.io/cluster-api-provider-azure => github.com/openshift/cluster-api-provider-azure v0.1.0-alpha.3.0.20210626224711-5d94c794092f // Indirect dependency through MAO from cluster API providers
)

// Pin MCO so it doesn't get downgraded
replace github.com/openshift/machine-config-operator => github.com/openshift/machine-config-operator v0.0.1-0.20201009041932-4fe8559913b8

// Needed because machine-api-operator uses a "later" v12 version, which is actually an earlier version.
// This should be kept in line with the k8s version used.
replace k8s.io/client-go => k8s.io/client-go v0.27.2

// Needed so that the InstallConfig CRD can be created. Later versions of controller-gen balk at using IPNet as a field.
replace sigs.k8s.io/controller-tools => sigs.k8s.io/controller-tools v0.3.1-0.20200617211605-651903477185

replace github.com/terraform-providers/terraform-provider-nutanix => github.com/nutanix/terraform-provider-nutanix v1.5.0

replace github.com/mattn/go-sqlite3 => github.com/mattn/go-sqlite3 v1.10.0

replace github.com/openshift/assisted-service/api => github.com/openshift/assisted-service/api v0.0.0-20230831114549-1922eda29cf8

replace github.com/openshift/assisted-service/client => github.com/openshift/assisted-service/client v0.0.0-20230831114549-1922eda29cf8

replace github.com/openshift/assisted-service/models => github.com/openshift/assisted-service/models v0.0.0-20230831114549-1922eda29cf8

// https://bugzilla.redhat.com/show_bug.cgi?id=2064702
replace golang.org/x/crypto => golang.org/x/crypto v0.0.0-20220315160706-3147a52a75dd

// https://bugzilla.redhat.com/show_bug.cgi?id=2100495
replace golang.org/x/text => golang.org/x/text v0.3.7

// https://issues.redhat.com/browse/OCPBUGS-5324
replace gopkg.in/yaml.v2 => gopkg.in/yaml.v2 v2.4.0

// https://issues.redhat.com/browse/OCPBUGS-5667
replace github.com/Masterminds/goutils => github.com/Masterminds/goutils v1.1.1

// https://bugzilla.redhat.com/show_bug.cgi?id=2045880
replace github.com/prometheus/client_golang => github.com/prometheus/client_golang v1.12.1

// https://issues.redhat.com/browse/OCPBUGS-6422
replace golang.org/x/net => golang.org/x/net v0.5.0

// https://issues.redhat.com/browse/OCPBUGS-8119
replace github.com/containerd/containerd => github.com/containerd/containerd v1.5.18

// https://issues.redhat.com/browse/OCPBUGS-8540
replace go.mongodb.org/mongo-driver => go.mongodb.org/mongo-driver v1.11.2
