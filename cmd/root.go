package cmd

import (
	_ "embed"
	"fmt"
	"os"
	"path/filepath"

	backme "github.com/dimus/backme/pkg"
	"github.com/gnames/gnsys"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

//go:embed backme.yaml
var configYAML string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "backme",
	Short: "Backup files organizer.",

	Long: `Quite often big files (like database dumps) are backed up on a
 regular basis. Such backups become a disk space hogs and when time comes to
 restoring data it is cumbersome to find the right file for the job. This
 script checks files of a certain pattern and sorts them by date into daily
 monthly, yearly directories, progressively deleting more and more of aged
 files. For example it keeps all files from the last 24 hours, but keeps only
 one file per day in the last 30 days, only one file per month for the first
 3 years, and after that only 1 file per year. As a result backup takes
 significantly less space and it is easier to find a file from a specific
 period of time.`,

	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		versionFlag(cmd)

		conf := getConfig()
		err := backme.Organize(conf)
		if err != nil {
			log.Fatal().Err(err).Msg("")
		}
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.
	rootCmd.PersistentFlags().BoolP("version", "V", false, "Show build version and date")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	//rootCmd.Flags().BoolP("version", "v", false, "Show build version and date")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	var configDir string
	var err error
	configFile := "backme"

	// Find config directory.
	configDir = "/etc"

	// Search config in home directory with name ".gnmatcher" (without extension).
	viper.AddConfigPath(configDir)
	viper.SetConfigName(configFile)

	configPath := filepath.Join(configDir, fmt.Sprintf("%s.yaml", configFile))

	exists, _ := gnsys.FileExists(configPath)
	if os.Getuid() != 0 && !exists {
		log.Warn().Msgf("Run at least once as a superuser to generate config file at '%s'", configPath)
		return
	}
	_ = touchConfigFile(configPath)

	// If a config file is found, read it in.
	err = viper.ReadInConfig()
	if err != nil {
		log.Warn().Msgf("Cannot use config file: %s", viper.ConfigFileUsed())
	}
}

func versionFlag(cmd *cobra.Command) {
	version, err := cmd.Flags().GetBool("version")
	if err != nil {
		log.Fatal().Err(err).Msg("")
	}
	if version {
		fmt.Printf("\nVersion: %s\n\nBuild:    %s\n\n",
			backme.Version, backme.Build)
		os.Exit(0)
	}
}

func getConfig() *backme.Config {
	conf := backme.NewConfig()
	err := viper.Unmarshal(conf)
	if err != nil {
		log.Fatal().Err(err).Msg("")
	}

	err = backme.CheckConfig(conf)
	if err != nil {
		log.Fatal().Err(err).Msg("")
	}

	log.Info().Msg("The following backup directories are in config:")
	for i, v := range conf.InputDirs {
		log.Info().Msgf("Backup dir %d: %s", i, v.Path)
		if _, err := os.Stat(v.Path); os.IsNotExist(err) {
			log.Fatal().Err(err).Msgf("Directory %s does not exist, exiting...", v.Path)
		}
	}
	log.Info().Msg(backme.LogSep)
	return conf
}

// touchConfigFile checks if config file exists, and if not, it gets created.
func touchConfigFile(configPath string) error {
	exists, _ := gnsys.FileExists(configPath)
	if exists {
		return nil
	}

	log.Info().Msgf("Creating config file: %s.", configPath)
	return createConfig(configPath)
}

// createConfig creates config file.
func createConfig(path string) error {
	err := gnsys.MakeDir(filepath.Dir(path))
	if err != nil {
		log.Warn().Err(err).Msgf("Cannot create dir %s.", path)
		return err
	}

	err = os.WriteFile(path, []byte(configYAML), 0644)
	if err != nil {
		log.Warn().Err(err).Msgf("Cannot write to file %s", path)
		return err
	}
	return nil
}
