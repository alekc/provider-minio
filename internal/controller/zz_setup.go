// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	bucket "github.com/alekc/provider-minio/internal/controller/bucket/bucket"
	providerconfig "github.com/alekc/provider-minio/internal/controller/providerconfig"
	user "github.com/alekc/provider-minio/internal/controller/user/user"
	userpolicyattachment "github.com/alekc/provider-minio/internal/controller/user/userpolicyattachment"
)

// Setup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		bucket.Setup,
		providerconfig.Setup,
		user.Setup,
		userpolicyattachment.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
