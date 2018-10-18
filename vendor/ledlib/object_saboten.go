package ledlib

func NewObjectSaboten() LedObject {
	paths := []string{
		"/asset/image/saboten/saboten0.png",
		"/asset/image/saboten/saboten1.png",
		"/asset/image/saboten/saboten2.png",
		"/asset/image/saboten/saboten3.png",
		"/asset/image/saboten/saboten3.png",
		"/asset/image/saboten/saboten2.png",
		"/asset/image/saboten/saboten1.png",
		"/asset/image/saboten/saboten0.png",
	}
	return NewObjectBitmap(paths)
}
