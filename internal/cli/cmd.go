package cli

import (
    "fmt"
    "os"

    "github.com/spf13/cobra"
    "github.com/spf13/viper"
    "go.uber.org/zap"
)

var (
    rootCmd = &cobra.Command{
        Use:     "cli",
        Long:    "A CLI tool for managing tp-link based smart devices.",
        Short:   "CLI to manage tp-link smart devices.",
        Version: cliVersion(),
    }

    rootLogger *zap.Logger
)

func cliVersion() string {
    return "DEV"
}

func Execute() {
    if err := rootCmd.Execute(); err != nil {
        _, _ = fmt.Fprintln(os.Stderr, err)
        os.Exit(1)
    }
}

func init() {
    cobra.OnInitialize(initConfig)

    rootCmd.AddCommand(testCmd)
}

func initConfig() {
    var err error

    rootLogger, err = zap.NewDevelopment()
    if err != nil {
        fmt.Printf("Failed to create new logger: %s", err)
        os.Exit(1)
    }

    viper.AddConfigPath(".")
    viper.SetConfigName("tplink")

    viper.AutomaticEnv()

    if err = viper.ReadInConfig(); err != nil {
        _, _ = fmt.Fprintf(os.Stderr,
            "Failed to load config file: %s\n", err)
        // os.Exit(1)
    }
}
