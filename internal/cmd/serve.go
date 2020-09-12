package cmd

import (
    "github.com/meroedu/lantern/internal/driver/configuration"
    "github.com/spf13/cobra"

    "github.com/meroedu/lantern/internal/cmd/daemon"
    "github.com/meroedu/lantern/internal/driver"
)

var serveCmd = &cobra.Command{
    Use:   "serve",
    Short: "Starts server and serve HTTP API",
    Run: func(cmd *cobra.Command, args []string) {
        var (
            c        = configuration.NewViperProvider()
            registry = driver.NewRegistryDefault().WithConfiguration(c)
        )

        daemon.Serve(registry)
    },
}

func init() {
    rootCmd.AddCommand(serveCmd)
}
