all: fullscreen space-switch resize scratchpad

fullscreen:
	go build -tags="-s -w" -o ./bin/yabai-fullscreen ./cmd/yabai-fullscreen

space-switch:
	go build -tags="-s -w" -o ./bin/yabai-space-switch ./cmd/yabai-space-switch

resize:
	go build -tags="-s -w" -o ./bin/yabai-resize ./cmd/yabai-resize

scratchpad:
	go build -tags="-s -w" -o ./bin/yabai-scratchpad ./cmd/yabai-scratchpad

scratchpad-min:
	go build -tags="-s -w" -o ./bin/yabai-scratchpad-min ./cmd/yabai-scratchpad-min

install:
	cp -f ./bin/* ${HOME}/.local/bin/
