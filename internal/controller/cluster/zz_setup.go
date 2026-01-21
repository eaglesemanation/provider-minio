// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	policy "github.com/eaglesemanation/provider-minio/internal/controller/cluster/iam/policy"
	serviceaccount "github.com/eaglesemanation/provider-minio/internal/controller/cluster/iam/serviceaccount"
	user "github.com/eaglesemanation/provider-minio/internal/controller/cluster/iam/user"
	userpolicyattachment "github.com/eaglesemanation/provider-minio/internal/controller/cluster/iam/userpolicyattachment"
	accesskey "github.com/eaglesemanation/provider-minio/internal/controller/cluster/minio/accesskey"
	providerconfig "github.com/eaglesemanation/provider-minio/internal/controller/cluster/providerconfig"
	bucket "github.com/eaglesemanation/provider-minio/internal/controller/cluster/s3/bucket"
)

// Setup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		policy.Setup,
		serviceaccount.Setup,
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

// SetupGated creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		policy.SetupGated,
		serviceaccount.SetupGated,
		user.SetupGated,
		userpolicyattachment.SetupGated,
		accesskey.SetupGated,
		providerconfig.SetupGated,
		bucket.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
