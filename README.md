# SARS-CoV-2 drug design

[![Gitter](https://badges.gitter.im/discoverai/community.svg)](https://gitter.im/discoverai/community?utm_source=badge&utm_medium=badge&utm_campaign=pr-badge)

This projects purpose is to create a drug (drug discovery and drug design) that is likely to bind with the spike
proteins of SARS-CoV-2 und not bind with any human cell receptors. This would render the virus ineffective inside the body.

There is a notebook with more [general information](https://github.com/DiscoverAI/laboratory/blob/master/notebooks/covid19/General%20Information.ipynb)
about the virus you can read.

## Setup
Before you start make sure that your public ssh key is available in Github. This is a git repository with submodules. To get the submodules run:
```bash
git submodule update --init --recursive
```

To pull the changes from the submodules run:
```bash
git pull --recurse-submodules && git submodule update --remote
```

If you've never worked with submodules before, active work that changes submodules should happen in those respective repositories. For example, if there are changes needed to `Pinkman`, you will need to clone the repository and make your changes there. Afterwards, you can update the submodules in this directory using `git pull --recurse-submodules && git submodule update --remote` to bring the `HEAD` commit up to date.

## Onboarding
Please visit our [Onboarding doc](ONBOARDING.md).

## Architecture
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

To create a possible candidate drug for SARS-CoV-2 there multiple steps that need to take place.
From a high level overview the [moses dataset](https://github.com/molecularsets/moses) is being used to create a
generative model that can be used for drug discovery. The drug is then being changed in order to bind more likely to 
SARS-CoV-2 with a reinforcement learning model. The result of these models is then being used to create the top 10 drugs
that are the likeliest candidates for further clinical trial testing.

There is a [notebook](https://github.com/DiscoverAI/laboratory/blob/master/notebooks/covid19/moses-dataset.ipynb) that
explores the moses dataset for more information.

In detail the steps are:
1. Normalize and split train and test dataset from moses (pinkman)
2. Use the train and test dataset to train a generative model (walter_white)
3. Use the generative model and a predefined environment to train a reinforcement learning agent (gustavo_fring)
4. Use the trained models and compute the best 10 drug candidates (tuco)

The reinforcement learning agent uses a reinforcement learning environment that uses a docking software to calculate the
docking affinity of the modified drugs.
