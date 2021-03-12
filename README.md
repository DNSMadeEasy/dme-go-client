# dme-go-client
 This repository contains the golang client SDK to interact with DNS Made Easy Platform using REST API calls. This SDK is used by [terraform-provider-dme](https://github.com/DNSMadeEasy/terraform-provider-dme).

## Installation ##

Use `go get` to retrieve the SDK to add it to your `GOPATH` workspace, or project's Go module dependencies.


```sh
$go get github.com/DNSMadeEasy/dme-go-client
```

There are no additional dependancies needed to be installed.

## Overview ##
  
* <strong>client</strong> :- This package contains the HTTP Client configuration as well as service methods which serves the CRUD operations on the DNS Objects in DNSMadeEasy platform.

* <strong>models</strong> :- This package contains all the models structs and utility methods for the same.

## How to Use ##

import the client in your go application and retrive the client object by calling client.GetClient() method.
```golang
import github.com/DNSMadeEasy/dme-go-client/client
client.GetClient("apikey", "secretkey")
// or
client.GetClient("apikey", "secretkey", client.Insecure(true/false),client.ProxyUrl(string),client.Sandbox(true/false))
```


Use that client object to call the service methods to perform the CRUD operations on the model objects.

Example,

```golang
    client.Save(obj<type interface>,endpoint<type string>)
```
