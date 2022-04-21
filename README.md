# workload-generator
Generate Random Workloads For Kubernetes Tool

## How To Use

### Help

```bash
workload-generator --help
```

### Generate Random Workloads

Generate random kubernetes workloads:
```bash 
workload-generator generate --kubeconfig <kubeconfig path> --namespace <namespace>
```

Control the different repos to deploy, and whether to deploy all their tags by using `--count` and `--all-tags`
```bash 
workload-generator generate --kubeconfig <kubeconfig path> --namespace <namespace> --count 3 --all-tags
```
This command will deploy workloads with all tags of 3 different repos

### Delete Generated Random Workloads
Delete the generated workloads created by `workload-generator`
```bash 
workload-generator delete --kubeconfig <kubeconfig path> --namespace
```