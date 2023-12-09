pipeline {
    // install golang 1.14 on Jenkins node
    agent { node { label 'uptycs' } } 
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
                            string(credentialsId: 'alphacentauri_UPTYCS_CI_SECRET', variable: 'UPTYCS_CI_SECRET'),
                            string(credentialsId: 'alphacentauri_UPTYCS_API_KEY', variable: 'UPTYCS_API_KEY'),
                            string(credentialsId: 'alphacentauri_UPTYCS_API_SECRET', variable: 'UPTYCS_API_SECRET'),
                            string(credentialsId: 'GITHUB_TOKEN', variable: 'GITHUB_TOKEN'),
                            string(credentialsId: 'JENKINS_TOKEN', variable: 'JENKINS_TOKEN'),

                        ]) {
                        // uptycs scanner and its parameters
                        def scannerImage = '267292272963.dkr.ecr.us-east-1.amazonaws.com/uptycs/uptycs-ci:66d564c51c1bea6f9e5f81d6e6a6de022b8e5eca'   
			
                        def scannerImageOpts = [
                        '--rm', '--restart no',
                        "--network host",
			"--env-file uptycs-env.txt",
                        "--env RUN_DISPLAY_URL=${RUN_DISPLAY_URL}",
                        '--volume /var/run/docker.sock:/var/run/docker.sock:rw',
                        '--volume ${WORKSPACE}:/opt/uptycs/cloud',
                        '--env JOB_NAME="${JOB_NAME}"'
                        ].join(' ')
                        // scanner options 
                        def scanArgs = [
                            "scan",
                            '--disable-secrets',
                            '--disable-vulnerabilities',
                            '--disable-malware',
                            '--disable-image-scan',
			    //"--github-checks",
                            //"--jenkins-checks",
                            '--jenkins-token ${JENKINS_TOKEN}',
                            "--image-id 'test:${BUILD_ID}'",
                            //"--insecure",
                            "--ci-runner-type jenkins",
                            '--api-key ${UPTYCS_API_KEY}',
                            '--api-secret ${UPTYCS_API_SECRET}',
                            '--github-token ${GITHUB_TOKEN}',
			    '--jenkins-token ${JENKINS_TOKEN}',
                            '--uptycs-secret ${UPTYCS_CI_SECRET}',
                            '--approved-email-domain uptycs.com',
                            '--config-file uptycs/.uptycs-ci.yml',
			    '--audit'
                        ].join(' ')
			docker.withRegistry('https://267292272963.dkr.ecr.us-east-1.amazonaws.com', 'ecr:us-east-1:uptycs-shared-jenkins') {
                        	sh (script: "docker run ${scannerImageOpts} ${scannerImage} ${scanArgs}")
			}
                    } //script
                } // withCredentials
            } //steps
        } //stage
    } //stages
} // pipeline
