
package main

import (
	"fmt"
	"github.com/dbcdk/k8s-ingress/kubeclient"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	// needed to enable oidc authentication
	_ "k8s.io/client-go/plugin/pkg/client/auth/oidc"
)


func fail(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	conf, err := kubeclient.GetKubeConfig("")
	fail(err)

	clients, err := kubeclient.GetKubeClient(conf)
	fail(err)

	ic := clients.NetworkingV1beta1().Ingresses("")
	ings, err := ic.List(v1.ListOptions{})
	fail(err)

	for _, i := range ings.Items {
		fmt.Printf("%s.%s", i.Namespace, i.Name)
	}
}
