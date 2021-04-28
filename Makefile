gen-proto:
	sh ./scripts/gen-proto.sh
clear-proto:
	sh ./scripts/clear-proto.sh
regen-proto: clear-proto gen-proto