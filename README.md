# Sedna-modified

## brief introduction

This project is based on Sedna(https://github.com/kubeedge/sedna).

We added the "jointmultiedgeservice" section to complete our functionality.

## Feature

We extended the "jointinferenceservice" section to implement the following features:

- Only cloudWorker or edgeWorker can be deployed separately
- Multiple Edgeworkers can be deployed at once
- Add the file field for file mounting
- Add the log_level field to match logs
- Mounts the ~/.kube/config of the cluster so that pod can obtain cluster information
- add ServiceConfig to use nodePort mode for communication



## Quick Start

We assume that you have finished the k8s and kubeedge installation

- git clone

  ```sh
  git clone https://github.com/AdaYangOlzz/sedna-modified.git
  ```

- install sedna

  ```sh
  curl https://raw.githubusercontent.com/kubeedge/sedna/main/scripts/installation/install.sh
  ```

  modify the `TMP_DIR`， `SEDNA_VERSION`  and image(`adayoung/sedna-gm:v0.3.12` and `adayoung/sedna-lc:v0.3.12`)

  ```sh
  SEDNA_ACTION=create bash install.sh
  ```

- yaml example([sedna-modified/build/crd-samples/sedna/jointmultiedgeservice_v1alpha1.yaml at master · AdaYangOlzz/sedna-modified (github.com)](https://github.com/AdaYangOlzz/sedna-modified/blob/master/build/crd-samples/sedna/jointmultiedgeservice_v1alpha1.yaml))

  ```sh
  kubectl apply -f <yaml-name>
  ```

- get infomation

  ```sh
  kubectl get pod -n <namespace-name>
  kubectl get deploy -n <namespace-name>
  kubectl get svc -n <namespace-name>
  kubectl get mulji -n <namespace-name>
  ```

- how to build gm/lc image
  ```sh
  make WHAT=gm
  make gmiamge
  make WHAT=lc
  docker buildx build --platform linux/arm64,linux/amd64 --build-arg GO_LDFLAGS="" -t repo:5000/adayoung/sedna-lc:v0.3.12 -f build/lc/Dockerfile . --push
  ```