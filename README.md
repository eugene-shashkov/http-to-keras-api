### HTTP TO KERAS API

- API Gives possibility to send json data to golang server and send it further to keras python script 
- Is model created checking method
- HTTP method to make prediction
- Docker image ready


> This method gives possibility to train model by sending json data in one request.
> Learning process will be started automaticly
```
POST /trainmodel
```

> "Is model exist" method is checking if learning finished and .h5 file created 
```
POST /ismodelexist
```

> Making prediction with existed model
```
POST /prediction
```