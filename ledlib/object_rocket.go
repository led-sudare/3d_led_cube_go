package ledlib

func NewObjectRocket() LedObject {
	paths := []string{
		"/asset/image/rocket/rocket1.png",
		"/asset/image/rocket/rocket2.png",
		"/asset/image/rocket/rocket3.png",
		"/asset/image/rocket/rocket4.png",
		"/asset/image/rocket/rocket4.png",
		"/asset/image/rocket/rocket3.png",
		"/asset/image/rocket/rocket2.png",
		"/asset/image/rocket/rocket1.png",
	}
	return NewObjectBitmap(paths)
}
