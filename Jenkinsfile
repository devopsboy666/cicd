pipeline {
    agent any

    environment {
        // Global environment variable for image tag
        IMAGE_TAG = 'v${BUILD_NUMBER}-dev' // Change this to your desired tag, BUID_NUMBER is number run pipeline
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
            environment {
                // Load Kubernetes token from Jenkins credentials (create form credentials -> kind secret txt)
                K8S_TOKEN = credentials('k8s-token')
                // Load kubernetes rootCA (create form credentials -> kind secret file)
                K8S_CA = credentials('ca-k8s')
                // Kubernetes Server
                K8S_SERVER = credentials('server-k8s') 
            }
            steps {
                script {
                    sh 'echo "Deploying to Kubernetes cluster"'
                    sh 'echo "K8S_TOKEN = ${K8S_TOKEN}"'
                    sh 'echo "K8S_CA = ${K8S_CA}"'
                    sh 'echo "K8S_SERVER = ${K8S_SERVER}"'

                    // Edit IMAGE & Label file deployment, service
                    sh """
                    sed -i "s|\${IMAGE}|${IMAGE_NAME}:${IMAGE_TAG}|g" k8s/deployment.yaml
                    sed -i "s|\${LABEL}|${IMAGE_TAG}|g" k8s/deployment.yaml
                    sed -i "s|\${LABEL}|${IMAGE_TAG}|g" k8s/service.yaml
                    """

                    // Use Kubernetes API with the service account token
                    sh """
                    kubectl apply -f k8s/deployment.yaml --token=${K8S_TOKEN} --server=${K8S_SERVER} --certificate-authority=${K8S_CA}
                    kubectl apply -f k8s/service.yaml --token=${K8S_TOKEN} --server=${K8S_SERVER} --certificate-authority=${K8S_CA}
                    kubectl apply -f k8s/ingress.yaml --token=${K8S_TOKEN} --server=${K8S_SERVER} --certificate-authority=${K8S_CA}
                    """
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
