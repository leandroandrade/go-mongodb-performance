# go-mongodb

### Simple test MongoDB performance.


## Scenario
For the performance test, were used:

 - loadtest: load test write to nodejs. https://www.npmjs.com/package/loadtest
 - 100.000 requests
 - 1000 concurrent

The tests was executed three times. Below the results:

## First block
##### Using: json.NewDecoder(request.Body).Decode()
|                |1 REQUEST|2 REQUEST|3 REQUEST|AVERAGE|
|----------------|-------------------------------|-----------------------------|----------------|----------------|
|Complete requests:|100.000|100.000|100.000|**100.000**|
|Total Time (s):|93,769887331|87,40585918|94,69504|**93,76989**|
|Req. per seconds:|1066|1144|1056|**1066**|
|Mean latency (ms):|929,7|867,4|940,1|**929,7**|

## Second block
##### Using: ioutil.ReadAll | json.Unmarshal
|                |1 REQUEST|2 REQUEST|3 REQUEST|AVERAGE|
|----------------|-------------------------------|-----------------------------|----------------|----------------|
|Complete requests:|100.000|100.000|100.000|**100.000**|
|Total Time (s):|93,902336384|92,574589311|92,574949702|**92,5749497**|
|Req. per seconds:|1065|1080|1080|**1080**|
|Mean latency (ms):|931.8|919|918,3|**918,65**|

## Third block
##### Using: io.Copy - json.Unmarshal
|                |1 REQUEST|2 REQUEST|3 REQUEST|AVERAGE|
|----------------|-------------------------------|-----------------------------|----------------|----------------|
|Complete requests:|100.000|100.000|100.000|**100.000**|
|Total Time (s):|89,82584|97,64605|89,51946|**89,82584**|
|Req. per seconds:|1113|1024|1117|**1113**|
|Mean latency (ms):|891,5|969,5|888,4|**891,5**|
