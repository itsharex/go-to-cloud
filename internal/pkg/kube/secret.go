package kube

import (
	"golang.org/x/net/context"
	core "k8s.io/api/core/v1"
	k8sErrors "k8s.io/apimachinery/pkg/api/errors"
	meta "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const AuthJson string = `
{
	"auths":{
		"registry":{
			"username":"username",
			"password":"password",
			"auth":"username:password | base64"
		}
	}
}
`

// CreateOrUpdateSecret 创建或更新密钥
// kubectl create secret docker-registry SECRET_KEY_NAME --docker-server=REGISTRY_URL --docker-username=USER_NAME --docker-password=PASSWORD --namespace=NAMESPACE
func (client *Client) CreateOrUpdateSecret(namespace *string) (secret *core.Secret, err error) {

	registrySecretName := RegistrySecretName
	client.getOrCreateNamespace(namespace)

	secret, err = client.clientSet.CoreV1().Secrets(*namespace).Get(context.TODO(), registrySecretName, meta.GetOptions{})
	if err != nil {
		if !k8sErrors.IsNotFound(err) {
			return nil, err
		}

		_, err = client.clientSet.CoreV1().Secrets(*namespace).Create(context.TODO(), &core.Secret{
			ObjectMeta: meta.ObjectMeta{
				Namespace: *namespace,
				Name:      registrySecretName,
			},
			Type: core.SecretTypeDockercfg,
			Data: map[string][]byte{
				".dockerconfigjson": []byte(AuthJson),
			},
		},
			meta.CreateOptions{})
	} else {
		if string(secret.Data[".dockerconfigjson"]) == AuthJson {
			return secret, nil
		}
		secret.Data = map[string][]byte{".dockerconfigjson": []byte(AuthJson)}
		_, err = client.clientSet.CoreV1().Secrets(*namespace).Update(context.TODO(), secret, meta.UpdateOptions{})
	}
	return
}
