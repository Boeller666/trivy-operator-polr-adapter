# Trivy Operator PolicyReport Adapter

Maps Trivy Operator CRDs into the unified PolicyReport and ClusterPolicyReport from the Kubernetes Policy Working Group. This makes it possible to use tooling like [Policy Reporter](https://github.com/kyverno/policy-reporter) for the different kinds of Trivy Reports.

## Pre Requirements

1. [Trivy Operator](https://github.com/aquasecurity/trivy-operator) with the related CRDs is installed and running
2. [PolicyReport CRDs](https://github.com/kubernetes-sigs/wg-policy-prototypes/tree/master/policy-report/crd/v1alpha2) are installed in your Cluster

## Installation via Helm

```bash
helm repo add trivy-operator-polr-adapter https://fjogeleit.github.io/trivy-operator-polr-adapter
helm install trivy-operator-polr-adapter trivy-operator-polr-adapter/trivy-operator-polr-adapter -n trivy-adapter --create-namespace
```

## Usage

Local usage with ConfigAuditReport and VulnerabilityReports mapping enabled.

```bash
./trivy-operator-polr-adapter run --kubeconfig ~/.kube/config --enable-config-audit --enable-vulnerability
```

## Configuration


| Argument                | Helm Value                             | Description                                                           | Default Helm Value |
|-------------------------|----------------------------------------|-----------------------------------------------------------------------|--------------------|
|--kubeconfig             |                                        | Path to the used kubeconfig, mainly for local development             |                    |
|--enable-vulnerability   |`adapters.vulnerabilityReports.enabled` | Enables the transformation of VulnerabilityReports into PolicyReports | `true`             |
|--enable-config-audit    |`adapters.configAuditReports.enabled`   | Enables the transformation of ConfigAuditReports into PolicyReports   | `true`             |
|--enable-rbac-assessment |`adapters.rbacAssessmentReports.enabled`| Enables the transformation of RbacAssessmentReport into PolicyReports and<br>ClusterRbacAssessmentReport into ClusterPolicyReports  | `false`             |
|--enable-exposed-secrets |`adapters.exposedSecretReports.enabled` | Enables the transformation of ExposedSecretReport into PolicyReports   | `false`             |
|--enable-compliance |`adapters.complianceReports.enabled` | Enables the transformation of ClusterComplianceDetailReport into ClusterPolicyReports<br>(Not implemented in Trivy Operator yet) | `false`             |
|--enable-kube-bench |`adapters.cisKubeBenchReports.enabled` | Enables the transformation of CISKubeBenchReports into ClusterPolicyReports<br>(Not available in newer version of Trivy Operator) | `false`             |

## Available Sources

Sources of the PolicyReportResults can be used to filter different Reports from metrics, views or notifications in Policy Reporter

| Source               | TrivyReport Report                                 |
|----------------------|----------------------------------------------------|
| Trivy ConfigAudit    | ConfigAuditReport                                  |
| Trivy Vulnerability  | VulnerabilityReport                                |
| Trivy ExposedSecrets | ExposedSecretReport                                |
| Trivy RbacAssessment | ClusterRbacAssessmentReport / RbacAssessmentReport |

## Integreted Adapters
### VulnerabilityReports

Maps VulnerabilityReports into PolicyReports with the relation 1:1. The PolicyReport is referenced with the scanned resource like the VulnerabilityReport itself.

```yaml
apiVersion: wgpolicyk8s.io/v1alpha2
kind: PolicyReport
metadata:
  labels:
    app.kubernetes.io/created-by: trivy-operator-polr-adapter
    trivy-operator.source: VulnerabilityReport
  name: trivy-vuln-polr-nginx-5fbc65fff
  namespace: test
  ownerReferences:
  - apiVersion: apps/v1
    blockOwnerDeletion: false
    controller: true
    kind: ReplicaSet
    name: nginx-5fbc65fff
    uid: 710f2142-7613-4cf5-aef7-dc65306626e2
  resourceVersion: "122118"
  uid: 2ea883ef-c060-4e80-ae34-3f9b527c02bc
results:
- category: Vulnerability Scan
  message: 'apt: integer overflows and underflows while parsing .deb packages'
  policy: CVE-2020-27350
  properties:
    artifact.repository: library/nginx
    artifact.tag: "1.17"
    fixedVersion: 1.8.2.2
    installedVersion: 1.8.2.1
    primaryLink: https://avd.aquasec.com/nvd/cve-2020-27350
    registry.server: index.docker.io
    resource: apt
    score: "5.7"
  resources:
  - apiVersion: apps/v1
    kind: ReplicaSet
    name: nginx-5fbc65fff
    namespace: test
    uid: 710f2142-7613-4cf5-aef7-dc65306626e2
  result: warn
  severity: medium
  source: Trivy Vulnerability
  timestamp:
    nanos: 0
    seconds: 1653395950
summary:
  error: 0
  fail: 109
  pass: 0
  skip: 1
  warn: 166
```

### ConfigAuditReports

Maps ConfigAuditReports into PolicyReports with the relation 1:1. The PolicyReport is referenced with the scanned resource like the ConfigAuditReport itself.

```yaml
apiVersion: wgpolicyk8s.io/v1alpha2
kind: PolicyReport
metadata:
  labels:
    app.kubernetes.io/created-by: trivy-operator-polr-adapter
    trivy-operator.source: ConfigAuditReport
  name: trivy-audit-polr-nginx-5fbc65fff
  namespace: test
  ownerReferences:
  - apiVersion: apps/v1
    blockOwnerDeletion: false
    controller: true
    kind: ReplicaSet
    name: nginx-5fbc65fff
    uid: 710f2142-7613-4cf5-aef7-dc65306626e2
results:
- category: Kubernetes Security Check
  message: Sysctls can disable security mechanisms or affect all containers on a host,
    and should be disallowed except for an allowed 'safe' subset. A sysctl is considered
    safe if it is namespaced in the container or the Pod, and it is isolated from
    other Pods or processes on the same Node.
  policy: Unsafe sysctl options set
  resources:
  - apiVersion: apps/v1
    kind: ReplicaSet
    name: nginx-5fbc65fff
    namespace: test
    uid: 710f2142-7613-4cf5-aef7-dc65306626e2
  result: pass
  rule: KSV026
  severity: medium
  source: Trivy ConfigAudit
  timestamp:
    nanos: 0
    seconds: 1653395950
summary:
  error: 0
  fail: 26
  pass: 42
  skip: 0
  warn: 0
```

### RbacAssessmentReport

Maps RbacAssessmentReport into PolicyReports.

```yaml
apiVersion: wgpolicyk8s.io/v1alpha2
kind: PolicyReport
metadata:
  labels:
    app.kubernetes.io/created-by: trivy-operator-polr-adapter
    trivy-operator.source: RbacAssessmentReport
  name: trivy-rbac-polr-role-57d656695f
  namespace: kyverno
  ownerReferences:
  - apiVersion: rbac.authorization.k8s.io/v1
    blockOwnerDeletion: false
    controller: true
    kind: Role
    name: kyverno:leaderelection
    uid: ea031ce4-9f63-4aa9-a68c-da42b523768d
results:
- category: Kubernetes Security Check
  message: Check whether role permits update/create of a malicious pod
  policy: Do not allow update/create of a malicious pod
  properties:
    1. message: Role permits create/update of a malicious pod
    resultID: 5d52ad869c9da5e8533ae31a62b8e5a8a2f1838f
  resources:
  - apiVersion: rbac.authorization.k8s.io/v1
    kind: Role
    name: kyverno:leaderelection
    namespace: kyverno
    uid: ea031ce4-9f63-4aa9-a68c-da42b523768d
  result: fail
  rule: KSV048
  severity: high
  source: Trivy RbacAssessment
  timestamp:
    nanos: 0
    seconds: 1661168982
- category: Kubernetes Security Check
  message: Check whether role permits allowing users in a rolebinding to add other
    users to their rolebindings
  policy: Do not allow users in a rolebinding to add other users to their rolebindings
  properties:
    resultID: 3de0c6a7f01df775fad425283b2cf56771e10902
  resources:
  - apiVersion: rbac.authorization.k8s.io/v1
    kind: Role
    name: kyverno:leaderelection
    namespace: kyverno
    uid: ea031ce4-9f63-4aa9-a68c-da42b523768d
  result: pass
  rule: KSV055
  severity: low
  source: Trivy RbacAssessment
  timestamp:
    nanos: 0
    seconds: 1661168982
summary:
  error: 0
  fail: 1
  pass: 1
  skip: 0
  warn: 0
```

### ClusterRbacAssessmentReport

Maps ClusterRbacAssessmentReport into PolicyReports.

```yaml
apiVersion: wgpolicyk8s.io/v1alpha2
kind: ClusterPolicyReport
metadata:
  creationTimestamp: "2022-08-22T18:15:30Z"
  generation: 2
  labels:
    app.kubernetes.io/created-by: trivy-operator-polr-adapter
    trivy-operator.source: ClusterRbacAssessmentReport
  name: trivy-rbac-cpolr-clusterrole-5585c7b9ff
  ownerReferences:
  - apiVersion: rbac.authorization.k8s.io/v1
    blockOwnerDeletion: false
    controller: true
    kind: ClusterRole
    name: system:certificates.k8s.io:kubelet-serving-approver
    uid: 21449ac8-2f58-4eff-8f3d-c9e4e0024821
  resourceVersion: "39436"
  uid: 2296a252-b108-4d4a-b705-4b8983babe2b
results:
- category: Kubernetes Security Check
  message: Some workloads leverage configmaps to store sensitive data or configuration
    parameters that affect runtime behavior that can be modified by an attacker or
    combined with another issue to potentially lead to compromise.
  policy: Do not allow management of configmaps
  properties:
    resultID: d06e66683ee5de1136d5996ae0f4e1ae9b5d85c7
  resources:
  - apiVersion: rbac.authorization.k8s.io/v1
    kind: ClusterRole
    name: system:certificates.k8s.io:kubelet-serving-approver
    uid: 21449ac8-2f58-4eff-8f3d-c9e4e0024821
  result: pass
  rule: KSV049
  severity: medium
  source: Trivy RbacAssessment
  timestamp:
    nanos: 0
    seconds: 1661165899
- category: Kubernetes Security Check
  message: Check whether role permits privilege escalation from node proxy
  policy: Do not allow privilege escalation from node proxy
  properties:
    resultID: 519454bf1ec35b55d0d8041fb191017bf83519d3
  resources:
  - apiVersion: rbac.authorization.k8s.io/v1
    kind: ClusterRole
    name: system:certificates.k8s.io:kubelet-serving-approver
    uid: 21449ac8-2f58-4eff-8f3d-c9e4e0024821
  result: pass
  rule: KSV047
  severity: high
  source: Trivy RbacAssessment
  timestamp:
    nanos: 0
    seconds: 1661165899
summary:
  error: 0
  fail: 0
  pass: 2
  skip: 0
  warn: 0

```

## Policy Reporter UI Screenshots

### VulnerabilityReports

![Policy Reporter UI - PolicyReport VulnerabilityReports Screenshot](https://github.com/fjogeleit/trivy-operator-polr-adapter/blob/main/screens/vulnr.png?raw=true)

### ConfigAuditReports

![Policy Reporter UI - PolicyReport ConfigAuditReports Screenshot](https://github.com/fjogeleit/trivy-operator-polr-adapter/blob/main/screens/config-audit.png?raw=true)

### CISKubeBenchReports

![Policy Reporter UI - PolicyReport CISKubeBenchReports Screenshot](https://github.com/fjogeleit/trivy-operator-polr-adapter/blob/main/screens/kube-bench.png?raw=true)
