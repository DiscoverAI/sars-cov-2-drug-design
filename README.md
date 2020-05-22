# SARS-CoV-2 drug design
This projects purpose is to create a drug that is likely to bind with the spike proteins of
SARS-CoV-2 and therefore renders it ineffective. 
 
This is a git repository with submodules. To get the submodules run:
```bash
git submodule update --init --recursive
```

# Architecture
This project consists of multiple jobs that process data, train a generative model, which
is then used to train a DQN to create a model which will likely bind to the spike protein of SARS-CoV-2.

This is done on AWS. You will also need to run and maintain a
[MLflow Tracking Server](https://mlflow.org/docs/latest/tracking.html#mlflow-tracking-servers) instance and a
[Apache Airflow](https://airflow.apache.org/) instance. These are already included, and you just need to deploy them
yourself (you can do this locally for example by just running it).

![architecture](https://github.com/DiscoverAI/sars-cov-2-drug-design/raw/master/docs/architecture.png)
