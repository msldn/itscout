# This Repository

### This repostory contains an example go applicaiton along with the helm charts required for deploying such applicaiton. This code for learning and training porpuses and is not intended to be used in production.
---
## ITScout Application
ITScout Is configuraiton management database for storing details about configuraiton items in a specific it environment. It consists of CRUD API layer that can be used to manipulate CI details. It uses the mux go library and is using a postgres database in the background

## Combatability
While this repostiory can be extended with any kubernetes falvor, it has been specifically designed and tested to wrok with local docker desktop kubernets cluster. because of the use of local-path volumes, this is not expected to work out of the box on public cloud clusters.


## Docker compose
While not being essential, an alternative of starting the applicaiton on k8s cluster is to use docker compose. when started using docker-compose, application is exposed under http://localhost:8000
- `docker-compose up --build`


## Runiung on kubernetes
### Prerequesites
- On your workstation, make sure to have connection to k8s cluter - Docker esktop one and that you are able to execute kubectl commands.
- Kubectl and Helm clients to be installed on the workstation.
- Set up the /etc/hosts to refer to the domain name confrigured in values.yaml . by default, it would be itscout.local
- Nginx ingress controller has to be enabled for docker desktop. this is not avaiablable by default. You may use ingresscontroller\deploy.yaml for deploying the ingress contreoller using this command 
  - `kubectl apply -f <Path_to_repo>/ingresscontroller/deploy.yaml`

### Installation
- `cd helm/itscout `
- `helm dependency update `
- `helm install itscout . `

### Uninstallation
- `helm delete itscout `

### Accessing the Applicaiton
#### Get the application URL by running these commands:
- http://chart-example.local/healthz
- http://chart-example.local/api/cis

## Limitations
### Here are few limitations that you should be aware of while using this code
* HTTPS is not implemented.
* Pstgresql is deployed as k8s deployment and not a statefullset
* Use of local-path.