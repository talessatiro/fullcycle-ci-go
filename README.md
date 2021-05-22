## Aplicação de exemplo contendo configurações para Integração Contínua usando o Github Actions

Aplicação go simples que contem a definição de uma função de soma e um teste para cobrir essa função.

### A configuração de CI (arquivo ci.yaml) inclui:

- Realizar o checkout do projeto
- Validar execução do arquivo math.go
- Validar execução dos testes
- Build para geração de uma imagem docker contendo a aplicação
- Push da imagem para o DockerHub (Necessário configuração das secrets DOCKERHUB_USERNAME e DOCKERHUB_TOKEN no repositório)

### A configuração do SonarQube inclue:

- Instalar SonarQube local via docker: **docker run -d --name sonarqube -e SONAR_ES_BOOTSTRAP_CHECKS_DISABLE=true -p 9000:9000 sonarqube:latest**
- Criar um novo projeto no SonarQube
- Configurar o sonar-scanner no seu ambiente
- Executar o sonar-scanner para varrer seu projeto e enviar os outputs de resultado para o serviço do Sonar
  - sonar-scanner -Dsonar.projectKey=projectKey -Dsonar.sources=. -Dsonar.host.url=http://localhost:9000 -Dsonar.login=projectToken