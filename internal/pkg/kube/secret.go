package kube

import (
	"golang.org/x/net/context"
	cv1 "k8s.io/api/core/v1"
	k8serrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const AuthJson string = `
{
	"auths":{
		"aliyun-fat-registry.cn-hongkong.cr.aliyuncs.com":{
			"username":"it.partner@yallatech.ae",
			"password":"zGUycFPXiH6YYw2k",
			"auth":"aXQucGFydG5lckB5YWxsYXRlY2guYWU6ekdVeWNGUFhpSDZZWXcyaw=="
		}
	}
}
`

// CreateOrUpdateSecret 创建或更新密钥
// kubectl create secret docker-registry SECRET_KEY_NAME --docker-server=REGISTRY_URL --docker-username=USER_NAME --docker-password=PASSWORD --namespace=NAMESPACE
func (client *Client) CreateOrUpdateSecret(namespace *string) (secret *cv1.Secret, err error) {

	registrySecretName := REGISTRY_SECRET_NAME
	if _, err := client.getOrCreateNamespace(namespace); err != nil {
		return nil, err
	}

	secret, err = client.clientSet.CoreV1().Secrets(*namespace).Get(context.TODO(), registrySecretName, metav1.GetOptions{})
	if err != nil {
		if !k8serrors.IsNotFound(err) {
			return nil, err
		}

		_, err = client.clientSet.CoreV1().Secrets(*namespace).Create(context.TODO(), &cv1.Secret{
			ObjectMeta: metav1.ObjectMeta{
				Namespace: *namespace,
				Name:      registrySecretName,
			},
			Type: cv1.SecretTypeDockercfg,
			Data: map[string][]byte{
				".dockerconfigjson": []byte(AuthJson),
			},
		},
			metav1.CreateOptions{})
	} else {
		if string(secret.Data[".dockerconfigjson"]) == AuthJson {
			return secret, nil
		}
		secret.Data = map[string][]byte{".dockerconfigjson": []byte(AuthJson)}
		_, err = client.clientSet.CoreV1().Secrets(*namespace).Update(context.TODO(), secret, metav1.UpdateOptions{})
	}
	return
}
