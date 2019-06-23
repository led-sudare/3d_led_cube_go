package ledlib

func NewObjectStickman() LedObject {
	paths := []string{
		"",
		"/asset/image/stickman/stickman2.png",
		"/asset/image/stickman/stickman3.png",
		"/asset/image/stickman/stickman4.png",
		"/asset/image/stickman/stickman4.png",
		"/asset/image/stickman/stickman5.png",
		"/asset/image/stickman/stickman2.png",
	}
	return NewObjectBitmap(paths)
}
