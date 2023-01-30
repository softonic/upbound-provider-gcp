/*
Copyright 2021 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/upbound/upjet/pkg/controller"

	projectiammember "github.com/upbound/provider-gcp/internal/controller/cloudplatform/projectiammember"
	serviceaccount "github.com/upbound/provider-gcp/internal/controller/cloudplatform/serviceaccount"
	serviceaccountiammember "github.com/upbound/provider-gcp/internal/controller/cloudplatform/serviceaccountiammember"
	serviceaccountkey "github.com/upbound/provider-gcp/internal/controller/cloudplatform/serviceaccountkey"
	address "github.com/upbound/provider-gcp/internal/controller/compute/address"
	backendservice "github.com/upbound/provider-gcp/internal/controller/compute/backendservice"
	globaladdress "github.com/upbound/provider-gcp/internal/controller/compute/globaladdress"
	globalforwardingrule "github.com/upbound/provider-gcp/internal/controller/compute/globalforwardingrule"
	healthcheck "github.com/upbound/provider-gcp/internal/controller/compute/healthcheck"
	managedsslcertificate "github.com/upbound/provider-gcp/internal/controller/compute/managedsslcertificate"
	sslcertificate "github.com/upbound/provider-gcp/internal/controller/compute/sslcertificate"
	targethttpsproxy "github.com/upbound/provider-gcp/internal/controller/compute/targethttpsproxy"
	urlmap "github.com/upbound/provider-gcp/internal/controller/compute/urlmap"
	cryptokey "github.com/upbound/provider-gcp/internal/controller/kms/cryptokey"
	cryptokeyiammember "github.com/upbound/provider-gcp/internal/controller/kms/cryptokeyiammember"
	keyring "github.com/upbound/provider-gcp/internal/controller/kms/keyring"
	keyringiammember "github.com/upbound/provider-gcp/internal/controller/kms/keyringiammember"
	keyringimportjob "github.com/upbound/provider-gcp/internal/controller/kms/keyringimportjob"
	secretciphertext "github.com/upbound/provider-gcp/internal/controller/kms/secretciphertext"
	providerconfig "github.com/upbound/provider-gcp/internal/controller/providerconfig"
	bucket "github.com/upbound/provider-gcp/internal/controller/storage/bucket"
	bucketiammember "github.com/upbound/provider-gcp/internal/controller/storage/bucketiammember"
)

// Setup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		projectiammember.Setup,
		serviceaccount.Setup,
		serviceaccountiammember.Setup,
		serviceaccountkey.Setup,
		address.Setup,
		backendservice.Setup,
		globaladdress.Setup,
		globalforwardingrule.Setup,
		healthcheck.Setup,
		managedsslcertificate.Setup,
		sslcertificate.Setup,
		targethttpsproxy.Setup,
		urlmap.Setup,
		cryptokey.Setup,
		cryptokeyiammember.Setup,
		keyring.Setup,
		keyringiammember.Setup,
		keyringimportjob.Setup,
		secretciphertext.Setup,
		providerconfig.Setup,
		bucket.Setup,
		bucketiammember.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
