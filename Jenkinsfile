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
            
                script {
		    sh (script: "set > uptycs-env.txt")
                    withCredentials([
                            string(credentialsId: 'UPTYCS_CI_SECRET', variable: 'UPTYCS_CI_SECRET'),
                            string(credentialsId: 'UPTYCS_API_KEY', variable: 'UPTYCS_API_KEY'),
                            string(credentialsId: 'UPTYCS_API_SECRET', variable: 'UPTYCS_API_SECRET'),
                            string(credentialsId: 'GITHUB_TOKEN', variable: 'GITHUB_TOKEN'),
                            string(credentialsId: 'JENKINS_TOKEN', variable: 'JENKINS_TOKEN'),

                        ]) {
                        // uptycs scanner and its parameters
                        def scannerImage = 'uptycs/uptycs-ci:latest-aarch64'             
                        def scannerImageOpts = [
                        '--rm', '--restart no',
			"--env-file uptycs-env.txt",
                        "--env RUN_DISPLAY_URL=${RUN_DISPLAY_URL}",
                        '--volume /var/run/docker.sock:/var/run/docker.sock:ro',
                        '--volume /Users/usirsiwal/work/uptycs/testrepo/uptycs:/opt/uptycs/cloud',
                        '--env JOB_NAME="${JOB_NAME}"'
                        ].join(' ')
                        // scanner options 
                        def scanArgs = [
                            "scan",
			                "--github-checks",
                            "--jenkins-checks",
                            '--jenkins-token ${JENKINS_TOKEN}',
                            "--image-id 'test:${BUILD_ID}'",
                            "--ci-runner-type jenkins",
                            '--api-key ${UPTYCS_API_KEY}',
                            '--api-secret ${UPTYCS_API_SECRET}',
                            '--github-token ${GITHUB_TOKEN}',
                            '--uptycs-secret ${UPTYCS_CI_SECRET}',
                        ].join(' ')
                        sh (script: "docker run ${scannerImageOpts} ${scannerImage} ${scanArgs}")
                    } //script
                } // withCredentials
            } //steps
        } //stage
    } //stages
} // pipeline
