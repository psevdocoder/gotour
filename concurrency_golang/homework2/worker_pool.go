package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"sync"
	"time"
)

var actions = []string{"logged in", "logged out", "created record", "deleted record", "updated account"}

type logItem struct {
	action    string
	timestamp time.Time
}

type User struct {
	id    int
	email string
	logs  []logItem
}

func (u User) getActivityInfo() string {
	output := fmt.Sprintf("UID: %d; Email: %s;\nActivity Log:\n", u.id, u.email)
	for index, item := range u.logs {
		output += fmt.Sprintf("%d. [%s] at %s\n", index, item.action, item.timestamp.Format(time.RFC3339))
	}

	return output
}

func main() {
	rand.New(rand.NewSource(time.Now().Unix()))

	startTime := time.Now()

	const userCount, workerCount = 100, 20

	users := make(chan User, userCount)

	wg := &sync.WaitGroup{}

	for i := 0; i < userCount; i++ {
		wg.Add(1)
		go generateUsers(i, users, wg)
	}
	wg.Wait()
	close(users)

	for i := 0; i < workerCount; i++ {
		wg.Add(1)
		go saveUserInfo(i, users, wg)
	}

	wg.Wait()

	fmt.Printf("DONE! Time Elapsed: %.2f seconds\n", time.Since(startTime).Seconds())
}

func saveUserInfo(workerID int, users <-chan User, wg *sync.WaitGroup) {
	defer wg.Done()
	for user := range users {
		fmt.Printf("WRITING FILE FOR UID %d\n", user.id)
		filename := fmt.Sprintf("users/uid%d.txt", user.id)
		file, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE, 0644)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("worker #%d finished\n", workerID)

		_, err = file.WriteString(user.getActivityInfo())
		if err != nil {
			return
		}
		time.Sleep(time.Second * 1)
	}
}

func generateUsers(count int, users chan<- User, wg *sync.WaitGroup) {
	defer wg.Done()

	users <- User{
		id:    count + 1,
		email: fmt.Sprintf("user%d@company.com", count+1),
		logs:  generateLogs(rand.Intn(1000)),
	}
	fmt.Printf("generated user %d\n", count+1)
	time.Sleep(time.Second * 1)

}

func generateLogs(count int) []logItem {
	logs := make([]logItem, count)
	for i := 0; i < count; i++ {
		logs[i] = logItem{
			action:    actions[rand.Intn(len(actions)-1)],
			timestamp: time.Now(),
		}
	}
	return logs
}
