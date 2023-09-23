// Code generated by cli_gen.go DO NOT EDIT.
// source: github.com/rancher/opni/plugins/metrics/apis/cortexops/cortexops.proto

package cortexops

import (
	context "context"
	errors "errors"
	cli "github.com/rancher/opni/internal/codegen/cli"
	compactor "github.com/rancher/opni/internal/cortex/config/compactor"
	querier "github.com/rancher/opni/internal/cortex/config/querier"
	runtimeconfig "github.com/rancher/opni/internal/cortex/config/runtimeconfig"
	storage "github.com/rancher/opni/internal/cortex/config/storage"
	validation "github.com/rancher/opni/internal/cortex/config/validation"
	cliutil "github.com/rancher/opni/pkg/opni/cliutil"
	driverutil "github.com/rancher/opni/pkg/plugins/driverutil"
	storage1 "github.com/rancher/opni/pkg/storage"
	flagutil "github.com/rancher/opni/pkg/util/flagutil"
	lo "github.com/samber/lo"
	cobra "github.com/spf13/cobra"
	pflag "github.com/spf13/pflag"
	errdetails "google.golang.org/genproto/googleapis/rpc/errdetails"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	proto "google.golang.org/protobuf/proto"
	protoiface "google.golang.org/protobuf/runtime/protoiface"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	strings "strings"
)

type contextKey_CortexOps_type struct{}

var contextKey_CortexOps contextKey_CortexOps_type

func ContextWithCortexOpsClient(ctx context.Context, client CortexOpsClient) context.Context {
	return context.WithValue(ctx, contextKey_CortexOps, client)
}

func CortexOpsClientFromContext(ctx context.Context) (CortexOpsClient, bool) {
	client, ok := ctx.Value(contextKey_CortexOps).(CortexOpsClient)
	return client, ok
}

var extraCmds_CortexOps []*cobra.Command

func addExtraCortexOpsCmd(custom *cobra.Command) {
	extraCmds_CortexOps = append(extraCmds_CortexOps, custom)
}

func BuildCortexOpsCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:               "ops",
		Short:             `The CortexOps service contains setup and configuration lifecycle actions for the managed Cortex cluster.`,
		Args:              cobra.NoArgs,
		ValidArgsFunction: cobra.NoFileCompletions,
	}

	cliutil.AddSubcommands(cmd, append([]*cobra.Command{
		BuildCortexOpsGetDefaultConfigurationCmd(),
		BuildCortexOpsSetDefaultConfigurationCmd(),
		BuildCortexOpsResetDefaultConfigurationCmd(),
		BuildCortexOpsGetConfigurationCmd(),
		BuildCortexOpsSetConfigurationCmd(),
		BuildCortexOpsResetConfigurationCmd(),
		BuildCortexOpsStatusCmd(),
		BuildCortexOpsInstallCmd(),
		BuildCortexOpsUninstallCmd(),
		BuildCortexOpsListPresetsCmd(),
		BuildCortexOpsConfigurationHistoryCmd(),
	}, extraCmds_CortexOps...)...)
	cli.AddOutputFlag(cmd)
	return cmd
}

var buildHooks_CortexOpsGetDefaultConfiguration []func(*cobra.Command)

func addBuildHook_CortexOpsGetDefaultConfiguration(hook func(*cobra.Command)) {
	buildHooks_CortexOpsGetDefaultConfiguration = append(buildHooks_CortexOpsGetDefaultConfiguration, hook)
}

func BuildCortexOpsGetDefaultConfigurationCmd() *cobra.Command {
	in := &driverutil.GetRequest{}
	cmd := &cobra.Command{
		Use:   "config get-default",
		Short: "Returns the default implementation-specific configuration, or one previously set.",
		Long: `
If a default configuration was previously set using SetDefaultConfiguration, it
returns that configuration. Otherwise, returns implementation-specific defaults.
An optional revision argument can be provided to get a specific historical
version of the configuration instead of the current configuration.

HTTP handlers for this method:
- GET /configuration/default
`[1:],
		Args:              cobra.NoArgs,
		ValidArgsFunction: cobra.NoFileCompletions,
		RunE: func(cmd *cobra.Command, args []string) error {
			client, ok := CortexOpsClientFromContext(cmd.Context())
			if !ok {
				cmd.PrintErrln("failed to get client from context")
				return nil
			}
			if in == nil {
				return errors.New("no input provided")
			}
			response, err := client.GetDefaultConfiguration(cmd.Context(), in)
			if err != nil {
				return err
			}
			cli.RenderOutput(cmd, response)
			return nil
		},
	}
	cmd.Flags().AddFlagSet(in.FlagSet())
	for _, hook := range buildHooks_CortexOpsGetDefaultConfiguration {
		hook(cmd)
	}
	return cmd
}

func BuildCortexOpsSetDefaultConfigurationCmd() *cobra.Command {
	in := &CapabilityBackendConfigSpec{}
	cmd := &cobra.Command{
		Use:   "config set-default",
		Short: "Sets the default configuration that will be used as the base for future configuration changes.",
		Long: `
If no custom default configuration is set using this method,
implementation-specific defaults may be chosen.
If all fields are unset, this will clear any previously-set default configuration
and revert back to the implementation-specific defaults.

This API is different from the SetConfiguration API, and should not be necessary
for most use cases. It can be used in situations where an additional persistence
layer that is not driver-specific is desired.

HTTP handlers for this method:
- PUT /configuration/default
`[1:],
		Args:              cobra.NoArgs,
		ValidArgsFunction: cobra.NoFileCompletions,
		RunE: func(cmd *cobra.Command, args []string) error {
			client, ok := CortexOpsClientFromContext(cmd.Context())
			if !ok {
				cmd.PrintErrln("failed to get client from context")
				return nil
			}
			if cmd.Flags().Lookup("interactive").Value.String() == "true" {
				client, ok := CortexOpsClientFromContext(cmd.Context())
				if !ok {
					cmd.PrintErrln("failed to get client from context")
					return nil
				}
				if curValue, err := client.GetDefaultConfiguration(cmd.Context(), &driverutil.GetRequest{}); err == nil {
					in = curValue
				}
				if edited, err := cliutil.EditInteractive(in); err != nil {
					return err
				} else {
					in = edited
				}
			} else if fileName := cmd.Flags().Lookup("file").Value.String(); fileName != "" {
				if err := cliutil.LoadFromFile(in, fileName); err != nil {
					return err
				}
			}
			if in == nil {
				return errors.New("no input provided")
			}
			_, err := client.SetDefaultConfiguration(cmd.Context(), in)
			if err != nil {
				return err
			}
			return nil
		},
	}
	cmd.Flags().StringP("file", "f", "", "path to a file containing the config, or - to read from stdin")
	cmd.Flags().BoolP("interactive", "i", false, "edit the config interactively in an editor")
	cmd.MarkFlagsMutuallyExclusive("file", "interactive")
	cmd.MarkFlagFilename("file")
	return cmd
}

func BuildCortexOpsResetDefaultConfigurationCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "config reset-default",
		Short: "Resets the default configuration to the implementation-specific defaults.",
		Long: `
If a custom default configuration was previously set using SetDefaultConfiguration,
this will clear it and revert back to the implementation-specific defaults.
Otherwise, this will have no effect.

HTTP handlers for this method:
- DELETE /configuration/default
`[1:],
		Args:              cobra.NoArgs,
		ValidArgsFunction: cobra.NoFileCompletions,
		RunE: func(cmd *cobra.Command, args []string) error {
			client, ok := CortexOpsClientFromContext(cmd.Context())
			if !ok {
				cmd.PrintErrln("failed to get client from context")
				return nil
			}
			_, err := client.ResetDefaultConfiguration(cmd.Context(), &emptypb.Empty{})
			if err != nil {
				return err
			}
			return nil
		},
	}
	return cmd
}

var buildHooks_CortexOpsGetConfiguration []func(*cobra.Command)

func addBuildHook_CortexOpsGetConfiguration(hook func(*cobra.Command)) {
	buildHooks_CortexOpsGetConfiguration = append(buildHooks_CortexOpsGetConfiguration, hook)
}

func BuildCortexOpsGetConfigurationCmd() *cobra.Command {
	in := &driverutil.GetRequest{}
	cmd := &cobra.Command{
		Use:   "config get",
		Short: "Gets the current configuration of the managed Cortex cluster.",
		Long: `
An optional revision argument can be provided to get a specific historical
version of the configuration instead of the current configuration.

HTTP handlers for this method:
- GET /configuration
`[1:],
		Args:              cobra.NoArgs,
		ValidArgsFunction: cobra.NoFileCompletions,
		RunE: func(cmd *cobra.Command, args []string) error {
			client, ok := CortexOpsClientFromContext(cmd.Context())
			if !ok {
				cmd.PrintErrln("failed to get client from context")
				return nil
			}
			if in == nil {
				return errors.New("no input provided")
			}
			response, err := client.GetConfiguration(cmd.Context(), in)
			if err != nil {
				return err
			}
			cli.RenderOutput(cmd, response)
			return nil
		},
	}
	cmd.Flags().AddFlagSet(in.FlagSet())
	for _, hook := range buildHooks_CortexOpsGetConfiguration {
		hook(cmd)
	}
	return cmd
}

var buildHooks_CortexOpsSetConfiguration []func(*cobra.Command)

func addBuildHook_CortexOpsSetConfiguration(hook func(*cobra.Command)) {
	buildHooks_CortexOpsSetConfiguration = append(buildHooks_CortexOpsSetConfiguration, hook)
}

func BuildCortexOpsSetConfigurationCmd() *cobra.Command {
	in := &CapabilityBackendConfigSpec{}
	cmd := &cobra.Command{
		Use:   "config set",
		Short: "Updates the configuration of the managed Cortex cluster to match the provided configuration.",
		Long: `
If the cluster is not installed, it will be configured, but remain disabled.
Otherwise, the already-installed cluster will be reconfigured.
The provided configuration will be merged with the default configuration
by directly overwriting fields. Slices and maps are overwritten and not combined.
Subsequent calls to this API will merge inputs with the current configuration,
not the default configuration.
When updating an existing configuration, the revision number in the updated configuration
must match the revision number of the existing configuration, otherwise a conflict
error will be returned. The timestamp field of the revision is ignored.

Note: some fields may contain secrets. The placeholder value "***" can be used to
keep an existing secret when updating the cluster configuration.

HTTP handlers for this method:
- PUT /configuration
`[1:],
		Args:              cobra.NoArgs,
		ValidArgsFunction: cobra.NoFileCompletions,
		RunE: func(cmd *cobra.Command, args []string) error {
			client, ok := CortexOpsClientFromContext(cmd.Context())
			if !ok {
				cmd.PrintErrln("failed to get client from context")
				return nil
			}
			if cmd.Flags().Lookup("interactive").Value.String() == "true" {
				client, ok := CortexOpsClientFromContext(cmd.Context())
				if !ok {
					cmd.PrintErrln("failed to get client from context")
					return nil
				}
				if curValue, err := client.GetConfiguration(cmd.Context(), &driverutil.GetRequest{}); err == nil {
					in = curValue
				}
				if edited, err := cliutil.EditInteractive(in); err != nil {
					return err
				} else {
					in = edited
				}
			} else if fileName := cmd.Flags().Lookup("file").Value.String(); fileName != "" {
				if err := cliutil.LoadFromFile(in, fileName); err != nil {
					return err
				}
			}
			if in == nil {
				return errors.New("no input provided")
			}
			_, err := client.SetConfiguration(cmd.Context(), in)
			if err != nil {
				return err
			}
			return nil
		},
	}
	cmd.Flags().StringP("file", "f", "", "path to a file containing the config, or - to read from stdin")
	cmd.Flags().BoolP("interactive", "i", false, "edit the config interactively in an editor")
	cmd.MarkFlagsMutuallyExclusive("file", "interactive")
	cmd.MarkFlagFilename("file")
	for _, hook := range buildHooks_CortexOpsSetConfiguration {
		hook(cmd)
	}
	return cmd
}

func BuildCortexOpsResetConfigurationCmd() *cobra.Command {
	in := &ResetRequest{}
	cmd := &cobra.Command{
		Use:   "config reset",
		Short: "Resets the configuration of the managed Cortex cluster to the current default configuration.",
		Long: `

The request may optionally contain a field mask to specify which fields should
be preserved. Furthermore, if a mask is set, the request may also contain a patch
object used to apply additional changes to the masked fields. These changes are
applied atomically at the time of reset. Fields present in the patch object, but
not in the mask, are ignored.

For example, with the following message:
  message Example {
    optional int32 a = 1;
    optional int32 b = 2;
    optional int32 c = 3;
  }

and current state:
  active:  { a: 1, b: 2, c: 3 }
  default: { a: 4, b: 5, c: 6 }

and reset request parameters:
{
  mask:    { paths: [ "a", "b" ] }
  patch:   { a: 100 }
}

The resulting active configuration will be:
 active:  {
   a: 100, // masked, set to 100 via patch
   b: 2,   // masked, but not set in patch, so left unchanged
   c: 6,   // not masked, reset to default
 }

HTTP handlers for this method:
- DELETE /configuration
`[1:],
		Args:              cobra.NoArgs,
		ValidArgsFunction: cobra.NoFileCompletions,
		RunE: func(cmd *cobra.Command, args []string) error {
			client, ok := CortexOpsClientFromContext(cmd.Context())
			if !ok {
				cmd.PrintErrln("failed to get client from context")
				return nil
			}
			if in == nil {
				return errors.New("no input provided")
			}
			_, err := client.ResetConfiguration(cmd.Context(), in)
			if err != nil {
				return err
			}
			return nil
		},
	}
	return cmd
}

func BuildCortexOpsStatusCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "status",
		Short: "Gets the current status of the managed Cortex cluster.",
		Long: `
The status includes the current install state, version, and metadata. If
the cluster is in the process of being reconfigured or uninstalled, it will
be reflected in the install state.
No guarantees are made about the contents of the metadata field; its
contents are strictly informational.

HTTP handlers for this method:
- GET /status
`[1:],
		Args:              cobra.NoArgs,
		ValidArgsFunction: cobra.NoFileCompletions,
		RunE: func(cmd *cobra.Command, args []string) error {
			client, ok := CortexOpsClientFromContext(cmd.Context())
			if !ok {
				cmd.PrintErrln("failed to get client from context")
				return nil
			}
			response, err := client.Status(cmd.Context(), &emptypb.Empty{})
			if err != nil {
				return err
			}
			cli.RenderOutput(cmd, response)
			return nil
		},
	}
	return cmd
}

func BuildCortexOpsInstallCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "install",
		Short: "Installs the managed Cortex cluster.",
		Long: `
The cluster will be installed using the current configuration, or the
default configuration if none is explicitly set.

HTTP handlers for this method:
- POST /install
`[1:],
		Args:              cobra.NoArgs,
		ValidArgsFunction: cobra.NoFileCompletions,
		RunE: func(cmd *cobra.Command, args []string) error {
			client, ok := CortexOpsClientFromContext(cmd.Context())
			if !ok {
				cmd.PrintErrln("failed to get client from context")
				return nil
			}
			_, err := client.Install(cmd.Context(), &emptypb.Empty{})
			if err != nil {
				return err
			}
			return nil
		},
	}
	return cmd
}

func BuildCortexOpsUninstallCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "uninstall",
		Short: "Uninstalls the managed Cortex cluster.",
		Long: `
Implementation details including error handling and system state requirements
are left to the cluster driver, and this API makes no guarantees about
the state of the cluster after the call completes (regardless of success).

HTTP handlers for this method:
- POST /uninstall
`[1:],
		Args:              cobra.NoArgs,
		ValidArgsFunction: cobra.NoFileCompletions,
		RunE: func(cmd *cobra.Command, args []string) error {
			client, ok := CortexOpsClientFromContext(cmd.Context())
			if !ok {
				cmd.PrintErrln("failed to get client from context")
				return nil
			}
			_, err := client.Uninstall(cmd.Context(), &emptypb.Empty{})
			if err != nil {
				return err
			}
			return nil
		},
	}
	return cmd
}

func BuildCortexOpsListPresetsCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "presets list",
		Short: "Returns a static list of presets that can be used as a base for configuring the managed Cortex cluster.",
		Long: `
There are several ways to use the presets, depending
on the desired behavior:
1. Set the default configuration to a preset spec, then use SetConfiguration
   to fill in any additional required fields (credentials, etc)
2. Add the required fields to the default configuration, then use
   SetConfiguration with a preset spec.
3. Leave the default configuration as-is, and use SetConfiguration with a
   preset spec plus the required fields.

HTTP handlers for this method:
- GET /presets
`[1:],
		Args:              cobra.NoArgs,
		ValidArgsFunction: cobra.NoFileCompletions,
		RunE: func(cmd *cobra.Command, args []string) error {
			client, ok := CortexOpsClientFromContext(cmd.Context())
			if !ok {
				cmd.PrintErrln("failed to get client from context")
				return nil
			}
			response, err := client.ListPresets(cmd.Context(), &emptypb.Empty{})
			if err != nil {
				return err
			}
			cli.RenderOutput(cmd, response)
			return nil
		},
	}
	return cmd
}

func BuildCortexOpsConfigurationHistoryCmd() *cobra.Command {
	in := &driverutil.ConfigurationHistoryRequest{}
	cmd := &cobra.Command{
		Use:   "config history",
		Short: "Get a list of all past revisions of the configuration.",
		Long: `
Will return the history for either the active or default configuration
depending on the specified target.
The entries are ordered from oldest to newest, where the last entry is
the current configuration.

HTTP handlers for this method:
- GET /configuration/history
`[1:],
		Args:              cobra.NoArgs,
		ValidArgsFunction: cobra.NoFileCompletions,
		RunE: func(cmd *cobra.Command, args []string) error {
			client, ok := CortexOpsClientFromContext(cmd.Context())
			if !ok {
				cmd.PrintErrln("failed to get client from context")
				return nil
			}
			if in == nil {
				return errors.New("no input provided")
			}
			response, err := client.ConfigurationHistory(cmd.Context(), in)
			if err != nil {
				return err
			}
			cli.RenderOutput(cmd, response)
			return nil
		},
	}
	cmd.Flags().AddFlagSet(in.FlagSet())
	cmd.RegisterFlagCompletionFunc("target", func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return []string{"ActiveConfiguration", "DefaultConfiguration"}, cobra.ShellCompDirectiveDefault
	})
	return cmd
}

func (in *CapabilityBackendConfigSpec) FlagSet(prefix ...string) *pflag.FlagSet {
	fs := pflag.NewFlagSet("CapabilityBackendConfigSpec", pflag.ExitOnError)
	fs.SortFlags = true
	if in.CortexConfig == nil {
		in.CortexConfig = &CortexApplicationConfig{}
	}
	fs.AddFlagSet(in.CortexConfig.FlagSet(append(prefix, "cortex-config")...))
	if in.Grafana == nil {
		in.Grafana = &GrafanaConfig{}
	}
	fs.AddFlagSet(in.Grafana.FlagSet(append(prefix, "grafana")...))
	return fs
}

func (in *CapabilityBackendConfigSpec) RedactSecrets() {
	if in == nil {
		return
	}
	in.CortexConfig.RedactSecrets()
}

func (in *CapabilityBackendConfigSpec) UnredactSecrets(unredacted *CapabilityBackendConfigSpec) error {
	if in == nil {
		return nil
	}
	var details []protoiface.MessageV1
	if err := in.CortexConfig.UnredactSecrets(unredacted.GetCortexConfig()); storage1.IsDiscontinuity(err) {
		for _, sd := range status.Convert(err).Details() {
			if info, ok := sd.(*errdetails.ErrorInfo); ok {
				info.Metadata["field"] = "cortexConfig." + info.Metadata["field"]
				details = append(details, info)
			}
		}
	}
	if len(details) == 0 {
		return nil
	}
	return lo.Must(status.New(codes.InvalidArgument, "cannot unredact: missing values for secret fields").WithDetails(details...)).Err()
}

func (in *CortexApplicationConfig) FlagSet(prefix ...string) *pflag.FlagSet {
	fs := pflag.NewFlagSet("CortexApplicationConfig", pflag.ExitOnError)
	fs.SortFlags = true
	if in.Limits == nil {
		in.Limits = &validation.Limits{}
	}
	fs.AddFlagSet(in.Limits.FlagSet(append(prefix, "limits")...))
	flagutil.SetDefValue(fs, strings.Join(append(prefix, "limits", "ingestion-rate"), "."), "600000")
	flagutil.SetDefValue(fs, strings.Join(append(prefix, "limits", "ingestion-burst-size"), "."), "1000000")
	flagutil.SetDefValue(fs, strings.Join(append(prefix, "limits", "compactor-blocks-retention-period"), "."), "seconds:2592000")
	if in.RuntimeConfig == nil {
		in.RuntimeConfig = &runtimeconfig.RuntimeConfigValues{}
	}
	fs.AddFlagSet(in.RuntimeConfig.FlagSet(append(prefix, "runtime-config")...))
	if in.Compactor == nil {
		in.Compactor = &compactor.Config{}
	}
	fs.AddFlagSet(in.Compactor.FlagSet(append(prefix, "compactor")...))
	if in.Querier == nil {
		in.Querier = &querier.Config{}
	}
	fs.AddFlagSet(in.Querier.FlagSet(append(prefix, "querier")...))
	if in.Storage == nil {
		in.Storage = &storage.Config{}
	}
	fs.AddFlagSet(in.Storage.FlagSet(append(prefix, "storage")...))
	fs.Var(flagutil.StringPtrValue(flagutil.Ptr("debug"), &in.LogLevel), strings.Join(append(prefix, "log-level"), "."), "")
	return fs
}

func (in *CortexApplicationConfig) RedactSecrets() {
	if in == nil {
		return
	}
	in.Storage.RedactSecrets()
}

func (in *CortexApplicationConfig) UnredactSecrets(unredacted *CortexApplicationConfig) error {
	if in == nil {
		return nil
	}
	var details []protoiface.MessageV1
	if err := in.Storage.UnredactSecrets(unredacted.GetStorage()); storage1.IsDiscontinuity(err) {
		for _, sd := range status.Convert(err).Details() {
			if info, ok := sd.(*errdetails.ErrorInfo); ok {
				info.Metadata["field"] = "storage." + info.Metadata["field"]
				details = append(details, info)
			}
		}
	}
	if len(details) == 0 {
		return nil
	}
	return lo.Must(status.New(codes.InvalidArgument, "cannot unredact: missing values for secret fields").WithDetails(details...)).Err()
}

func (in *GrafanaConfig) FlagSet(prefix ...string) *pflag.FlagSet {
	fs := pflag.NewFlagSet("GrafanaConfig", pflag.ExitOnError)
	fs.SortFlags = true
	fs.Var(flagutil.BoolPtrValue(flagutil.Ptr(false), &in.Enabled), strings.Join(append(prefix, "enabled"), "."), "Whether to deploy a managed Grafana instance.")
	fs.Var(flagutil.StringPtrValue(flagutil.Ptr("latest"), &in.Version), strings.Join(append(prefix, "version"), "."), "The version of Grafana to deploy.")
	fs.Var(flagutil.StringPtrValue(nil, &in.Hostname), strings.Join(append(prefix, "hostname"), "."), "")
	return fs
}

func (in *CapabilityBackendConfigSpec) DeepCopyInto(out *CapabilityBackendConfigSpec) {
	out.Reset()
	proto.Merge(out, in)
}

func (in *CapabilityBackendConfigSpec) DeepCopy() *CapabilityBackendConfigSpec {
	return proto.Clone(in).(*CapabilityBackendConfigSpec)
}

func (in *CortexWorkloadsConfig) DeepCopyInto(out *CortexWorkloadsConfig) {
	out.Reset()
	proto.Merge(out, in)
}

func (in *CortexWorkloadsConfig) DeepCopy() *CortexWorkloadsConfig {
	return proto.Clone(in).(*CortexWorkloadsConfig)
}

func (in *CortexWorkloadSpec) DeepCopyInto(out *CortexWorkloadSpec) {
	out.Reset()
	proto.Merge(out, in)
}

func (in *CortexWorkloadSpec) DeepCopy() *CortexWorkloadSpec {
	return proto.Clone(in).(*CortexWorkloadSpec)
}

func (in *CortexApplicationConfig) DeepCopyInto(out *CortexApplicationConfig) {
	out.Reset()
	proto.Merge(out, in)
}

func (in *CortexApplicationConfig) DeepCopy() *CortexApplicationConfig {
	return proto.Clone(in).(*CortexApplicationConfig)
}

func (in *GrafanaConfig) DeepCopyInto(out *GrafanaConfig) {
	out.Reset()
	proto.Merge(out, in)
}

func (in *GrafanaConfig) DeepCopy() *GrafanaConfig {
	return proto.Clone(in).(*GrafanaConfig)
}

func (in *PresetList) DeepCopyInto(out *PresetList) {
	out.Reset()
	proto.Merge(out, in)
}

func (in *PresetList) DeepCopy() *PresetList {
	return proto.Clone(in).(*PresetList)
}

func (in *Preset) DeepCopyInto(out *Preset) {
	out.Reset()
	proto.Merge(out, in)
}

func (in *Preset) DeepCopy() *Preset {
	return proto.Clone(in).(*Preset)
}

func (in *DryRunRequest) DeepCopyInto(out *DryRunRequest) {
	out.Reset()
	proto.Merge(out, in)
}

func (in *DryRunRequest) DeepCopy() *DryRunRequest {
	return proto.Clone(in).(*DryRunRequest)
}

func (in *DryRunResponse) DeepCopyInto(out *DryRunResponse) {
	out.Reset()
	proto.Merge(out, in)
}

func (in *DryRunResponse) DeepCopy() *DryRunResponse {
	return proto.Clone(in).(*DryRunResponse)
}

func (in *ConfigurationHistoryResponse) DeepCopyInto(out *ConfigurationHistoryResponse) {
	out.Reset()
	proto.Merge(out, in)
}

func (in *ConfigurationHistoryResponse) DeepCopy() *ConfigurationHistoryResponse {
	return proto.Clone(in).(*ConfigurationHistoryResponse)
}

func (in *ResetRequest) DeepCopyInto(out *ResetRequest) {
	out.Reset()
	proto.Merge(out, in)
}

func (in *ResetRequest) DeepCopy() *ResetRequest {
	return proto.Clone(in).(*ResetRequest)
}
