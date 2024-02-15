pipeline {
  agent any
  stages {
    stage('Checkout code') {
      parallel {
        stage('Checkout code') {
          steps {
            git(url: 'https://github.com/terryhycheng/go-htmx-todo-list', branch: 'main')
          }
        }

        stage('') {
          steps {
            withSonarQubeEnv('Synology Sonar Server') {
              waitForQualityGate true
            }

          }
        }

      }
    }

    stage('Run Test') {
      steps {
        dockerNode(image: 'golang:1.22.0-alpine3.18') {
          sh 'go test'
        }

      }
    }

  }
}