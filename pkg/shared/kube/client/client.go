/*
Copyright 2021 The KodeRover Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package client

import (
	"fmt"

	"github.com/koderover/zadig/pkg/config"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/koderover/zadig/pkg/setting"
	aslanClient "github.com/koderover/zadig/pkg/shared/client/aslan"
	"github.com/koderover/zadig/pkg/tool/kube/multicluster"
)

func GetKubeClient(hubserverAddr, clusterID string) (client.Client, error) {
	if clusterID == setting.LocalClusterID || clusterID == "" {
		if clusterID == setting.LocalClusterID {
			clusterID = ""
		}

		return multicluster.GetKubeClient(hubserverAddr, clusterID)
	}
	cluster, err := aslanClient.New(config.AslanServiceAddress()).GetClusterInfo(clusterID)
	if err != nil {
		return nil, err
	}

	switch cluster.Type {
	case setting.AgentClusterType, "":
		return multicluster.GetKubeClient(hubserverAddr, clusterID)
	case setting.KubeConfigClusterType:
		return multicluster.GetKubeClientFromKubeConfig(clusterID, cluster.KubeConfig)
	default:
		return nil, fmt.Errorf("failed to create kubeclient: unknown cluster type: %s", cluster.Type)
	}
}

func GetKubeClientSet(hubServerAddr, clusterID string) (*kubernetes.Clientset, error) {
	if clusterID == setting.LocalClusterID || clusterID == "" {
		if clusterID == setting.LocalClusterID {
			clusterID = ""
		}

		return multicluster.GetKubeClientSet(hubServerAddr, clusterID)
	}
	cluster, err := aslanClient.New(config.AslanServiceAddress()).GetClusterInfo(clusterID)
	if err != nil {
		return nil, err
	}
	switch cluster.Type {
	case setting.AgentClusterType, "":
		return multicluster.GetKubeClientSet(hubServerAddr, clusterID)
	case setting.KubeConfigClusterType:
		return multicluster.GetKubeClientSetFromKubeConfig(clusterID, cluster.KubeConfig)
	default:
		return nil, fmt.Errorf("failed to create kubeclient: unknown cluster type: %s", cluster.Type)
	}
}

func GetDynamicKubeClient(hubserverAddr, clusterID string) (dynamic.Interface, error) {
	if clusterID == setting.LocalClusterID || clusterID == "" {
		if clusterID == setting.LocalClusterID {
			clusterID = ""
		}

		return multicluster.GetDynamicKubeclient(hubserverAddr, clusterID)
	}
	cluster, err := aslanClient.New(config.AslanServiceAddress()).GetClusterInfo(clusterID)
	if err != nil {
		return nil, err
	}
	switch cluster.Type {
	case setting.AgentClusterType, "":
		return multicluster.GetDynamicKubeclient(hubserverAddr, clusterID)
	case setting.KubeConfigClusterType:
		return multicluster.GetDynamicKubeclientFromKubeConfig(clusterID, cluster.KubeConfig)
	default:
		return nil, fmt.Errorf("failed to create kubeclient: unknown cluster type: %s", cluster.Type)
	}
}

func GetKubeAPIReader(hubServerAddr, clusterID string) (client.Reader, error) {
	if clusterID == setting.LocalClusterID || clusterID == "" {
		if clusterID == setting.LocalClusterID {
			clusterID = ""
		}
		return multicluster.GetKubeAPIReader(hubServerAddr, clusterID)
	}
	cluster, err := aslanClient.New(config.AslanServiceAddress()).GetClusterInfo(clusterID)
	if err != nil {
		return nil, err
	}

	switch cluster.Type {
	case setting.AgentClusterType, "":
		return multicluster.GetKubeAPIReader(hubServerAddr, clusterID)
	case setting.KubeConfigClusterType:
		return multicluster.GetKubeClientFromKubeConfig(clusterID, cluster.KubeConfig)
	default:
		return nil, fmt.Errorf("failed to create kubeclient: unknown cluster type: %s", cluster.Type)
	}
}

func GetRESTConfig(hubServerAddr, clusterID string) (*rest.Config, error) {
	if clusterID == setting.LocalClusterID || clusterID == "" {
		if clusterID == setting.LocalClusterID {
			clusterID = ""
		}
		return multicluster.GetRESTConfig(hubServerAddr, clusterID)
	}
	cluster, err := aslanClient.New(config.AslanServiceAddress()).GetClusterInfo(clusterID)
	if err != nil {
		return nil, err
	}

	switch cluster.Type {
	case setting.AgentClusterType, "":
		return multicluster.GetRESTConfig(hubServerAddr, clusterID)
	case setting.KubeConfigClusterType:
		return multicluster.GetRestConfigFromKubeConfig(clusterID, cluster.KubeConfig)
	default:
		return nil, fmt.Errorf("failed to create kubeclient: unknown cluster type: %s", cluster.Type)
	}
}

func GetClientset(hubServerAddr, clusterID string) (kubernetes.Interface, error) {
	if clusterID == setting.LocalClusterID || clusterID == "" {
		if clusterID == setting.LocalClusterID {
			clusterID = ""
		}
		return multicluster.GetClientset(hubServerAddr, clusterID)
	}
	cluster, err := aslanClient.New(config.AslanServiceAddress()).GetClusterInfo(clusterID)
	if err != nil {
		return nil, err
	}

	switch cluster.Type {
	case setting.AgentClusterType, "":
		return multicluster.GetClientset(hubServerAddr, clusterID)
	case setting.KubeConfigClusterType:
		return multicluster.GetClientSetFromKubeConfig(clusterID, cluster.KubeConfig)
	default:
		return nil, fmt.Errorf("failed to create kubeclient: unknown cluster type: %s", cluster.Type)
	}
}
