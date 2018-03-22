node {
  def project = 'cd-jenkins-193814'
  def appName = 'gowebapp'
  def feSvcName = "${appName}"
  def imageTag = "gcr.io/${project}/${appName}:${env.BRANCH_NAME}.${env.BUILD_NUMBER}"

  checkout scm

  stage 'Build image'
  sh("docker build -t ${imageTag} .")
  
  stage 'Run Go tests'
  sh("docker run -d --net='host' ${imageTag} go run GoWebApp.go")
  sh("docker run --net='host' ${imageTag} go test") 
  
  stage 'Push image to registry'
  sh("gcloud docker -- push ${imageTag}")

  stage "Deploy Application"
  switch (env.BRANCH_NAME) {
    // Roll out to canary environment
    case "canary":
        // Change deployed image in canary to the one we just built
        sh("sed -i.bak 's#gcr.io/cloud-solutions-images/gceme:1.0.0#${imageTag}#' ./k8s/canary/*.yaml")
        sh("kubectl --namespace=production apply -f k8s/services/")
        sh("kubectl --namespace=production apply -f k8s/canary/")
        sh("echo http://`kubectl --namespace=production get service/${feSvcName} --output=json | jq -r '.status.loadBalancer.ingress[0].ip'` > ${feSvcName}")
        break

    // Roll out to production
    case "master":
        sh("kubectl get ns production || kubectl create ns production")
   
        // Change deployed image in canary to the one we just built
        sh("sed -i.bak 's#gcr.io/cloud-solutions-images/ceme:1.0.0g#${imageTag}#' ./k8s/production/*.yaml")
        sh("kubectl --namespace=production apply -f k8s/services/")
        sh("kubectl --namespace=production apply -f k8s/production/")
        sh("echo http://`kubectl --namespace=production get service/${feSvcName} --output=json | jq -r '.status.loadBalancer.ingress[0].ip'` > ${feSvcName}")
        break

    // Roll out a dev environment
    //default:
    case "development":
        sh("sed -i.bak 's#gcr.io/cloud-solutions-images/gceme:1.0.0#${imageTag}#' ./k8s/dev/*.yaml")
        sh("kubectl --namespace=${env.BRANCH_NAME} apply -f k8s/services/")
        sh("kubectl --namespace=${env.BRANCH_NAME} apply -f k8s/dev/frontend-dev.yaml")
        //echo 'To access your environment run `kubectl proxy`'
        //echo "Then access your service via http://localhost:8001/api/v1/proxy/namespaces/${env.BRANCH_NAME}/services/${feSvcName}:80/"
        sh("echo http://`kubectl --namespace=development get service/${feSvcName} --output=json | jq -r '.status.loadBalancer.ingress[0].ip'` > ${feSvcName}")
  }
}
