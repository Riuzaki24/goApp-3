package main

import "fmt"

type bookmarkMap = map[string]string 

func main() {
	
	

	bookmarks := bookmarkMap{}
	fmt.Println("Приложение для закладок")
	Menu:
	for {
		variant := getMenu()
		switch variant {
		case 1:
			printBookmarks(bookmarks)
		case 2:
			addBookmarks(bookmarks)
		case 3:
			deleteBookmarks(bookmarks)
		case 4:
			break Menu
		}
	}

}


func getMenu() int{
	var variant int
	fmt.Println("Выберете вариант")
	fmt.Println("1. Посмотреть закладки")
	fmt.Println("2. Добавить закладку")
	fmt.Println("3. Удалить закладку")
	fmt.Println("4. Выход")
	fmt.Scan(&variant)
	return variant
}

func printBookmarks(bookmarks bookmarkMap) {
	if len(bookmarks) == 0 {
		fmt.Println("Заметки отсутсвует")
	}
	for key, value := range bookmarks {
		fmt.Println(key, ":", value)
	}
}

func addBookmarks(bookmarks bookmarkMap){
	var newBookmarkKey string
	var newBookmarkValue string
	fmt.Println("Введите название: ")
	fmt.Scan(&newBookmarkKey)
	fmt.Println("Введите ссылку: ")
	fmt.Scan(&newBookmarkValue)
	bookmarks[newBookmarkKey] = newBookmarkValue
}

func deleteBookmarks(bookmarks bookmarkMap){
	var bookmarkKeyToDelete string 
	fmt.Print("Введите название: ")
	fmt.Scan(&bookmarkKeyToDelete)
	delete(bookmarks, bookmarkKeyToDelete)
}
                