# abnml-gtm-interview - Abnormal Security Take Home Exam
This is a take home exam for Abnormal Security, specifically for a backend API developer.

## Introduction
Abnormal Security has a successful product market fit of securing cloud email 
systems of our customers.

The intent of this exercise is to figure out the technical skill level of the 
candidate, and the design decisions, and choices made.  We are keen to see the
use of unit testing, event driven thinking and leverage the use of docker in the
process.

## Problem Introduction - License Management

The sales team in Abnormal use Salesforce to track opportunities and update 
the system when clients move to purchase our premier email product offering.

To automate the flow of data from salesforce to Abnormal's systems, we use 
Salesforce webhook APIs to read for events triggered when opportunity data is 
changed.  

These events are filtered, and the licenses of the customer in our 
license management system is then updated to reflect the product Stock Keeping
Units (SKUs) purchased by the customer.

## Instructions
Candidates are required to:
1. [ ] Fork this repository to their personal github space
2. [ ] Add the code necessary to complete the task
3. [ ] Ensure that the code performs as required
4. [ ] Ensure that unit tests meets at least 50% code coverage
5. [ ] Be prepared to explain decisions taken and possible improvements
6. [ ] Do not spend more than 2 hours on the take home test
7. [ ] Submit by submitting a PR on this repository

The application attached ships with a postgresql, and a kafka broker running
in the docker compose environment.

Candidates are required to run `docker-compose up -d` before starting
development to ensure they understand the tasks required.

## Problem
We provided a simple scaffolding for the web server image, and a simple kafka 
producer and consumer library

As input, we run a simple random data generator that generates some simple 
data to simulate salesforce opportunity data being populated.

Your task include:

1. parse the salesforce opportunity data for license updates, and produce the 
   formatted data on kafka

2. read the update license message from the kafka broker, and update the 
   license for the client

3. 

## Simulated API
The code in simulator runs a random data generator to simulate roughly 6 
changes per minute.

The data follows the salesforce API type data, and should be received by the 
URL `http://web/salesforce/webhook`

## Technical Stack
You can choose to complete the test with either python or golang, and should 
not be doing both.

### Common
A `Dockerfile` and `docker-compose.yaml` to setup and simulate the environment.

The candidate should be able to run `docker-compose up -d` to simulate running 
the environment.

The simulator will start sending messages to the web component that forms part
of this test

### Python 3
The stack consists of:
1. python==3.11.1
2. django==4.1
3. psycopg2
4. pytest

A requirements.txt file is provide, and you should set up with a virtual env
and install with the `pip install -r requirements.txt` command to have the basic
libraries installed.

### Evaluation
1. We will require that testing be incorporated to meet at least 50% of code 
   coverage in unit testing.
2. 


## Troubleshooting
### Common Issues

### Python 3
1. Error with installing `psycopg2` on Mac
   Refer to this (thread)[https://stackoverflow.com/questions/66777470/psycopg2-installed-on-m1-running-macos-big-sur-but-unable-to-run-local-server] on stackoverflow for solutions
   ```
   brew install openssl
   export LDFLAGS="-L$(brew --prefix openssl)/lib"
   pip install psycopg2
   ```


