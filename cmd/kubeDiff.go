package cmd

import (
	"github.com/spf13/cobra"
)

//var clientSet;

func init() {
	rootCmd.AddCommand(kubePlan)
}

var kubePlan = &cobra.Command{
	Use:   "plan",
	Short: `Emulates "kubectl apply --dry-run" and creates diff files`,
	Run: func(cmd *cobra.Command, args []string) {
		runDiff()
	},
}

func runDiff() {
	println("Creating dry-run diffs....")
}

/* func initKube() {
	var kubeconfig *string
	if home := homedir.HomeDir(); home != "" {
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	} else {
		kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	}
	flag.Parse()

	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		panic(err)
	}
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err)
	}

} */
