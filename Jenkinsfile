pipeline {
    agent any

    stages {
        stage('init') {
            steps {
                echo 'starting'
                echo "${env.BUILD_ID}"
                echo "${env.JENKINS_URL}"
            }
        }
        stage('build') {
            steps {
                echo 'Doing a thing'
                echo 'yay'
            }
        }
        stage('finish') {
            steps {
                echo 'that went well'
            }
        }
    }
}
