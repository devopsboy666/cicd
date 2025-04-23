@Library('my-sharelib') _

pipeline {

    parameters {
        string(name: 'GIT_URL', defaultValue: 'https://github.com/devopsboy666/cicd.git', description: 'Git URL')
        string(name: 'BRANCH_NAME', defaultValue: 'main', description: 'Git branch to build')
        string(name: 'IMAGE_NAME', defaultValue: 'goapp', description: 'Image Name for Application')
        string(name: 'IMAGE_TAG', defaultValue: '1.0.0', description: 'Image Tag for Application')
        string(name: 'RELEASE', defaultValue: 'goapp', description: 'Helm Release Name')
    }

    agent { label 'k8s' }

    stages {

        stage('Init') {
            steps {
                script {
                    def gitUrl = params.GIT_URL
                    def repoName = gitUrl.tokenize('/').last().replace('.git', '')
                    def pathHelm = repoName + "/helm"

                    // สร้าง object แล้วเก็บใส่ env หรือ global var ก็ได้
                    env.REPO_NAME = repoName
                    env.PATH_HELM = pathHelm

                    // หรือสร้าง object แล้ว pass ต่อใน stage ถัดไป
                    // ex: currentBuild.description = repoName
                }
            }
        }

        stage('Clone Repo') {
            steps {
                script {
                    def gitClone = new com.demo.CloneRepo([steps: this, branch: params.BRANCH_NAME, repoUrl: params.GIT_URL])
                    gitClone.gitCloneRepo()
                }
            }
        }

        stage('Build Image') {
            steps {
                script {
                    def buildImage = new com.demo.BuildImage([
                        steps: this,
                        imageName: params.IMAGE_NAME,
                        imageTag: params.IMAGE_TAG,
                        dockerfilePath: env.REPO_NAME
                    ])
                    buildImage.dockerBuild()
                }
            }
        }

        stage('Deploy') {
            steps {
                script {
                    try {
                        input message: "Deploy App?", ok: "Approve"
                    } catch (err) {
                        echo "User ยกเลิกการ Deploy: ${err}"
                        currentBuild.result = 'ABORTED'
                        return
                    }

                    def deployApp = new com.demo.DeployApp([
                        steps: this,
                        releaseName: params.RELEASE,
                        imageName: params.IMAGE_NAME,
                        imageTag: params.IMAGE_TAG,
                        pathHelm: env.PATH_HELM
                    ])
                    deployApp.helmDeploy()
                }
            }
        }
    }

    post {
        always {
            echo "🧹 Cleaning up workspace..."
            cleanWs()
        }
    }
}
