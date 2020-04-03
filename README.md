#### `docker build -t varyumin/hetdata_uid .` 
#### `docker push varyumin/hetdata_uid` 
#### `helm delete netdata --purge `
#### `helm upgrade netdata -i ./helm/tets_uid --namespace stage-devops  --debug`