package storage

func Seed() {
	var group = Group{Title: "Favorite"}
	Storage.db.FirstOrCreate(&group)
	var text = Text{Value: "Hello world!", GroupId: group.ID}
	Storage.db.FirstOrCreate(&text)
}
