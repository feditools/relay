pipeline {
  environment {
    BUILD_IMAGE = 'gobuild:1.17'
    BUILD_ARGS = '-e GOCACHE=/gocache -e HOME=${WORKSPACE} -v /var/lib/jenkins/gocache:/gocache -v /var/lib/jenkins/go/pkg:/go/pkg'
    PATH = '/go/bin:~/go/bin:/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin:/usr/games:/usr/local/games:/snap/bin:/usr/local/go/bin'
  }

  agent any

  stages {

    stage('Test') {
      agent {
        docker {
          image "${BUILD_IMAGE}"
          args "${BUILD_ARGS}"
          reuseNode true
        }
      }
      steps {
        script {
          withCredentials([
            string(credentialsId: 'codecov-feditools-go-lib', variable: 'CODECOV_TOKEN')
          ]) {
            sh """#!/bin/bash
            go get -t -v ./...
            go test -race -coverprofile=coverage.txt -covermode=atomic ./...
            RESULT=\$?
            bash <(curl -s https://codecov.io/bash)
            exit \$RESULT
            """
          }
        }
      }
    }

//    stage('Check Formatting') {
//      agent {
//        docker {
//          image "${BUILD_IMAGE}"
//          args "${BUILD_ARGS}"
//          reuseNode true
//        }
//      }
//      steps {
//        script {
//          def ret = sh(
//            script: "make check",
//            returnStatus: true
//          )
//          if (ret != 0) {
//            currentBuild.result = 'UNSTABLE'
//          }
//        }
//      }
//    }

  }

}
