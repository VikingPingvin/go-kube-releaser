package cmd

import (
	"context"
	"flag"
	"fmt"
	"path/filepath"

	"github.com/spf13/cobra"
	appsv1 "k8s.io/api/apps/v1"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
)

//var clientSet;

var namespace = "appservices-gateway-sandbox"
var kubecontext = "dev-appservices-gateway-deployer"
var clientset *kubernetes.Clientset

func init() {
	rootCmd.AddCommand(kubePlan)
	initKube()
}

var kubePlan = &cobra.Command{
	Use:   "plan",
	Short: `Emulates "kubectl apply --dry-run" and creates diff files`,
	Run: func(cmd *cobra.Command, args []string) {
		runPlan()
	},
}

func runPlan() {
	println("Creating dry-run diffs....")
	testDeployments(clientset)
	testDryRunClient(clientset)
}

func initKube() {
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
	cs, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err)
	}
	clientset = cs

}

func testDeployments(clientset *kubernetes.Clientset) {
	ctx := context.Background()

	list, err := clientset.AppsV1().Deployments(namespace).List(ctx, metav1.ListOptions{})
	if err != nil {
		fmt.Println(err)
	} else {
		for _, item := range list.Items {
			fmt.Printf("%+v\n", item)
		}
	}
}

func testDryRunClient(clientset *kubernetes.Clientset) {
	dryRunOpts := []string{"All"}
	//deployment := v1.
	deploy, err := clientset.AppsV1().Deployments(namespace).Create(context.Background(), testDeployManifest, metav1.CreateOptions{DryRun: dryRunOpts})
	if err != nil {
		println(err)
	}

	println(deploy)
}

func int32Ptr(i int32) *int32 { return &i }

var testDeployManifest = &appsv1.Deployment{
	ObjectMeta: metav1.ObjectMeta{
		Name: "go-kube-releaser-test-deploy",
	},
	Spec: appsv1.DeploymentSpec{
		Replicas: int32Ptr(2),
		Selector: &metav1.LabelSelector{
			MatchLabels: map[string]string{
				"app": "demo",
			},
		},
		Template: v1.PodTemplateSpec{
			ObjectMeta: metav1.ObjectMeta{
				Labels: map[string]string{
					"app": "demo",
				},
			},
			Spec: v1.PodSpec{
				Containers: []v1.Container{
					{
						Name:  "go-kube-releaser-test-deploy",
						Image: "alpine:3.16",
					},
				},
			},
		},
	},
}

/* `apiVersion: apps/v1
kind: Deployment
metadata:
  labels:
  	applied-by: go-kube-releaser
  name: go-kube-releaser-test-deploy
  namespace: appservices-gateway-sandbox
spec:
  replicas: 1
  selector:
    matchLabels:
		applied-by: go-kube-releaser
  template:
    metadata:
      labels:
        applied-by: go-kube-releaser
    spec:
      containers:
        image: alpine
        imagePullPolicy: Always
        name: go-kube-releaser-test-deploy
		` */
