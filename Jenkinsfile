pipeline { 
    agent any 
    stages {
        stage('Build') { 
            steps {
                checkout scm
                def customImage = docker.build("my-image:latest")
            }
        }
        stage('Deploy') {
            steps {
                sh 'make publish'
            }
        }
    }
}