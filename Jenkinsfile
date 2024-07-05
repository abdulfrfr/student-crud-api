pipeline{
    agents any

    stages{
        stage('checkout'){

        }

        stage('build'){
            steps{
                sh '''
                docker build -t student-crud-api .
                '''
            }
            
        }
    }
}