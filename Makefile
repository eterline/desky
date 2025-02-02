.PHONY: test

backend-app = desky-backend
front-app = desky-front

clear:
	rm -rf desky* web *.json app logging trace.log run.sh

clone:
	git clone 'https://github.com/eterline/$(backend-app).git'
	git clone 'https://github.com/eterline/$(front-app).git'

del-git:
	rm -rf $(backend-app)
	rm -rf $(front-app)

build:
	./builder.sh
	

init: clear clone build del-git

deploy: init
	./deploy.sh

.DEFAULT_GOAL := test