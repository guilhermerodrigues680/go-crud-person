hello:
	@echo "Use \`make help\` para obter instruções."

build:
	scripts/build.sh

cross:
	scripts/build-cross.sh

help:
	@echo
	@echo " Comandos disponiveis"
	@echo
	@echo " \`make build\` Constroe apenas o binário do Linux."
	@echo " \`make cross\` Constroe todos os binários para todas as plataformas."
	@echo

# disallow any parallelism (-j) for Make. This is necessary since some
# commands during the build process create temporary files that collide
# under parallel conditions.
.NOTPARALLEL:

.PHONY: hello build cross help
