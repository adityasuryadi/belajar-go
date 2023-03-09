pipeline {
    agent {
        docker {
            image 'golang:alpine'
        }
    }
    stages {
        stage('Build') {
            steps {
                sh './build.sh'
            }
        }
        stage('Test') { 
            steps {
                sh './test.sh' 
            }
        }
        stage('Manual Approval'){
            steps{
                input message: 'Lanjutkan ke tahap Deploy?'
            }
        }
        stage('Deploy') {
            steps{
                sh './deploy.sh'
            }
        }
    }
}