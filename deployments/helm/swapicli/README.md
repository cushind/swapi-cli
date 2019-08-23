# swapicli

Helm chart source can be found at: https://github.com/bambash/helm-cronjobs

## Minikube
With your minikube cluster running, install tiller with:
```
helm init --history-max 200
```

This chart will run as a cronjob. Set the cron schedule in `values.yaml` (note UTC time is 6/7 hours ahead of mountain time)

Then install the chart with:
```
helm install --name swapi-cli deployments/helm/swapicli
```

## Getting the job output
Watch for new pods with `kubectl get pods -w` and upon the creation of the new pod get the logs with `kubectl logs <pod-name> -f`