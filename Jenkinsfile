@Library('my-sharelib') _


// ประกาศตัวแปร
def repoUrl = "https://github.com/devopsboy666/cicd.git"
def repoName = repoUrl.tokenize('/').last().replace('.git', '')


// ประการ Class
def gitClone = new com.demo.CloneRepo([steps: this, branch: 'main', repoUrl: repoUrl])
def buildImage = new com.demo.BuildImage([steps: this, imageName: 'goapp', imageTag: 'demo', dockerfilePath: repoName])


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
    }
}