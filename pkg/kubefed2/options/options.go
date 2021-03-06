/*
Copyright 2018 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package options

import (
	"github.com/pkg/errors"
	"github.com/spf13/pflag"

	"github.com/kubernetes-sigs/federation-v2/pkg/controller/util"
)

// GlobalSubcommandOptions holds the configuration required by the subcommands of
// `kubefed2`.
type GlobalSubcommandOptions struct {
	HostClusterContext  string
	FederationNamespace string
	Kubeconfig          string
	DryRun              bool
}

// GlobalSubcommandBind adds the global subcommand flags to the flagset passed in.
func (o *GlobalSubcommandOptions) GlobalSubcommandBind(flags *pflag.FlagSet) {
	flags.StringVar(&o.Kubeconfig, "kubeconfig", "", "Path to the kubeconfig file to use for CLI requests.")
	flags.StringVar(&o.HostClusterContext, "host-cluster-context", "", "Host cluster context")
	flags.StringVar(&o.FederationNamespace, "federation-namespace", util.DefaultFederationSystemNamespace,
		"Namespace in the host cluster where the federation system components are installed.  This namespace will also be the target of propagation if the controller manager is configured with --limited-scope and clusters are joined with --limited-scope.")
	flags.BoolVar(&o.DryRun, "dry-run", false,
		"Run the command in dry-run mode, without making any server requests.")
}

// CommonSubcommandOptions holds the common configuration required by some of
// the subcommands of `kubefed2`.
type CommonSubcommandOptions struct {
	ClusterName      string
	ClusterContext   string
	ClusterNamespace string
	HostClusterName  string
}

// CommonSubcommandBind adds the common subcommand flags to the flagset passed in.
func (o *CommonSubcommandOptions) CommonSubcommandBind(flags *pflag.FlagSet) {
	flags.StringVar(&o.ClusterContext, "cluster-context", "",
		"Name of the cluster's context in the local kubeconfig. Defaults to cluster name if unspecified.")
	flags.StringVar(&o.ClusterNamespace, "registry-namespace", util.MulticlusterPublicNamespace,
		"Namespace in the host cluster where clusters are registered")
	flags.StringVar(&o.HostClusterName, "host-cluster-name", "",
		"If set, overrides the use of host-cluster-context name in resource names created in the target cluster. This option must be used when the context name has characters invalid for kubernetes resources like \"/\" and \":\".")
}

// SetName sets the name from the args passed in for the required positional
// argument.
func (o *CommonSubcommandOptions) SetName(args []string) error {
	if len(args) == 0 {
		return errors.New("NAME is required")
	}

	o.ClusterName = args[0]
	return nil
}
