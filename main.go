package main

//Import the packages we need
import (
    "fmt"
    "os"
    "time"
    "io"

    "github.com/sensu/sensu-go/types"
    "github.com/shirou/gopsutil/host"
    "github.com/spf13/cobra"
)

//Set up some variables. Most notably, warning and critical as time durations
var (
    warning, critical int64 
    stdin   *os.File
)

//Start our main function
func main() {
    rootCmd := configureRootCommand()
    if err := rootCmd.Execute(); err != nil {
        fmt.Fprintf(os.Stderr, "error: %v\n", err)
        os.Exit(1)
    }
}

//Set up our flags for the command. Note that we have time duration defaults for warning & critical
func configureRootCommand() *cobra.Command {
    cmd := &cobra.Command{
        Use:   "sensu-go-uptime-status",
        Short: "The Sensu Go check for system uptime",
        RunE:  run,
    }

    cmd.Flags().Int64VarP(&warning,
        "warning",
        "w",
        60,
        "Warning value in days. Default is 60")

    cmd.Flags().Int64VarP(&critical,
        "critical",
        "c",
        90,
        "Critical value in days. Default is 90")
		
    return cmd
}

func run(cmd *cobra.Command, args []string) error {

    if len(args) != 0 {
        _ = cmd.Help()
        return fmt.Errorf("invalid argument(s) received")
    }

    if stdin == nil {
        stdin = os.Stdin
    }
	
    event := &types.Event{}
	
    return checkUptime(event)
}

//Here we start the meat of what we do.
func checkUptime(event *types.Event) error {
	
    //Setting "CheckUptime" as a constant
    const checkName = "CheckUptime"
    const metricName = "current_system_uptime"
    //Setting uptime as the value retrieved from gopsutil
    uptime, err := host.Uptime()
	
    //Let's set up some error handling
    if err != nil {
        msg := fmt.Sprintf("Failed to determine uptime %s", err.Error())
        io.WriteString(os.Stdout, msg)
        os.Exit(3)
    }
	
    //Add a variable for uptimeSecs, which converts uptime to a duration
    uptimeSecs := time.Duration(uptime)*time.Second
    uptimeDays := int64(uptimeSecs.Seconds())/86400
    fmt.Sprintf("%d is uptime days",uptimeDays)
    //Sets up conditionss for a comparison
    if uptimeDays > critical {
        msg := fmt.Sprintf("%s CRITICAL - value = %d days | %s=%d days\n", checkName, uptimeDays, metricName, uptimeDays)
        io.WriteString(os.Stdout, msg)
        os.Exit(2)
    } else if uptimeDays >= warning && uptimeDays <= critical {
        msg := fmt.Sprintf("%s WARNING - value = %d days | %s=%d days\n", checkName, uptimeDays, metricName, uptimeDays)
        io.WriteString(os.Stdout, msg)
        os.Exit(1)
    } else {
        msg := fmt.Sprintf("%s OK - value = %d days | %s=%d days\n", checkName, uptimeDays, metricName, uptimeDays)
        io.WriteString(os.Stdout, msg)
        os.Exit(0)
    }
    return nil
}



