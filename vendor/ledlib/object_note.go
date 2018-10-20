package ledlib

func NewObjectNote() LedObject {
	paths := []string{
		"/asset/image/note/eighth1.png",
		"/asset/image/note/eighth2.png",
		"/asset/image/note/eighth3.png",
		"/asset/image/note/eighth4.png",
		"/asset/image/note/eighth4.png",
		"/asset/image/note/eighth3.png",
		"/asset/image/note/eighth2.png",
		"/asset/image/note/eighth1.png",
	}
	return NewObjectBitmap(paths)
}
