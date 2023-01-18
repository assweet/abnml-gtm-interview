# abnml-gtm-interview - Abnormal Security Take Home Exam
This is a take home exam for Abnormal Security, specifically for a backend API developer.

## Introduction
Abnormal Security has a successful product that secures cloud email systems of 
our customers.

The intent of this exercise is to figure out the technical skill level of the 
candidate, and the design decisions, and choices made.  We are keen to see the
use of unit testing, event driven thinking and leverage the use of docker in the
process.

## Problem Introduction - License Management

The sales team in Abnormal use Salesforce to track opportunities and update 
the system when clients move to purchase our premier email product offering.

To automate the flow of data from salesforce to Abnormal's systems, we use 
Salesforce webhook APIs to read for events triggered when customer data is 
changed.

In this problem, you will need to assume that we have made the arrangement to
subscribe to salesforce customer data change events, and you should provide 
the endpoint that receives incoming data change events from salesforce, and 
perform some simple processing.

You will need to filter the events, and provide a system that maintains 
licenses and stock keeping units (SKUs) purchased by the customer.  Other 
applications in Abnormal will need to be able to read the service provided to
ascertain if the feature is made available to the user logged in.

## Instructions
Candidates are required to:
1. [ ] Fork this repository to their personal github space
2. [ ] Add the code necessary to complete the task
3. [ ] Ensure that the code performs as required
4. [ ] Ensure that unit tests meets at least 50% code coverage
5. [ ] Be prepared to explain decisions taken and possible improvements
6. [ ] Do not spend more than 2 hours on the take home test
7. [ ] Submit by submitting a PR on this repository

The application attached ships with a postgresql DB server, and a kafka 
broker running in the docker compose environment.

Candidates are required to run `docker-compose up -d` before starting
development to ensure they have a working environment.

One of the images spun up is for a golang application that simulates salesforce
periodically sending customer data to the webhook endpoint the candidate must 
code and provide: `http://web/salesforce/webhook`

### Salesforce Customer Data
For the purposes of this take home, please assume that the necessary work has
been done to subscribe to salesforce customer change events, and it has been 
configured to send these events via RESTful POST calls to 
`http://web/salesforce/webhook`

```
type CustomerData struct {
	AccountManager  string
	Id              string
	Name            string
	SalesEngineerId string
	SignedTos       bool
  SKU1            bool
  SKU2            bool
  SKU3            bool
}
```

## Problem(s)
Your task include:

1. [ ] prepare an web service that will service the endpoint 
   `http://web/salesforce/webhook`.
2. [ ] queue a message with an appropriate payload on the kafka broker
3. [ ] read the customer change data from the kafka broker
4. [ ] design the license and SKU data models
5. [ ] save the customer change data in the provided postgresql db
6. [ ] ensure at least 50% code coverage for code developed
7. [ ] (optional) provide an API to allow abnormal applications to read 
   the license data

## Technical Stack
You can choose to complete the test with either python or golang, and should 
not be doing both.

### Common

#### Docker environment
A `Dockerfile` and `docker-compose.yaml` to setup and simulate the environment.

The candidate should be able to run `docker-compose up -d` to simulate running 
the environment.

#### Simulator
A simulator has been added to the repository and docker environment.  This will 
start sending messages to the web component at the endpoint 
`http://web/salesforce/webhook` with HTTP POST, and with the data structure 
provide in the `Salesforce Customer Data` section above.

### Python 3
The stack consists of:
1. python==3.11.1
2. django==4.1
3. psycopg2
4. pytest

A requirements.txt file is provide, and you should set up with a virtual env
and install with the `pip install -r requirements.txt` command to have the basic
libraries installed.

The candidate may choose any framework he is familiar with, and not limit 
himself/herself to using django.

### Golang
To be provided at a later stage

### Evaluation
1. We will require that testing be incorporated to meet at least 50% of code 
   coverage in unit testing.
2. Candidates must submit their work via a github PR on this repository

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


