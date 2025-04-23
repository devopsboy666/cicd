@Library('my-sharelib') _


// Parameter
parameters {
    string(name: 'GIT_URL', defaultValue: 'https://github.com/devopsboy666/cicd.git', description: 'Git URL')
    string(name: 'BRANCH_NAME', defaultValue: 'main', description: 'Git branch to build')
    string(name: 'IMAGE_NAME', defaultValue: 'goapp', description: 'Image Name for Application')
    string(name: 'IMAGE_TAG', defaultValue: '1.0.0', description: 'Image Name for Application')
    string(name: 'RELEASE', defaultValue: 'goapp', description: 'Helm Release Name')
}



// ประกาศตัวแปร
def gitUrl = params.GIT_URL
def repoName = gitUrl.tokenize('/').last().replace('.git', '')
def pathHelm = repoName + "/helm"



// ประการ Class
def gitClone = new com.demo.CloneRepo([steps: this, branch: 'main', repoUrl: repoUrl])
def buildImage = new com.demo.BuildImage([steps: this, imageName: params.IMAGE_NAME, imageTag: params.IMAGE_TAG, dockerfilePath: repoName])
def deployApp = new com.demo.DeployApp([steps: this, releaseName: params.RELEASE ,imageName: params.IMAGE_NAME, imageTag: params.IMAGE_TAG, pathHelm: pathHelm])



pipeline {
    agent { label 'k8s' }
    stages {

        stage('Check Connection Libs'){
            steps {
                script {
                    try {
                        gitClone.gitCloneRepo()
                    } catch (err) {
                        error(err.message)
                    }
                }
            }
        }

        stage('Build Image'){
            steps {
                script {
                    try {
                        buildImage.dockerBuild()
                    } catch(err) {
                        error(err.message)
                    }
                }
            }
        }

        stage('Deploy App'){
            steps {
                script {
                    try {
                        input message: "Deploy App?", ok: "Approve"
                    } catch (err) {
                        echo "User กดยกเลิกการ Deploy: ${err}"
                        currentBuild.result = 'ABORTED'  // ไม่ถือว่า fail แค่ยกเลิกเฉยๆ
                        return
                    }

                    try {
                        deployApp.helmDeploy()
                    } catch (err) {
                        error(err.message)
                    }
                }
            }
        }   
    }
}