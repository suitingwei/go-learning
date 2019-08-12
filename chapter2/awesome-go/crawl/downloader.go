package crawl

import (
	"encoding/json"
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/http"
	"strings"
	"time"
)

const (
	StatusInitialized = iota //the main process is not running
	StatusRunning            //the main process is running
	StatusResume             //the main process is running
	StatusStopped
)

//The downloader is used to download the data from the web page.
//The main routine will create the go routines to do the real download process.
//And the stat routine to monitor the download process.
//Main-Routine
//      |
//      |------ stat-routine
//      |------ download-routines
//      			|------ download-routine-1
//      			|------ download-routine-2
//      			|------ download-routine-3
type Downloader struct {
	AutoRun               bool //whether auto run the main process during the construct
	saver                 Saver
	channel               chan int //the signal used to notify the running goroutine
	status                int      //the current status of the downloader, should be one of the signals
	ProjectsCount         int      //the total count of the awesome-go projects, should be the length of Downloader.Projects
	FinishedProjectsCount int      //the total count of the finished projects.
	Projects              []*AwesomeGoData
	CrawlersCount         int
	crawlerChannels       []chan int
}

//awesome represents a github go project's basic information
type AwesomeGoData struct {
	Url              string    //the url of the awesome-go project on github.com
	Name             string    //the name of the project
	Desc             string    //the description of the project
	Stars            int       //the stars of the project
	Commits          int       //the number of the commits
	LastUpdateTime   time.Time //the last update time of the project
	ContributorCount int       //the number of the contributor who contributed to the project
	HadDownloaded    bool      //whether the project has been downloaded successfully
}

type ControlRequestParams struct {
	Action int `json:"action"` //the action could be: start,stop,resume,terminate,restart
}

type StatRequestParams struct {
	Format string `json:"format"` //the format should be one of: simple,detail
}

func NewDownloader(saver Saver, autoRun bool, crawlersCount int) *Downloader {
	downloader := &Downloader{
		saver:         saver,
		channel:       make(chan int),
		AutoRun:       autoRun,
		status:        StatusInitialized,
		CrawlersCount: crawlersCount,
	}

	//Auto run the main loop to download the awesome go project information by go routine
	//NOTICE, at this moment, the main-entry data may not be ready.
	if downloader.AutoRun {
		downloader.initCrawlers()
	}

	return downloader
}

func (d *Downloader) Start() {
	log.Println("[Downloader] Start")

	//initialize the data to be crawled  by go routines
	d.initData()

	//initialize the crawlers
	d.initCrawlers()

	//Send the start signal to the channel
	for i := 0; i < d.CrawlersCount; i++ {
		d.crawlerChannels[i] <- StatusRunning
	}
}

func (d *Downloader) initCrawlers() {
	//The downloader can be run by mainly two ways.
	//1.Auto run when the downloader created and the autoRun parameter is true
	//2.After the creation process, the client can call the Start method to activate the main loop.
	if d.status == StatusInitialized {
		//update the downloader status
		d.status = StatusRunning

		//declare an array of channels.
		//NOTICE, after this expression, go will create an array of nil channel
		d.crawlerChannels = make([]chan int, d.CrawlersCount)

		//MUST use the for-loop to initialize the channel in array
		for i := 0; i < d.CrawlersCount; i++ {
			d.crawlerChannels[i] = make(chan int)
		}

		//initialized the crawlers
		for i := 0; i < d.CrawlersCount; i++ {
			go d.mainLoop(i, d.crawlerChannels[i])
		}
	}
}

//mainLoop the crawl process
func (d *Downloader) mainLoop(crawlerNum int, controlChannel <-chan int) {
	status := StatusStopped

	for {
		select {
		case status = <-controlChannel:
			switch status {
			case StatusStopped:
				log.Println("[Downloader] Stop running process...")
				status = StatusStopped
			case StatusResume:
				log.Println("[Downloader] Resume running process...")
				status = StatusRunning
			}
		default:
			log.Printf("[Crawler:%d] running status:%d\n", crawlerNum, status)
			time.Sleep(time.Second * 1)

			if status == StatusStopped {
				break
			}
			d.doCrawl()
		}
	}
}

func (d *Downloader) Stop() {
	log.Println("[Downloader] Stop")

	//Send the start signal to the channel
	for i := 0; i < d.CrawlersCount; i++ {
		d.crawlerChannels[i] <- StatusStopped
	}
}

func (d *Downloader) Resume() {
	log.Println("[Downloader] Resume")

	for i := 0; i < d.CrawlersCount; i++ {
		d.crawlerChannels[i] <- StatusResume
	}
}

//Serve the http api to monitor and control the downloader behaviour.
func (d *Downloader) Serve(port string) {

	http.HandleFunc("/control", d.control)

	http.HandleFunc("/stat", d.stat)

	log.Fatalln(http.ListenAndServe(port, nil))
}

func (d *Downloader) control(w http.ResponseWriter, r *http.Request) {
	control := &ControlRequestParams{}

	log.Println("Receive the http-api:control")

	err := json.NewDecoder(r.Body).Decode(control)

	if err != nil {
		log.Println("[Downloader] control: request body is not valid:" + err.Error())
		return
	}

	switch control.Action {
	case StatusInitialized, StatusRunning:
		d.Start()
	case StatusStopped:
		d.Stop()
	case StatusResume:
		d.Resume()
	}

	_, err = w.Write([]byte("OK"))

	if err != nil {
		log.Fatalln("Failed to send response:" + err.Error())
	}

	return
}

func (d *Downloader) stat(w http.ResponseWriter, r *http.Request) {
	log.Println("Receive the http-api:stat")
}

//Do the real download process.
//TODO use to config to set up the concurrency level.
func (d *Downloader) doCrawl() {
	log.Printf("[Do Crawl] running,[%d/%d]", d.FinishedProjectsCount, d.ProjectsCount)

}

//initialize the main entry data of the awesome-go projects from the
//@link https://awesome-go.com
func (d *Downloader) initData() {
	log.Println("[InitData] Start fetching data from awesome-go.com")

	resp, err := http.Get("https://awesome-go.com/")

	if err != nil {
		log.Fatalln("Failed to download the main page:" + err.Error())
	}

	defer resp.Body.Close()

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(resp.Body)

	if err != nil {
		log.Fatalln("Failed to create go query from response body:" + err.Error())
	}

	// Find the review items
	doc.Find("#content > ul").Each(func(groupIndex int, s *goquery.Selection) {
		// For each item found, get the band and title
		s.Find("li > a").Each(func(projectIndex int, project *goquery.Selection) {
			href, ok := project.Attr("href")

			if strings.HasPrefix(href, "https://github.com/") {
				projectName := project.Text()
				if ok {
					//log.Printf("[Info][%d-%d] %s : %s\n", groupIndex, projectIndex, projectName, href)

					_ = append(d.Projects, &AwesomeGoData{
						Url:  href,
						Name: projectName,
					})

					d.ProjectsCount++
				}
			}
		})

	})

	log.Println("[InitData] Finished")
	return
}
