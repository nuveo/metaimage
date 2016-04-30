package metaimage

import "testing"

func TestGetMetadata(t *testing.T) {
	_, err := GetMetadata("./tests/funny_lazy_cat-wallpaper-1280x1024.jpg")
	if err != nil {
		t.Logf("%#v\n", err)
	}

	metadata, err := GetMetadata("./tests/IGP2768W.jpg")
	if err != nil {
		t.Log(err)
	}
	t.Log(metadata)
}
