// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	bucket "github.com/alekc/provider-minio/internal/controller/bucket/bucket"
	notification "github.com/alekc/provider-minio/internal/controller/bucket/notification"
	policy "github.com/alekc/provider-minio/internal/controller/bucket/policy"
	replication "github.com/alekc/provider-minio/internal/controller/bucket/replication"
	retention "github.com/alekc/provider-minio/internal/controller/bucket/retention"
	serversideencryption "github.com/alekc/provider-minio/internal/controller/bucket/serversideencryption"
	accesskey "github.com/alekc/provider-minio/internal/controller/iam/accesskey"
	group "github.com/alekc/provider-minio/internal/controller/iam/group"
	groupmembership "github.com/alekc/provider-minio/internal/controller/iam/groupmembership"
	grouppolicy "github.com/alekc/provider-minio/internal/controller/iam/grouppolicy"
	grouppolicyattachment "github.com/alekc/provider-minio/internal/controller/iam/grouppolicyattachment"
	groupuserattachment "github.com/alekc/provider-minio/internal/controller/iam/groupuserattachment"
	ldapgrouppolicyattachment "github.com/alekc/provider-minio/internal/controller/iam/ldapgrouppolicyattachment"
	ldapuserpolicyattachment "github.com/alekc/provider-minio/internal/controller/iam/ldapuserpolicyattachment"
	policyiam "github.com/alekc/provider-minio/internal/controller/iam/policy"
	serviceaccount "github.com/alekc/provider-minio/internal/controller/iam/serviceaccount"
	user "github.com/alekc/provider-minio/internal/controller/iam/user"
	userpolicyattachment "github.com/alekc/provider-minio/internal/controller/iam/userpolicyattachment"
	policyilm "github.com/alekc/provider-minio/internal/controller/ilm/policy"
	tier "github.com/alekc/provider-minio/internal/controller/ilm/tier"
	key "github.com/alekc/provider-minio/internal/controller/kms/key"
	providerconfig "github.com/alekc/provider-minio/internal/controller/providerconfig"
)

// Setup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		bucket.Setup,
		notification.Setup,
		policy.Setup,
		replication.Setup,
		retention.Setup,
		serversideencryption.Setup,
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
