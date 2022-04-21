package deployer

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	// import auth plugin package to support gke
	_ "k8s.io/client-go/plugin/pkg/client/auth"
)

const (
	appLabel = "workload-generator"
)

type WorkloadsDeployer struct {
	WorkloadType       string
	WorkloadNamePrefix string
	Namespace          string
	k8sClient          *kubernetes.Clientset
}

func NewWorkloadsDeployer(workloadType string, workloadNamePrefix string, kubeConfigPath string, namespace string) (*WorkloadsDeployer, error) {
	k8sClient, err := GetKubernetesClient(kubeConfigPath)
	if err != nil {
		return nil, err
	}

	return &WorkloadsDeployer{
		WorkloadType:       workloadType,
		WorkloadNamePrefix: workloadNamePrefix,
		Namespace:          namespace,
		k8sClient:          k8sClient,
	}, nil
}

func GetKubernetesClient(kubeConfigPath string) (*kubernetes.Clientset, error) {
	config, err := clientcmd.BuildConfigFromFlags("", kubeConfigPath)
	if err != nil {
		return nil, err
	}
	k8sClientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		return nil, err
	}

	return k8sClientset, nil
}

func (deployer *WorkloadsDeployer) DeployWorkload(imageFullTag string) error {
	switch deployer.WorkloadType {
	case "Pod":
		return deployer.deployPodWorkload(imageFullTag, deployer.WorkloadNamePrefix)
	}
	return nil
}

func (deployer *WorkloadsDeployer) deployPodWorkload(imageFullTag, workloadNamePrefix string) error {
	workloadName := fmt.Sprintf("%s-%s", workloadNamePrefix, uuid.New().String()[:6])
	pod := deployer.createPodResource(imageFullTag, workloadName)
	podCreated, err := deployer.k8sClient.CoreV1().Pods(deployer.Namespace).Create(context.TODO(), pod, metav1.CreateOptions{})
	if err != nil {
		return err
	}

	logrus.Infof("pod %s created", podCreated.Name)
	return nil
}

func (deployer *WorkloadsDeployer) createPodResource(imageFullTag, workloadName string) *v1.Pod {
	k8sPod := &v1.Pod{
		TypeMeta: metav1.TypeMeta{
			Kind: "Pod",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      workloadName,
			Namespace: deployer.Namespace,
			Labels:    map[string]string{"app": appLabel},
		},
		Spec: v1.PodSpec{
			Containers: []v1.Container{
				{Name: workloadName, Image: imageFullTag, Command: []string{"tail"}, Args: []string{"-f", "/dev/null"}},
			},
		},
	}

	return k8sPod
}

func (deployer *WorkloadsDeployer) DeleteWorkloads() error {
	switch deployer.WorkloadType {
	case "Pod":
		return deployer.deletePodWorkloads()
	}
	return nil
}

func (deployer *WorkloadsDeployer) deletePodWorkloads() error {
	label := fmt.Sprintf("app=%s", appLabel)
	return deployer.k8sClient.CoreV1().Pods(deployer.Namespace).DeleteCollection(context.Background(), metav1.DeleteOptions{}, metav1.ListOptions{LabelSelector: label})
}
