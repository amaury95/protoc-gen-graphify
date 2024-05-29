# Release new version

GIT_VERSION=$(shell git describe --tags)
GIT_NEXT_PATCH=$(shell echo $(GIT_VERSION) | awk -F. '{print $$1"."$$2"."$$3+1}')
GIT_NEXT_MINOR=$(shell echo $(GIT_VERSION) | awk -F. '{print $$1"."$$2+1".0"}')
GIT_NEXT_MAJOR=v$(shell echo $(GIT_VERSION) | awk -F. '{print $$1+1".0.0"}')

commit:
	@git commit -am "Release $(version)"

tag:
	@git tag $(version)

push:
	@git push origin master $(version)

release: commit tag push

install: release
	@go install github.com/amaury95/protoc-gen-graphify@$(version)

# Bug fixes
patch:
	@make install version=${GIT_NEXT_PATCH}

# Minor changes: Does not break the API
minor:
	@make install version=${GIT_NEXT_MINOR}

# Major changes: Breaks the API
major:
	@make install version=${GIT_NEXT_MAJOR}