package z

import (
  "fmt"
  "github.com/spf13/cobra"
  "os"
)

var developer string

var rootCmd = &cobra.Command{
  Use:   "zeit",
  Short: "Command line Zeiterfassung",
  Long:  `A command line time tracker.`,
}

func Execute() {
  if err := rootCmd.Execute(); err != nil {
    fmt.Println(err)
    os.Exit(1)
  }
}

func init() {
  cobra.OnInitialize(initConfig)
}

func initConfig() {
}
