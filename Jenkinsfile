pipeline{
    agents any

    stages{
        stage('checkout'){

        }

        stage('build'){
            sh '''
                docker build -t student-crud-api .
            '''
        }
    }
}