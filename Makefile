build:
	go build -o dist/tufin main.go

helm-install:
	helm install my-wordpress-sql ./wordpress-sql-chart --kubeconfig=/etc/rancher/k3s/k3s.yaml

helm-upgrade:
	helm upgrade my-wordpress-sql ./wordpress-sql-chart --kubeconfig=/etc/rancher/k3s/k3s.yaml

helm-uninstall:
	helm uninstall my-wordpress-sql --kubeconfig=/etc/rancher/k3s/k3s.yaml

helm-list:
	helm list --kubeconfig=/etc/rancher/k3s/k3s.yaml