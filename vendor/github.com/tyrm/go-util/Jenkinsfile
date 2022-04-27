pipeline {
  agent any

  stages {

    stage('Test') {
      agent {
        docker {
          image 'gobuild:1.17'
          args '-e GOCACHE=/gocache -e HOME=${WORKSPACE} -v /var/lib/jenkins/gocache:/gocache -v /var/lib/jenkins/go:/go'
        }
      }
      steps {
        script {
          withCredentials([
            string(credentialsId: 'codecov-tyrm-go-util', variable: 'CODECOV_TOKEN')
          ]) {
            sh """#!/bin/bash
            go get -t -v ./...
            go test --tags=integration -race -coverprofile=coverage.txt -covermode=atomic ./...
            RESULT=\$?
            gosec -fmt=junit-xml -out=gosec.xml  ./...
            bash <(curl -s https://codecov.io/bash)
            exit \$RESULT
            """
          }
          junit allowEmptyResults: true, checksName: 'Security', testResults: "gosec.xml"
        }
      }
    }

  }

}
