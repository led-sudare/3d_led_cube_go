package ledlib

func NewObjectSnowman() LedObject {
	paths := []string{
		"/asset/image/snowman/snowman1.png",
		"/asset/image/snowman/snowman2.png",
		"/asset/image/snowman/snowman3.png",
		"/asset/image/snowman/snowman3.png",
		"/asset/image/snowman/snowman3.png",
		"/asset/image/snowman/snowman3.png",
		"/asset/image/snowman/snowman4.png",
		"/asset/image/snowman/snowman5.png",
	}
	return NewObjectBitmap(paths)
}
