# Tashilkar Healthchecker

this project is responsible for checking some api's health beside on their configuration.
also you can create new api to be checked with this system.
in the case of danger, we have another service(webhook) on port 3001, that can make us aware of the situtation.(this is telegram webhook so you must have VPN and register on ```tashilkar``` bot on telegram).

## Installation and Runnig

First of all you must have docker installed on your local machine.
then follow these steps(it might take a little bit longer for the first time):

```bash
 docker-compose up -d 
```
after running the above command, the services will be created and you
can register new endpoint.

after registering new endpoint, you can start console command for checking the registered endpoint(s).
simply you can start console command with this line:
```bash
docker exec -d -it tashilkar_health_checker  ./command health_checker 
```

this console command will alert you on a telegram bot when it could not request to registered endpoint.

you must register to that bot:
```https://t.me/tashilkar_bot```

<span style="color:#e0093b; font-weight: bold; font-size: large">
remember, every time you add new endpoint, you must reset the console command.
</span>

# Services tech description
after installation and running, the service will be up on port 3000.
this service has for apis and one console command:

### health checker console:
check the apis health within their interval and save the result in the checked_endpoints collection.
### Set status API:
for start/stopping health checker command.
when you send 0 as status it will stop the console command.
when you send anything but 0 as status it will start the console command.
this staus will be saved in health_checker_availability collection.
### Delete:
for deleting APIs.(will be saved in endpoints collection)
### List:
for getting list of APIs.(will be saved in endpoints collection)
### Register:
for registering new API.(will be saved in endpoints collection)


```json lines
the postman collcation also has been added to service. 
its name is: Tashilkar.postman_collection.json
```

### Unit tests
we have some unit test for testing out critical unit methdos(logics).
```bash
docker exec  -it tashilkar_health_checker go test logic/tests/endpoint_test.go
```

### Database
As we discussed;
the APIs and result of health check would be saved into mongo db.
we have 3 collection inside tashilkar database:
###### endpoints: for created, read, delete APIs.
###### checked_endpoints: the result of every single health checking would be saves on this collection.
###### health_checker_availability: this collection holds the result of manually stop/start health checker.

```
We have config.yaml file on both services.
for example in web hook we can change webhook chat id or token or ...
```
<span style="color:#09e009">
Thank you for reading the readme file.
</span>


## Author
Saeed Rasooli [Linkedin](linkedin.com/in/saeed-rasooli-029527101/)
