{{/*
Expand the name of the chart.
*/}}
{{- define "scylla-operator.name" -}}
{{- default .Chart.Name .Values.nameOverride | trunc 63 | trimSuffix "-" }}
{{- end }}

{{/*
Create a default fully qualified app name.
We truncate at 63 chars because some Kubernetes name fields are limited to this (by the DNS naming spec).
If release name contains chart name it will be used as a full name.
*/}}
{{- define "scylla-operator.fullname" -}}
{{- if .Values.fullnameOverride }}
{{- .Values.fullnameOverride | trunc 63 | trimSuffix "-" }}
{{- else }}
{{- $name := default .Chart.Name .Values.nameOverride }}
{{- if contains $name .Release.Name }}
{{- .Release.Name | trunc 63 | trimSuffix "-" }}
{{- else }}
{{- printf "%s-%s" .Release.Name $name | trunc 63 | trimSuffix "-" }}
{{- end }}
{{- end }}
{{- end }}

{{/*
Create chart name and version as used by the chart label.
*/}}
{{- define "scylla-operator.chart" -}}
{{- printf "%s-%s" .Chart.Name .Chart.Version | replace "+" "_" | trunc 63 | trimSuffix "-" }}
{{- end }}

{{/*
Selector labels
*/}}
{{- define "scylla-operator.selectorLabels" -}}
app.kubernetes.io/name: scylla-operator
app.kubernetes.io/instance: scylla-operator
{{- end }}

{{/*
Common labels
*/}}
{{- define "scylla-operator.labels" -}}
{{ include "scylla-operator.selectorLabels" . }}
{{- end }}

{{/*
Create the name of the service account to use
*/}}
{{- define "scylla-operator.serviceAccountName" -}}
{{- if .Values.serviceAccount.create }}
{{- default (include "scylla-operator.fullname" .) .Values.serviceAccount.name }}
{{- else }}
{{- default "default" .Values.serviceAccount.name }}
{{- end }}
{{- end }}

{{/*
Name of the cluster.
*/}}
{{- define "scylla.name" -}}
{{- .name | trunc 63 | trimSuffix "-" }}
{{- end }}

{{/*
Namespace of the cluster.
*/}}
{{- define "scylla.namespace" -}}
{{- .namespace | trunc 63 | trimSuffix "-" }}
{{- end }}

{{/*
Common cluster labels
*/}}
{{- define "scylla.labels" -}}
{{ include "scylla.selectorLabels" . }}
{{- end }}

{{/*
Selector cluster labels
*/}}
{{- define "scylla.selectorLabels" -}}
app.kubernetes.io/name: scylla
app.kubernetes.io/instance: {{ include "scylla.name" . }}
{{- end }}

{{/*
Cluster service account to use
*/}}
{{- define "scylla.serviceAccountName" -}}
{{- printf "%s-%s" (include "scylla.name" .) "member" }}
{{- end }}
