# trivy-extractor

Enriches trivy metric data with a namespace team mapping. This allows a Security team to easily identify what vulnerabilities belong to what team.

# Configuration

Simply add a new configmap to the k8s namespace and update the `trivyExtractor.namespacesTeamConfigMapName` value in the `vaules.yaml` file.


Example config map:

```
apiVersion: v1
kind: ConfigMap
metadata:
  name: namespaces
data: 
  namespaces.csv: |
    TEAM 1,app-1
    TEAM 2,app-2
    TEAM 3,app-3
```

