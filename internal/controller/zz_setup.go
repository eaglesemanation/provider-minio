// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	policy "github.com/eaglesemanation/provider-minio/internal/controller/iam/policy"
	user "github.com/eaglesemanation/provider-minio/internal/controller/iam/user"
	userpolicyattachment "github.com/eaglesemanation/provider-minio/internal/controller/iam/userpolicyattachment"
	accesskey "github.com/eaglesemanation/provider-minio/internal/controller/minio/accesskey"
	providerconfig "github.com/eaglesemanation/provider-minio/internal/controller/providerconfig"
	bucket "github.com/eaglesemanation/provider-minio/internal/controller/s3/bucket"
)

// Setup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		policy.Setup,
		user.Setup,
		userpolicyattachment.Setup,
		accesskey.Setup,
		providerconfig.Setup,
		bucket.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
