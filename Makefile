# Generates mocks for interfaces
INTERFACES_GO_FILES := $(shell find internal -name "interfaces.go")
INTERFACES_GEN_GO_FILES := $(INTERFACES_GO_FILES:%.go=%.mock.gen.go)

# Generate mocks for interfaces
init: $(INTERFACES_GEN_GO_FILES)
$(INTERFACES_GEN_GO_FILES): %.mock.gen.go: %.go
	mockgen -source=$< -destination=$@ -package=$(shell basename $(dir $<))

clean:
	find . -name "*.mock.gen.go" -type f -delete

test_unit:
	go test -short -coverprofile=c.out ./$(file) && go tool cover -html=c.out