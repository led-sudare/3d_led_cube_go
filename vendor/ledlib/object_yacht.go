package ledlib

func NewObjectYacht() LedObject {
	paths := []string{
		"/asset/image/yacht/yacht1.png",
		"/asset/image/yacht/yacht2.png",
		"/asset/image/yacht/yacht3.png",
		"/asset/image/yacht/yacht4.png",
		"/asset/image/yacht/yacht4.png",
		"/asset/image/yacht/yacht3.png",
		"/asset/image/yacht/yacht2.png",
		"/asset/image/yacht/yacht1.png",
	}
	return NewObjectBitmap(paths)
}
