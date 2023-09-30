package photo

type PhotoFormatter struct {
	ID       int           `json:"id"`
	Title    string        `json:"title"`
	Caption  string        `json:"caption"`
	PhotoURL string        `json:"photo_url"`
	User     UserFormatter `json:"user"`
}

type UserFormatter struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func FormatPhoto(photo Photo) PhotoFormatter {
	photoFormatter := PhotoFormatter{}
	photoFormatter.ID = photo.ID
	photoFormatter.Title = photo.Title
	photoFormatter.Caption = photo.Caption
	photoFormatter.PhotoURL = photo.PhotoURL

	user := photo.User
	userFormatter := UserFormatter{}
	userFormatter.ID = user.ID
	userFormatter.Name = user.Name

	photoFormatter.User = userFormatter

	return photoFormatter
}

func FormatPhotos(photos []Photo) []PhotoFormatter {

	photosFormatter := []PhotoFormatter{}

	for _, photo := range photos {
		photoFormatter := FormatPhoto(photo)
		photosFormatter = append(photosFormatter, photoFormatter)
	}

	return photosFormatter

}