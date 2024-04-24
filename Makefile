#Easy way for commit & push
#make git msg="your commit message"
git:
	git add .
	git commit -m "$m"
	git push

#Display help options and discriptions
#make help
help:
	go run main.go help

#Start the project manually
#make start
start:
	go run main.go serve

#For runing the current program in a developer mode and make changes
#make developer mode
developer mode:
	CompileDaemon -command="go run main.go serve"

#Runs the Migration process
#make migrate
migrate:
	go run main.go migrate