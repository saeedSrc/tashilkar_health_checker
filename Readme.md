# Tashilkar Healthchecker

this project is responsible for checking some api's health beside on their configuration.
also you can create new api to be checked with this system.
in the case of danger, we have another service(webhook) on port 3001, that can make us aware of the situtation.(this is telegram webhook so you must have VPN and register on ```tashilkar``` bot on telegram).

## Installation and Runnig

first of all you must have docker installed on your local machine.
then follow these steps:

```bash
pip install foobar
```
# Services tech description
after installation and running, the service will be up on port 3000.
this service has for apis and one console command:

##### health checker console: check the apis health within their interval.
##### Set status API: for start/stopping health checker command.
##### Delete: for deleting APIs.
##### List: for getting list of APIs.
##### Register:  for registering new API.


```json lines
the postman collcation also has been added to service. 
its name is: Tashilkar.postman_collection.json
```

Also we have some unit test for testing out critical unit methdos.

The APIs and result of health check would be saved into mongo db.
we have 3 collection inside tashilkar database:

###### endpoints: for created, read, delete APIs.
###### checked_endpoints: the result of every single health checking would be saves on this collection.
###### health_checker_availability: this collection holds the result of manually stop/start health checker.

```
We have config.yaml file on both services.
for example in web hook we can change webhook chat id or token or ...
```
<span style="color:red">

</span>


Thank you for reading the readme file.

## Author
Saeed Rasooli [Linkedin](linkedin.com/in/saeed-rasooli-029527101/)
