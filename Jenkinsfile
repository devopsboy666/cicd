pipeline {
    agent any 

    stages {
        stage('Clone Repository') {
            steps {
                // คำสั่งสำหรับ clone repository จาก GitHub
                git 'https://github.com/pakawat116688/cicd.git'
            }
        }

        stage('Build') {
            steps {
                // คำสั่งสำหรับ build Go application
                sh 'go build -o myapp'
            }
        }

        stage('Test') {
            steps {
                // คำสั่งสำหรับ run tests
                sh 'go test ./...'
            }
        }

        stage('Deploy') {
            steps {
                // คำสั่งสำหรับ deploy application หรือการย้ายไฟล์ไปยังเซิร์ฟเวอร์
                sh 'scp -r ./myapp user@server:/path/to/deploy'
            }
        }
    }

    post {
        always {
            // คำสั่งที่รันทุกครั้ง ไม่ว่าจะสำเร็จหรือไม่ เช่น การทำความสะอาด (cleanup)
            cleanWs()
        }
    }
}
