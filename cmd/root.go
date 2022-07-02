package cmd

import (
	"log"
	"os"

	"github.com/nenecchuu/lizbeth-be-core/cmd/http"
	"github.com/nenecchuu/lizbeth-be-core/cmd/telegram_listener"

	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:   "Lizbeth",
		Short: "Lizbeth - Backend Service",
		Long:  "Lizbeth - Backend Service",
	}
)

func Execute() {
	rootCmd.AddCommand(http.ServeHTTPCmd(), telegram_listener.ServeTelegramCmd())
	http.ServeHTTPCmd().Flags().StringP("config", "c", "config/file", "Config URL dir i.e. config/file")
	telegram_listener.ServeTelegramCmd().Flags().StringP("config", "c", "config/file", "Config URL dir i.e. config/file")

	if err := rootCmd.Execute(); err != nil {
		log.Fatalln("Error: \n", err.Error())
		os.Exit(-1)
	}
}
