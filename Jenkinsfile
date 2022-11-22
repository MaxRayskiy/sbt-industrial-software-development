String repoUrl = "https://github.com/MaxRayskiy/sbt-industrial-software-development"

pipeline {
    agent any
    tools { go '1.19.3' }
    stages {
        stage('Clone') {
            steps {
                git url: repoUrl, branch: "task3"
            }
        }
        stage("build") {
            steps {
                dir("issue-tracker") {
                    echo 'BUILD EXECUTION STARTED'
                    sh 'whoami && go version'
                    sh 'docker compose up -d'
                    sh 'docker compose down'
                }
            }
        }
    }
}
