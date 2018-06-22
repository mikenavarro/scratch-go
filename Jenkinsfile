pipeline {
node {
    checkout scm

    def customImage = docker.build("my-image:latest")

    customImage.inside {
        sh 'make test'
    }
}
}