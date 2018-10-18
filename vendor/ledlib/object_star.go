package ledlib

func NewObjectStar() LedObject {
	paths := []string{
		"",
		"/asset/image/star/star2.png",
		"/asset/image/star/star3.png",
		"/asset/image/star/star4.png",
		"/asset/image/star/star4.png",
		"/asset/image/star/star3.png",
		"/asset/image/star/star2.png",
		"",
	}
	return NewObjectBitmap(paths)
}
