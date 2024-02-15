pipeline {
  agent any  
  tools { go '1.19' }

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
          withSonarQubeEnv('sonarqube') {
              sh "${scannerHome}/bin/sonar-scanner"
          }
          timeout(time: 10, unit: 'MINUTES') {
              waitForQualityGate abortPipeline: true
          }
      }
    }
    stage('Test') {
      steps {
        // Output will be something like "go version go1.19 darwin/arm64"
        sh 'go version'
        sh 'go test ./... -v -cover'
      }
    }
    stage('Build') {
      agent { dockerfile true }
      steps {
          docker.build("terryhycheng/go-hello-world:latest")
        }
      }
  }
}