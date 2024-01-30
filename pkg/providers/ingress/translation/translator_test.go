// Licensed to the Apache Software Foundation (ASF) under one or more
// contributor license agreements.  See the NOTICE file distributed with
// this work for additional information regarding copyright ownership.
// The ASF licenses this file to You under the Apache License, Version 2.0
// (the "License"); you may not use this file except in compliance with
// the License.  You may obtain a copy of the License at
//
//	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
package translation

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/informers"
	"k8s.io/client-go/kubernetes/fake"
	"k8s.io/client-go/tools/cache"

	"github.com/apache/apisix-ingress-controller/pkg/providers/apisix/translation"
)

const _ingressKey string = "kubernetes.io/ingress.class"

func TestTranslateIngressTLS(t *testing.T) {
	secretName := "secret"
	cert := `
	-----BEGIN CERTIFICATE-----
	MIIEojCCAwqgAwIBAgIJAK253pMhgCkxMA0GCSqGSIb3DQEBCwUAMFYxCzAJBgNV
	BAYTAkNOMRIwEAYDVQQIDAlHdWFuZ0RvbmcxDzANBgNVBAcMBlpodUhhaTEPMA0G
	A1UECgwGaXJlc3R5MREwDwYDVQQDDAh0ZXN0LmNvbTAgFw0xOTA2MjQyMjE4MDVa
	GA8yMTE5MDUzMTIyMTgwNVowVjELMAkGA1UEBhMCQ04xEjAQBgNVBAgMCUd1YW5n
	RG9uZzEPMA0GA1UEBwwGWmh1SGFpMQ8wDQYDVQQKDAZpcmVzdHkxETAPBgNVBAMM
	CHRlc3QuY29tMIIBojANBgkqhkiG9w0BAQEFAAOCAY8AMIIBigKCAYEAyCM0rqJe
	cvgnCfOw4fATotPwk5Ba0gC2YvIrO+gSbQkyxXF5jhZB3W6BkWUWR4oNFLLSqcVb
	VDPitz/Mt46Mo8amuS6zTbQetGnBARzPLtmVhJfoeLj0efMiOepOSZflj9Ob4yKR
	2bGdEFOdHPjm+4ggXU9jMKeLqdVvxll/JiVFBW5smPtW1Oc/BV5terhscJdOgmRr
	abf9xiIis9/qVYfyGn52u9452V0owUuwP7nZ01jt6iMWEGeQU6mwPENgvj1olji2
	WjdG2UwpUVp3jp3l7j1ekQ6mI0F7yI+LeHzfUwiyVt1TmtMWn1ztk6FfLRqwJWR/
	Evm95vnfS3Le4S2ky3XAgn2UnCMyej3wDN6qHR1onpRVeXhrBajbCRDRBMwaNw/1
	/3Uvza8QKK10PzQR6OcQ0xo9psMkd9j9ts/dTuo2fzaqpIfyUbPST4GdqNG9NyIh
	/B9g26/0EWcjyO7mYVkaycrtLMaXm1u9jyRmcQQI1cGrGwyXbrieNp63AgMBAAGj
	cTBvMB0GA1UdDgQWBBSZtSvV8mBwl0bpkvFtgyiOUUcbszAfBgNVHSMEGDAWgBSZ
	tSvV8mBwl0bpkvFtgyiOUUcbszAMBgNVHRMEBTADAQH/MB8GA1UdEQQYMBaCCHRl
	c3QuY29tggoqLnRlc3QuY29tMA0GCSqGSIb3DQEBCwUAA4IBgQAHGEul/x7ViVgC
	tC8CbXEslYEkj1XVr2Y4hXZXAXKd3W7V3TC8rqWWBbr6L/tsSVFt126V5WyRmOaY
	1A5pju8VhnkhYxYfZALQxJN2tZPFVeME9iGJ9BE1wPtpMgITX8Rt9kbNlENfAgOl
	PYzrUZN1YUQjX+X8t8/1VkSmyZysr6ngJ46/M8F16gfYXc9zFj846Z9VST0zCKob
	rJs3GtHOkS9zGGldqKKCj+Awl0jvTstI4qtS1ED92tcnJh5j/SSXCAB5FgnpKZWy
	hme45nBQj86rJ8FhN+/aQ9H9/2Ib6Q4wbpaIvf4lQdLUEcWAeZGW6Rk0JURwEog1
	7/mMgkapDglgeFx9f/XztSTrkHTaX4Obr+nYrZ2V4KOB4llZnK5GeNjDrOOJDk2y
	IJFgBOZJWyS93dQfuKEj42hA79MuX64lMSCVQSjX+ipR289GQZqFrIhiJxLyA+Ve
	U/OOcSRr39Kuis/JJ+DkgHYa/PWHZhnJQBxcqXXk1bJGw9BNbhM=
	-----END CERTIFICATE-----
	
	`
	key := `
	-----BEGIN RSA PRIVATE KEY-----
	MIIG5AIBAAKCAYEAyCM0rqJecvgnCfOw4fATotPwk5Ba0gC2YvIrO+gSbQkyxXF5
	jhZB3W6BkWUWR4oNFLLSqcVbVDPitz/Mt46Mo8amuS6zTbQetGnBARzPLtmVhJfo
	eLj0efMiOepOSZflj9Ob4yKR2bGdEFOdHPjm+4ggXU9jMKeLqdVvxll/JiVFBW5s
	mPtW1Oc/BV5terhscJdOgmRrabf9xiIis9/qVYfyGn52u9452V0owUuwP7nZ01jt
	6iMWEGeQU6mwPENgvj1olji2WjdG2UwpUVp3jp3l7j1ekQ6mI0F7yI+LeHzfUwiy
	Vt1TmtMWn1ztk6FfLRqwJWR/Evm95vnfS3Le4S2ky3XAgn2UnCMyej3wDN6qHR1o
	npRVeXhrBajbCRDRBMwaNw/1/3Uvza8QKK10PzQR6OcQ0xo9psMkd9j9ts/dTuo2
	fzaqpIfyUbPST4GdqNG9NyIh/B9g26/0EWcjyO7mYVkaycrtLMaXm1u9jyRmcQQI
	1cGrGwyXbrieNp63AgMBAAECggGBAJM8g0duoHmIYoAJzbmKe4ew0C5fZtFUQNmu
	O2xJITUiLT3ga4LCkRYsdBnY+nkK8PCnViAb10KtIT+bKipoLsNWI9Xcq4Cg4G3t
	11XQMgPPgxYXA6m8t+73ldhxrcKqgvI6xVZmWlKDPn+CY/Wqj5PA476B5wEmYbNC
	GIcd1FLl3E9Qm4g4b/sVXOHARF6iSvTR+6ol4nfWKlaXSlx2gNkHuG8RVpyDsp9c
	z9zUqAdZ3QyFQhKcWWEcL6u9DLBpB/gUjyB3qWhDMe7jcCBZR1ALyRyEjmDwZzv2
	jlv8qlLFfn9R29UI0pbuL1eRAz97scFOFme1s9oSU9a12YHfEd2wJOM9bqiKju8y
	DZzePhEYuTZ8qxwiPJGy7XvRYTGHAs8+iDlG4vVpA0qD++1FTpv06cg/fOdnwshE
	OJlEC0ozMvnM2rZ2oYejdG3aAnUHmSNa5tkJwXnmj/EMw1TEXf+H6+xknAkw05nh
	zsxXrbuFUe7VRfgB5ElMA/V4NsScgQKBwQDmMRtnS32UZjw4A8DsHOKFzugfWzJ8
	Gc+3sTgs+4dNIAvo0sjibQ3xl01h0BB2Pr1KtkgBYB8LJW/FuYdCRS/KlXH7PHgX
	84gYWImhNhcNOL3coO8NXvd6+m+a/Z7xghbQtaraui6cDWPiCNd/sdLMZQ/7LopM
	RbM32nrgBKMOJpMok1Z6zsPzT83SjkcSxjVzgULNYEp03uf1PWmHuvjO1yELwX9/
	goACViF+jst12RUEiEQIYwr4y637GQBy+9cCgcEA3pN9W5OjSPDVsTcVERig8++O
	BFURiUa7nXRHzKp2wT6jlMVcu8Pb2fjclxRyaMGYKZBRuXDlc/RNO3uTytGYNdC2
	IptU5N4M7iZHXj190xtDxRnYQWWo/PR6EcJj3f/tc3Itm1rX0JfuI3JzJQgDb9Z2
	s/9/ub8RRvmQV9LM/utgyOwNdf5dyVoPcTY2739X4ZzXNH+CybfNa+LWpiJIVEs2
	txXbgZrhmlaWzwA525nZ0UlKdfktdcXeqke9eBghAoHARVTHFy6CjV7ZhlmDEtqE
	U58FBOS36O7xRDdpXwsHLnCXhbFu9du41mom0W4UdzjgVI9gUqG71+SXrKr7lTc3
	dMHcSbplxXkBJawND/Q1rzLG5JvIRHO1AGJLmRgIdl8jNgtxgV2QSkoyKlNVbM2H
	Wy6ZSKM03lIj74+rcKuU3N87dX4jDuwV0sPXjzJxL7NpR/fHwgndgyPcI14y2cGz
	zMC44EyQdTw+B/YfMnoZx83xaaMNMqV6GYNnTHi0TO2TAoHBAKmdrh9WkE2qsr59
	IoHHygh7Wzez+Ewr6hfgoEK4+QzlBlX+XV/9rxIaE0jS3Sk1txadk5oFDebimuSk
	lQkv1pXUOqh+xSAwk5v88dBAfh2dnnSa8HFN3oz+ZfQYtnBcc4DR1y2X+fVNgr3i
	nxruU2gsAIPFRnmvwKPc1YIH9A6kIzqaoNt1f9VM243D6fNzkO4uztWEApBkkJgR
	4s/yOjp6ovS9JG1NMXWjXQPcwTq3sQVLnAHxZRJmOvx69UmK4QKBwFYXXjeXiU3d
	bcrPfe6qNGjfzK+BkhWznuFUMbuxyZWDYQD5yb6ukUosrj7pmZv3BxKcKCvmONU+
	CHgIXB+hG+R9S2mCcH1qBQoP/RSm+TUzS/Bl2UeuhnFZh2jSZQy3OwryUi6nhF0u
	LDzMI/6aO1ggsI23Ri0Y9ZtqVKczTkxzdQKR9xvoNBUufjimRlS80sJCEB3Qm20S
	wzarryret/7GFW1/3cz+hTj9/d45i25zArr3Pocfpur5mfz3fJO8jg==
	-----END RSA PRIVATE KEY-----
	
	`
	secretNamespace := "test"
	//create the secret which will be used
	secret := &corev1.Secret{
		TypeMeta: metav1.TypeMeta{},
		ObjectMeta: metav1.ObjectMeta{
			Name:      secretName,
			Namespace: secretNamespace,
		},
		Data: map[string][]byte{
			"cert": []byte(cert),
			"key":  []byte(key),
		},
	}
	host := "test.com" //sni and host
	ingName := "apisix"
	//create secret lister
	client := fake.NewSimpleClientset()
	informersFactory := informers.NewSharedInformerFactory(client, 0)
	secretInformer := informersFactory.Core().V1().Secrets().Informer()
	secretLister := informersFactory.Core().V1().Secrets().Lister()

	processCh := make(chan struct{})
	secretInformer.AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc: func(obj interface{}) {
			processCh <- struct{}{}
		},
	})

	stopCh := make(chan struct{})
	defer close(stopCh)
	go secretInformer.Run(stopCh)
	cache.WaitForCacheSync(stopCh, secretInformer.HasSynced)

	_, err := client.CoreV1().Secrets(secretNamespace).Create(context.Background(), secret, metav1.CreateOptions{})
	assert.Nil(t, err)
	tr := &translator{
		ApisixTranslator: translation.NewApisixTranslator(&translation.TranslatorOptions{
			SecretLister: secretLister,
		}, translator{}),
	}
	<-processCh
	ssl, err := tr.TranslateIngressTLS(secretNamespace, ingName, secretName, []string{host})
	assert.Nil(t, err)

	assert.Equal(t, ssl.Cert, cert)
	assert.Equal(t, ssl.Key, key)
	assert.Equal(t, ssl.Snis[0], host)
}
