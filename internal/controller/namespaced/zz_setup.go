// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	bucket "github.com/alekc/provider-minio/internal/controller/namespaced/bucket/bucket"
	notification "github.com/alekc/provider-minio/internal/controller/namespaced/bucket/notification"
	object "github.com/alekc/provider-minio/internal/controller/namespaced/bucket/object"
	policy "github.com/alekc/provider-minio/internal/controller/namespaced/bucket/policy"
	replication "github.com/alekc/provider-minio/internal/controller/namespaced/bucket/replication"
	retention "github.com/alekc/provider-minio/internal/controller/namespaced/bucket/retention"
	serversideencryption "github.com/alekc/provider-minio/internal/controller/namespaced/bucket/serversideencryption"
	versioning "github.com/alekc/provider-minio/internal/controller/namespaced/bucket/versioning"
	accesskey "github.com/alekc/provider-minio/internal/controller/namespaced/iam/accesskey"
	group "github.com/alekc/provider-minio/internal/controller/namespaced/iam/group"
	groupmembership "github.com/alekc/provider-minio/internal/controller/namespaced/iam/groupmembership"
	grouppolicy "github.com/alekc/provider-minio/internal/controller/namespaced/iam/grouppolicy"
	grouppolicyattachment "github.com/alekc/provider-minio/internal/controller/namespaced/iam/grouppolicyattachment"
	groupuserattachment "github.com/alekc/provider-minio/internal/controller/namespaced/iam/groupuserattachment"
	ldapgrouppolicyattachment "github.com/alekc/provider-minio/internal/controller/namespaced/iam/ldapgrouppolicyattachment"
	ldapuserpolicyattachment "github.com/alekc/provider-minio/internal/controller/namespaced/iam/ldapuserpolicyattachment"
	policyiam "github.com/alekc/provider-minio/internal/controller/namespaced/iam/policy"
	serviceaccount "github.com/alekc/provider-minio/internal/controller/namespaced/iam/serviceaccount"
	user "github.com/alekc/provider-minio/internal/controller/namespaced/iam/user"
	userpolicyattachment "github.com/alekc/provider-minio/internal/controller/namespaced/iam/userpolicyattachment"
	policyilm "github.com/alekc/provider-minio/internal/controller/namespaced/ilm/policy"
	tier "github.com/alekc/provider-minio/internal/controller/namespaced/ilm/tier"
	key "github.com/alekc/provider-minio/internal/controller/namespaced/kms/key"
	providerconfig "github.com/alekc/provider-minio/internal/controller/namespaced/providerconfig"
)

// Setup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		bucket.Setup,
		notification.Setup,
		object.Setup,
		policy.Setup,
		replication.Setup,
		retention.Setup,
		serversideencryption.Setup,
		versioning.Setup,
		accesskey.Setup,
		group.Setup,
		groupmembership.Setup,
		grouppolicy.Setup,
		grouppolicyattachment.Setup,
		groupuserattachment.Setup,
		ldapgrouppolicyattachment.Setup,
		ldapuserpolicyattachment.Setup,
		policyiam.Setup,
		serviceaccount.Setup,
		user.Setup,
		userpolicyattachment.Setup,
		policyilm.Setup,
		tier.Setup,
		key.Setup,
		providerconfig.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupGated creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		bucket.SetupGated,
		notification.SetupGated,
		object.SetupGated,
		policy.SetupGated,
		replication.SetupGated,
		retention.SetupGated,
		serversideencryption.SetupGated,
		versioning.SetupGated,
		accesskey.SetupGated,
		group.SetupGated,
		groupmembership.SetupGated,
		grouppolicy.SetupGated,
		grouppolicyattachment.SetupGated,
		groupuserattachment.SetupGated,
		ldapgrouppolicyattachment.SetupGated,
		ldapuserpolicyattachment.SetupGated,
		policyiam.SetupGated,
		serviceaccount.SetupGated,
		user.SetupGated,
		userpolicyattachment.SetupGated,
		policyilm.SetupGated,
		tier.SetupGated,
		key.SetupGated,
		providerconfig.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
