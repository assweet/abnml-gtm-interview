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


## Problem
The application attached ships with a simple postgresql based docker image, 
and a kafka broker.

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


