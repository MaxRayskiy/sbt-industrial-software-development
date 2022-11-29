String repoUrl = "https://github.com/MaxRayskiy/sbt-industrial-software-development"

pipeline {
    agent any
    tools {
        go '1.19.3'
    }
    stages {
        stage('Clone') {
            steps {
                git url: repoUrl, branch: "task3"
            }
        }
        stage("buid-lint") {
            steps {
                dir("issue-tracker") {
                    sh 'whoami && go version'
                }
            }
        }
        stage("test") {
            steps {
                dir("issue-tracker") {
                    echo 'BUILD EXECUTION STARTED'
                    sh 'docker compose up -d'
                }
                dir("tests") {
                    sh 'python3 -m pip install -r ./requirements.txt'
                    sh 'python3 -m pytest --alluredir=./reports test.py'
                }
                allure([
              	   includeProperties: false,
              	   reportBuildPolicy: 'ALWAYS',
              	   results: [[path: 'tests/reports']]
            	])
                dir("issue-tracker") {
                    sh 'docker compose down'
                }
            }
        }
    }
}
