module github.com/cfergeau/go-snc

go 1.16

require (
	github.com/imdario/mergo v0.3.7 // indirect
	github.com/openshift/client-go v0.0.0-20210521082421-73d9475a9142
	k8s.io/apimachinery v0.21.1
	k8s.io/client-go v0.21.1
)

replace (
	github.com/apcera/gssapi => github.com/openshift/gssapi v0.0.0-20161010215902-5fb4217df13b
	k8s.io/apimachinery => github.com/openshift/kubernetes-apimachinery v0.0.0-20210521074607-b6b98f7a1855
	k8s.io/client-go => github.com/openshift/kubernetes-client-go v0.0.0-20210521075216-71b63307b5df
)
