pipeline {
    agent any

    environment {
        // Global environment variable for image tag
        IMAGE_TAG = "v${BUILD_NUMBER}" // Change this to your desired tag, BUID_NUMBER is number run pipeline
        NEXUS_URL = 'http://192.168.1.215:9876'
        IMAGE_NAME = '192.168.1.215:9876/go/gofiber'
        // Load Kubernetes token from Jenkins credentials (create form credentials -> kind secret txt)
    }

    stages {
        stage('Clone Repository') {
            steps {
                // Clone the Git repository
                git branch: 'main', url: 'https://github.com/pakawat116688/cicd.git'
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
                    cat k8s/deployment.yaml
                    sed -i "s|LABEL|${IMAGE_TAG}|g" k8s/service.yaml
                    cat k8s/service.yaml
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
