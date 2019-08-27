package cmd

import (
	"fmt"

	"github.com/LINBIT/linstor-iscsi/pkg/crmcontrol"
	"github.com/LINBIT/linstor-iscsi/pkg/iscsi"
	"github.com/LINBIT/linstor-iscsi/pkg/linstorcontrol"
	"github.com/LINBIT/linstor-iscsi/pkg/targetutil"
	"github.com/logrusorgru/aurora"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

const (
	longStatusOk       = "Ok"
	longStatusBad      = "Stopped"
	longStatusUnknown  = "Unknown"
	longStatusDegraded = "Degraded"
)

func stateToLongStatus(state crmcontrol.LrmRunState) string {
	var str string
	switch state {
	case crmcontrol.Running:
		str = longStatusOk
	case crmcontrol.Stopped:
		str = longStatusBad
	default:
		str = longStatusUnknown
	}

	return stateToColor(state)(str).String()
}

func linstorStateToLongStatus(state linstorcontrol.ResourceState) string {
	var str string
	switch state {
	case linstorcontrol.Ok:
		str = longStatusOk
	case linstorcontrol.Degraded:
		str = longStatusDegraded
	case linstorcontrol.Bad:
		str = longStatusBad
	default:
		str = longStatusUnknown
	}
	return linstorStateToColor(state)(str).String()
}

// statusCmd represents the status command
var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "Displays the status of an iSCSI target or logical unit",
	Long: `Triggers Pacemaker to probe the resoruce primitives of this iSCSI target.
That means Pacemaker will run the status operation on the nodes where the
resource can run.
This makes sure that Pacemakers view of the world is updated to the state
of the world.

For example:
linstor-iscsi status --iqn=iqn.2019-08.com.linbit:example --lun=0`,
	Args: cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		if !cmd.Flags().Changed("controller") {
			foundIP, err := crmcontrol.FindLinstorController()
			if err == nil { // it might be ok to not find it...
				controller = foundIP
			}
		}
		targetCfg := targetutil.TargetConfig{
			IQN:  iqn,
			LUNs: []*targetutil.LUN{&targetutil.LUN{ID: uint8(lun)}},
		}
		target := targetutil.NewTargetMust(targetCfg)
		targetName, err := targetutil.ExtractTargetName(targetCfg.IQN)
		if err != nil {
			log.Fatal(err)
		}
		linstorCfg := linstorcontrol.Linstor{
			Loglevel:     log.GetLevel().String(),
			ControllerIP: controller,
			ResourceName: linstorcontrol.ResourceNameFromLUN(targetName, uint8(lun)),
		}
		iscsiCfg := &iscsi.ISCSI{Linstor: linstorCfg, Target: target}

		state, err := iscsiCfg.ProbeResource()
		if err != nil {
			log.Fatal(err)
		}

		linstorState, err := linstorCfg.AggregateResourceState()
		if err != nil {
			log.Warning(err)
			linstorState = linstorcontrol.Unknown
		}

		fmt.Printf("Status of target %s, LUN %d:\n", aurora.Cyan(iqn),
			aurora.Cyan(lun))
		fmt.Printf("- Target: %s\n", stateToLongStatus(state.TargetState))
		fmt.Printf("- LU: %s\n", stateToLongStatus(state.LUStates[uint8(lun)]))
		fmt.Printf("- IP: %s\n", stateToLongStatus(state.IPState))
		fmt.Printf("- LINSTOR: %s\n", linstorStateToLongStatus(linstorState))
	},
}

func init() {
	rootCmd.AddCommand(statusCmd)

	statusCmd.MarkPersistentFlagRequired("iqn")
	statusCmd.MarkPersistentFlagRequired("lun")
}