package pages

import "fmt"

func htmlTemplates(pageName string) []string {
	page := fmt.Sprintf("./pages/html/%s.html", pageName)
	return []string{"./pages/html/header.html", "./pages/html/nav.html", page, "./pages/html/footer.html"}
}
