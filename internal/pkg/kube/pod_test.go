package kube

import (
	"context"
	"github.com/stretchr/testify/assert"
	meta "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestGetPod(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}

	kubeconfigPath := filepath.Join(os.Getenv("HOME"), ".kube", "config")
	restcfg, err := clientcmd.BuildConfigFromFlags("", kubeconfigPath)
	assert.NoError(t, err)
	kubeClient, err := kubernetes.NewForConfig(restcfg)
	assert.NoError(t, err)

	client := &Client{
		clientSet: kubeClient,
		defaultApplyOptions: &meta.ApplyOptions{
			FieldManager: "application/apply-patch+yaml",
			Force:        true,
		},
	}

	const NodeSelectorLabel = "gotocloud-builder"
	const BuildIdSelectorLabel = "build"

	pods, err := client.GetPods(context.TODO(), "kube-system", "", "")
	assert.NoError(t, err)
	assert.NotEmpty(t, pods)

	pod := func() *PodDescription {
		for i, p := range pods {
			if p.Name == "kube-apiserver-docker-desktop" {
				return &pods[i]
			}
		}
		return nil
	}()
	containerName := "kube-apiserver"
	var tailLine int64 = 1024

	logBuilder := strings.Builder{}
	logBuf, err := client.GetPodLogs(context.TODO(), "kube-system", pod.Name, containerName, &tailLine, false)
	assert.NoError(t, err)
	assert.NotEmpty(t, logBuf)
	logBuilder.WriteString(string(logBuf) + "\n")

	//log, err := client.GetPodStreamLogs(context.TODO(), "kube-system", pod.Name, containerName, &tailLine, false, false)
	//assert.NoError(t, err)
	//defer log.Close()
	//buf := new(bytes.Buffer)
	//io.Copy(buf, log)
	//logBuilder.WriteString(buf.String())
	//
	//content := make([]byte, 1024)
	//for {
	//	n, err := log.Read(content)
	//	assert.NoError(t, err)
	//	msg := string(content[:n])
	//	assert.NotNil(t, msg)
	//	logBuilder.WriteString(msg + "\n")
	//}
}
