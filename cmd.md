Dir webservice :  represents the web service api  can handle the users, add and remove for example
Dir customer : represents the ways how the users call terraform to talk with services 
Dir provider : represents the custom terraform provider that will let terraform talk to web service. 


docker build -t timage .
docker run -it --volume .:/workspace --name tbox timage
# if you stop the container and want to restart it later, run: docker start -ai tbox

go run main.go &
curl -X POST -d "Jane" http://localhost:6251/2
curl http://localhost:6251/2


dependency provider terraform -> in 3_provider

the conf in .terraformrc represents the project is in local and not typically in terraform registry

In docker container -> cp /workspace/.terraformrc /root/  -> show to Terraform where is the provider 


cd /workspace/3_provider
go mod tidy # download dependencies
go build -o ./bin/terraform-provider-myuserprovider

cd /workspace/2_customer
terraform plan
terraform apply -auto-approve
