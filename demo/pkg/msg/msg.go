package msg

// needs to be accessed from teh very root, look in the gomod file to see the name at the root is coursecontent
// vscode will generate this for you if you just start typing and save
import "coursecontent/demo/pkg/display"

// public function (capitalised)
func Hi() {
	display.Display("Hi")
}
