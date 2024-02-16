pipeline {
  agent any

  tools {
    go 'go-1.22'
  }

  stages {
    stage('Checkout') {
      steps {
        checkout scm
      }
    }
    
    stage('Create coverage file & sonarqube') {
      environment {
          scannerHome = tool 'SonarQubeScanner'
          PROJECT_NAME = 'go-todo-list-scan'
      }
      steps {
          sh 'go version'
          sh 'go test -coverprofile=coverage.out ./...'
          
          withSonarQubeEnv('Synology Sonar Server') {
            sh "${scannerHome}/bin/sonar-scanner -Dsonar.projectKey=$PROJECT_NAME -Dsonar.sources=. -Dsonar.language=go -Dsonar.go.coverage.reportPaths=./coverage.out -Dsonar.coverage.exclusions=**/*_test.go -Dsonar.exclusions=**/*_templ.go,tailwind.config.js"
          }
          
          timeout(time: 10, unit: 'MINUTES') {
              waitForQualityGate abortPipeline: true
          }
      }
    }
    stage('Build') {
      steps {
          withCredentials([string(credentialsId: 'docker-hub-credentials', variable: 'DOCKER_HUB_CREDENTIALS')]) {
            script {
              docker.withRegistry('https://registry.hub.docker.com', "$DOCKER_HUB_CREDENTIALS") {
                docker.build("docker.io/terryhycheng/go-web:latest").push('latest')
              }
            }
          }
      }
    }
  }
}