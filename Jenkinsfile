pipeline {
    agent any

    environment {
        IMAGE_TAG = "v${BUILD_NUMBER}" // Change this to your desired tag, BUID_NUMBER is number run pipeline
        NEXUS_URL = 'http://192.168.1.215:9876'
        IMAGE_NAME = '192.168.1.215:9876/go/gofiber'
    }

    stages {
        stage('Clone Repository') {
            steps {
                // Clone the Git repository
                git branch: 'main', url: 'https://github.com/pakawat116688/cicd.git'
            }
        }

        stage('OWASP Dependency Check') {
            steps {
                script {
                    dependencyCheck additionalArguments: '--project "cicd" --out . --format HTML', 
                                    odcInstallation: 'OWASP-Dependency-Check-Vulnerabilities', 
                                    scanpath: './'
                }
                sh 'pwd'
                sh 'ls -ltr'
            }
            post {
                always {
                    script {
                        // Move the report to /tmp directory
                        // sh 'pwd'
                        sh 'ls -ltr'
                        sh "mv dependency-check-report.html /tmp/dependency-check-report-${BUILD_NUMBER}.html"
                        // sh "mv dependency-check-report.html /tmp/dependency-check-report-${BUILD_NUMBER}.html || true"
                        sh 'ls -ltr /tmp'
                    }

                    // Archive the report as an artifact and publish HTML report
                    archiveArtifacts artifacts: "/tmp/dependency-check-report-${BUILD_NUMBER}.html", allowEmptyArchive: true
                    publishHTML target: [
                        allowMissing: true,
                        alwaysLinkToLastBuild: false,
                        keepAll: true,
                        reportDir: '/tmp/',
                        reportFiles: "dependency-check-report-${BUILD_NUMBER}.html",
                        reportName: 'OWASP Dependency Check Report'
                    ]
                }
            }
        }

        stage('Build Docker Image') {
            steps {
                script {
                    sh 'echo "Building Docker image: ${IMAGE_NAME}:${IMAGE_TAG}"'
                    // Build Docker image with tag
                    docker.build("${IMAGE_NAME}:${IMAGE_TAG}")
                }
            }
        }

        stage('Push Docker Image') {
            steps {
                script {
                    sh 'echo "Push Docker image url: ${NEXUS_URL}"'
                    // Login to Nexus registry
                    docker.withRegistry("${NEXUS_URL}", 'nexus-credentials') {
                        // Push Docker image to Nexus
                        docker.image("${IMAGE_NAME}:${IMAGE_TAG}").push("${IMAGE_TAG}")
                    }
                }
            }
        }

        stage('Deploy to Kubernetes') {
            steps {
                script {

                    sh """
                    sed -i "s|image: ''|image: ${IMAGE_NAME}:${IMAGE_TAG}|g" k8s/deployment.yaml
                    sed -i "s|LABEL|${IMAGE_TAG}|g" k8s/deployment.yaml
                    sed -i "s|LABEL|${IMAGE_TAG}|g" k8s/service.yaml
                    """

                    withKubeConfig([credentialsId: 'context']) {
                        sh 'kubectl apply -f k8s/deployment.yaml'
                        sh 'kubectl apply -f k8s/service.yaml'
                        sh 'kubectl apply -f k8s/ingress.yaml'
                    }
                }
            }
            
        }
    }

    post {
        always {
            // Clean up
            cleanWs()
        }
    }
}
