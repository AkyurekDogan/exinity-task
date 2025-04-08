# infra/modules/ingress/main.tf

# Deploy NGINX Ingress Controller in the EKS cluster
resource "helm_release" "nginx_ingress" {
  name       = "nginx-ingress"
  namespace  = "kube-system"
  repository = "https://kubernetes.github.io/ingress-nginx"
  chart      = "ingress-nginx"
  version    = "4.0.13" # Choose the latest stable version

  values = [
    <<EOF
controller:
  replicaCount: 2
  ingressClass: nginx
  service:
    externalTrafficPolicy: Local
EOF
  ]
}

# Create an Ingress Resource for the API
resource "kubernetes_ingress" "api_ingress" {
  metadata {
    name      = "api-ingress"
    namespace = "default"
    annotations = {
      "nginx.ingress.kubernetes.io/rewrite-target" = "/"
    }
  }

  spec {
    rule {
      host = "www.exinity-task.com"

      http {
        path {
          path    = "/"
          backend {
            service {
              name = "api-service"
              port {
                number = 80
              }
            }
          }
        }
      }
    }
  }
}
