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
        sh("sed -i.bak 's#imageversion#${imageTag}#' ./k8s/frontend-canary.yaml")
        sh("kubectl --namespace=production apply -f k8s/services.yaml")
        sh("kubectl --namespace=production apply -f k8s/frontend-canary.yaml")
        sh("echo http://`kubectl --namespace=production get service/${feSvcName} --output=json | jq -r '.status.loadBalancer.ingress[0].ip'` > ${feSvcName}")
        break

    // Roll out to production
    case "master":
        sh("sed -i.bak 's#imageversion#${imageTag}#' ./k8s/frontend-production.yaml")
        sh("kubectl --namespace=production apply -f k8s/services.yaml")
        sh("kubectl --namespace=production apply -f k8s/frontend-production.yaml")
        sh("echo http://`kubectl --namespace=production get service/${feSvcName} --output=json | jq -r '.status.loadBalancer.ingress[0].ip'` > ${feSvcName}")
        break

    // Roll out a dev environment
    case "development":
        sh("sed -i.bak 's#imageversion#${imageTag}#' ./k8s/frontend-dev.yaml")
        sh("kubectl --namespace=${env.BRANCH_NAME} apply -f k8s/services.yaml")
        sh("kubectl --namespace=${env.BRANCH_NAME} apply -f k8s/frontend-dev.yaml")
        sh("echo http://`kubectl --namespace=${env.BRANCH_NAME} get service/${feSvcName} --output=json | jq -r '.status.loadBalancer.ingress[0].ip'` > ${feSvcName}")
  }
}
