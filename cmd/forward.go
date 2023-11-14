/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"log"
	"os/exec"
	"time"

	"github.com/spf13/cobra"
)

// // forwardCmd represents the forward command
// var forwardCmd = &cobra.Command{
// 	Use:   "forward",
// 	Short: "A brief description of your command",
// 	Long: `A longer description that spans multiple lines and likely contains examples
// and usage of using your command. For example:

// Cobra is a CLI library for Go that empowers applications.
// This application is a tool to generate the needed files
// to quickly create a Cobra application.`,
// 	Run: func(cmd *cobra.Command, args []string) {
// 		fmt.Println("forward called")
// 	},
// }

// forwardCmd represents the forward command
var forwardCmd = &cobra.Command{
    Use:   "forward",
    Short: "Forward a port from a Kubernetes service",
    Long: `Forward a port from a Kubernetes service. 
           Use the -d flag for default settings or provide [service] [localPort:remotePort] as arguments.`,
    Args: cobra.MaximumNArgs(2),
    Run: func(cmd *cobra.Command, args []string) {
        var service, namespace, portMapping string
        useDefault, _ := cmd.Flags().GetBool("default")

        if useDefault {
            // Set the default values
            service = "svc/argocd-server"
            namespace = "argocd"
            portMapping = "10000:443"
        } else if len(args) == 2 {
            // Use provided arguments
            service = args[0]
            portMapping = args[1]
            namespace, _ = cmd.Flags().GetString("namespace")
        } else {
            log.Fatal("Invalid arguments. Use -d for default settings or provide [service] [localPort:remotePort]")
        }

        for {
            log.Printf("Starting port-forward for service %s on namespace %s with port mapping %s\n", service, namespace, portMapping)
            if err := startPortForward(namespace, service, portMapping); err != nil {
                log.Printf("Error in port forwarding: %v, retrying in 5 seconds...", err)
                time.Sleep(5 * time.Second)
                continue
            }
            // Implement a check to see if the port-forwarding is still active
        }
    },
}

func startPortForward(namespace, service, portMapping string) error {
    cmd := exec.Command("kubectl", "port-forward", service, "-n", namespace, portMapping)
    return cmd.Run()
}


func init() {
	rootCmd.AddCommand(forwardCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	forwardCmd.PersistentFlags().String("namespace", "default", "k8s namespace")
	forwardCmd.PersistentFlags().BoolP("default", "d", false, "Use default port forwarding settings")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// forwardCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
