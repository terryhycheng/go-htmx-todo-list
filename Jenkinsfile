node {
    def registry = 'docker.io/terryhycheng/go-web'
    def dockerImage

    stage('Checkout') {
        checkout scm
    }

    stage('Create coverage file & sonarqube') {
        def scannerHome
        def PROJECT_NAME = 'go-todo-list-scan'
        withEnv(['PATH+SONARQUBE_SCANNER=/path/to/sonar-scanner/bin']) {
            sh 'go version'
            sh 'go test -coverprofile=coverage.out ./...'

            sh "${scannerHome}/bin/sonar-scanner -Dsonar.projectKey=$PROJECT_NAME -Dsonar.sources=. -Dsonar.language=go -Dsonar.go.coverage.reportPaths=./coverage.out -Dsonar.coverage.exclusions=**/*_test.go -Dsonar.exclusions=**/*_templ.go,tailwind.config.js"

            timeout(time: 10, unit: 'MINUTES') {
                def qg = waitForQualityGate()
                if (qg.status != 'OK') {
                    error 'Quality Gate failed: ${qg.status}'
                }
            }
        }
    }

    stage('Build image') {
        dockerImage = docker.build(registry + ":latest")
    }

    stage('Push to Docker Hub') {
        withCredentials([string(credentialsId: 'docker-hub-credentials', variable: 'DOCKER_HUB_CREDENTIALS')]) {
            docker.withRegistry(registry, "$DOCKER_HUB_CREDENTIALS") {
                dockerImage.push()
            }
        }
    }
}
