package main

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"net/url"
	"time"

	clientset "github.com/openshift/client-go/config/clientset/versioned"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	restclient "k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

func kubernetesClient(ip string, kubeconfigFilePath string) (*clientset.Clientset, error) {
	config, err := kubernetesClientConfiguration(ip, kubeconfigFilePath)
	if err != nil {
		return nil, err
	}
	return clientset.NewForConfig(config)
}

func kubernetesClientConfiguration(ip string, kubeconfigFilePath string) (*restclient.Config, error) {
	config, err := clientcmd.BuildConfigFromFlags("", kubeconfigFilePath)
	if err != nil {
		return nil, err
	}
	// override dial to directly use the IP of the VM
	config.Dial = func(ctx context.Context, network, address string) (net.Conn, error) {
		var d net.Dialer
		return d.DialContext(ctx, "tcp", fmt.Sprintf("%s:6443", ip))
	}
	// discard any proxy configuration of the host
	config.Proxy = func(request *http.Request) (*url.URL, error) {
		return nil, nil
	}
	config.Timeout = 5 * time.Second
	return config, nil
}

func listClusterOperators(ctx context.Context, ip string, kubeconfigFilePath string) error {
	client, err := kubernetesClient(ip, kubeconfigFilePath)
	if err != nil {
		return err
	}
	lister := client.ConfigV1().ClusterOperators()
	co, err := lister.List(ctx, metav1.ListOptions{})
	if err != nil {
		return err
	}

	for _, c := range co.Items {
		fmt.Printf("operator %s\n", c.ObjectMeta.Name)
	}
	return nil
}

func main() {
	err := listClusterOperators(context.Background(), "api.crc.testing", "kubeconfig")
	if err != nil {
		fmt.Println("Error listing cluster operators:", err)
	}
}
