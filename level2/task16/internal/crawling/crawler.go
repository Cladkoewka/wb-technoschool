package crawling

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"sync"

	"golang.org/x/net/html"
)

// Crawler - struct for downloading web page
type Crawler struct {
	root *url.URL
	depth int 
	outDir string 
	visited map[string]bool 
	mu sync.Mutex
	wg sync.WaitGroup
	client *http.Client
}

// NewCrawler - returns new crawler
func NewCrawler(root string, depth int, outDir string) (*Crawler, error) {
	u, err := url.Parse(root)
	if err != nil {
		return nil, err
	}
	return &Crawler{
		root:    u,
		depth:   depth,
		outDir:  outDir,
		visited: make(map[string]bool),
		client:  &http.Client{},
	}, nil
}

// Start - start crawl web page
func (c *Crawler) Start() {
	c.wg.Add(1)
	go c.crawl(c.root, 0)
	c.wg.Wait()
}

func (c *Crawler) crawl(u *url.URL, depth int) {
	defer c.wg.Done()

	if depth > c.depth {
		return
	}
	if u.Host != c.root.Host {
		return
	}

	c.mu.Lock()
	if c.visited[u.String()] {
		c.mu.Unlock()
		return
	}
	c.visited[u.String()] = true
	c.mu.Unlock()

	fmt.Printf("Downloading: %s\n", u)

	resp, err := c.client.Get(u.String())
	if err != nil {
		fmt.Fprintf(os.Stderr, "error fetching %s: %v\n", u, err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		fmt.Fprintf(os.Stderr, "bad status %s: %d\n", u, resp.StatusCode)
		return
	}

	ct := resp.Header.Get("Content-Type")
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error reading %s: %v\n", u, err)
		return
	}

	localPath := c.urlToPath(u)
	err = os.MkdirAll(filepath.Dir(localPath), 0755)
	if err != nil {
		fmt.Fprintf(os.Stderr, "mkdir error: %v\n", err)
		return
	}

	// if HTML - parse
	if strings.HasPrefix(ct, "text/html") {
		newData, links := c.processHTML(u, data)
		err = os.WriteFile(localPath, newData, 0644)
		if err != nil {
			fmt.Fprintf(os.Stderr, "write error: %v\n", err)
			return
		}

		for _, link := range links {
			c.wg.Add(1)
			go c.crawl(link, depth+1)
		}
	} else {
		err = os.WriteFile(localPath, data, 0644)
		if err != nil {
			fmt.Fprintf(os.Stderr, "write error: %v\n", err)
		}
	}
}

func (c *Crawler) urlToPath(u *url.URL) string {
	path := u.Path
	if path == "" || strings.HasSuffix(path, "/") {
		path = path + "index.html"
	}
	return filepath.Join(c.outDir, u.Host, path)
}

func (c *Crawler) processHTML(base *url.URL, data []byte) ([]byte, []*url.URL) {
	doc, err := html.Parse(strings.NewReader(string(data)))
	if err != nil {
		return data, nil
	}
	var links []*url.URL

	var f func(*html.Node)
	f = func(n *html.Node) {
		if n.Type == html.ElementNode {
			var attrKey string
			switch n.Data {
			case "a", "link":
				attrKey = "href"
			case "img", "script":
				attrKey = "src"
			}
			if attrKey != "" {
				for i, a := range n.Attr {
					if a.Key == attrKey {
						link, err := base.Parse(a.Val)
						if err == nil && link.Scheme != "" && link.Host == c.root.Host {
							links = append(links, link)
							n.Attr[i].Val = c.relativeLocalPath(link)
						}
					}
				}
			}
		}
		for ch := n.FirstChild; ch != nil; ch = ch.NextSibling {
			f(ch)
		}
	}
	f(doc)

	var buf strings.Builder
	html.Render(&buf, doc)
	return []byte(buf.String()), links
}

func (c *Crawler) relativeLocalPath(u *url.URL) string {
	local := c.urlToPath(u)
	rel, err := filepath.Rel(filepath.Join(c.outDir, c.root.Host), local)
	if err != nil {
		return local
	}
	return rel
}