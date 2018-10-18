package ledlib

func NewObjectTulip() LedObject {
	paths := []string{
		"",
		"/asset/image/tulip/tulip1.png",
		"/asset/image/tulip/tulip2.png",
		"/asset/image/tulip/tulip3.png",
		"/asset/image/tulip/tulip3.png",
		"/asset/image/tulip/tulip4.png",
		"/asset/image/tulip/tulip5.png",
		"",
	}
	return NewObjectBitmap(paths)
}
