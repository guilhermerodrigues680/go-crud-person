hello:
	@echo "Use \`make help\` para obter instruções."

build:
	scripts/build.sh

help: Makefile
	@echo
	@echo " Comandos disponiveis"
	@echo
	@echo " make build"
	@echo

# disallow any parallelism (-j) for Make. This is necessary since some
# commands during the build process create temporary files that collide
# under parallel conditions.
.NOTPARALLEL:

.PHONY: hello build help
