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
	SignalStart     = iota //the signal which used to notify the download-routines to begin downloading
	SignalStop             //when the download routines receive the stop-signal,it will stop the download process
	SignalResume           //resume the download process
	SignalRestart          //restart the whole download process,all the data downloaded will be DELETED
	SignalTerminate        //terminated the whole download process, the data downloaded will be RETAINED

)

const (
	StatusInitialized = iota //the main process is not running
	StatusRunning            //the main process is running
	StatusStopped
	StatusFinished
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

func NewDownloader(saver Saver, autoRun bool) *Downloader {
	downloader := &Downloader{
		saver:   saver,
		channel: make(chan int),
		AutoRun: autoRun,
		status:  StatusInitialized,
	}

	//Auto run the main loop to download the awesome go project information by go routine
	//NOTICE, at this moment, the main-entry data may not be ready.
	if downloader.AutoRun {
		go downloader.MainLoop()
	}

	return downloader
}

func (d *Downloader) Start() {
	log.Println("[Downloader] Start")

	//initialize the data to be crawled  by go routines
	d.initData()

	//The downloader can be run by mainly two ways.
	//1.Auto run when the downloader created and the autoRun parameter is true
	//2.After the creation process, the client can call the Start method to activate the main loop.
	if d.status == StatusInitialized {
		go d.MainLoop()
	}

	//Send the start signal to the channel
	d.channel <- SignalStart
}

//MainLoop the crawl process
func (d *Downloader) MainLoop() {
	for {
		select {
		case signal := <-d.channel:
			switch signal {
			case SignalStart:
				log.Println("[Downloader] Start running process...")
			case SignalStop:
				log.Println("[Downloader] Stop running process...")
			case SignalResume:
				log.Println("[Downloader] Resume running process...")
			case SignalTerminate:
				log.Println("[Downloader] Terminate running process...")
			case SignalRestart:
				log.Println("[Downloader] Restart running process...")
			}
		default:
			if d.status != StatusRunning {
				d.status = StatusRunning
			}
			d.DoCrawl()
		}
	}
}

func (d *Downloader) Stop() {
	log.Println("[Downloader] Stop")

	d.channel <- SignalStop
}

func (d *Downloader) Resume() {
	log.Println("[Downloader] Resume")

	d.channel <- SignalResume
}

func (d *Downloader) Terminate() {
	log.Println("[Downloader] Terminate")

	d.channel <- SignalTerminate
}

func (d *Downloader) Restart() {
	log.Println("[Downloader] Restart")

	d.channel <- SignalRestart
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
	case SignalStart:
		d.Start()
	case SignalStop:
		d.Stop()
	case SignalResume:
		d.Resume()
	case SignalTerminate:
		d.Terminate()
	case SignalRestart:
		d.Restart()
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
func (d *Downloader) DoCrawl() {
	log.Printf("[Do Crawl] running,[%d/%d]", d.FinishedProjectsCount, d.ProjectsCount)

	time.Sleep(time.Second * 1)
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
