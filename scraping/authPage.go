package scraping

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/sclevine/agouti"
)

func Attendance() {
	//webドライバー準備
	if err := godotenv.Load("/Users/n/go/src/go_webscraping_practice/.env"); err != nil {
		log.Fatal(err)
	}

	loginpage := os.Getenv("LOGINPAGE")
	id := os.Getenv("LOGINID")
	password := os.Getenv("LOGINPASSWORD")

	driver := agouti.ChromeDriver(agouti.ChromeOptions(
		"args", []string{
			"--headless",
		}),
	)

	//Driverの開始と新規ページの作成
	if err := driver.Start(); err != nil {
		log.Fatalf("driver : %v", err)
	}
	defer driver.Stop()

	page, err := driver.NewPage()
	if err != nil {
		log.Fatal(err)
	}
	page.Size(1280, 800)

	//ログインページに遷移
	if err := page.Navigate(loginpage); err != nil {
		log.Fatalf("navigate : %v", err)
	}

	if err := page.FindByID("textFormId").Fill(id); err != nil {
		log.Fatalf("navigate : %v", err)
	}

	if err := page.FindByID("textFormPass").Fill(password); err != nil {
		log.Fatalf("navigate : %v", err)
	}

	if err := page.FindByClass("submitForm").Click(); err != nil {
		log.Fatalf("navigate : %v", err)
	}

	//ページ遷移を終えるまで待機
	time.Sleep(2 * time.Second)

	//Homeページから勤怠ページに遷移し出勤報告
	if err := page.FindByClass("kintaibutton").Click(); err != nil {
		log.Fatalf("navigate : %v", err)
	}

	time.Sleep(2 * time.Second)

	if err := page.FindByButton("報告").Click(); err != nil {
		fmt.Print(err)
	}

	time.Sleep(2 * time.Second)

	cd := time.Now().Local().String()
	if err := page.Screenshot(fmt.Sprintf("/Users/n/go/src/go_webscraping_practice/screenshot/%s出勤完了.png", cd)); err != nil {
		log.Fatalf("navigate : %v", err)
	}

}
