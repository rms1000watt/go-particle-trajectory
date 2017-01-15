package cmd

import (
	"log"
	"os"
	"strings"

	"github.com/rms1000watt/go-particle-trajectory/src"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

var (
	verbose    bool
	iterations int
)

// startCmd represents the start command
var startCmd = &cobra.Command{
	Use:   "start",
	Short: "TODO short explanation",
	Long:  `TODO long explanation`,
	Run: func(cmd *cobra.Command, args []string) {
		cfg := src.Config{
			Verbose:    verbose,
			Iterations: iterations,
		}
		log.Print("Config", cfg)
		src.Start(cfg)
	},
}

func init() {
	RootCmd.AddCommand(startCmd)
	startCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "Enable verbose output")
	startCmd.PersistentFlags().IntVarP(&iterations, "iterations", "i", 5, "iterations to loop")

	startCmd.PersistentFlags().VisitAll(func(f *pflag.Flag) {
		key := strings.ToUpper(strings.Replace(f.Name, "-", "_", -1))
		if val := os.Getenv(key); val != "" {
			if err := startCmd.PersistentFlags().Set(f.Name, val); err != nil {
				log.Println("ERROR", err)
			}
		}
	})
}
