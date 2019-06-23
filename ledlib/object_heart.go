package ledlib

func NewObjectHeart() LedObject {
	paths := []string{
		"/asset/image/heart/heart1.png",
		"/asset/image/heart/heart2.png",
		"/asset/image/heart/heart3.png",
		"/asset/image/heart/heart4.png",
		"/asset/image/heart/heart4.png",
		"/asset/image/heart/heart3.png",
		"/asset/image/heart/heart2.png",
		"/asset/image/heart/heart1.png",
	}
	return NewObjectBitmap(paths)
}
