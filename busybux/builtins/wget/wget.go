package wget

import (
	"flag"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"time"
)

type Option struct {
	helpFlag    bool
	serverFlag  bool
	timeoutFlag int
	outfileFlag string
	uaFlag      string
}

var option = &Option{}

func Main(w io.Writer, args []string) error {
	flagSet := flag.NewFlagSet("wc", flag.ContinueOnError)
	flagSet.BoolVar(&option.helpFlag, "help", false, "show this message.")
	flagSet.BoolVar(&option.serverFlag, "S", false, "print server response")
	flagSet.StringVar(&option.outfileFlag, "O", "", "write documents to FILE")
	flagSet.IntVar(&option.timeoutFlag, "T", 10, "set all timeout values to SECONDS")
	flagSet.StringVar(&option.uaFlag, "U", "Wget/0.0.1", "identify as AGENT instead of Wget/VERSION")
	if err := flagSet.Parse(args); err != nil {
		return err
	}

	if option.helpFlag {
		flagSet.Usage()
		return nil
	}

	return wget(w, flagSet.Args())
}

func wget(w io.Writer, urls []string) error {
	for _, url := range urls {
		f, err := getOutput(w, url)
		if err != nil {
			fmt.Fprintln(w, "wget:", err)
			continue
		}
		defer f.Close()
		r, err := fetch(url)
		if err != nil {
			fmt.Fprintln(w, "wget:", err)
			continue
		}
		defer r.Body.Close()
		if option.serverFlag {
			fmt.Println(option.serverFlag)
			fmt.Fprintf(w, "%s %s\n", r.Proto, r.Status)
			for k, v := range r.Header {
				fmt.Fprintf(w, "%s: %s\n", k, v)
			}
		}
		if _, err := io.Copy(f, r.Body); err != nil {
			fmt.Fprintln(w, "wget:", err)
			continue
		}
	}
	return nil
}

func fetch(url string) (*http.Response, error) {
	timeout := time.Second * time.Duration(option.timeoutFlag)
	client := &http.Client{Timeout: timeout}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("User-Agent", option.uaFlag)
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func getOutput(w io.Writer, uri string) (*os.File, error) {
	if option.outfileFlag == "-" {
		return w.(*os.File), nil
	}
	if option.outfileFlag != "" {
		return os.Create(option.outfileFlag)
	}
	u, err := url.Parse(uri)
	if err != nil {
		return nil, err
	}
	name := filepath.Base(u.Path)
	if len(name) == 0 || name == "." {
		name = "index.html"
	}
	return os.Create(name)
}
