package structs

type Window struct {
	ID    int    `json:"id"`
	Pid   int    `json:"pid"`
	App   string `json:"app"`
	Title string `json:"title"`
	Frame struct {
		X float64 `json:"x"`
		Y float64 `json:"y"`
		W float64 `json:"w"`
		H float64 `json:"h"`
	} `json:"frame"`
	Role               string  `json:"role"`
	Subrole            string  `json:"subrole"`
	Tags               string  `json:"tags"`
	Display            int     `json:"display"`
	Space              int     `json:"space"`
	Level              int     `json:"level"`
	Opacity            float64 `json:"opacity"`
	SplitType          string  `json:"split-type"`
	SplitChild         string  `json:"split-child"`
	StackIndex         int     `json:"stack-index"`
	CanMove            bool    `json:"can-move"`
	CanResize          bool    `json:"can-resize"`
	HasFocus           bool    `json:"has-focus"`
	HasShadow          bool    `json:"has-shadow"`
	HasBorder          bool    `json:"has-border"`
	HasParentZoom      bool    `json:"has-parent-zoom"`
	HasFullscreenZoom  bool    `json:"has-fullscreen-zoom"`
	IsNativeFullscreen bool    `json:"is-native-fullscreen"`
	IsVisible          bool    `json:"is-visible"`
	IsMinimized        bool    `json:"is-minimized"`
	IsHidden           bool    `json:"is-hidden"`
	IsFloating         bool    `json:"is-floating"`
	IsSticky           bool    `json:"is-sticky"`
	IsTopmost          bool    `json:"is-topmost"`
	IsGrabbed          bool    `json:"is-grabbed"`
}

type Space struct {
	ID               int    `json:"id"`
	Label            string `json:"label"`
	Index            int    `json:"index"`
	Display          int    `json:"display"`
	Windows          []int  `json:"windows"`
	Type             string `json:"type"`
	Visible          int    `json:"visible"`
	Focused          int    `json:"focused"`
	NativeFullscreen int    `json:"native-fullscreen"`
	FirstWindow      int    `json:"first-window"`
	LastWindow       int    `json:"last-window"`
}

type Config struct {
	Fullscreen struct {
		Full struct {
			Commands []string `yaml:"commands"`
		} `yaml:"full"`
		Restore struct {
			Commands []string `yaml:"commands"`
		} `yaml:"restore"`
	} `yaml:"fullscreen"`
	Close struct {
		Commands struct {
			Pre  []string `yaml:"pre"`
			Post []string `yaml:"post"`
		} `yaml:"commands"`
	} `yaml:"close"`
	Switch struct {
		IfFullscreen struct {
			Commands []string `yaml:"commands"`
		} `yaml:"if-fullscreen"`
		Else struct {
			Commands []string `yaml:"commands"`
		} `yaml:"else"`
	} `yaml:"switch"`
	Scratchpads         []Scratchpad `yaml:"scratchpads"`
	ScratchpadWorkspace int          `yaml:"scratchpad-workspace"`
}

type Scratchpad struct {
	Name        string `yaml:"name"`
	Title       string `yaml:"title"`
	App         string `yaml:"app"`
	OpenCommand string `yaml:"open-command"`
}
