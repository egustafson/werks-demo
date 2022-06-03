package cmd

var (
	verboseFlag   = false
	profileFlag   = ""
	extensionFlag = ""
)

func init() {
	rootCmd.PersistentFlags().StringVarP(&profileFlag, "profile", "p", "",
		"config profile to load")
	rootCmd.PersistentFlags().StringVarP(&extensionFlag, "ext", "x", "",
		"config extension to load")
	rootCmd.PersistentFlags().BoolVarP(&verboseFlag, "verbose", "v", false,
		"enable verbose output")
}
