package main

import (
	"errors"
	"fmt"
	"sync"
	"time"

	batchv1 "k8s.io/api/batch/v1"
	"k8s.io/client-go/informers"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/tools/clientcmd"
)

// watch for delete operation
func main() {
	ns := "{your-namespace}"
	name := "{your-name}"

	config, err := clientcmd.BuildConfigFromFlags("")
	if err != nil {
		panic(err)
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	successCh := make(chan bool)
	defer close(successCh)

	sharedInformerfactory := informers.NewSharedInformerFactoryWithOptions(clientset, time.Second, informers.WithNamespace(ns))
	informer := sharedInformerfactory.Batch().V1().Jobs().Informer()
	informer.AddEventHandler(cache.ResourceEventHandlerFuncs{
		DeleteFunc: func(obj interface{}) {
			item := obj.(*batchv1.Job)
			if item.GetName() == name {
				successCh <- true
			}
		},
	})

	var wg sync.WaitGroup
	wg.Add(1)

	stopCh := make(chan struct{})
	go func() {
		informer.Run(stopCh)
		wg.Done()
	}()

	var res error
	timeout := make(chan bool)
	go func() {
		time.Sleep(time.Second * 30)
		timeout <- true
	}()

LOOP:
	for {
		select {
		case <-successCh:
			close(stopCh)
			break LOOP
		case <-timeout:
			res = errors.New("delete job failed: timeout")
			close(stopCh)
			break LOOP
		default:
			time.Sleep(time.Second * 5)
		}
	}

	wg.Wait()
	fmt.Println(res)
}
