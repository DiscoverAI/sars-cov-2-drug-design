# SARS-CoV-2 drug design
This projects purpose is to create a drug that is likely to bind with the spike proteins of
SARS-CoV-2 and therefore renders it ineffective. 
 
This is a git repository with submodules. To get the submodules run:
```bash
git submodule update --init --recursive
```

To pull the changes from the submodules run:
```bash
git pull --recurse-submodules
git submodule update --remote
```

# Architecture
This project consists of multiple jobs that process data, train a generative model, which
is then used to train a DQN to create a model which will likely bind to the spike protein of SARS-CoV-2.

This is done on AWS. You will also need to run and maintain a
[MLflow Tracking Server](https://mlflow.org/docs/latest/tracking.html#mlflow-tracking-servers) instance and a
[Apache Airflow](https://airflow.apache.org/) instance. These are already included, and you just need to deploy them
yourself (you can also do this locally, see __Run observation services locally__ section below).

![architecture](https://github.com/DiscoverAI/sars-cov-2-drug-design/raw/master/docs/architecture.png)

### Run observation services locally
To start ehrmantraut and saul-goodman locally run:
```bash
docker-compose up -d
```
This will start ehrmantraut locally on [http://localhost:5000](http://localhost:5000) and saul-goodman on
 [http://localhost:8080](http://localhost:8080).
 
To configure the services, please make sure a .env file exists with all necessary environment variables (just take a
look at the [docker-compose.yml](https://raw.githubusercontent.com/DiscoverAI/sars-cov-2-drug-design/master/docker-compose.yml))

## Datapipeline
![architecture](https://github.com/DiscoverAI/sars-cov-2-drug-design/raw/master/docs/datapipeline.png)
