pipeline {
  agent any

  stages {
    stage('Checkout') {
      steps {
        checkout scm
      }
    }
    stage('Sonarqube') {
      environment {
          scannerHome = tool 'SonarQubeScanner'
      }
      steps {
          withSonarQubeEnv('Synology Sonar Server') {
              sh "${scannerHome}/bin/sonar-scanner"
          }
          timeout(time: 10, unit: 'MINUTES') {
              waitForQualityGate abortPipeline: true
          }
      }
    }
    stage('Test') {
      steps {
        script {
        docker.image('golang:1.22.0-alpine3.19').inside {
          sh 'go version'
          sh 'go test ./... -v -cover'
        }
        }
      }
    }
    stage('Build') {
      steps {
        script {
          docker.build("terryhycheng/go-hello-world:latest")
        }
      }
      }
  }
}