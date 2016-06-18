package main
import (
    "net"
    "fmt"
    "log"
    "github.com/spf13/cobra"
)

func main() {
    var prot string
    var remote string

    //Gets the IP of the host, if it can.
    var resolveUDP = &cobra.Command{
        Use:   "check [validate a remote connection.]",
        Short: "Validate a remote connection.",
        Long: "Validate a remote connection.",
        Run: func(cmd *cobra.Command, args []string) {


        addr, err := net.ResolveUDPAddr(prot, remote)
        if err != nil {
    				log.Printf("There was a problem: %s\n", err)
    			} else {
          fmt.Println("Connection successful!  Remote IP is: ", addr)
        }
      },
    }

    resolveUDP.PersistentFlags().StringVarP(&prot, "protocol", "p", "", "Protocol to use (only UDP for now)")
    resolveUDP.PersistentFlags().StringVarP(&remote, "remote", "r", "", "Fully qualified host/IP and port (host:port)")

    var rootCmd = &cobra.Command{Use: "cerebro"}
    //rootCmd.AddCommand(cmdPrint, cmdEcho)
    //cmdEcho.AddCommand(cmdTimes)
    rootCmd.AddCommand(resolveUDP)
    rootCmd.Execute()
}
