pipeline {
    // install golang 1.14 on Jenkins node
    agent any
    tools {
        go 'go1.20'
    }
    environment {
        GO114MODULE = 'on'
        CGO_ENABLED = 0 
        GOPATH = "${JENKINS_HOME}/jobs/${JOB_NAME}/builds/${BUILD_ID}"
    }
    stages {
        stage("build") {
            steps {
                echo 'BUILD image STARTED'
                sh 'docker build --tag test:${BUILD_ID} . '
            }
        }
    }
}