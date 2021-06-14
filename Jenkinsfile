pipeline {
  environment {
    PROJECT = "mytime-316618"
    APP_NAME = "mytimeapi"
    FE_SVC_NAME = "fullstack-app-mysql" //"${APP_NAME}-frontend"
    CLUSTER = "jenkins-cd"
    CLUSTER_ZONE = "us-east1-d"
    IMAGE_TAG = "gcr.io/${PROJECT}/${APP_NAME}:${env.BRANCH_NAME}.${env.BUILD_NUMBER}"
    //IMAGE_TAG = "gcr.io/${PROJECT}/${APP_NAME}:latest"
    JENKINS_CRED = "${PROJECT}"
  }

  agent {
    kubernetes {
      label 'sample-app'
      defaultContainer 'jnlp'
      yaml """
apiVersion: v1
kind: Pod
metadata:
labels:
  component: ci
spec:
  # Use service account that can deploy to all namespaces
  serviceAccountName: cd-jenkins
  containers:
  - name: golang
    image: golang:1.10
    command:
    - cat
    tty: true
  - name: gcloud
    image: gcr.io/cloud-builders/gcloud
    command:
    - cat
    tty: true
  - name: kubectl
    image: gcr.io/cloud-builders/kubectl
    command:
    - cat
    tty: true
"""
}
  }
  stages {
    stage('Build and push image with Container Builder') {
      steps {
        container('gcloud') {
          dir("apitasks") {
            sh "pwd"
            sh "ls"
           // sh "PYTHONUNBUFFERED=1 gcloud builds submit -t ${IMAGE_TAG} ."
        
          }
        }
      }
    }
    stage('Deploy dev') {
      // Canary branch
      when { branch 'develop' } /*or when {branch 'master'}*/
      steps {
        container('kubectl') {
            dir("apitasks") {
               sh "pwd"
               sh "ls"
              // Change deployed image in canary to the one we just built
              sh("sed -i.bak 's#gcr.io/mytime-316618/mytimeapi:latest#${IMAGE_TAG}#' ./k8s/dev/*.yaml")
              // step([$class: 'KubernetesEngineBuilder', namespace:'production', projectId: env.PROJECT, clusterName: env.CLUSTER, zone: env.CLUSTER_ZONE, manifestPattern: 'k8s/services', credentialsId: env.JENKINS_CRED, verifyDeployments: false])
              step([$class: 'KubernetesEngineBuilder', namespace:'default', projectId: env.PROJECT, clusterName: env.CLUSTER, zone: env.CLUSTER_ZONE, manifestPattern: 'k8s/dev', credentialsId: env.JENKINS_CRED, verifyDeployments: false])
              //sh("export LB_IP=`kubectl get  service/fullstack-mysql \ -o jsonpath='{.status.loadBalancer.ingress[].ip}'` ")
              sh("kubectl get services")
            }
        }
      }
    }
    stage('Deploy Production') {
      // Canary branch
      when { branch 'master' }
      steps {
        container('kubectl') {
            dir("apitasks") {
               sh "pwd"
               sh "ls"
              // Change deployed image in canary to the one we just built
              sh("sed -i.bak 's#gcr.io/mytime-316618/mytimeapi:latest#${IMAGE_TAG}#' ./k8s/production/*.yaml")
              // step([$class: 'KubernetesEngineBuilder', namespace:'production', projectId: env.PROJECT, clusterName: env.CLUSTER, zone: env.CLUSTER_ZONE, manifestPattern: 'k8s/services', credentialsId: env.JENKINS_CRED, verifyDeployments: false])
              step([$class: 'KubernetesEngineBuilder', namespace:'default', projectId: env.PROJECT, clusterName: env.CLUSTER, zone: env.CLUSTER_ZONE, manifestPattern: 'k8s/production', credentialsId: env.JENKINS_CRED, verifyDeployments: false])
              //step([$class: 'KubernetesEngineBuilder', namespace:'default', projectId: env.PROJECT, clusterName: env.CLUSTER, zone: env.CLUSTER_ZONE, manifestPattern: 'k8s/services/prod', credentialsId: env.JENKINS_CRED, verifyDeployments: false])

              //sh("export LB_IP=`kubectl get  service/fullstack-mysql \ -o jsonpath='{.status.loadBalancer.ingress[].ip}'` ")
              sh("kubectl get services")
              //sh("export LB_IP=`kubectl get  service/fullstack-mysql \\-o jsonpath='{.status.loadBalancer.ingress[].ip}'`")
            //   env.LB_IP = sh("kubectl get services fullstack-mysql \\-o jsonpath='{.status.loadBalancer.ingress[].ip}'")
             // sh("echo ${LB_IP}")
             // sh("kubectl set env deployment/fullstack-app-mysql  --overwrite DB_HOST=${LB_IP}")
              script{
                sh("kubectl get services fullstack-mysql \\-o jsonpath='{.status.loadBalancer.ingress[].ip}'>lbip.txt")
                def LB_IP = readFile(file: 'lbip.txt')
                sh("echo ${LB_IP}")
                sh("kubectl set env deployment/fullstack-app-mysql  --overwrite DB_HOST=${LB_IP}")
              }
            
            
            }
        }
      }
    }
    /*stage('Deploy Production') {
      // Production branch
      when { branch 'master' }
      steps{
        container('kubectl') {
        // Change deployed image in canary to the one we just built
          sh("sed -i.bak 's#gcr.io/cloud-solutions-images/gceme:1.0.0#${IMAGE_TAG}#' ./apitasks/k8s/production/*.yaml")
          step([$class: 'KubernetesEngineBuilder', namespace:'production', projectId: env.PROJECT, clusterName: env.CLUSTER, zone: env.CLUSTER_ZONE, manifestPattern: 'apitasks/k8s/services', credentialsId: env.JENKINS_CRED, verifyDeployments: false])
          step([$class: 'KubernetesEngineBuilder', namespace:'production', projectId: env.PROJECT, clusterName: env.CLUSTER, zone: env.CLUSTER_ZONE, manifestPattern: 'apitasks/k8s/production', credentialsId: env.JENKINS_CRED, verifyDeployments: true])
          sh("echo http://`kubectl --namespace=production get service/${FE_SVC_NAME} -o jsonpath='{.status.loadBalancer.ingress[0].ip}'` > ${FE_SVC_NAME}")
        }
      }
    }
    stage('Deploy Dev') {
      // Developer Branches
      when {
        not { branch 'master' }
        not { branch 'canary' }
      }
      steps {
        container('kubectl') {
          // Create namespace if it doesn't exist
          sh("kubectl get ns ${env.BRANCH_NAME} || kubectl create ns ${env.BRANCH_NAME}")
          // Don't use public load balancing for development branches
          sh("sed -i.bak 's#LoadBalancer#ClusterIP#' ./k8s/services/frontend.yaml")
          sh("sed -i.bak 's#gcr.io/cloud-solutions-images/gceme:1.0.0#${IMAGE_TAG}#' ./apitasks/k8s/dev/*.yaml")
          step([$class: 'KubernetesEngineBuilder', namespace: "${env.BRANCH_NAME}", projectId: env.PROJECT, clusterName: env.CLUSTER, zone: env.CLUSTER_ZONE, manifestPattern: 'apitasks/k8s/services', credentialsId: env.JENKINS_CRED, verifyDeployments: false])
          step([$class: 'KubernetesEngineBuilder', namespace: "${env.BRANCH_NAME}", projectId: env.PROJECT, clusterName: env.CLUSTER, zone: env.CLUSTER_ZONE, manifestPattern: 'apitasks/k8s/dev', credentialsId: env.JENKINS_CRED, verifyDeployments: true])
          echo 'To access your environment run `kubectl proxy`'
          echo "Then access your service via http://localhost:8001/api/v1/proxy/namespaces/${env.BRANCH_NAME}/services/${FE_SVC_NAME}:80/"
        }
      }
    }*/
  }
}
