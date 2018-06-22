pipeline { 
    agent any 
    stages {
        stage('Build') { 
            steps {
                checkout scm
                sh 'docker build -t mikenavarro/scratch-go .'
            }
        }
        stage('Deploy') {
            steps {
                sh 'make publish'
            }
        }
    }
}