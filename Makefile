#Easy way for commit & push
#make git msg="your commit message"
git:
	git add .
	git commit -m "$m"
	git push

#Start the project manually
#make start
start:
	go run main.go

#For runing the current program in a developer mode and make changes
#make developer mode
developer mode:
	CompileDaemon -command="./go-blog"