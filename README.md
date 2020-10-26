Curso de Microsserviços:
# Desenvolvimento de Aplicações Modernas e Escaláveis com Microsserviços

## Módulo: DEVOPS

### Exercício: Kubernetes e hpa
* Autor: Rafael Goulart
  * E-mail cadastrado: rafaelgoulart@residuall.org
  * E-mail pessoal: faelplg@gmail.com
* Nome do artefato: `cm-go-hpa`
* [Repositório do GitHub (este)](https://github.com/faelplg/cm-go-hpa)
* [Imagem no Docker Hub](https://hub.docker.com/r/faelplg/go-hpa)

#### Tarefa 1: Algoritmo em Go, CI e deploy no K8s
Resolução:
1. Criação do projeto `cm-go-hpa`;
2. Implementação do programa do `php-apache-hpa` em go;
3. Desenvolvimento dos testes baseado na função do loop;
4. Criação do repositório e habilitação do gatilho no Cloud Build;
5. Criação do processo de CI executando os testes no Cloud Build ao enviar um PR;
6. Build do programa em Go e push da imagem no Docker Hub;
7. Configuração de um cluster no Kubernetes Engine;
8. Criação do arquivo `go-hpa.yaml` com deployment e service para deploy no K8s;
    - 6 replicas;
    - CPU requests: 50;
    - CPU limits: 100;

#### Tarefas 2 e 3: Implementação do hpa e verificação do autoscaler 
Resolução:
1. Criação do hpa com as características abaixo:
    - A partir de 15% de CPU;
    - Mínimo de pods: 1;
    - Máximo de pods: 6;
2. Criação de POD para verificação se o autoscaler está funcionando corretamente;
    - Comandos testados:

```bash
kubectl run -it loader --image=busybox /bin/sh
```

```bash
while true; do wget -q -O- http://go-hpa.default.svc.cluster.local; done;
```

```bash
watch kubectl get hpa
```