package main

import (
	"log"
	"os"
	"text/template"
)

type Pod struct {
	Name      string
	Namespace string
	Container Container
}

type Container struct {
	Name     string
	Image    string
	Resource Resource
}

type Resource struct {
	Memory string
}

func main() {
	podConfig := preparePodsConfig()

	pod, err := template.ParseFiles("pods.yml")
	if err != nil {
		log.Fatalln("Error parsing base pods", err)
	}

	err = pod.ExecuteTemplate(os.Stdout, "pods.yml", podConfig)
	if err != nil {
		log.Fatalln("error executing base template", err)
	}

}

func preparePodsConfig() Pod {
	r := Resource{Memory: "200Mi"}
	c := Container{Name: "memory-demo", Image: "polinux/stress", Resource: r}
	p := Pod{Name: "memory-demo-pod", Namespace: "KubeCf", Container: c}
	return p
}
