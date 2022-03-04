package controllers_test

import (
	"context"

	. "github.com/kralicky/kmatch"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	cfgv1beta1 "github.com/rancher/opni-monitoring/pkg/config/v1beta1"
	"github.com/rancher/opni-monitoring/pkg/noauth"
	"github.com/rancher/opni-monitoring/pkg/sdk/api/v1beta1"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	rbacv1 "k8s.io/api/rbac/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var _ = Describe("Gateway Controller", func() {
	When("creating a gateway resource", func() {
		var gw *v1beta1.Gateway
		It("should succeed", func() {
			gw = &v1beta1.Gateway{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "test",
					Namespace: "default",
				},
				Spec: v1beta1.GatewaySpec{
					Auth: v1beta1.AuthSpec{
						Provider: cfgv1beta1.AuthProviderNoAuth,
						Noauth:   &noauth.ServerConfig{},
					},
				},
			}
			Expect(k8sClient.Create(context.Background(), gw)).To(Succeed())
			Eventually(Object(gw)).Should(Exist())
		})

		It("should create the gateway deployment", func() {
			Eventually(Object(&appsv1.Deployment{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "opni-gateway",
					Namespace: gw.Namespace,
				},
			})).Should(ExistAnd(
				HaveOwner(gw),
				HaveMatchingContainer(And(
					HaveImage("rancher/opni-gateway:latest"),
					HavePorts(
						"http",
						"management-grpc",
						"management-http",
						"management-web",
						"noauth",
					),
					HaveVolumeMounts(
						"config",
						"certs",
						"cortex-client-certs",
						"cortex-server-cacert",
					),
				)),
				HaveMatchingVolume(And(
					HaveName("config"),
					HaveVolumeSource("ConfigMap"),
				)),
				HaveMatchingVolume(And(
					HaveName("certs"),
					HaveVolumeSource("Secret"),
				)),
				HaveMatchingVolume(And(
					HaveName("cortex-client-certs"),
					HaveVolumeSource("Secret"),
				)),
				HaveMatchingVolume(And(
					HaveName("cortex-server-cacert"),
					HaveVolumeSource("Secret"),
				)),
			))
		})

		It("should create the gateway service", func() {
			Eventually(Object(&corev1.Service{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "opni-gateway",
					Namespace: gw.Namespace,
				},
			})).Should(ExistAnd(
				HaveOwner(gw),
				HavePorts(
					"http",
					"management-grpc",
					"management-http",
					"management-web",
				),
				HaveType(corev1.ServiceTypeClusterIP),
			))
		})
		It("should create the gateway configmap", func() {
			Eventually(Object(&corev1.ConfigMap{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "opni-gateway",
					Namespace: gw.Namespace,
				},
			})).Should(ExistAnd(
				HaveOwner(gw),
				HaveData("config.yaml", nil),
			))
		})

		It("should create gateway rbac", func() {
			Eventually(Object(&corev1.ServiceAccount{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "opni-monitoring",
					Namespace: gw.Namespace,
				},
			})).Should(ExistAnd(
				HaveOwner(gw),
			))
			Eventually(Object(&rbacv1.Role{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "opni-monitoring-crd",
					Namespace: gw.Namespace,
				},
			})).Should(ExistAnd(
				HaveOwner(gw),
			))
			Eventually(Object(&rbacv1.RoleBinding{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "opni-monitoring-crd",
					Namespace: gw.Namespace,
				},
			})).Should(ExistAnd(
				HaveOwner(gw),
			))
		})
	})
})
