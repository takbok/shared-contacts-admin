GOAPP := /home/kumar/personal/BhaiData/GAE/go_appengine/goapp

deploy:
	$(GOAPP) deploy app.yaml

serve:
	$(GOAPP) serve app.yaml

# admin.user@your-domain.com, password
